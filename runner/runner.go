package runner

import "context"

type Runner interface {
	Run(context context.Context) error
}
