package enqueueimpl

import (
	"context"
	asynqClient "github.com/arsyadarmawan/asynq-distributed-task/client"
	"github.com/arsyadarmawan/asynq-distributed-task/common/task"
)

type Enqueuer struct {
	client *asynqClient.Client
}

func NewEnqueuer(client *asynqClient.Client) Enqueuer {
	return Enqueuer{client: client}
}

func (e Enqueuer) Enqueue(ctx context.Context, name string, payload any) error {
	task, err := task.NewAsynqTask(name, payload, e.client.Options...)
	if err != nil {
		return err
	}
	_, err = e.client.EnqueueContext(ctx, task)
	if err != nil {
		return err
	}
	return nil
}
