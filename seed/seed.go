package seed

import (
	"crypto/rand"
	"fmt"
	"log"
	"math/big"

	"github.com/8BITS-COLAB/ballot-box/candidate"
	"github.com/8BITS-COLAB/ballot-box/db"
	"github.com/8BITS-COLAB/ballot-box/peer"
	"github.com/8BITS-COLAB/ballot-box/vote"
	"github.com/8BITS-COLAB/ballot-box/voter"
	"github.com/jaswdr/faker"
)

var f = faker.New()

var d = db.New()

func Up() {
	d.AutoMigrate(&peer.Peer{})
	d.AutoMigrate(&candidate.Candidate{})
	d.AutoMigrate(&voter.Voter{})
	d.AutoMigrate(&vote.Vote{})

	var cs []candidate.Candidate

	for i := 0; i < 10; i++ {
		cs = append(cs, candidate.Candidate{
			Name:  f.Person().FirstName(),
			Party: f.Lorem().Word(),
		})
	}

	for _, c := range cs {
		c = *candidate.New(c.Name, c.Party)
		log.Printf("candidate %s with code %s created\n", c.Name, c.Code)

	}

	log.Println("Creating voter...")
	i, _ := rand.Int(rand.Reader, big.NewInt(1000000))
	r := fmt.Sprintf("%d", i)
	k := "123456"
	v, ks, err := voter.New(r, k)

	if err != nil {
		log.Fatalf("failed to create voter: %s", err)
	}

	log.Printf("Voter created:\n- registry: %s\n- address: %s\n- secret key: %s\n- private key: %s\n", v.Registry, v.Address, k, ks.PrivateKey)
	log.Println("Creating voter...done")
	log.Println()

	log.Println("Setting up ballot box...done")
}
