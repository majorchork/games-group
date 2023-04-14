package main

import (
	"fmt"
	"github.com/majorchork/tech-crib-africa/cmd/server"
	"log"
)

func main() {
	fmt.Println("Hello, world!")

	server.Start()
	log.Println("successfully disconnected")
}
