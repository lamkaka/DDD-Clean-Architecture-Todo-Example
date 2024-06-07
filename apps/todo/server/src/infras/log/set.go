package log

import (
	"github.com/google/wire"
)

var Set = wire.NewSet(NewRootLogger)
