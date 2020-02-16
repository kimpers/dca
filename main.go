package main

import (
	"log"
	"math/rand"
	"os"
	//"strconv"
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

		err := client.Buy()

		if err != nil {
			log.Printf("Failed to complete order. Error: %v", err)
			return
		}

		//amount, _ := strconv.ParseFloat(order.Funds, 64)

		//log.Printf("Order completed of %s for %s (%s)", strconv.FormatFloat(amount, 'f', -1, 64), order.ProductId, order.Id)

}
