package health

import (
	"github.com/google/wire"
)

var Set = wire.NewSet(NewController, NewLogFilter)
