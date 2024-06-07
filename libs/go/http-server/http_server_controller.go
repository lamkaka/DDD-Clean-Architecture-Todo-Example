package http_server

// A Controller is responsible for handling http requests that are going to specific path(s).
type Controller interface {
	RegisterRoutes(router Router)
}

type ControllerFunc func(router Router)

func (f ControllerFunc) RegisterRoutes(router Router) {
	f(router)
}
