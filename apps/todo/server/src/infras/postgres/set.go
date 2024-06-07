package postgres

import (
	postgres_client "todo-server/infras/postgres/client"

	"github.com/google/wire"
)

var Set = wire.NewSet(
	postgres_client.NewClient,
)
