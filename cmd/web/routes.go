package main

import (
	"github.com/bmizerany/pat"
	"github.com/ismail118/simple-chat/internal/handlers"
	"net/http"
)

func routes() http.Handler {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.Home))
	mux.Get("/ws", http.HandlerFunc(handlers.WsEndpoint))

	fileSever := http.FileServer(http.Dir("./static/"))
	mux.Get("/asset/", http.StripPrefix("/asset/", fileSever))

	return mux
}
