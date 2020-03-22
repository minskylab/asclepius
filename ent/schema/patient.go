package schema

import (
	"time"

	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/google/uuid"
)

// Patient holds the schema definition for the Patient entity.
type Patient struct {
	ent.Schema
}

func (Patient) Indexes() []ent.Index {
	return []ent.Index{
		// index.Fields("facebookID", "watsonID", "phone", "email"),
	}
}

// Fields of the Patient.
func (Patient) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}),
		field.String("name"),
		field.String("phone").Unique(),
		field.Int("age").Positive().Optional(),
		field.String("email").Unique().Optional(),
		field.String("password").Sensitive().MinLen(6).Optional(),
		field.String("facebookID").Unique().Optional(),
		field.String("watsonID").Unique().Optional(),
		field.Time("first_contact").Default(func() time.Time {
			return time.Now()
		}),
		field.Strings("conditions").Optional(),
	}
}

// Edges of the Patient.
func (Patient) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("history", History.Type).Required().Unique(),
		edge.To("schedule", Schedule.Type).Unique(),
	}
}
