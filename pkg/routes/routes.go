package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// register the routes
func RegisterRoutes(r *gin.Engine) {
	// ping
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	//
}
