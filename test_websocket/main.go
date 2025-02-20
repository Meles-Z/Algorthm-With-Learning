package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// create send websocket
var upgrade = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		allowedOrigins := []string{
			"http://localhost:8080",
			"http://localhost:80",
		}
		origin := r.Header.Get("Origin")
		for _, allowedOrg := range allowedOrigins {
			if origin == allowedOrg {
				return true
			}
		}
		return false
	},
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrade.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Upgrade error:", err)
		return
	}
	defer conn.Close()

	for {
		msgType, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Error to read message:", err)
			break
		}
		log.Printf("Recieved message:%s\n", string(message))
		if err := conn.WriteMessage(msgType, message); err != nil {
			log.Println("Error to write message:", err)
			break
		}

	}
}

func main() {
	http.HandleFunc("/echo", echoHandler)
	log.Println("Server started on 8080")
	http.ListenAndServe(":8080", nil)

}
