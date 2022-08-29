package main

import (
	"fmt"
	"github.com/vivcis/library-app/server"
	"log"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.Lmicroseconds)
	fmt.Println("Starting... Library")

	err := server.Start()
	if err != nil {
		log.Fatalf("server error: %v", err)
		return
	}
}
