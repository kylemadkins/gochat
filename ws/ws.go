package ws

import (
	"log"

	"github.com/gorilla/websocket"
)

type ConnectionManager struct {
	connections map[*websocket.Conn]string
}

func NewConnectionManager() *ConnectionManager {
	cm := ConnectionManager{}
	cm.connections = make(map[*websocket.Conn]string)
	return &cm
}

func (cm *ConnectionManager) Add(conn *websocket.Conn, v string) {
	cm.connections[conn] = v
}

func (cm *ConnectionManager) Delete(conn *websocket.Conn) {
	delete(cm.connections, conn)
}

func (cm *ConnectionManager) Broadcast(json interface{}) {
	for conn := range cm.connections {
		err := conn.WriteJSON(json)
		if err != nil {
			log.Printf("Unable to broadcast message: %v\n", err)
			_ = conn.Close()
			cm.Delete(conn)
		}
	}
}
