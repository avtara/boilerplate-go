package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (so *svObject) initRoute() {
	so.Service.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	users := so.Service.Group("/users")

	users.GET("/last-login", so.handlerGetLastLogin)
	users.POST("/register", so.handlerRegister)

}
