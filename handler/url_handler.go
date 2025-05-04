package handler

import "net/http"

type URLHandler struct{}

func NewURLHandler() *URLHandler {
	return &URLHandler{}
}

func (h *URLHandler) URLShortener(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello World..!!!"))
}
