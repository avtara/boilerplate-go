package asynq

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hibiken/asynq"
)

func (so *aqObject) handlerProcessTaskSendEmail(ctx context.Context, task *asynq.Task) error {
	type PayloadSendEmailWelcome struct {
		Email string `json:"email"`
	}

	var payload PayloadSendEmailWelcome

	if err := json.Unmarshal(task.Payload(), &payload); err != nil {
		return fmt.Errorf("failed to unmarshal payload: %w", asynq.SkipRetry)
	}

	fmt.Println("ini dari handler", payload.Email)
	return nil
}
