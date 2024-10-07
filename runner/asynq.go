package runner

import (
	"context"
	"github.com/arsyadarmawan/asynq-distributed-task/registry"
	"github.com/hibiken/asynq"
)

type AsynqRunnerOpts struct {
	AsynqServer *asynq.Server
	Mux         *asynq.ServeMux
}

type AsynqRunner struct {
	opts AsynqRunnerOpts
}

func NewAsynqRunner(opts AsynqRunnerOpts) *AsynqRunner {
	return &AsynqRunner{opts: opts}
}

func (r AsynqRunner) AddHandlerRegistry(registry registry.HandlerRegistry) {
	registry.RegisterRoutesTo(r.opts.Mux)
}

func (a AsynqRunner) Run(context context.Context) error {
	err := a.opts.AsynqServer.Run(a.opts.Mux)
	if err != nil {
		return err
	}
	return nil
}
