package todo_rest

import (
	"libs/http_server"
)

type SucessResponseBody struct {
	Success bool `json:"success"`
}

func (ctrl controller) deleteByID(reqCtx *http_server.RequestContext) error {
	ctx := reqCtx.UserContext()

	id := reqCtx.Params("todoID")

	err := ctrl.writeService.DeleteByID(ctx, id)
	if err != nil {
		return err
	}
	resBody := SucessResponseBody{Success: true}
	return reqCtx.Status(200).JSON(resBody)
}
