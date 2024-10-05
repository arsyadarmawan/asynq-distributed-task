package runner

import (
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

func NewAsynqRunner(opts AsynqRunnerOpts, mux *asynq.ServeMux) *AsynqRunner {
	return &AsynqRunner{opts: opts, mux: mux}
}

func (a AsynqRunner) Run(context context.Context) error {
	return a.opts.AsynqServer.Run(a.mux)
}
