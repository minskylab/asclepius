package schema

import (
	"time"

	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/google/uuid"
)

// MedicalNote holds the schema definition for the MedicalNote entity.
type MedicalNote struct {
	ent.Schema
}

// Fields of the MedicalNote.
func (MedicalNote) Fields() []ent.Field {
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

// Edges of the MedicalNote.
func (MedicalNote) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("history", History.Type).Unique().Ref("notes"),
		edge.From("owner", Doctor.Type).Unique().Ref("notes"),
	}
}
