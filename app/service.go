package app

import (
	"github.com/avtara/boilerplate-go/internal/service/delivery/asynq"
	"github.com/avtara/boilerplate-go/internal/service/delivery/server"
	"github.com/avtara/boilerplate-go/internal/service/repository/postgres"
	"github.com/avtara/boilerplate-go/internal/service/usecase"
)

func (cfg *App) InitService() (err error) {
	userRepository := postgres.NewUserRepository(cfg.DB)
	userUsecase := usecase.NewUserUseCase(userRepository, cfg.Asynq)
	server.NewServerHandler(cfg.Server, userUsecase)
	asynq.NewAsyncHandler(cfg.Asynq, cfg.AsynqServer, cfg.AsynqMux)

	return
}
