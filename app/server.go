package app

import (
	"github.com/gin-gonic/gin"
)

func (cfg *App) InitServer() (err error) {
	cfg.Server = gin.Default()

	return
}
