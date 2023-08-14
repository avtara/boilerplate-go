package asynq

import (
	"context"
	"encoding/json"
	"github.com/avtara/boilerplate-go/internal/service"
	"github.com/hibiken/asynq"
)

type asynqRepository struct {
	conn *asynq.Client
}

func NewAsynqRepository(
	conn *asynq.Client,
) service.AsynqRepository {
	return &asynqRepository{
		conn: conn,
	}
}

func (a *asynqRepository) Publish(ctx context.Context, typename string, payload interface{}) (taskInfo *asynq.TaskInfo, err error) {
	var payloadJSON []byte
	payloadJSON, err = json.Marshal(payload)
	if err != nil {
		return
	}

	task := asynq.NewTask(typename, payloadJSON)
	taskInfo, err = a.conn.EnqueueContext(ctx, task)
	if err != nil {
		return
	}

	return
}
