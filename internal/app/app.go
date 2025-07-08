package app

import (
	"myproject/internal/app/api"

	"github.com/go-chi/chi/v5"
	"github.com/gin-gonic/gin"
)

func NewApp() (*chi.Mux, error) {
	r := api.NewChiRouter()
	api.RegisterChiRoutes(r)

	return r, nil
}

func NewGinApp() (*gin.Engine, error) {
	r := api.NewGinRouter()
	api.RegisterGinRoutes(r)

	return r, nil
}
