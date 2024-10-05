package message

import (
	"context"
)

type Command struct {
	Title   string
	Message string
}

const (
	MessageDeliveryName = "message_delivery_asynq"
)

type Message interface {
	Send(ctx context.Context, asynqName string, command Command) error
	Receive(ctx context.Context, command Command) (Command, error)
}
