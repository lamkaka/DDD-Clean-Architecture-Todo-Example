package schema

import (
	"time"
	"todo-server/domains"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/samber/lo"
)

// TodoDoc holds the schema definition for the TodoDoc entity.
type TodoDoc struct {
	ent.Schema
}

func (TodoDoc) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "todos"},
	}
}

// Fields of the TodoDoc
func (TodoDoc) Fields() []ent.Field {

	statusValues := lo.Map(domains.GetTodoAllStatuses(), func(status domains.TodoStatus, index int) string {
		return string(status)
	})

	return []ent.Field{
		field.String("id"),
		field.String("name"),
		field.String("desc"),
		field.Time("due_at"),
		field.Enum("status").Values(statusValues...),
		field.Time("created_at").Default(time.Now()),
	}
}

// Edges of the TodoDoc
func (TodoDoc) Edges() []ent.Edge {
	return []ent.Edge{}
}

func (TodoDoc) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name"),
		index.Fields("due_at"),
		index.Fields("status"),
	}
}
