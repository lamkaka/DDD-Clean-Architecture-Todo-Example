package todo_rest

import (
	"libs/http_server"
	todo_applications "todo-server/applications/todo"
	"todo-server/domains"
)

type ListQuery struct {
	GroupID *string `query:"groupID, omitempty"`
}

func QueryToFilter(query ListQuery) todo_applications.QueryFilter {
	return todo_applications.QueryFilter{}
}

type ListResponseBody struct {
	Data []ResponseBody `json:"data"`
}

func (ctrl controller) list(reqCtx *http_server.RequestContext) error {
	ctx := reqCtx.UserContext()

	var query ListQuery
	err := reqCtx.QueryParser(&query)
	if err != nil {
		ctrl.logger.Warn(ctx, err)
		return err
	}

	filter := QueryToFilter(query)

	todos, err := ctrl.readService.List(ctx, filter)
	if err != nil {
		return err
	}

	resBody, err := EntitiesToResponse(todos)
	if err != nil {
		ctrl.logger.Error(ctx, err)
		return err
	}

	return reqCtx.Status(200).JSON(resBody)
}

func EntitiesToResponse(todos domains.Todos) (ListResponseBody, error) {
	listResBody := ListResponseBody{Data: []ResponseBody{}}

	for _, todo := range todos {
		resBody, err := EntityToResponse(todo)
		if err != nil {
			return ListResponseBody{}, err
		}
		listResBody.Data = append(listResBody.Data, resBody)
	}

	return listResBody, nil
}
