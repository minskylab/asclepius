package schema

import (
	"time"

	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/google/uuid"
)

// Doctor holds the schema definition for the Doctor entity.
type Doctor struct {
	ent.Schema
}

// Fields of the Doctor.
func (Doctor) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}),
		field.Strings("name"),
		field.String("email").Unique(),
		field.String("phone").Unique(),
		field.Enum("state").Values("idle", "working", "unavailable"),
		field.Time("lastConnection").Optional().Default(func() time.Time {
			return time.Now()
		}),
		field.Bool("volunteer").Default(true),
	}
}

// Edges of the Doctor.
func (Doctor) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("notes", MedicalNote.Type),
		edge.To("responses", TaskResponse.Type),
		edge.From("tasks", Task.Type).Ref("responsible"),
	}
}
