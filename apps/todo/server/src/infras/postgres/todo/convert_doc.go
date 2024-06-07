package todo_postgres

import (
	"todo-server/domains"
	"todo-server/infras/postgres/client/ent"
)

func docToEntity(todoDoc *ent.TodoDoc) (domains.Todo, error) {
	return domains.NewTodo(
		todoDoc.ID,
		todoDoc.Name,
		todoDoc.Desc,
		todoDoc.DueAt,
		domains.TodoStatus(todoDoc.Status),
	)
}

func docsToEntities(todoDocs []*ent.TodoDoc) ([]domains.Todo, error) {
	var todos domains.Todos
	for _, todoDoc := range todoDocs {
		todo, err := docToEntity(todoDoc)
		if err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}

	return todos, nil
}
