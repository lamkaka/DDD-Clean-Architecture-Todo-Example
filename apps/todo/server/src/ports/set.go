package ports

import (
	"todo-server/ports/health"
	"todo-server/ports/http"

	"github.com/google/wire"
)

var Set = wire.NewSet(
	health.Set,
	http.NewServer,
	NewLogFilters,
	NewErrorHandler)
