package todo_applications

import (
	"context"
	"time"
	"todo-server/domains"
	"todo-server/infras/log"
)

type WriteService interface {
	Create(ctx context.Context, cmd CreateCommand) (domains.Todo, error)
	UpdateByID(ctx context.Context, id string, cmd UpdateCommand) (domains.Todo, error)
	DeleteByID(ctx context.Context, id string) error
}

func NewWriteService(rootLogger log.RootLogger, repository Repository) WriteService {
	logger := rootLogger.Child("todoWriteService")
	return writeService{logger, repository}
}

type writeService struct {
	logger     log.Logger
	repository Repository
}

type CreateCommand struct {
	Name        *string
	Description *string
	DueAt       *time.Time
}

func (svc writeService) Create(ctx context.Context, cmd CreateCommand) (domains.Todo, error) {
	svc.logger.Info(ctx, "Creating todo %v", cmd)

	var name string
	if cmd.Name != nil {
		name = *cmd.Name
	}
	var desc string
	if cmd.Description != nil {
		desc = *cmd.Description
	}
	var dueAt time.Time
	if cmd.DueAt != nil {
		dueAt = *cmd.DueAt
	}
	todo, err := domains.NewTodo(
		"",
		name,
		desc,
		dueAt,
		"",
	)
	if err != nil {
		svc.logger.Error(ctx, err)
		return nil, err
	}

	err = svc.repository.Create(ctx, todo)
	if err != nil {
		return nil, err
	}

	return todo, nil
}

func (svc writeService) DeleteByID(ctx context.Context, id string) error {
	svc.logger.Info(ctx, "Deleting todo %s", id)

	err := svc.repository.DeleteByID(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
