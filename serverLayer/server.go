package main

import (
	"golang.org/x/net/websocket"
	"net/http"
)

func main() {
	// Define new Server mux
	muxServer := http.NewServeMux()

	muxServer.Handle("/", websocket.Handler(func(ws *websocket.Conn) {
	}))
	server := http.Server{
		Addr:    "127.0.0.1:8085",
		Handler: muxServer,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
