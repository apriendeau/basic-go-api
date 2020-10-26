package handlers

import (
	"encoding/json"
	"net/http"
)

type Handler struct{}

type Message struct {
	Message string `json:"message"`
}

func New() *Handler {
	return &Handler{}
}

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
