package main

import (
	"github.com/gorilla/websocket"

	//"github.com/gorilla/websocket"
	"net/http"
)

var upgrader = websocket.Upgrader{}

func main() {
	http.HandleFunc("/select", selectPort)
	http.ListenAndServe(":8086", nil)
}

func selectPort(writer http.ResponseWriter, req *http.Request) {
	var conn, _ = upgrader.Upgrade(writer, req, nil)

	go func(conn *websocket.Conn) {
		for {
			mType, msg, _ := conn.ReadMessage()
			conn.WriteMessage(mType, msg)
		}
	}(conn)
}
