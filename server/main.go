package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/websocket"
)

func main() {
	var port = flag.Int("port", 8080, "Port that server listens on")
	flag.Parse()

	http.HandleFunc("/ws", wsHandler)
	http.HandleFunc("/", rootHandler)

	err := http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
	if err != nil {
		log.Println("Error serving web app.", err)
	}
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	var upgrader = websocket.Upgrader{}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket connection error.", err)
		return
	}
	defer conn.Close()

	if r.Header == nil || r.Header["Content-Type"] == nil || !contains(r.Header["Content-Type"], "text/plain") {
		log.Println("Content-Type is not correct. Header:", r.Header)
		return
	}
	log.Println("Client created a WebSocket.")

	waitForRequests(conn)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Works")
	w.Write([]byte("Works"))
}

func waitForRequests(conn *websocket.Conn) {
	for {
		mType, data, err := conn.ReadMessage()
		if err != nil {
			log.Println("Failed to read message.", err)
			return
		}
		if mType != websocket.TextMessage {
			log.Printf("Message was not text. Type: %d.\n", mType)
			return
		}

		message := string(data)
		if message != "" {
			log.Println("Message received:", message)
			replacedMessage := replaceQuestionMarks(message)
			conn.WriteMessage(websocket.TextMessage, []byte(replacedMessage))
		}
	}
}

// replaceQuestionMarks replaces ? with !
func replaceQuestionMarks(msg string) string {
	return strings.Replace(msg, "?", "!", -1)
}

func contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}
