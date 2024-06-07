package todo_rest

import (
	"libs/http_server"
	"time"
	todo_applications "todo-server/applications/todo"
	"todo-server/domains"
)

type UpdateRequestBody struct {
	Name        *string             `json:"name"`
	Description *string             `json:"description"`
	DueAt       *time.Time          `json:"dueAt"`
	Status      *domains.TodoStatus `json:"status"`
}

func (ctrl controller) updateByID(reqCtx *http_server.RequestContext) error {
	ctx := reqCtx.UserContext()
	id := reqCtx.Params("todoID")

	var reqBody UpdateRequestBody
	err := reqCtx.BodyParser(&reqBody)
	if err != nil {
		err = http_server.NewError(400, err.Error())
		ctrl.logger.Error(ctx, err)
		return err
	}

	cmd := updateRequestBodyToCmd(reqBody)

	todo, err := ctrl.writeService.UpdateByID(ctx, id, cmd)
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

func updateRequestBodyToCmd(reqBody UpdateRequestBody) todo_applications.UpdateCommand {
	command := todo_applications.UpdateCommand{
		Name:        reqBody.Name,
		Description: reqBody.Description,
		DueAt:       reqBody.DueAt,
		Status:      reqBody.Status,
	}
	return command
}
