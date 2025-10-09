package router

import (
	"github.com/gorilla/mux"
	"github.com/josearpaiaq/shortener/handlers"
	"github.com/josearpaiaq/shortener/middlewares"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()

	r.Use(middlewares.LoggingMiddleware)

	r.HandleFunc("/", handlers.GetHomeRoute).Methods("GET")
	r.HandleFunc("/shorten", handlers.ShortenURL).Methods("POST")
	r.HandleFunc("/{url}", handlers.GoToShortenURL).Methods("GET")

	return r
}