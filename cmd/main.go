package main

import (
	. "chat-example/app"
	. "chat-example/pkg/handler"
	"log"
	"net/http"
)

func main() {
	hub := NewHub()
	go hub.Run()
	http.HandleFunc("/messages", func(w http.ResponseWriter, r *http.Request) {
		ServeWs(hub, w, r)
	})

	http.HandleFunc("/", GetMainPage)

	if err := http.ListenAndServe(":8090", nil); err != nil {
		log.Fatal(err)
	}
}
