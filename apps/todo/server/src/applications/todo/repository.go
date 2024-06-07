package todo_applications

import (
	"context"
	"todo-server/domains"
)

type Repository interface {
	GetByID(ctx context.Context, id string) (domains.Todo, error)
	List(ctx context.Context, filter QueryFilter) (domains.Todos, error)
}
