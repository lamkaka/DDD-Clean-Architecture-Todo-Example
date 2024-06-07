package todo_rest

import (
	"libs/http_server"
	"time"
	todo_applications "todo-server/applications/todo"
)

type CreateRequestBody struct {
	Name        *string    `json:"name"`
	Description *string    `json:"description"`
	DueAt       *time.Time `json:"dueAt"`
}

func (ctrl controller) create(reqCtx *http_server.RequestContext) error {
	ctx := reqCtx.UserContext()

	var reqBody CreateRequestBody
	err := reqCtx.BodyParser(&reqBody)
	if err != nil {
		err = http_server.NewError(400, err.Error())
		ctrl.logger.Error(ctx, err)
		return err
	}

	cmd := createRequestBodyToCmd(reqBody)
	todo, err := ctrl.writeService.Create(ctx, cmd)
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

func createRequestBodyToCmd(reqBody CreateRequestBody) todo_applications.CreateCommand {
	command := todo_applications.CreateCommand{
		Name:        reqBody.Name,
		Description: reqBody.Description,
		DueAt:       reqBody.DueAt,
	}
	return command
}
