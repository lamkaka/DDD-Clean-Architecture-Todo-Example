package env

import (
	"github.com/google/wire"
)

var Set = wire.NewSet(
	NewParser,
	NewLogConfig,
	NewHttpConfig,
	NewPostgresConfig,
)
