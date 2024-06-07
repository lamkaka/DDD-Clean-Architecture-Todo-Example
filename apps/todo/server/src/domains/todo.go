package domains

import (
	"fmt"
	"libs/errors"
	"time"

	"github.com/lucsky/cuid"
)

type TodoStatus string

const (
	TodoNotStartedStatus TodoStatus = "NotStarted"
	TodoInProgressStatus TodoStatus = "InProgress"
	TodoCompletedStatus  TodoStatus = "Completed"
)

const TODO_NAME_LENGTH_MIN = int(3)

type Todo interface {
	ID() string
	Name() string
	Description() string
	DueAt() time.Time
	Status() TodoStatus
	SetName(name string) error
	SetDescription(desc string) error
	SetDueAt(dueAt time.Time) error
	SetStatus(status TodoStatus) error
}
type Todos = []Todo

func NewTodo(id, name, desc string, dueAt time.Time, status TodoStatus) (Todo, error) {
	if id == "" {
		id = cuid.New()
	}
	if err := validateName(name); err != nil {
		return nil, err
	}
	if err := validateDueAt(dueAt); err != nil {
		return nil, err
	}
	if status == "" {
		status = TodoNotStartedStatus
	}

	return &todo{
		id:     id,
		name:   name,
		desc:   desc,
		dueAt:  dueAt,
		status: status,
	}, nil
}

type todo struct {
	id     string
	name   string
	desc   string
	dueAt  time.Time
	status TodoStatus
}

func (t *todo) ID() string {
	return t.id
}

func (t *todo) Name() string {
	return t.name
}

func (t *todo) SetName(name string) error {
	if err := validateName(name); err != nil {
		return err
	}
	t.name = name
	return nil
}

func (t *todo) Description() string {
	return t.desc
}

func (t *todo) SetDescription(desc string) error {
	t.desc = desc
	return nil
}

func (t *todo) DueAt() time.Time {
	return t.dueAt
}

func (t *todo) SetDueAt(dueAt time.Time) error {
	if err := validateDueAt(dueAt); err != nil {
		return err
	}
	t.dueAt = dueAt
	return nil
}

func (t *todo) Status() TodoStatus {
	return t.status
}

func (t *todo) SetStatus(status TodoStatus) error {
	t.status = status
	return nil
}

func GetTodoAllStatuses() []TodoStatus {
	return []TodoStatus{
		TodoNotStartedStatus,
		TodoInProgressStatus,
		TodoCompletedStatus,
	}
}

func validateName(name string) error {
	if len(name) < TODO_NAME_LENGTH_MIN {
		err := errors.New(errors.ErrorEntityValidation, fmt.Sprintf("Name must have at least %d characters", TODO_NAME_LENGTH_MIN))
		return err
	}
	return nil
}

func validateDueAt(dueAt time.Time) error {
	if !dueAt.IsZero() && dueAt.Before(time.Now()) {
		err := errors.New(errors.ErrorEntityValidation, "DueAt must be time after now. ")
		return err
	}
	return nil
}
