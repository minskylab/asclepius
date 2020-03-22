package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/google/uuid"
)

// EpidemiologicResults holds the schema definition for the EpidemiologicResults entity.
type EpidemiologicResults struct {
	ent.Schema
}

// Fields of the EpidemiologicResults.
func (EpidemiologicResults) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}),
		field.Strings("visitedPlaces").Optional(),
		field.Bool("infectedFamily").Optional(),
		field.Int("fromInfectedPlace").Min(0).Max(5).Optional(),
		field.Int("toInfectedPlace").Min(0).Max(5).Optional(),
	}
}

// Edges of the EpidemiologicResults.
func (EpidemiologicResults) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("test", Test.Type).Ref("epidemiologic").Unique(),
	}
}
