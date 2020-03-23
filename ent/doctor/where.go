// Code generated by entc, DO NOT EDIT.

package doctor

import (
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/minskylab/asclepius/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id uuid.UUID) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
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
func IDNotIn(ids ...uuid.UUID) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
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
func IDGT(id uuid.UUID) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Email applies equality check predicate on the "email" field. It's identical to EmailEQ.
func Email(v string) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEmail), v))
	})
}

// Phone applies equality check predicate on the "phone" field. It's identical to PhoneEQ.
func Phone(v string) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPhone), v))
	})
}

// LastConnection applies equality check predicate on the "lastConnection" field. It's identical to LastConnectionEQ.
func LastConnection(v time.Time) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLastConnection), v))
	})
}

// Volunteer applies equality check predicate on the "volunteer" field. It's identical to VolunteerEQ.
func Volunteer(v bool) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldVolunteer), v))
	})
}

// EmailEQ applies the EQ predicate on the "email" field.
func EmailEQ(v string) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEmail), v))
	})
}

// EmailNEQ applies the NEQ predicate on the "email" field.
func EmailNEQ(v string) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEmail), v))
	})
}

// EmailIn applies the In predicate on the "email" field.
func EmailIn(vs ...string) predicate.Doctor {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Doctor(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldEmail), v...))
	})
}

// EmailNotIn applies the NotIn predicate on the "email" field.
func EmailNotIn(vs ...string) predicate.Doctor {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Doctor(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldEmail), v...))
	})
}

// EmailGT applies the GT predicate on the "email" field.
func EmailGT(v string) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEmail), v))
	})
}

// EmailGTE applies the GTE predicate on the "email" field.
func EmailGTE(v string) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEmail), v))
	})
}

// EmailLT applies the LT predicate on the "email" field.
func EmailLT(v string) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEmail), v))
	})
}

// EmailLTE applies the LTE predicate on the "email" field.
func EmailLTE(v string) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEmail), v))
	})
}

// EmailContains applies the Contains predicate on the "email" field.
func EmailContains(v string) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldEmail), v))
	})
}

// EmailHasPrefix applies the HasPrefix predicate on the "email" field.
func EmailHasPrefix(v string) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldEmail), v))
	})
}

// EmailHasSuffix applies the HasSuffix predicate on the "email" field.
func EmailHasSuffix(v string) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldEmail), v))
	})
}

// EmailEqualFold applies the EqualFold predicate on the "email" field.
func EmailEqualFold(v string) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldEmail), v))
	})
}

// EmailContainsFold applies the ContainsFold predicate on the "email" field.
func EmailContainsFold(v string) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldEmail), v))
	})
}

// PhoneEQ applies the EQ predicate on the "phone" field.
func PhoneEQ(v string) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPhone), v))
	})
}

// PhoneNEQ applies the NEQ predicate on the "phone" field.
func PhoneNEQ(v string) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPhone), v))
	})
}

// PhoneIn applies the In predicate on the "phone" field.
func PhoneIn(vs ...string) predicate.Doctor {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Doctor(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPhone), v...))
	})
}

// PhoneNotIn applies the NotIn predicate on the "phone" field.
func PhoneNotIn(vs ...string) predicate.Doctor {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Doctor(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPhone), v...))
	})
}

// PhoneGT applies the GT predicate on the "phone" field.
func PhoneGT(v string) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPhone), v))
	})
}

// PhoneGTE applies the GTE predicate on the "phone" field.
func PhoneGTE(v string) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPhone), v))
	})
}

// PhoneLT applies the LT predicate on the "phone" field.
func PhoneLT(v string) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPhone), v))
	})
}

// PhoneLTE applies the LTE predicate on the "phone" field.
func PhoneLTE(v string) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPhone), v))
	})
}

// PhoneContains applies the Contains predicate on the "phone" field.
func PhoneContains(v string) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPhone), v))
	})
}

// PhoneHasPrefix applies the HasPrefix predicate on the "phone" field.
func PhoneHasPrefix(v string) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPhone), v))
	})
}

