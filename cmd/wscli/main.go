package main

import (
	"flag"
	"log"

	"github.com/los-dogos-studio/gurian-belote/wscli"
)

func main() {
	userId := flag.String("userId", "", "User ID")
	ip := flag.String("ip", "localhost", "IP address of the server")
	port := flag.Int("port", 8080, "Port of the server")
	flag.Parse()

	if *userId == "" {
		log.Fatal("User ID is required")
	}

	client := wscli.NewWsCli(*userId, *ip, *port)

	log.Fatal(client.Run())
}
