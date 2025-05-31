package main

import (
	"Go-WebSockets/socket"
	"golang.org/x/net/websocket"
	"net/http"
)

func main() {

	config := socket.NewConfig()
	// Define new Server mux
	muxServer := http.NewServeMux()

	muxServer.Handle("/", websocket.Handler(func(wsConnection *websocket.Conn) {
		webSocketHandler(wsConnection, config)
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

func webSocketHandler(c *websocket.Conn, config *socket.Config) {
	go config.RunSocket()

	config.ClientRegisterChan <- c

	for {
		var message socket.Message
		err := websocket.JSON.Receive(c, &message)
		if err != nil {
			config.UnregisterClient <- c
			continue
		}
		config.MessageData <- message
	}
}
