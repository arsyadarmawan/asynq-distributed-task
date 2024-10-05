package registry

import "github.com/hibiken/asynq"

type HandlerRegistry interface {
	RegisterRoutesTo(mux *asynq.ServeMux)
}
