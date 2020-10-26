package handlers

import (
	"encoding/json"
	"net/http"
)

// Handler is for storing any secrets or services.
type Handler struct{}

// Message is a basic JSON structure response.
type Message struct {
	Message string `json:"message"`
}

// New is the intializer.
func New() *Handler {
	return &Handler{}
}

// Index is the index page.
func (h *Handler) Index(w http.ResponseWriter, r *http.Request) {
	b, err := json.Marshal(&Message{"hello"})
	if err != nil {
		w.WriteHeader(500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(b)
}
