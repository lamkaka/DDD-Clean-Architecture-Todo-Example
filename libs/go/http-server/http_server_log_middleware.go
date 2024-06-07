package http_server

import (
	"fmt"
	"io"

	fiber_logger "github.com/gofiber/fiber/v2/middleware/logger"
)

// Filter whether the request should be logged, returns true if should be logged
type LogFilter func(*RequestContext) bool

func newLogMiddleware(config Config) Controller {
	return logMiddleware{config}
}

type logMiddleware struct {
	config Config
}

func (m logMiddleware) RegisterRoutes(router Router) {
	router.Use(func(requestCtx *RequestContext) error {
		if !m.config.LogFilter(requestCtx) {
			return requestCtx.Next()
		}

		fmt.Printf(
			"Received request: %s %vB",
			getBaseLog(requestCtx),
			requestCtx.Request().Header.ContentLength(),
		)
		return requestCtx.Next()
	})

	router.Use(fiber_logger.New(fiber_logger.Config{
		Output: io.Discard,
		Next: func(requestCtx *RequestContext) bool {
			return !m.config.LogFilter(requestCtx)
		},
		Done: func(requestCtx *RequestContext, _ []byte) {
			fmt.Printf(
				"Sent response: %s %v %vB",
				getBaseLog(requestCtx),
				requestCtx.Response().StatusCode(),
				len(requestCtx.Response().Body()),
			)
		},
	}))
}

func getBaseLog(requestCtx *RequestContext) string {
	return fmt.Sprintf(
		"%s:%s %s %s %s",
		requestCtx.IP(),
		requestCtx.Port(),
		requestCtx.Method(),
		requestCtx.Path(),
		requestCtx.Context().QueryArgs().QueryString(),
	)
}
