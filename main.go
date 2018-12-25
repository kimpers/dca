package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kimpers/dca/lib"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Incorrect params. Correct usage: dca [path to config.yaml]")
	}

	configPath := os.Args[1]

	config := lib.Config{}
	config.ReadFile(configPath)

	client := lib.NewClient(&config)

	fmt.Printf("--- m:\n%v\n\n", config)
	order, err := client.Buy("ETH")
	log.Printf("%v", err)
	log.Printf("%v", order)

}
