package todo_applications

import (
	"context"
	"time"
	"todo-server/domains"
	"todo-server/infras/log"
)

type ReadService interface {
	GetByID(ctx context.Context, id string) (domains.Todo, error)
	List(ctx context.Context, filter QueryFilter) (domains.Todos, error)
}

func NewReadService(rootLogger log.RootLogger, repository Repository) ReadService {
	logger := rootLogger.Child("todoReadService")
	return readService{logger, repository}
}

type readService struct {
	logger     log.Logger
	repository Repository
}

func (svc readService) GetByID(ctx context.Context, id string) (domains.Todo, error) {
	svc.logger.Info(ctx, "Getting todo %s", id)

	todo, err := svc.repository.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return todo, nil
}

func (svc readService) List(ctx context.Context, filter QueryFilter) (domains.Todos, error) {
	svc.logger.Info(ctx, "Listing todos")

	todos, err := svc.repository.List(ctx, filter)
	if err != nil {
		return nil, err
	}

	return todos, nil
}

type QueryFilter struct {
	Statuses   []domains.TodoStatus
	DueAtAfter *time.Time
}
