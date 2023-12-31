package handlers

import (
	"log"
	"net/http"

	"github.com/CloudyKit/jet/v6"
	"github.com/gorilla/websocket"

	"github.com/kylemadkins/gochat/chat"
)

var views = jet.NewSet(
	jet.NewOSFileSystemLoader("./templates"),
	jet.InDevelopmentMode(),
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		accepted := []string{"http://localhost:8000"}
		for _, v := range accepted {
			if v == r.Header["Origin"][0] {
				return true
			}
		}
		return false
	},
}

func Home(w http.ResponseWriter, r *http.Request) {
	err := renderPage(w, "index.jet", nil)
	if err != nil {
		log.Println(err)
	}
}

func Ws(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	chat.ConnManager.Set(ws, "")
	log.Println("Client connected")
	resp := chat.ChatResponse{Message: "Connected to server!"}

	err = ws.WriteJSON(resp)
	if err != nil {
		log.Println(err)
	}

	go chat.ListenForMessages(ws)
}

func renderPage(w http.ResponseWriter, tmpl string, vars jet.VarMap) error {
	view, err := views.GetTemplate(tmpl)
	if err != nil {
		log.Println(err)
		return err
	}

	err = view.Execute(w, vars, nil)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
