// Code generated by entc, DO NOT EDIT.

package history

import (
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/minskylab/asclepius/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id uuid.UUID) predicate.History {
	return predicate.History(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.History {
	return predicate.History(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.History {
	return predicate.History(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.History {
	return predicate.History(func(s *sql.Selector) {
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
func IDNotIn(ids ...uuid.UUID) predicate.History {
	return predicate.History(func(s *sql.Selector) {
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
func IDGT(id uuid.UUID) predicate.History {
	return predicate.History(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.History {
	return predicate.History(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.History {
	return predicate.History(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.History {
	return predicate.History(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// HasPatient applies the HasEdge predicate on the "patient" edge.
func HasPatient() predicate.History {
	return predicate.History(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PatientTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, PatientTable, PatientColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPatientWith applies the HasEdge predicate on the "patient" edge with a given conditions (other predicates).
func HasPatientWith(preds ...predicate.Patient) predicate.History {
	return predicate.History(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PatientInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, PatientTable, PatientColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasTests applies the HasEdge predicate on the "tests" edge.
func HasTests() predicate.History {
	return predicate.History(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TestsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, TestsTable, TestsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTestsWith applies the HasEdge predicate on the "tests" edge with a given conditions (other predicates).
func HasTestsWith(preds ...predicate.Test) predicate.History {
	return predicate.History(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TestsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, TestsTable, TestsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasNotes applies the HasEdge predicate on the "notes" edge.
func HasNotes() predicate.History {
	return predicate.History(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(NotesTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, NotesTable, NotesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasNotesWith applies the HasEdge predicate on the "notes" edge with a given conditions (other predicates).
func HasNotesWith(preds ...predicate.MedicalNote) predicate.History {
	return predicate.History(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(NotesInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, NotesTable, NotesColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.History) predicate.History {
	return predicate.History(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.History) predicate.History {
	return predicate.History(func(s *sql.Selector) {
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
func Not(p predicate.History) predicate.History {
	return predicate.History(func(s *sql.Selector) {
		p(s.Not())
	})
}
