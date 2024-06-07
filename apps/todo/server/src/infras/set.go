package infras

import (
	"todo-server/infras/env"
	"todo-server/infras/log"

	"github.com/google/wire"
)

var Set = wire.NewSet(
	env.Set,
	log.Set,
)
