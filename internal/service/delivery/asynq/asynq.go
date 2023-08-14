package asynq

import (
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

	asynqMux.HandleFunc("task:send_email_welcome", obj.handlerProcessTaskSendEmail)
	asynqServer.Start(asynqMux)
}
