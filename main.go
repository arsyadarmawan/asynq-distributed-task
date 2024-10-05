package main

import (
	"asynq-distributed-task/client"
	"asynq-distributed-task/enqueue/enqueueimpl"
	"asynq-distributed-task/message"
	"asynq-distributed-task/message/messageimpl"
	"asynq-distributed-task/runner"
	"asynq-distributed-task/server"
	"context"
	"fmt"
	"github.com/hibiken/asynq"
)

func main() {
	c := client.InitConfiguration(client.RedisConnection{
		Addr: "localhost:6379",
		Port: "6379",
		DB:   0,
	})
	enqueuer := enqueueimpl.NewEnqueuer(c)
	msg := messageimpl.NewMessage(messageimpl.MessageOpts{
		enqueuer,
	})

	ctx := context.Background()
	errMessage := msg.Send(ctx, message.MessageDeliveryName, message.Command{
		Title:   "Weekly Reports",
		Message: "Hello Sir, This is my current report",
	})
	if errMessage != nil {
		fmt.Sprintf(errMessage.Error())
	}

	srv := server.InitServer(server.RedisServerConfig{
		Address: "localhost:6379",
		DB:      0,
	}, asynq.Config{
		Queues: map[string]int{
			"critical": 6,
			"default":  3,
			"low":      1,
		},
	})

	mux := asynq.NewServeMux()
	run := runner.NewAsynqRunner(runner.AsynqRunnerOpts{
		AsynqServer: srv.Server,
		Mux:         mux,
	})

	mux.HandleFunc(message.MessageDeliveryName, msg.HandleEmailDeliveryTask)
	run.Run(ctx)
}
