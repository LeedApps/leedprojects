package routes

import (
	"github.com/gin-gonic/gin"
)

// LPRouter defines the application routes
func LPRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	return r
}
