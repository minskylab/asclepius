package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/google/uuid"
)

// ClinicalResults holds the schema definition for the ClinicalResults entity.
type ClinicalResults struct {
	ent.Schema
}

// Fields of the ClinicalResults.
func (ClinicalResults) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}),
		field.Bool("generalDiscomfort"),
		field.Bool("fever"),
		field.Bool("thirdAge"),
		field.Bool("dyspnea"),
	}
}

// Edges of the ClinicalResults.
func (ClinicalResults) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("test", Test.Type).Ref("clinical").Unique(),
	}
}
