package main

import (
	"log"
	"math/rand"
	"os"
	"sync"
	"time"

	"github.com/kimpers/dca/lib"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Incorrect params. Correct usage: dca [path to config.yaml]")
	}

	// Seed randomness for picking coin to purchase
	rand.Seed(time.Now().UnixNano())

	configPath := os.Args[1]

	config := lib.Config{}
	config.ReadFile(configPath)

	client := lib.NewClient(&config)

	scheduledBuy := func() {
		order, err := client.Buy()

		if err != nil {
			log.Printf("Failed to complete order. Error: %v", err)
			return
		}

		log.Printf("Order completed %v", order)
	}

	lib.Schedule(config.Schedule, scheduledBuy)

	// Wait forever
	wg := &sync.WaitGroup{}
	wg.Add(1)
	wg.Wait()
}
