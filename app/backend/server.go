// app/backend/server.go
package backend

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
)

// SetupRouter initializes Gin router
func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Define a basic API endpoint

	// 允许所有来源的跨域请求
	router.Use(cors.Default())

	router.GET("/api/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// Another example of an API endpoint
	router.POST("/api/data", func(c *gin.Context) {
		var jsonData map[string]interface{}
		if err := c.BindJSON(&jsonData); err == nil {
			c.JSON(http.StatusOK, gin.H{
				"received": jsonData,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid data",
			})
		}
	})

	return router
}
