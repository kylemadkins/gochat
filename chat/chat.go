package chat

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"

	"github.com/kylemadkins/gochat/ws"
)

type ChatResponse struct {
	Action      string `json:"action"`
	Message     string `json:"message"`
	MessageType string `json:"message_type"`
}

type ChatPayload struct {
	Username string `json:"username"`
	Action   string `json:"action"`
	Message  string `json:"message"`
}

var messages chan ChatPayload

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
		if err == nil {
			messages <- payload
		}
	}
}

func BroadcastMessages() {
	var resp ChatResponse
	for {
		e := <-messages
		resp.Action = "Got here"
		resp.Message = fmt.Sprintf("Some message. Action was %s", e.Action)
		ConnManager.Broadcast(resp)
	}
}
