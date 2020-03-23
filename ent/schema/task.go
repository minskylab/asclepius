package schema

import (
	"time"

	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/google/uuid"
)

// Task holds the schema definition for the Task entity.
type Task struct {
	ent.Schema
}

// Fields of the Task.
func (Task) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}),
		field.String("title").Immutable(),
		field.Time("startAt").Default(func() time.Time {
			return time.Now()
		}),
		field.Time("endsAt").Default(func() time.Time {
			return time.Now().Add(15*time.Minute)
		}),
		field.Strings("description").Optional(),
	}
}

// Edges of the Task.
func (Task) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("responsible", Doctor.Type),
		edge.To("responses", TaskResponse.Type),
		edge.From("schedule", Schedule.Type).Ref("tasks").Unique(),
	}
}
