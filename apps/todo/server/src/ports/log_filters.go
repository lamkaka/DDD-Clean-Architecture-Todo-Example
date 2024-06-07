package ports

import (
	"libs/http_server"
	"todo-server/ports/health"
)

func NewLogFilters(
	healthLogFilter health.LogFilter,
) http_server.LogFilter {
	var filters []http_server.LogFilter = []http_server.LogFilter{
		http_server.LogFilter(healthLogFilter),
	}

	return func(requestCtx *http_server.RequestContext) bool {
		for _, filter := range filters {
			if !filter(requestCtx) {
				return false
			}
		}

		return true
	}
}
