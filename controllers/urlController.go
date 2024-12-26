package controllers

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"url-shortener/database"

	"github.com/gorilla/mux"
)

type URLRequest struct {
	OriginalURL string `json:"original_url"`
}

type URLResponse struct {
	ShortURL string `json:"short_url"`
}

func ShortenURL(w http.ResponseWriter, r *http.Request) {
	var request URLRequest
	json.NewDecoder(r.Body).Decode(&request)

	shortURL := generateShortURL()
	err := database.SaveURL(shortURL, request.OriginalURL)
	if err != nil {
		http.Error(w, "Failed to save URL", http.StatusInternalServerError)
		return
	}
	baseURL := os.Getenv("BASE_URL") // Read the base URL from an environment variable
	if baseURL == "" {
		baseURL = "http://localhost:8080" // Default fallback
	}
	response := URLResponse{ShortURL: fmt.Sprintf("%s/%s", baseURL, shortURL)}
	json.NewEncoder(w).Encode(response)
}

func RedirectURL(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	shortURL := params["shortURL"]

	originalURL, err := database.GetURL(shortURL)
	if err != nil {
		http.Error(w, "URL not found", http.StatusNotFound)
	}

	http.Redirect(w, r, originalURL, http.StatusFound)
}

func generateShortURL() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const length = 8
	var shortURL string
	for i := 0; i < length; i++ {
		shortURL += string(charset[rand.Intn(len(charset))])
	}
	return shortURL
}
