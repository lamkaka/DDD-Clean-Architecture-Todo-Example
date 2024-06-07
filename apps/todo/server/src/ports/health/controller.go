package health

import (
	"libs/http_server"
)

type Controller http_server.Controller

const ROUTE_PREFIX = "/healthz"

func NewController() Controller {
	return &controller{}
}

type controller struct {
}

func (ctrl controller) RegisterRoutes(router http_server.Router) {
	routes := router.Group(ROUTE_PREFIX)
	routes.Get("/liveness", ctrl.healthCheckLiveness)
	routes.Get("/readiness", ctrl.healthCheckReadiness)
}

func (ctrl controller) healthCheckLiveness(ctx *http_server.RequestContext) error {
	return ctx.Status(200).JSON("OK")
}

func (ctrl controller) healthCheckReadiness(ctx *http_server.RequestContext) error {
	return ctx.Status(200).JSON("OK")
}
