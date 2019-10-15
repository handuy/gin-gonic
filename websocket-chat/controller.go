package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type Client struct {
	ws   *websocket.Conn
	send chan []byte
}

type Hub struct {
	Clients    map[*Client]bool
	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan string
	Content    string
}

var NewHub = Hub{
	Clients:    make(map[*Client]bool),
	Register:   make(chan *Client),
	Unregister: make(chan *Client),
	Broadcast:  make(chan string),
	Content:    "",
}

func (hub *Hub) Run() {
	for {
		select {
		case newClient := <-hub.Register:
			// đăng kí client mới với Hub
			hub.Clients[newClient] = true
			// gửi message hiện đang được lưu trong Hub đến client mới
			newClient.send <- []byte(hub.Content)
			break
		case removedClient := <-hub.Unregister:
			delete(hub.Clients, removedClient)
			close(removedClient.send)
			break
		// channel hub.Broadcast nhận được message mới từ 1 client
		// thông qua method client.readMessageFromSocketConnection()
		case newMessage := <-hub.Broadcast:
			// Lưu message mới nhận vào hub.Content
			hub.Content = newMessage
			// nó sẽ gửi message này đến các client được lưu trong trường Clients của hub
			hub.sendToOtherClients()
			break
		}
	}
}

func (hub *Hub) sendToOtherClients() {
	log.Println("Chạy hàm send to other clients")

	for client := range hub.Clients {
		select {
		// Đẩy message từ Hub Content vào channel send của mỗi client
		// để từ đó gọi method client.writeMessage()
		case client.send <- []byte(hub.Content):
			break
		// Trường hợp không thể kết nối đến client
		default:
			delete(hub.Clients, client)
			close(client.send)
		}

	}
}

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func Wshandler(w http.ResponseWriter, r *http.Request) {
	conn, err := wsupgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Failed to set websocket upgrade: %+v", err)
		return
	}

	client := &Client{
		send: make(chan []byte),
		ws:   conn,
	}

	// Đăng kí client mới với Hub
	NewHub.Register <- client

	go client.writeMessage()

	client.readMessageFromSocketConnection()
}

func (client *Client) readMessageFromSocketConnection() {
	log.Println("Chạy hàm read message")
	// Trường hợp gặp lỗi kết nối đến browser/tab
	defer func() {
		NewHub.Unregister <- client
		client.ws.Close()
	}()

	// Vòng lặp for vô tận liên tục đọc message từ socket connection
	// để gửi vào channel Broadcast của Hub
	for {
		_, msg, err := client.ws.ReadMessage()
		if err != nil {
			break
		}
		NewHub.Broadcast <- string(msg)
	}
}

func (client *Client) writeMessage() {
	log.Println("Chạy hàm write message")

	defer func() {
		client.ws.Close()
	}()

	for {
		select {
		// channel send của client khi nhận được message từ Hub.sendToOtherClients()
		// client sẽ viết message này vào socket connection để trả về cho browser/tab
		case message, ok := <-client.send:
			if !ok {
				client.ws.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			err := client.ws.WriteMessage(websocket.TextMessage, message)
			if err != nil {
				return
			}
		}
	}
}
