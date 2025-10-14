package routes

import (
	"github.com/gorilla/mux"
	"github.com/josearpaiaq/shortener/handlers"
	"github.com/josearpaiaq/shortener/middlewares"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()

	// TODO: prefix all routes with /api
	// r.PathPrefix("/api").Subrouter()

	// Add middleware
	r.Use(middlewares.LoggingMiddleware)

	// root route
	r.HandleFunc("/", handlers.GetHomeRoute).Methods("GET")

	// stats route
	r.HandleFunc("/stats", handlers.Stats).Methods("GET")

	// shorten URL route
	r.HandleFunc("/shorten", handlers.ShortenURL).Methods("POST")
	
	// get all URLs route
	r.HandleFunc("/urls", handlers.GetAllURLs).Methods("GET")
	
	// get one URL route
	r.HandleFunc("/urls/{url}", handlers.GetURL).Methods("GET")

	// This is a catch-all route for the shortened URLs
	r.HandleFunc("/s/{short_url}", handlers.RedirectToOriginalURL).Methods("GET")

	return r
}