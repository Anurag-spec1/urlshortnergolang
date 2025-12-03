package handler  

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type URL struct {
	ID           string    `json:"id"`
	OriginalURL  string    `json:"original_url"`
	ShortURL     string    `json:"short_url"`
	CreationDate time.Time `json:"creation_date"`
}

var urlDb = make(map[string]URL)

func generateShortURL(originalURL string) string {
	hasher := md5.New()
	hasher.Write([]byte(originalURL))
	data := hasher.Sum(nil)
	hash := hex.EncodeToString(data)
	return hash[:8]
}

func createURL(originalURL string) string {
	shortURL := generateShortURL(originalURL)
	id := shortURL
	urlDb[id] = URL{
		ID:           id,
		OriginalURL:  originalURL,
		ShortURL:     shortURL,
		CreationDate: time.Now(),
	}
	return shortURL
}

func getURL(id string) (URL, error) {
	url, ok := urlDb[id]
	if !ok {
		return URL{}, errors.New("URL NOT FOUND")
	}
	return url, nil
}

func Handler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		fmt.Fprintf(w, "Anurag's Territory")
	case "/shorten":
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var data struct {
			URL string `json:"url"`
		}
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		
		shortURL := createURL(data.URL)
		response := struct {
			ShortURL string `json:"short_url"`
		}{ShortURL: shortURL}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	default:
		if len(r.URL.Path) > 9 && r.URL.Path[:9] == "/redirect" {
			id := r.URL.Path[len("/redirect/"):]
			if id == "" {
				http.Error(w, "Invalid request", http.StatusBadRequest)
				return
			}
			
			url, err := getURL(id)
			if err != nil {
				http.Error(w, "URL not found", http.StatusNotFound)
				return
			}
			http.Redirect(w, r, url.OriginalURL, http.StatusFound)
		} else {
			http.Error(w, "Not Found", http.StatusNotFound)
		}
	}
}