package postgres

import (
	postgres_client "todo-server/infras/postgres/client"
	todo_postgres "todo-server/infras/postgres/todo"

	"github.com/google/wire"
)

var Set = wire.NewSet(
	todo_postgres.NewRepository,
	postgres_client.NewClient,
)
