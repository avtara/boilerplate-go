package app

import (
	"github.com/avtara/boilerplate-go/deliveries/server"
)

func (cfg *App) InitService() (err error) {
	server.NewServerHandler(cfg.Server)

	return
}
