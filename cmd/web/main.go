package main

import (
	"github.com/ismail118/simple-chat/internal/handlers"
	"log"
	"net/http"
)

func main() {
	// websocket channel listener
	log.Println("Starting channel listener")
	go handlers.ListenToWsChannel()

	// http server
	mux := routes()

	log.Println("start http server on port 8080")

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
