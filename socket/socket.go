package socket

import (
	"fmt"
	"golang.org/x/net/websocket"
)

type Message struct {
	Subject string `json:"subject"`
}
type Config struct {
	// Clients is a map of client identifiers to their WebSocket connections.
	Clients map[string]*websocket.Conn

	// ClientRegisterChan is a channel used to register new WebSocket connections.
	ClientRegisterChan chan *websocket.Conn

	// UnregisterClient is a channel used to unregister and remove WebSocket connections.
	UnregisterClient chan *websocket.Conn

	// MessageData is a channel used to send and receive messages.
	MessageData chan Message
}

// NewConfig returns a new instance of the Config struct.
// Config is a struct that holds the relevant information for a WebSocket server.
// It contains maps of client identifiers to their WebSocket connections and channels used for registering, unregistering, and sending messages.
func NewConfig() *Config {
	return &Config{
		Clients:            make(map[string]*websocket.Conn),
		ClientRegisterChan: make(chan *websocket.Conn),
		UnregisterClient:   make(chan *websocket.Conn),
		MessageData:        make(chan Message),
	}
}

// RegisterClient registers a new WebSocket connection with the Config struct.
func (c *Config) RegisterClient(client *websocket.Conn) {
	c.Clients[client.RemoteAddr().String()] = client

	fmt.Println("Client registered:", c.Clients)
}

// RemoveClient removes a WebSocket connection from the Config struct.
func (c *Config) RemoveClient(client *websocket.Conn) {
	delete(c.Clients, client.RemoteAddr().String())
	fmt.Println("Client removed:", c.Clients)
	fmt.Println("Clients:", c.Clients)
}

// BroadcastMessage sends a message to all registered WebSocket connections.
func (c *Config) BroadcastMessage(message Message) {
	for _, client := range c.Clients {
		err := websocket.JSON.Send(client, message)
		if err != nil {
			fmt.Println("Error sending message:", err)
		}
	}
}

// RunSocket Method for registering clients
func (c *Config) RunSocket(client *websocket.Conn) {
	for {
		select {
		case ClientRegisterChan := <-c.ClientRegisterChan:
			c.RegisterClient(ClientRegisterChan)
		case UnregisterClient := <-c.UnregisterClient:
			c.RemoveClient(UnregisterClient)
		case MessageData := <-c.MessageData:
			c.BroadcastMessage(MessageData)
		}
	}
}
