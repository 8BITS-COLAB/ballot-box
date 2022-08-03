package peer

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/8BITS-COLAB/ballot-box/candidate"
	"github.com/8BITS-COLAB/ballot-box/db"
	"github.com/aymerick/raymond"
)

type Peer struct {
	Addr string `json:"addr" gorm:"primaryKey;uniqueIndex"`
}

func init() {
	d, sql := db.New()
	defer sql.Close()

	d.AutoMigrate(&Peer{})
}

func Listen(port string) {
	server, err := net.Listen("tcp", fmt.Sprintf(":%s", port))

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	d, sql := db.New()
	defer sql.Close()
	defer server.Close()

	peer := Peer{Addr: server.Addr().String()}

	if err := d.Create(&peer).Error; err != nil {
		log.Fatalf("failed to create peer: %v", err)
	}

	fmt.Printf("listening on %s\n", port)

	// go func() {
	// 	for {
	// 		conn, err := server.Accept()
	// 		if err != nil {
	// 			log.Fatalf("failed to accept: %v", err)
	// 		}

	// 		go func(conn net.Conn) {
	// 			defer conn.Close()

	// 			buf := make([]byte, 1024)
	// 			for {
	// 				n, err := conn.Read(buf)
	// 				if err != nil {
	// 					log.Println(err)
	// 					return
	// 				}

	// 				fmt.Println(string(buf[:n]))
	// 			}
	// 		}(conn)
	// 	}
	// }()

	source := `
	<div class="entry" style="margin: 0 auto; display: grid; grid-template-columns: repeat(3, 1fr); gap: 10px;">
		{{#each candidates}}
			<div class="candidate" style="padding: 10px; border: 1px solid lightgray;">
				<div class="name">{{name}}</div>
				<div class="party">{{party}}</div>
				<div class="party">{{code}}</div>
			</div>
		{{/each}}
	</div>
	`
	var cs []candidate.Candidate

	if err := d.Find(&cs).Error; err != nil {
		log.Fatalf("failed to find candidates: %v", err)
	}

	go func() {
		result := raymond.MustRender(source, map[string]interface{}{
			"candidates": cs,
		})

		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte(result))
		})

		http.Serve(server, nil)
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	<-ch
	d.Delete(&peer)
	fmt.Println("Ending...")
}

func Connect() {
	d, sql := db.New()
	defer sql.Close()

	var peers []Peer

	if err := d.Find(&peers).Error; err != nil {
		log.Fatalf("failed to find peers: %v", err)
	}

	for _, peer := range peers {
		fmt.Printf("connecting to %s\n", peer.Addr)
		conn, err := net.Dial("tcp", peer.Addr)
		if err != nil {
			d.Delete(&peer)
			log.Fatalf("failed to dial: %v", err)
		}

		go func(conn net.Conn) {
			for {
				rw := bufio.NewReadWriter(bufio.NewReader(os.Stdin), bufio.NewWriter(conn))

				for {
					fmt.Printf("Enter message: ")
					msg, err := rw.ReadString('\n')
					if err != nil {
						log.Fatalf("failed to read: %v", err)
					}

					if _, err := rw.WriteString(msg); err != nil {
						log.Fatalf("failed to write: %v", err)
					}

					if err := rw.Flush(); err != nil {
						log.Fatalf("failed to flush: %v", err)
					}
				}
			}
		}(conn)
	}
}
