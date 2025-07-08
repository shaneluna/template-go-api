package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewGinRouter() *gin.Engine {
	// set gin to release mode for production
	gin.SetMode(gin.ReleaseMode)
	
	r := gin.New()

	// global middleware
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(JSONContentTypeMiddleware())

	// custom 404 handler
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, NewErrorResponse404())
	})

	return r
}

func RegisterGinRoutes(r *gin.Engine) {
	// TODO
}

// JSONContentTypeMiddleware sets the Content-Type header to application/json
func JSONContentTypeMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		c.Next()
	}
}