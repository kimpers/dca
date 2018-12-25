package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
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

	fmt.Printf("--- m:\n%v\n\n", config)
	order, err := client.Buy()
	log.Printf("%v", err)
	log.Printf("%v", order)

}
