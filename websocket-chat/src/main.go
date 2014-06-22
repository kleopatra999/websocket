package main

import (
	"log"
	//"net/http"

	"chat"
)

func init() {
	log.SetFlags(log.Lshortfile)

	// websocket server
	server := chat.NewServer("/entry")
	go server.Listen()

	// static files
	// http.Handle("/", http.FileServer(http.Dir("webroot")))

	// log.Fatal(http.ListenAndServe(":8080", nil))
}
