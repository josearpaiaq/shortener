package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type ShortenURLRequest struct {
	URL string `json:"url"`
}

type ShortenURLResponse struct {
	OriginalURL string `json:"original_url"`
	ShortURL string `json:"short_url"`
}

func GetHomeRoute(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the URL Shortener Service!"))
}

func ShortenURL(w http.ResponseWriter, r *http.Request) {
	fmt.Print("Shorten URL endpoint")
	w.Header().Set("Content-Type", "application/json")
	var response ShortenURLResponse
	_ = json.NewDecoder(r.Body).Decode(&response)
	// TODO: implement shorten URL logic
	// get original URL from request body
	// var data ShortenURLRequest
	// err := json.NewDecoder(r.Body).Decode(&data)

	// if err != nil {
	// 	fmt.Println("Error decoding JSON:", err)
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	return
	// }

	// fmt.Println(data)

	json.NewEncoder(w).Encode(&response)
}

func GoToShortenURL(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	url := vars["url"]
	w.Write([]byte("Go to shorten URL " + url))
}