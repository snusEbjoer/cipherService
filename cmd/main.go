package main

import (
	"log"

	"github.com/snusEbjoer/cipherService/internal/server"
)

func main() {
	log.Println("Server started")
	server.NewServer().Run()
}
