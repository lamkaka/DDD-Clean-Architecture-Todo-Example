package openapi_rest

import (
	"libs/http_server"
)

type Controller http_server.Controller

func NewController() Controller {
	return http_server.ControllerFunc(func(router http_server.Router) {
		router.Get("/openapi", func(requestCtx *http_server.RequestContext) error {
			return requestCtx.SendFile("/var/openapi.yaml")
		})
	})
}
