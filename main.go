package main

import (
	"github.com/vivcis/library-app/server"
	"log"
)

func main() {

	err := server.Start()
	if err != nil {
		log.Fatalf("server error: %v", err)
	}
}
