package main

import (
	"fmt"
	"time"

	"github.com/8BITS-COLAB/ballot-box/vote"
)

func main() {
	fmt.Println("Agent is starting...")
	for range time.Tick(time.Second * 10) {
		vote.CheckIntegrity(func(v vote.Vote) {
			fmt.Printf("Vote %d is valid\n", v.Index)
		})

		fmt.Println()
	}
}
