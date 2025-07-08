package app

import (
	"myproject/internal/app/api"

	"github.com/go-chi/chi/v5"
)

func NewApp() (*chi.Mux, error) {
	r := api.NewChiRouter()
	api.RegisterRoutes(r)

	return r, nil
}
