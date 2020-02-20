package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/kimpers/dca/lib"
)

func main() {
	numArgs := len(os.Args)
	if numArgs < 2 {
		log.Fatal("ERROR: Please specify command serve or encode-config")
	}

	command := os.Args[1]
	if command == "serve" {
		http.HandleFunc("/", RequestHandler)
		log.Println("Server running on port 8080")
		http.ListenAndServe(":8080", nil)
	} else if command == "encode-config" {
		if numArgs < 3 {
			log.Fatal("USAGE: iftb encode-config path/to/config.json")
		}
		configPath := os.Args[2]

		config := lib.Config{}
		encodedConfig := config.EncodeBase64ConfigFile(configPath)
		fmt.Println("export CONFIG_BASE64=" + encodedConfig)
	}
}

func RequestHandler(w http.ResponseWriter, r *http.Request) {
	// Seed randomness for picking coin to purchase
	rand.Seed(time.Now().UnixNano())

	configBase64 := os.Getenv("CONFIG_BASE64")

	config := lib.Config{}
	// Use CONFIG_BASE64 environment if available
	if len(os.Args) == 3 {
		configPath := os.Args[2]
		config.ReadFile(configPath)
	} else if len(configBase64) > 0 {
		config.DecodeBase64Config(configBase64)
	} else {
		// Try to read local file
		w.WriteHeader(500)
		fmt.Fprintf(w, "unexpected_error")

		log.Fatal("Neither CONFIG_BASE64 env variable or path/to/config.json specified. Please specify either of them")
	}

	// Validate url parameter token
	token := r.URL.Query().Get("token")
	if token != config.Token {
		fmt.Printf("Invalid request token %s\n", token)
		w.WriteHeader(400)
		fmt.Fprintf(w, "invalid_request")
		return
	}

	client := lib.NewClient(&config)

	err := client.Buy()

	if err != nil {
		log.Printf("Failed to complete order. Error: %v", err)
		w.WriteHeader(500)
		fmt.Fprint(w, "Error")
		return
	}

	w.WriteHeader(200)
	fmt.Fprint(w, "success")
}
