package peer

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/8BITS-COLAB/ballot-box/candidate"
	"github.com/8BITS-COLAB/ballot-box/db"
	"github.com/8BITS-COLAB/ballot-box/vote"
	"github.com/8BITS-COLAB/ballot-box/voter"
	"github.com/bytedance/sonic"
)

type Peer struct {
	Addr string `json:"addr" gorm:"primaryKey;uniqueIndex"`
}

var d = db.New()

func Listen(port string) {
	server, err := net.Listen("tcp", fmt.Sprintf(":%s", port))

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

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
	// var cs []candidate.Candidate

	// if err := d.Find(&cs).Error; err != nil {
	// 	log.Fatalf("failed to find candidates: %v", err)
	// }

	go func() {
		fs := http.FileServer(http.Dir("view"))

		http.Handle("/", fs)

		http.HandleFunc("/api/peers/", func(w http.ResponseWriter, r *http.Request) {
			if r.Method == http.MethodGet {
				var ps []Peer

				if err := d.Find(&ps).Error; err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					return
				}

				jason, _ := sonic.Marshal(ps)

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				w.Write(jason)
			}
		})

		http.HandleFunc("/api/candidates/", func(w http.ResponseWriter, r *http.Request) {
			if r.Method == http.MethodGet {
				code := strings.TrimPrefix(r.URL.Path, "/api/candidates/")

				if code != "" {
					c, err := candidate.GetByCode(code)

					if err != nil {
						w.WriteHeader(http.StatusInternalServerError)
						return
					}

					jason, _ := sonic.Marshal(c)

					w.Header().Set("Content-Type", "application/json")
					w.WriteHeader(http.StatusOK)
					w.Write(jason)
					return
				}

				var cs []candidate.Candidate

				if err := d.Find(&cs).Error; err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					return
				}

				jason, _ := sonic.Marshal(cs)

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				w.Write(jason)
			}
		})

		http.HandleFunc("/api/status/", func(w http.ResponseWriter, r *http.Request) {
			s := vote.Status()

			jason, _ := sonic.Marshal(s)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(jason)
		})

		http.HandleFunc("/api/votes", func(w http.ResponseWriter, r *http.Request) {
			if r.Method == http.MethodPost {
				var body map[string]string

				pvk := r.Header.Get("X-Voter-Private-Key")
				sk := r.Header.Get("X-Voter-Secret-Key")

				if pvk == "" || sk == "" {
					w.WriteHeader(http.StatusBadRequest)
					return
				}

				if err := sonic.ConfigDefault.NewDecoder(r.Body).Decode(&body); err != nil {
					w.WriteHeader(http.StatusBadRequest)
					return
				}

				v, err := vote.New(pvk, body["candidate_code"], sk)

				if err != nil {
					w.WriteHeader(http.StatusBadRequest)
					return
				}

				jason, _ := sonic.Marshal(v)

				w.WriteHeader(http.StatusCreated)
				w.Write(jason)
			}
		})

		http.HandleFunc("/api/login/", func(w http.ResponseWriter, r *http.Request) {
			if r.Method == http.MethodPost {
				var body map[string]string

				if err := sonic.ConfigDefault.NewDecoder(r.Body).Decode(&body); err != nil {
					w.WriteHeader(http.StatusBadRequest)
					return
				}

				pvk := body["private_key"]
				sk := body["secret_key"]

				v, err := voter.Show(pvk, sk)

				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					return
				}

				jason, _ := sonic.Marshal(v)

				w.WriteHeader(http.StatusOK)
				w.Write(jason)
			}
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
	var peers []Peer
	if err := d.Find(&peers).Error; err != nil {
		log.Fatalf("failed to find peers: %v", err)
	}

	for _, peer := range peers {
		fmt.Printf("connecting to %s\n", peer.Addr)
		_, err := net.Dial("tcp", peer.Addr)
		if err != nil {
			log.Fatalf("failed to dial: %v", err)
		}

		// go func(conn net.Conn) {
		// 	for {
		// 		rw := bufio.NewReadWriter(bufio.NewReader(os.Stdin), bufio.NewWriter(conn))

		// 		for {
		// 			fmt.Printf("Enter message: ")
		// 			msg, err := rw.ReadString('\n')
		// 			if err != nil {
		// 				log.Fatalf("failed to read: %v", err)
		// 			}

		// 			if _, err := rw.WriteString(msg); err != nil {
		// 				log.Fatalf("failed to write: %v", err)
		// 			}

		// 			if err := rw.Flush(); err != nil {
		// 				log.Fatalf("failed to flush: %v", err)
		// 			}
		// 		}
		// 	}
		// }(conn)
	}
}
