package runner

import (
	"asynq-distributed-task/registry"
	"context"
	"github.com/hibiken/asynq"
)

type AsynqRunnerOpts struct {
	AsynqServer *asynq.Server
}

type AsynqRunner struct {
	opts AsynqRunnerOpts
	mux  *asynq.ServeMux
}

func NewAsynqRunner(opts AsynqRunnerOpts) *AsynqRunner {
	return &AsynqRunner{opts: opts, mux: asynq.NewServeMux()}
}

func (r AsynqRunner) AddHandlerRegistry(registry registry.HandlerRegistry) {
	registry.RegisterRoutesTo(r.mux)
}

func (a AsynqRunner) Run(context context.Context) error {
	return a.opts.AsynqServer.Run(a.mux)
}
