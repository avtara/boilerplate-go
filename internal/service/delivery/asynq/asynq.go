package asynq

import (
	"github.com/avtara/boilerplate-go/internal/models"
	"github.com/hibiken/asynq"
)

type aqObject struct {
	client      *asynq.Client
	asynqServer *asynq.Server
	asynqMux    *asynq.ServeMux
}

func NewAsyncHandler(
	client *asynq.Client,
	asynqServer *asynq.Server,
	asynqMux *asynq.ServeMux,
) {
	obj := &aqObject{
		asynqMux:    asynqMux,
		client:      client,
		asynqServer: asynqServer,
	}

	asynqMux.HandleFunc(models.TypeNameTaskSendEmailWelcome, obj.handlerProcessTaskSendEmail)
	asynqServer.Start(asynqMux)
}
