package applications

import (
	todo_applications "todo-server/applications/todo"

	"github.com/google/wire"
)

var Set = wire.NewSet(
	todo_applications.Set,
)
