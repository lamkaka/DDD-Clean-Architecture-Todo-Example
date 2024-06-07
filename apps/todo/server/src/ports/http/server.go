package http

import (
	"libs/http_server"
	"todo-server/infras/log"
	"todo-server/ports/health"
)

type Config = http_server.Config
type Server = http_server.Server

func NewServer(
	config Config,
	healthController health.Controller,
	logFilter http_server.LogFilter,
	errHandler http_server.ErrorHandler,
	rootLogger log.RootLogger,
) (Server, error) {
	config.ErrorHandler = errHandler
	config.LogFilter = logFilter
	return http_server.NewServer(
		http_server.Config(config),
		healthController,
	)
}
