package app

import (
	"github.com/avtara/boilerplate-go/internal/service/delivery/server"
	"github.com/avtara/boilerplate-go/internal/service/repository/postgres"
	"github.com/avtara/boilerplate-go/internal/service/usecase"
)

func (cfg *App) InitService() (err error) {
	userRepository := postgres.NewUserRepository(cfg.DB)
	userUsecase := usecase.NewUserUseCase(userRepository)
	server.NewServerHandler(cfg.Server, userUsecase)

	return
}
