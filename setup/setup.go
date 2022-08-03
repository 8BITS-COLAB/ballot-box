package setup

import (
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	"os/exec"
	"time"

	"github.com/8BITS-COLAB/ballot-box/seed"
	"github.com/8BITS-COLAB/ballot-box/voter"
)

func Run() {
	fmt.Println("Setting up ballot box...")
	fmt.Println()

	fmt.Println("Droping database...")
	_, err := exec.Command("/bin/sh", "-c", "docker-compose down").Output()
	fmt.Println("Droping database...done")
	fmt.Println("Up database...")
	_, err = exec.Command("/bin/sh", "-c", "docker-compose up --build -d").Output()
	fmt.Println("Up database...done")

	if err != nil {
		log.Fatalf("failed to initialize database: %s", err)
	}

	time.Sleep(time.Second)

	fmt.Println("Seeding database...")
	seed.Up()
	fmt.Println("Seeding database...done")

	fmt.Println("Creating voter...")
	i, _ := rand.Int(rand.Reader, big.NewInt(1000000))
	r := fmt.Sprintf("%d", i)
	k := "123456"
	v, ks, err := voter.New(r, k)

	if err != nil {
		log.Fatalf("failed to create voter: %s", err)
	}

	fmt.Printf("Voter created:\n- registry: %s\n- address: %s\n- secret key: %s\n- private key: %s\n", v.Registry, v.Address, k, ks.PrivateKey)
	fmt.Println("Creating voter...done")
	fmt.Println()

	fmt.Println("Setting up ballot box...done")
}
