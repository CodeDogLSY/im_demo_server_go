// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"log"
)

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	clients map[*Client]bool

	// Inbound messages from the clients.
	broadcast chan []byte

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client
}

func newHub() *Hub {
	return &Hub{
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
			//返回所有用户列表
			list := make([]user, 0)
			for client := range h.clients {
				userObj := user{}
				userObj.Id = client.id
				userObj.Name = client.name
				list = append(list, userObj)
			}
			for client := range h.clients {
				dataObj := data{
					DataType:    2,
					DataContent: list,
				}
				usermsg, err := json.Marshal(dataObj)
				if (err != nil) {
					log.Printf("error : v%", err)
					return
				}
				select {
				case client.send <- usermsg:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast:
			//此处增加过滤,如果指定了收件人，就选择对应的收件人，否则发给所有人
			contentmsg := string(message)
			dataObj := data{}
			json.Unmarshal([]byte(contentmsg), &dataObj)
			if dataObj.DataType==1 {
				//tempContent := dataObj.DataContent.(string)
				msgObg, ok := (dataObj.DataContent).(map[string]interface{})
				if ok{
					toid:= msgObg["to_id"]
					to_id:= toid.(string)
					if len(to_id) == 0 {
						for client := range h.clients {
							select {
							case client.send <- message:
							default:
								close(client.send)
								delete(h.clients, client)
							}
						}
					} else {
						for client := range h.clients {
							if client.id == to_id {
								select {
								case client.send <- message:
								default:
									close(client.send)
									delete(h.clients, client)
								}
								break
							}
						}
					}
				}

			}
		}
	}
}
