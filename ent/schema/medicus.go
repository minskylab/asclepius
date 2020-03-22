package schema

import (
	"time"

	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/google/uuid"
)

// Medicus holds the schema definition for the Medicus entity.
type Medicus struct {
	ent.Schema
}

// Fields of the Medicus.
func (Medicus) Fields() []ent.Field {
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

// Edges of the Medicus.
func (Medicus) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("notes", MedicalNote.Type),
	}
}
