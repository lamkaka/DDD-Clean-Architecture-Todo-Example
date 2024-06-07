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

func (repo repository) List(ctx context.Context, filter todo_applications.QueryFilter) (domains.Todos, error) {
	repo.logger.Info(ctx, "Listing todos")

	todoDocQuery := repo.client.TodoDoc.Query()

	if filter.Statuses != nil {
		statuses := lo.Map(filter.Statuses, func(status domains.TodoStatus, index int) tododoc.Status {
			return tododoc.Status(status)
		})
		todoDocQuery = todoDocQuery.Where(tododoc.StatusIn(statuses...))
	}
	if filter.DueAtAfter != nil {
		todoDocQuery = todoDocQuery.Where(tododoc.CreatedAtGT(*filter.DueAtAfter))
	}

	todoDocs, err := todoDocQuery.All(ctx)
	if err != nil {
		repo.logger.Error(ctx, err)
		return nil, err
	}

	todos, err := docsToEntities(todoDocs)
	if err != nil {
		repo.logger.Error(ctx, err)
		return nil, err
	}

	return todos, nil
}

func (repo repository) Create(ctx context.Context, todo domains.Todo) error {
	repo.logger.Info(ctx, "Creating todo %s", todo.ID())

	docCreate := repo.client.TodoDoc.Create().
		SetID(todo.ID()).
		SetName(todo.Name()).
		SetDesc(todo.Description()).
		SetDueAt(todo.DueAt()).
		SetStatus(tododoc.Status(todo.Status()))

	err := docCreate.Exec(ctx)
	if err != nil {
		repo.logger.Error(ctx, err)
		return err
	}

	return nil
}
func (repo repository) handleEntNotFoundError(ctx context.Context, id string) error {
	err := errors.New(errors.ErrorEntityNotFound, fmt.Sprintf("Todo %s not found", id))
	repo.logger.Warn(ctx, err)
	return err
}
