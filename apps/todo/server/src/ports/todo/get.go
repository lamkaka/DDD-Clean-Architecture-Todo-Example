package todo_rest

import (
	"libs/http_server"
	"time"
	"todo-server/domains"
)

type ResponseBody struct {
	ID          string             `json:"id"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	DueAt       time.Time          `json:"dueAt"`
	Status      domains.TodoStatus `json:"status"`
}

func (ctrl controller) getByID(reqCtx *http_server.RequestContext) error {
	ctx := reqCtx.UserContext()

	id := reqCtx.Params("todoID")

	todo, err := ctrl.readService.GetByID(ctx, id)
	if err != nil {
		return err
	}

	resBody, err := EntityToResponse(todo)
	if err != nil {
		ctrl.logger.Error(ctx, err)
		return err
	}

	return reqCtx.Status(200).JSON(resBody)
}

func EntityToResponse(todo domains.Todo) (ResponseBody, error) {
	resBody := ResponseBody{
		ID:          todo.ID(),
		Name:        todo.Name(),
		Description: todo.Description(),
		DueAt:       todo.DueAt(),
		Status:      todo.Status(),
	}

	return resBody, nil
}
