package enqueue

import "context"

//go:generate mockgen --source=enqueue.go --destination=mock/enqueue.go --package=queue
type AsynqEnqueuer interface {
	Enqueue(ctx context.Context, name string, payload any) error
}
