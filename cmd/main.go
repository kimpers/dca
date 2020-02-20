package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	iftb "github.com/kimpers/dca"
	"github.com/kimpers/dca/lib"
)

func main() {
	numArgs := len(os.Args)
	if numArgs < 2 {
		log.Fatal("ERROR: Please specify command serve or encode-config")
	}

	command := os.Args[1]
	if command == "serve" {
		http.HandleFunc("/", iftb.RequestHandler)
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
