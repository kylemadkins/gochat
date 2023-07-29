package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kylemadkins/gochat/chat"
)

const (
	PORT = "8000"
)

func main() {
	log.Printf("Starting server on port %s...\n", PORT)

	go chat.BroadcastMessages()

	err := http.ListenAndServe(fmt.Sprintf(":%s", PORT), routes())
	if err != nil {
		log.Fatalf("Unable to start server on port %s: %v\n", PORT, err)
	}
}
