package app

import (
	"fmt"
	"github.com/avtara/boilerplate-go/utils"
	"github.com/gin-gonic/gin"
	"github.com/hibiken/asynq"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"os"
)

type App struct {
	Hostname    string
	Server      *gin.Engine
	DB          *sqlx.DB
	Asynq       *asynq.Client
	AsynqServer *asynq.Server
	AsynqMux    *asynq.ServeMux
}

func New() App {
	var cfg App

	cfg.Hostname, _ = os.Hostname()
	most(cfg.InitViper())
	most(cfg.InitLogrus())
	most(cfg.InitAsynq())
	most(cfg.InitServer())
	most(cfg.InitPostgres())

	most(cfg.InitService())

	return cfg
}

func (cfg *App) Start() (err error) {
	ch := make(chan bool)
	go func() {
		cfg.Server.Run(fmt.Sprintf("%s:%s", cfg.Hostname, utils.GetConfig("ports.server", "8000")))

		ch <- false
	}()

	<-ch
	return
}

func most(err error) {
	if err != nil {
		log.Panic(err)
	}
}
