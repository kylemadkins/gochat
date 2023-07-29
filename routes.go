package main

import (
	"net/http"

	"github.com/bmizerany/pat"

	"github.com/kylemadkins/gochat/handlers"
)

func routes() http.Handler {
	handler := pat.New()
	handler.Get("/", http.HandlerFunc(handlers.Home))
	handler.Get("/ws", http.HandlerFunc(handlers.Ws))
	return handler
}
