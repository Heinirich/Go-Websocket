package main

import (
	"fmt"
	"golang.org/x/net/websocket"
	"log"
	"math/rand"
	"time"
)

func main() {
	connection, err := websocket.Dial("ws://localhost:8085", "", createIP())

	if err != nil {
		log.Fatal(err)
	}

}

func createIP() string {
	var ip [4]int

	for i := 0; i < len(ip); i++ {
		rand.Seed(time.Now().UnixNano())

		ip[i] = rand.Intn(255)
	}
	return fmt.Sprintf("http://%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
}
