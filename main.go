package main

import (
	"net/http"

	"github.com/apriendeau/basic-go-api/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	h := handlers.New()
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", h.Index)

	if err := http.ListenAndServe(":3000", r); err != nil {
		panic(err)
	}
}
