package main

import (
	"log"

	"github.com/los-dogos-studio/gurian-belote/server"
)

func main() {
	log.Println("Welcome to Gurian Belote!")
	log.Println("Starting server...")
	log.Fatal(server.NewServer().Start())
}
