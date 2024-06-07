package ports

import (
	"todo-server/ports/health"
	"todo-server/ports/http"
	openapi_rest "todo-server/ports/openapi"
	todo_rest "todo-server/ports/todo"

	"github.com/google/wire"
)

var Set = wire.NewSet(
	health.Set,
	todo_rest.NewController,
	openapi_rest.NewController,
	http.NewServer,
	NewLogFilters,
	NewErrorHandler)
