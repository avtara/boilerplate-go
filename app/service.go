package app

import (
	"github.com/avtara/boilerplate-go/internal/deliveries/server"
)

func (cfg *App) InitService() (err error) {
	server.NewServerHandler(cfg.Server)

	return
}
