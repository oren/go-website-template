package main

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

const listenAddr = "localhost:4000"

func main() {
	http.HandleFunc("/", serveHome)
	http.HandleFunc("/static/", static)
	http.HandleFunc("/ws", serveWs)
	err := http.ListenAndServe(listenAddr, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func static(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, r.URL.Path[1:])
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func serveWs(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		if _, ok := err.(websocket.HandshakeError); !ok {
			log.Println(err)
		}
		return
	}

	message := []byte{'h', 'e', 'l', 'l', 'o'}

	if err = ws.WriteMessage(websocket.TextMessage, message); err != nil {
		return
	}

	for {
		messageType, p, err := ws.ReadMessage()
		if err != nil {
			return
		}

		s := string(p[:])
		log.Println(s)

		if err = ws.WriteMessage(messageType, p); err != nil {
			return
		}
	}
}
