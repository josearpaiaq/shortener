package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/josearpaiaq/shortener/utils"
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
	w.Header().Set("Content-Type", "application/json")
	var body ShortenURLRequest
	var response ShortenURLResponse
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	response.OriginalURL = body.URL
	originalURL := utils.GenerateShortURL(body.URL)
	response.ShortURL = originalURL
	json.NewEncoder(w).Encode(&response)
}

func GoToShortenURL(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	url := vars["url"]
	w.Write([]byte("Go to shorten URL " + url))
}