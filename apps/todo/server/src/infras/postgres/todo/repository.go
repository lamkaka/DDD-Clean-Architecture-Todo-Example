package todo_postgres

import (
	"context"
	"fmt"
	"libs/errors"
	todo_applications "todo-server/applications/todo"
	"todo-server/domains"
	"todo-server/infras/log"
	postgres_client "todo-server/infras/postgres/client"
	"todo-server/infras/postgres/client/ent"
	"todo-server/infras/postgres/client/ent/tododoc"

	"github.com/samber/lo"
)

func NewRepository(rootLogger log.RootLogger, client postgres_client.Client) todo_applications.Repository {
	logger := rootLogger.Child("todoPostgresRepository")
	return repository{logger, client}
}

type repository struct {
	logger log.Logger
	client postgres_client.Client
}

func (repo repository) GetByID(ctx context.Context, id string) (domains.Todo, error) {
	repo.logger.Info(ctx, "Getting todo %s", id)

	todoDoc, err := repo.client.TodoDoc.Query().Where(tododoc.ID(id)).Only(ctx)
	if ent.IsNotFound(err) {
		return nil, repo.handleEntNotFoundError(ctx, id)
	}
	if err != nil {
		repo.logger.Error(ctx, err)
		return nil, err
	}

	todo, err := docToEntity(todoDoc)
	if err != nil {
		repo.logger.Error(ctx, err)
		return nil, err
	}

	return todo, nil
}

func (repo repository) handleEntNotFoundError(ctx context.Context, id string) error {
	err := errors.New(errors.ErrorEntityNotFound, fmt.Sprintf("Todo %s not found", id))
	repo.logger.Warn(ctx, err)
	return err
}
