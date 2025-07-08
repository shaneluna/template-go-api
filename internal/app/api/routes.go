package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewChiRouter() *chi.Mux {
	r := chi.NewRouter()

	// common middleware
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.SetHeader("Content-Type", "application/json"))

	// custom 404 handler
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		resp := NewErrorResponse404()
		err := json.NewEncoder(w).Encode(resp)
		if err != nil {
			log.Printf("Failed to encode error response: %v", err)
		}
	})

	return r
}

func RegisterRoutes(*chi.Mux) {
	// TODO
}
