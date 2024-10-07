package messageimpl

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/arsyadarmawan/asynq-distributed-task/enqueue/enqueueimpl"
	"github.com/arsyadarmawan/asynq-distributed-task/message"
	"github.com/hibiken/asynq"
	log "github.com/sirupsen/logrus"
)

type MessageOpts struct {
	enqueueimpl.Enqueuer
}

type Message struct {
	Opts MessageOpts
}

func NewMessage(opts MessageOpts) *Message {
	return &Message{Opts: opts}
}

func (m Message) Send(ctx context.Context, asynqName string, command message.Command) error {
	return m.Opts.Enqueue(ctx, asynqName, command)
}

func (m Message) Receive(ctx context.Context, command message.Command) (message.Command, error) {
	log.WithContext(ctx).WithFields(log.Fields{
		"mesage": command,
	})
	log.Info("asynq schema works")

	return command, nil
}

func (m Message) HandleEmailDeliveryTask(ctx context.Context, t *asynq.Task) error {
	var p message.Command
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}

	msg, err := m.Receive(ctx, p)
	if err != nil {
		return err
	}
	fmt.Println(msg)
	return nil
}
