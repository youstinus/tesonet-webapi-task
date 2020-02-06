package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
)

func main() {
	address := "ws://localhost:8080/ws"
	header := http.Header{}
	header.Set("Content-Type", "text/plain")
	c, _, err := websocket.DefaultDialer.Dial(address, header)
	if err != nil {
		log.Println("Could not create a WebSocket.", err)
		return
	}
	defer c.Close()

	scanner := bufio.NewScanner(os.Stdin)
	go func() {
		for {
			scanner.Scan()
			text := scanner.Text()
			bytes := []byte(text)
			c.WriteMessage(websocket.TextMessage, bytes)
		}
	}()

	for {
		_, bytes, err := c.ReadMessage()
		if err != nil {
			log.Println("Error reading message.", err)
			return
		}
		fmt.Println(string(bytes))
	}
}
