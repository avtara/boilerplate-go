package server

import (
	"github.com/avtara/boilerplate-go/internal/service/delivery/server/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (so *svObject) initRoute() {
	so.Service.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	{
		protected := so.Service.Group("/protect")
		protected.Use(middleware.JWTAuthMiddleware())
		protected.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
	}

	users := so.Service.Group("/users")
	{
		users.GET("/last-login", so.handlerGetLastLogin)
		users.POST("/register", so.handlerRegister)
		users.POST("/login", so.handlerLogin)
	}
}
