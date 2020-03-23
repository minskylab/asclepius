package schema

import (
	"time"

	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/google/uuid"
)

// TaskResponse holds the schema definition for the TaskResponse entity.
type TaskResponse struct {
	ent.Schema
}

// Fields of the TaskResponse.
func (TaskResponse) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}),
		field.Time("at").Default(func() time.Time {
			return time.Now()
		}),
		field.Time("lastChange").Default(func() time.Time {
			return time.Now()
		}),
		field.Strings("observations"),
		field.Strings("meta").Optional(),
	}
}

// Edges of the TaskResponse.
func (TaskResponse) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("author", Doctor.Type).Required().Unique().Ref("responses"),
		edge.From("task", Task.Type).Ref("responses").Required().Unique(),
	}
}
