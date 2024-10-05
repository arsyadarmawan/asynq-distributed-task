package server

import "github.com/hibiken/asynq"

type (
	Server struct {
		*asynq.Server
		Options []asynq.Option
	}

	RedisServerConfig struct {
		Address, Username, Password string
		DB                          int
	}
)

func InitServer(config RedisServerConfig, AsynqConfig asynq.Config) *Server {
	server := asynq.NewServer(asynq.RedisClientOpt{
		Addr:     config.Address,
		Username: config.Username,
		Password: config.Password,
		DB:       config.DB,
	}, AsynqConfig)
	srv := &Server{
		Server:  server,
		Options: nil,
	}
	return srv
}
