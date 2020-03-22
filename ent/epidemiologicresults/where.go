// Code generated by entc, DO NOT EDIT.

package epidemiologicresults

import (
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/minskylab/asclepius/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id uuid.UUID) predicate.EpidemiologicResults {
	return predicate.EpidemiologicResults(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.EpidemiologicResults {
	return predicate.EpidemiologicResults(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.EpidemiologicResults {
	return predicate.EpidemiologicResults(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.EpidemiologicResults {
	return predicate.EpidemiologicResults(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.EpidemiologicResults {
	return predicate.EpidemiologicResults(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.EpidemiologicResults {
	return predicate.EpidemiologicResults(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.EpidemiologicResults {
	return predicate.EpidemiologicResults(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.EpidemiologicResults {
	return predicate.EpidemiologicResults(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.EpidemiologicResults {
	return predicate.EpidemiologicResults(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// InfectedFamily applies equality check predicate on the "infectedFamily" field. It's identical to InfectedFamilyEQ.
func InfectedFamily(v bool) predicate.EpidemiologicResults {
	return predicate.EpidemiologicResults(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldInfectedFamily), v))
	})
}

// FromInfectedPlace applies equality check predicate on the "fromInfectedPlace" field. It's identical to FromInfectedPlaceEQ.
func FromInfectedPlace(v int) predicate.EpidemiologicResults {
	return predicate.EpidemiologicResults(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFromInfectedPlace), v))
	})
}

// ToInfectedPlace applies equality check predicate on the "toInfectedPlace" field. It's identical to ToInfectedPlaceEQ.
func ToInfectedPlace(v int) predicate.EpidemiologicResults {
	return predicate.EpidemiologicResults(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldToInfectedPlace), v))
	})
}

// VisitedPlacesIsNil applies the IsNil predicate on the "visitedPlaces" field.
func VisitedPlacesIsNil() predicate.EpidemiologicResults {
	return predicate.EpidemiologicResults(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldVisitedPlaces)))
	})
}

// VisitedPlacesNotNil applies the NotNil predicate on the "visitedPlaces" field.
func VisitedPlacesNotNil() predicate.EpidemiologicResults {
	return predicate.EpidemiologicResults(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldVisitedPlaces)))
	})
}

// InfectedFamilyEQ applies the EQ predicate on the "infectedFamily" field.
func InfectedFamilyEQ(v bool) predicate.EpidemiologicResults {
	return predicate.EpidemiologicResults(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldInfectedFamily), v))
	})
}

// InfectedFamilyNEQ applies the NEQ predicate on the "infectedFamily" field.
func InfectedFamilyNEQ(v bool) predicate.EpidemiologicResults {
	return predicate.EpidemiologicResults(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldInfectedFamily), v))
	})
}

// InfectedFamilyIsNil applies the IsNil predicate on the "infectedFamily" field.
func InfectedFamilyIsNil() predicate.EpidemiologicResults {
	return predicate.EpidemiologicResults(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldInfectedFamily)))
	})
}

// InfectedFamilyNotNil applies the NotNil predicate on the "infectedFamily" field.
func InfectedFamilyNotNil() predicate.EpidemiologicResults {
	return predicate.EpidemiologicResults(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldInfectedFamily)))
	})
}

// FromInfectedPlaceEQ applies the EQ predicate on the "fromInfectedPlace" field.
func FromInfectedPlaceEQ(v int) predicate.EpidemiologicResults {
	return predicate.EpidemiologicResults(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFromInfectedPlace), v))
	})
}

// FromInfectedPlaceNEQ applies the NEQ predicate on the "fromInfectedPlace" field.
func FromInfectedPlaceNEQ(v int) predicate.EpidemiologicResults {
	return predicate.EpidemiologicResults(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldFromInfectedPlace), v))
	})
}

// FromInfectedPlaceIn applies the In predicate on the "fromInfectedPlace" field.
func FromInfectedPlaceIn(vs ...int) predicate.EpidemiologicResults {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.EpidemiologicResults(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldFromInfectedPlace), v...))
	})
}

// FromInfectedPlaceNotIn applies the NotIn predicate on the "fromInfectedPlace" field.
func FromInfectedPlaceNotIn(vs ...int) predicate.EpidemiologicResults {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.EpidemiologicResults(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldFromInfectedPlace), v...))
	})
}

// FromInfectedPlaceGT applies the GT predicate on the "fromInfectedPlace" field.
func FromInfectedPlaceGT(v int) predicate.EpidemiologicResults {
	return predicate.EpidemiologicResults(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldFromInfectedPlace), v))
	})
}

