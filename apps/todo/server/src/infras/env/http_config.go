package env

import (
	"context"
	"fmt"

	"todo-server/ports/http"
)

func NewHttpConfig(ctx context.Context, parser Parser) (http.Config, error) {
	var e httpConfigEnv
	err := parser.Parse(ctx, &e)
	if err != nil {
		return http.Config{}, err
	}

	return http.Config{ServerPort: fmt.Sprint(e.ServerPort)}, nil
}

type httpConfigEnv struct {
	ServerPort uint16 `env:"HTTP_SERVER_PORT"`
}
