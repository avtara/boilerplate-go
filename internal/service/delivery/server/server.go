package server

import (
	"github.com/avtara/boilerplate-go/internal/service"
	"github.com/avtara/boilerplate-go/utils"
	"github.com/gin-gonic/gin"
	"github.com/hibiken/asynq"
)

type svObject struct {
	Service     *gin.Engine
	UserUsecase service.UserUsecase
	Asynq       *asynq.Client

	IsSystemMaintenance bool
}

func NewServerHandler(
	svc *gin.Engine,
	UserUsecase service.UserUsecase,
) {
	obj := &svObject{
		Service:             svc,
		UserUsecase:         UserUsecase,
		IsSystemMaintenance: utils.GetConfig("is_system_maintenance", "FALSE") == "TRUE",
	}

	obj.initRoute()
}
