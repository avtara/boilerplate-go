package app

import (
	"github.com/avtara/boilerplate-go/internal/service/delivery/asynq"
	"github.com/avtara/boilerplate-go/internal/service/delivery/server"
	asynq2 "github.com/avtara/boilerplate-go/internal/service/repository/asynq"
	"github.com/avtara/boilerplate-go/internal/service/repository/postgres"
	"github.com/avtara/boilerplate-go/internal/service/usecase"
)

func (cfg *App) InitService() (err error) {
	userRepository := postgres.NewUserRepository(cfg.DB)
	asynqRepository := asynq2.NewAsynqRepository(cfg.Asynq)
	userUsecase := usecase.NewUserUseCase(userRepository, asynqRepository)
	server.NewServerHandler(cfg.Server, userUsecase)
	asynq.NewAsyncHandler(cfg.Asynq, cfg.AsynqServer, cfg.AsynqMux)

	return
}
