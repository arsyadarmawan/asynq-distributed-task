package client

import (
	"github.com/hibiken/asynq"
	"time"
)

type ClientOpts func(*Client)

func WithMaxRetry(max int) ClientOpts {
	return func(c *Client) {
		c.Options = append(c.Options, asynq.MaxRetry(max))
	}
}

func WithProcessIn(duration time.Duration) ClientOpts {
	return func(c *Client) {
		c.Options = append(c.Options, asynq.ProcessIn(duration))
	}
}

func WithTimeOut(duration time.Duration) ClientOpts {
	return func(c *Client) {
		c.Options = append(c.Options, asynq.Timeout(duration))
	}
}
