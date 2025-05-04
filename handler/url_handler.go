package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rahulshewale153/infra-url-shortener/model"
	"github.com/rahulshewale153/infra-url-shortener/service"
)

const (
	SERVICE_BASE_URL = "http://localhost:8080"
	CONTENT_TYPE     = "Content-Type"
	APPLICATION_JSON = "application/json"
)

type URLHandler struct {
	urlService service.URLService
}

func NewURLHandler(urlService service.URLService) *URLHandler {
	return &URLHandler{urlService: urlService}
}

// URLShortener handles the URL shortening request
func (h *URLHandler) URLShortener(w http.ResponseWriter, r *http.Request) {
	var req model.URLShortenerRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	shortURL, err := h.urlService.GetURLShortener(r.Context(), req.URL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Construct the full short URL
	response := model.URLShortenerResponse{
		ShortURL: fmt.Sprintf("%s/%s", SERVICE_BASE_URL, shortURL),
	}

	// Set the response header and status code
	w.Header().Set(CONTENT_TYPE, APPLICATION_JSON)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// GetOriginalURL handles the request to retrieve the original URL
func (h *URLHandler) GetOriginalURL(w http.ResponseWriter, r *http.Request) {
	shortURLID := mux.Vars(r)["short_url_id"]
	originalURL, err := h.urlService.GetOriginalURL(r.Context(), shortURLID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	http.Redirect(w, r, originalURL, http.StatusFound)
}

// Get Top 3 Domain
func (h *URLHandler) GetTop3Domain(w http.ResponseWriter, r *http.Request) {
	domains, err := h.urlService.GetTop3Domain(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set(CONTENT_TYPE, APPLICATION_JSON)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(domains)
}