// FromInfectedPlaceGTE applies the GTE predicate on the "fromInfectedPlace" field.
func FromInfectedPlaceGTE(v int) predicate.EpidemiologicResults {
	return predicate.EpidemiologicResults(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldFromInfectedPlace), v))
	})
}

// FromInfectedPlaceLT applies the LT predicate on the "fromInfectedPlace" field.
func FromInfectedPlaceLT(v int) predicate.EpidemiologicResults {
	return predicate.EpidemiologicResults(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldFromInfectedPlace), v))
	})
}

// FromInfectedPlaceLTE applies the LTE predicate on the "fromInfectedPlace" field.
func FromInfectedPlaceLTE(v int) predicate.EpidemiologicResults {
	return predicate.EpidemiologicResults(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldFromInfectedPlace), v))
	})
}

// FromInfectedPlaceIsNil applies the IsNil predicate on the "fromInfectedPlace" field.
func FromInfectedPlaceIsNil() predicate.EpidemiologicResults {
	return predicate.EpidemiologicResults(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldFromInfectedPlace)))
	})
}

// FromInfectedPlaceNotNil applies the NotNil predicate on the "fromInfectedPlace" field.
func FromInfectedPlaceNotNil() predicate.EpidemiologicResults {
	return predicate.EpidemiologicResults(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldFromInfectedPlace)))
	})
}

// ToInfectedPlaceEQ applies the EQ predicate on the "toInfectedPlace" field.
func ToInfectedPlaceEQ(v int) predicate.EpidemiologicResults {
	return predicate.EpidemiologicResults(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldToInfectedPlace), v))
	})
}

// ToInfectedPlaceNEQ applies the NEQ predicate on the "toInfectedPlace" field.
func ToInfectedPlaceNEQ(v int) predicate.EpidemiologicResults {
	return predicate.EpidemiologicResults(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldToInfectedPlace), v))
	})
}

// ToInfectedPlaceIn applies the In predicate on the "toInfectedPlace" field.
func ToInfectedPlaceIn(vs ...int) predicate.EpidemiologicResults {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.EpidemiologicResults(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldToInfectedPlace), v...))
	})
}

// ToInfectedPlaceNotIn applies the NotIn predicate on the "toInfectedPlace" field.
func ToInfectedPlaceNotIn(vs ...int) predicate.EpidemiologicResults {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.EpidemiologicResults(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldToInfectedPlace), v...))
	})
}

// ToInfectedPlaceGT applies the GT predicate on the "toInfectedPlace" field.
func ToInfectedPlaceGT(v int) predicate.EpidemiologicResults {
	return predicate.EpidemiologicResults(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldToInfectedPlace), v))
	})
}

// ToInfectedPlaceGTE applies the GTE predicate on the "toInfectedPlace" field.
func ToInfectedPlaceGTE(v int) predicate.EpidemiologicResults {
	return predicate.EpidemiologicResults(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldToInfectedPlace), v))
	})
}

// ToInfectedPlaceLT applies the LT predicate on the "toInfectedPlace" field.
func ToInfectedPlaceLT(v int) predicate.EpidemiologicResults {
	return predicate.EpidemiologicResults(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldToInfectedPlace), v))
	})
}

// ToInfectedPlaceLTE applies the LTE predicate on the "toInfectedPlace" field.
func ToInfectedPlaceLTE(v int) predicate.EpidemiologicResults {
	return predicate.EpidemiologicResults(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldToInfectedPlace), v))
	})
}

// ToInfectedPlaceIsNil applies the IsNil predicate on the "toInfectedPlace" field.
func ToInfectedPlaceIsNil() predicate.EpidemiologicResults {
	return predicate.EpidemiologicResults(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldToInfectedPlace)))
	})
}

// ToInfectedPlaceNotNil applies the NotNil predicate on the "toInfectedPlace" field.
func ToInfectedPlaceNotNil() predicate.EpidemiologicResults {
	return predicate.EpidemiologicResults(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldToInfectedPlace)))
	})
}

// HasTest applies the HasEdge predicate on the "test" edge.
func HasTest() predicate.EpidemiologicResults {
	return predicate.EpidemiologicResults(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TestTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, TestTable, TestColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTestWith applies the HasEdge predicate on the "test" edge with a given conditions (other predicates).
func HasTestWith(preds ...predicate.Test) predicate.EpidemiologicResults {
	return predicate.EpidemiologicResults(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TestInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, TestTable, TestColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.EpidemiologicResults) predicate.EpidemiologicResults {
	return predicate.EpidemiologicResults(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.EpidemiologicResults) predicate.EpidemiologicResults {
	return predicate.EpidemiologicResults(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.EpidemiologicResults) predicate.EpidemiologicResults {
	return predicate.EpidemiologicResults(func(s *sql.Selector) {
		p(s.Not())
	})
}