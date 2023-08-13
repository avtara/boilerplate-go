package server

import (
	"github.com/avtara/boilerplate-go/utils"
	"github.com/gin-gonic/gin"
)

type svObject struct {
	Service *gin.Engine

	IsSystemMaintenance bool
}

func NewServerHandler(
	svc *gin.Engine,
) {
	obj := &svObject{
		Service: svc,

		IsSystemMaintenance: utils.GetConfig("is_system_maintenance", "FALSE") == "TRUE",
	}

	obj.initRoute()
}