// PhoneHasSuffix applies the HasSuffix predicate on the "phone" field.
func PhoneHasSuffix(v string) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPhone), v))
	})
}

// PhoneEqualFold applies the EqualFold predicate on the "phone" field.
func PhoneEqualFold(v string) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPhone), v))
	})
}

// PhoneContainsFold applies the ContainsFold predicate on the "phone" field.
func PhoneContainsFold(v string) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPhone), v))
	})
}

// StateEQ applies the EQ predicate on the "state" field.
func StateEQ(v State) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldState), v))
	})
}

// StateNEQ applies the NEQ predicate on the "state" field.
func StateNEQ(v State) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldState), v))
	})
}

// StateIn applies the In predicate on the "state" field.
func StateIn(vs ...State) predicate.Doctor {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Doctor(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldState), v...))
	})
}

// StateNotIn applies the NotIn predicate on the "state" field.
func StateNotIn(vs ...State) predicate.Doctor {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Doctor(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldState), v...))
	})
}

// LastConnectionEQ applies the EQ predicate on the "lastConnection" field.
func LastConnectionEQ(v time.Time) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLastConnection), v))
	})
}

// LastConnectionNEQ applies the NEQ predicate on the "lastConnection" field.
func LastConnectionNEQ(v time.Time) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLastConnection), v))
	})
}

// LastConnectionIn applies the In predicate on the "lastConnection" field.
func LastConnectionIn(vs ...time.Time) predicate.Doctor {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Doctor(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldLastConnection), v...))
	})
}

// LastConnectionNotIn applies the NotIn predicate on the "lastConnection" field.
func LastConnectionNotIn(vs ...time.Time) predicate.Doctor {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Doctor(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldLastConnection), v...))
	})
}

// LastConnectionGT applies the GT predicate on the "lastConnection" field.
func LastConnectionGT(v time.Time) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLastConnection), v))
	})
}

// LastConnectionGTE applies the GTE predicate on the "lastConnection" field.
func LastConnectionGTE(v time.Time) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLastConnection), v))
	})
}

// LastConnectionLT applies the LT predicate on the "lastConnection" field.
func LastConnectionLT(v time.Time) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLastConnection), v))
	})
}

// LastConnectionLTE applies the LTE predicate on the "lastConnection" field.
func LastConnectionLTE(v time.Time) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLastConnection), v))
	})
}

// LastConnectionIsNil applies the IsNil predicate on the "lastConnection" field.
func LastConnectionIsNil() predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldLastConnection)))
	})
}

// LastConnectionNotNil applies the NotNil predicate on the "lastConnection" field.
func LastConnectionNotNil() predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldLastConnection)))
	})
}

// VolunteerEQ applies the EQ predicate on the "volunteer" field.
func VolunteerEQ(v bool) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldVolunteer), v))
	})
}

// VolunteerNEQ applies the NEQ predicate on the "volunteer" field.
func VolunteerNEQ(v bool) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldVolunteer), v))
	})
}

// HasNotes applies the HasEdge predicate on the "notes" edge.
func HasNotes() predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(NotesTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, NotesTable, NotesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasNotesWith applies the HasEdge predicate on the "notes" edge with a given conditions (other predicates).
func HasNotesWith(preds ...predicate.MedicalNote) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
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

// HasResponses applies the HasEdge predicate on the "responses" edge.
func HasResponses() predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ResponsesTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ResponsesTable, ResponsesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasResponsesWith applies the HasEdge predicate on the "responses" edge with a given conditions (other predicates).
func HasResponsesWith(preds ...predicate.TaskResponse) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ResponsesInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ResponsesTable, ResponsesColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasTasks applies the HasEdge predicate on the "tasks" edge.
func HasTasks() predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TasksTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, TasksTable, TasksPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTasksWith applies the HasEdge predicate on the "tasks" edge with a given conditions (other predicates).
func HasTasksWith(preds ...predicate.Task) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TasksInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, TasksTable, TasksPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Doctor) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Doctor) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
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
func Not(p predicate.Doctor) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		p(s.Not())
	})
}
