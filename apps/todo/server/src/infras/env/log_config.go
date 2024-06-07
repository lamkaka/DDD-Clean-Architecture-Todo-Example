package env

import (
	"context"
	"todo-server/infras/log"
)

func NewLogConfig(ctx context.Context, parser Parser) (log.Config, error) {
	var e logConfigEnv
	err := parser.Parse(ctx, &e)
	if err != nil {
		return log.Config{}, err
	}

	return log.Config(e), nil
}

type logConfigEnv struct {
	Level log.Level `env:"LOG_LEVEL"`
}
