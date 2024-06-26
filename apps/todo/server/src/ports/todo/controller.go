package todo_rest

import (
	"libs/http_server"
	todo_applications "todo-server/applications/todo"
	"todo-server/infras/log"
)

type Controller http_server.Controller

func NewController(
	rootLogger log.RootLogger,
	readService todo_applications.ReadService,
	writeService todo_applications.WriteService,
) Controller {
	logger := rootLogger.Child("todoRestController")
	return controller{logger, readService, writeService}
}

type controller struct {
	logger       log.Logger
	readService  todo_applications.ReadService
	writeService todo_applications.WriteService
}

func (ctrl controller) RegisterRoutes(router http_server.Router) {
	prefix := "/todos"
	routeGroup := router.Group(prefix)
	routeGroup.Get("", ctrl.list)
	routeGroup.Get("/:todoID", ctrl.getByID)
	routeGroup.Post("", ctrl.create)
	routeGroup.Patch("/:todoID", ctrl.updateByID)
	routeGroup.Delete("/:todoID", ctrl.deleteByID)
}
