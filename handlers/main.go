package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/josearpaiaq/shortener/db"
	"github.com/josearpaiaq/shortener/models"
	"github.com/josearpaiaq/shortener/utils"
	"gorm.io/gorm"
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
	var response models.URL
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	response.OriginalURL = body.URL
	originalURL := utils.GenerateShortURL(body.URL)
	response.ShortURL = originalURL

	// Save the URL to the database
	db.Connection.Create(&response)

	json.NewEncoder(w).Encode(&response)
}

// TODO: fix this function, aparently every request is being calling twice, updating the clicks twice instead of once per request
func RedirectToOriginalURL(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	shortUrl := vars["short_url"]

	var url models.URL
	// Check if the short URL exists
	if err := db.Connection.First(&url, "short_url = ?", shortUrl).Error; err != nil {
		http.Error(w, fmt.Sprintf("URL %s not found", shortUrl), http.StatusNotFound)
		return
	}

	// Safely increment clicks using a DB expression to avoid race conditions
	if err := db.Connection.Model(&models.URL{}).Where("id = ?", url.ID).UpdateColumn("clicks", gorm.Expr("clicks + ?", 1)).Error; err != nil {
		// Log the error but continue with the redirect
		fmt.Println("failed to increment clicks:", err)
	}

	original := url.OriginalURL
	if original == "" {
		http.Error(w, "original URL is empty", http.StatusInternalServerError)
		return
	}

	if !strings.HasPrefix(original, "http://") && !strings.HasPrefix(original, "https://") {
		original = "https://" + original
	}

	// Redirect to the original URL
	http.Redirect(w, r, original, http.StatusFound)
}

// TODO: get number of clicks across all URLs and return a JSON object like this:
// {
//   "url": "number of clicks"
// }
func Stats(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	stats := models.URL{}
	db.Connection.Model(&models.URL{}).Select("clicks").Scan(&stats)
	json.NewEncoder(w).Encode(stats)
}

func GetAllURLs(w http.ResponseWriter, r *http.Request) {
	var urls []models.URL
	db.Connection.Find(&urls)
	json.NewEncoder(w).Encode(urls)
}

func GetURL(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	urlParam := vars["url"]

	var url models.URL
	if err := db.Connection.First(&url, "short_url = ?", urlParam).Error; err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if url.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(fmt.Sprintf("URL %s not found", urlParam)))
		return
	}

	json.NewEncoder(w).Encode(url)
}