package todo_applications

import "github.com/google/wire"

var Set = wire.NewSet(
	NewReadService,
)
