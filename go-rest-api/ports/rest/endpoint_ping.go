package rest

import "github.com/gin-gonic/gin"

func (s server) PingEndpoint(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
