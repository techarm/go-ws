package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/techarm/go-ws/internal/handlers"
)

func routes() http.Handler {
	mux := pat.New()
	mux.Get("/", http.HandlerFunc(handlers.Home))
	mux.Get("/ws", http.HandlerFunc(handlers.WSEndPoint))

	return mux
}
