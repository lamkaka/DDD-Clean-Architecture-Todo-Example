package todo_rest

import (
	"libs/http_server"
	"time"
	todo_applications "todo-server/applications/todo"
	"todo-server/domains"
)

type ListRequestQuery struct {
	DueAtAfter *time.Time `query:"dueAtAfter"`
	Statuses   []string   `query:"statuses, omitempty"`
}

func QueryToFilter(query ListRequestQuery) todo_applications.QueryFilter {
	var statuses []domains.TodoStatus
	for _, reqStatus := range query.Statuses {
		status := domains.TodoStatus(reqStatus)
		statuses = append(statuses, status)
	}
	return todo_applications.QueryFilter{
		DueAtAfter: query.DueAtAfter,
		Statuses:   statuses,
	}
}

type ListResponseBody struct {
	Data []ResponseBody `json:"data"`
}

func (ctrl controller) list(reqCtx *http_server.RequestContext) error {
	ctx := reqCtx.UserContext()

	var query ListRequestQuery
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
