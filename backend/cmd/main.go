package main

import (
	"log"
	"lyked-backend/cmd/server"
)

func main() {
	if err := server.InitServer(); err != nil {
		log.Fatal("Failed to initialize server:", err)
	}

}
