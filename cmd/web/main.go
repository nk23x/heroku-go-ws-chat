package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/SajjadManafi/ws-chat/internal/handlers"
)

func main() {

	port := os.Getenv("PORT")

	if len(port) == 0 {
		port = "8080"
	}

	mux := routes()

	log.Println("Starting channel listener...")
	go handlers.ListenToWsChannel()

	log.Printf("Starting server on port %s\n", port)

	_ = http.ListenAndServe(fmt.Sprintf(":%s", port), mux)
}
