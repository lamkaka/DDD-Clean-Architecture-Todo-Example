//go:build wireinject

package main

import (
	"context"

	"todo-server/infras"
	"todo-server/ports"
	"todo-server/ports/http"

	"github.com/google/wire"
)

func initializeServer(ctx context.Context) (http.Server, error) {
	wire.Build(
		infras.Set,
		ports.Set,
	)
	return nil, nil
}
