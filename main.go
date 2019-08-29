package main

import (
	"flag"
	"log"

	//"github.com/gorilla/websocket"
	"net/http"
)

var addr = flag.String("addr", ":8080", "http service address")
func main() {
	flag.Parse()
	hub := newHub()
	go hub.run()
	http.HandleFunc("/ws",  func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})
	//http.HandleFunc("/alluser", func(w http.ResponseWriter, r *http.Request) {
	//	returnAlluser(hub, w, r)
	//})

	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
