package main

import (
	. "chat-example/app"
	. "chat-example/pkg/handler"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	hub := NewHub()
	go hub.Run()
	http.HandleFunc("/messages", func(w http.ResponseWriter, r *http.Request) {
		ServeWs(hub, w, r)
	})

	http.HandleFunc("/", GetMainPage)

	address := fmt.Sprintf(":%s", os.Getenv("APP_PORT"))
	log.Print(address)

	if err := http.ListenAndServe(address, nil); err != nil {
		log.Fatal(err)
	}
}
