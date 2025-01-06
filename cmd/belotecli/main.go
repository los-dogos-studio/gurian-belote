package main

import (
	"fmt"
	"github.com/los-dogos-studio/gurian-belote/cli"
)

func main() {
	fmt.Println("Welcome to Gurian Belote!")

	cli.StartBeloteGame(200)
}
