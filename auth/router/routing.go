package router

import (
	"net/http"

	"jwt-todo/helper/redis"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	client := router.Group("/api")
	{
		client.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"data": "hello world"})
		})
		client.POST("/login", redis.Login)
	}
	return router
}
