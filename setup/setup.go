package setup

import (
	"log"
	"time"

	"github.com/8BITS-COLAB/ballot-box/seed"
	"github.com/magefile/mage/sh"
)

func Run() {
	log.Println("Setting up ballot box...")
	log.Println()

	log.Println("Droping database...")

	if err := sh.RunV("docker-compose", "down"); err != nil {
		log.Fatal(err)
	}
	log.Println("Droping database...done")
	log.Println("Up database...")

	if err := sh.RunV("docker-compose", "up", "--build", "-d"); err != nil {
		log.Fatal(err)
	}
	log.Println("Up database...done")

	time.Sleep(time.Second * 1)

	log.Println("Seeding database...")
	seed.Up()
	log.Println("Seeding database...done")
}
