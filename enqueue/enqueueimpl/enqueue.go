package enqueueimpl

import (
	asynqClient "asynq-distributed-task/client"
	"asynq-distributed-task/common/task"
	"context"
)

type AsynqEnqueue struct {
	client *asynqClient.Client
}

func NewEnqueuer(client *asynqClient.Client) *AsynqEnqueue {
	return &AsynqEnqueue{client: client}
}

func (e AsynqEnqueue) Enqueue(ctx context.Context, name string, payload any) error {
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
