package task

import (
	"encoding/json"
	"github.com/hibiken/asynq"
)

func NewAsynqTask(name string, request any, options ...asynq.Option) (*asynq.Task, error) {
	payload, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	return asynq.NewTask(name, payload, options...), nil
}
