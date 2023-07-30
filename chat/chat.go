package chat

import (
	"log"

	"github.com/gorilla/websocket"

	"github.com/kylemadkins/gochat/ws"
)

type event string

const (
	None      event = ""
	Username  event = "username"
	ListUsers event = "list_users"
)

type ChatResponse struct {
	Event          event    `json:"event"`
	Message        string   `json:"message"`
	MessageType    string   `json:"message_type"`
	ConnectedUsers []string `json:"connected_users"`
}

type ChatPayload struct {
	Username string          `json:"username"`
	Event    event           `json:"event"`
	Message  string          `json:"message"`
	Conn     *websocket.Conn `json:"-"`
}

var messages = make(chan ChatPayload)

var ConnManager = ws.NewConnectionManager()

func ListenForMessages(conn *websocket.Conn) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Error: %v\n", r)
		}
	}()

	var payload ChatPayload
	for {
		err := conn.ReadJSON(&payload)
		if err != nil {
			ConnManager.Delete(conn)
			break
		} else {
			payload.Conn = conn
			messages <- payload
		}
	}
}

func BroadcastMessages() {
	var resp ChatResponse
	resp.ConnectedUsers = make([]string, 0)
	for {
		e := <-messages
		switch e.Event {
		case Username:
			ConnManager.Set(e.Conn, e.Username)
			resp.Event = ListUsers
			resp.ConnectedUsers = ConnManager.GetUsers()
			ConnManager.Broadcast(resp)
		}
	}
}
