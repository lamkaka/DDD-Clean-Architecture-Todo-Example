package health

import (
	"libs/http_server"
	"strings"
)

type LogFilter http_server.LogFilter

func NewLogFilter() LogFilter {
	return func(requestCtx *http_server.RequestContext) bool {
		return !strings.HasPrefix(requestCtx.Path(), ROUTE_PREFIX)
	}
}
