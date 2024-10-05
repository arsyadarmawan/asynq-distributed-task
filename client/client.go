package client

import (
	"github.com/hibiken/asynq"
	"time"
)

type RedisConnection struct {
	Addr     string
	Port     string
	UserName string
	Password string
	DB       int
}

type (
	Client struct {
		*asynq.Client
		Options []asynq.Option
	}
)

var (
	defaultOpts = []ClientOpts{
		WithMaxRetry(3),
		WithTimeOut(5 * time.Minute),
		WithProcessIn(1 * time.Minute),
	}
)

func InitConfiguration(connection RedisConnection, opts ...ClientOpts) *Client {
	c := &Client{
		Client: asynq.NewClient(asynq.RedisClientOpt{
			Addr:     connection.Addr,
			Username: connection.UserName,
			Password: connection.Password,
			DB:       connection.DB,
		}),
	}
	c.applyOptions(defaultOpts)
	c.applyOptions(opts)
	return c
}

func (c *Client) applyOptions(options []ClientOpts) {
	for _, opt := range options {
		opt(c)
	}
}
