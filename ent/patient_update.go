// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/google/uuid"
	"github.com/minskylab/asclepius/ent/history"
	"github.com/minskylab/asclepius/ent/patient"
	"github.com/minskylab/asclepius/ent/predicate"
	"github.com/minskylab/asclepius/ent/schedule"
)

// PatientUpdate is the builder for updating Patient entities.
type PatientUpdate struct {
	config
	name            *string
	phone           *string
	age             *int
	addage          *int
	clearage        bool
	email           *string
	clearemail      bool
	password        *string
	clearpassword   bool
	facebookID      *string
	clearfacebookID bool
	watsonID        *string
	clearwatsonID   bool
	first_contact   *time.Time
	conditions      *[]string
	clearconditions bool
	history         map[uuid.UUID]struct{}
	schedule        map[uuid.UUID]struct{}
	clearedHistory  bool
	clearedSchedule bool
	predicates      []predicate.Patient
}

// Where adds a new predicate for the builder.
func (pu *PatientUpdate) Where(ps ...predicate.Patient) *PatientUpdate {
	pu.predicates = append(pu.predicates, ps...)
	return pu
}

// SetName sets the name field.
func (pu *PatientUpdate) SetName(s string) *PatientUpdate {
	pu.name = &s
	return pu
}

// SetPhone sets the phone field.
func (pu *PatientUpdate) SetPhone(s string) *PatientUpdate {
	pu.phone = &s
	return pu
}

// SetAge sets the age field.
func (pu *PatientUpdate) SetAge(i int) *PatientUpdate {
	pu.age = &i
	pu.addage = nil
	return pu
}

// SetNillableAge sets the age field if the given value is not nil.
func (pu *PatientUpdate) SetNillableAge(i *int) *PatientUpdate {
	if i != nil {
		pu.SetAge(*i)
	}
	return pu
}

// AddAge adds i to age.
func (pu *PatientUpdate) AddAge(i int) *PatientUpdate {
	if pu.addage == nil {
		pu.addage = &i
	} else {
		*pu.addage += i
	}
	return pu
}

// ClearAge clears the value of age.
func (pu *PatientUpdate) ClearAge() *PatientUpdate {
	pu.age = nil
	pu.clearage = true
	return pu
}

// SetEmail sets the email field.
func (pu *PatientUpdate) SetEmail(s string) *PatientUpdate {
	pu.email = &s
	return pu
}

// SetNillableEmail sets the email field if the given value is not nil.
func (pu *PatientUpdate) SetNillableEmail(s *string) *PatientUpdate {
	if s != nil {
		pu.SetEmail(*s)
	}
	return pu
}

// ClearEmail clears the value of email.
func (pu *PatientUpdate) ClearEmail() *PatientUpdate {
	pu.email = nil
	pu.clearemail = true
	return pu
}

// SetPassword sets the password field.
func (pu *PatientUpdate) SetPassword(s string) *PatientUpdate {
	pu.password = &s
	return pu
}

// SetNillablePassword sets the password field if the given value is not nil.
func (pu *PatientUpdate) SetNillablePassword(s *string) *PatientUpdate {
	if s != nil {
		pu.SetPassword(*s)
	}
	return pu
}

// ClearPassword clears the value of password.
func (pu *PatientUpdate) ClearPassword() *PatientUpdate {
	pu.password = nil
	pu.clearpassword = true
	return pu
}

// SetFacebookID sets the facebookID field.
func (pu *PatientUpdate) SetFacebookID(s string) *PatientUpdate {
	pu.facebookID = &s
	return pu
}

// SetNillableFacebookID sets the facebookID field if the given value is not nil.
func (pu *PatientUpdate) SetNillableFacebookID(s *string) *PatientUpdate {
	if s != nil {
		pu.SetFacebookID(*s)
	}
	return pu
}

// ClearFacebookID clears the value of facebookID.
func (pu *PatientUpdate) ClearFacebookID() *PatientUpdate {
	pu.facebookID = nil
	pu.clearfacebookID = true
	return pu
}

// SetWatsonID sets the watsonID field.
func (pu *PatientUpdate) SetWatsonID(s string) *PatientUpdate {
	pu.watsonID = &s
	return pu
}

// SetNillableWatsonID sets the watsonID field if the given value is not nil.
func (pu *PatientUpdate) SetNillableWatsonID(s *string) *PatientUpdate {
	if s != nil {
		pu.SetWatsonID(*s)
	}
	return pu
}

// ClearWatsonID clears the value of watsonID.
func (pu *PatientUpdate) ClearWatsonID() *PatientUpdate {
	pu.watsonID = nil
	pu.clearwatsonID = true
	return pu
}

// SetFirstContact sets the first_contact field.
func (pu *PatientUpdate) SetFirstContact(t time.Time) *PatientUpdate {
	pu.first_contact = &t
	return pu
}

// SetNillableFirstContact sets the first_contact field if the given value is not nil.
func (pu *PatientUpdate) SetNillableFirstContact(t *time.Time) *PatientUpdate {
	if t != nil {
		pu.SetFirstContact(*t)
	}
	return pu
}

// SetConditions sets the conditions field.
func (pu *PatientUpdate) SetConditions(s []string) *PatientUpdate {
	pu.conditions = &s
	return pu
}

// ClearConditions clears the value of conditions.
func (pu *PatientUpdate) ClearConditions() *PatientUpdate {
	pu.conditions = nil
	pu.clearconditions = true
	return pu
}

// SetHistoryID sets the history edge to History by id.
func (pu *PatientUpdate) SetHistoryID(id uuid.UUID) *PatientUpdate {
	if pu.history == nil {
		pu.history = make(map[uuid.UUID]struct{})
	}
	pu.history[id] = struct{}{}
	return pu
}

// SetHistory sets the history edge to History.
func (pu *PatientUpdate) SetHistory(h *History) *PatientUpdate {
	return pu.SetHistoryID(h.ID)
}

// SetScheduleID sets the schedule edge to Schedule by id.
func (pu *PatientUpdate) SetScheduleID(id uuid.UUID) *PatientUpdate {
	if pu.schedule == nil {
		pu.schedule = make(map[uuid.UUID]struct{})
	}
	pu.schedule[id] = struct{}{}
	return pu
}

// SetNillableScheduleID sets the schedule edge to Schedule by id if the given value is not nil.
func (pu *PatientUpdate) SetNillableScheduleID(id *uuid.UUID) *PatientUpdate {
	if id != nil {
		pu = pu.SetScheduleID(*id)
	}
	return pu
}

// SetSchedule sets the schedule edge to Schedule.
func (pu *PatientUpdate) SetSchedule(s *Schedule) *PatientUpdate {
	return pu.SetScheduleID(s.ID)
}

// ClearHistory clears the history edge to History.
func (pu *PatientUpdate) ClearHistory() *PatientUpdate {
	pu.clearedHistory = true
	return pu
}

// ClearSchedule clears the schedule edge to Schedule.
func (pu *PatientUpdate) ClearSchedule() *PatientUpdate {
	pu.clearedSchedule = true
	return pu
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (pu *PatientUpdate) Save(ctx context.Context) (int, error) {
	if pu.age != nil {
		if err := patient.AgeValidator(*pu.age); err != nil {
			return 0, fmt.Errorf("ent: validator failed for field \"age\": %v", err)
		}
	}
	if pu.password != nil {
		if err := patient.PasswordValidator(*pu.password); err != nil {
			return 0, fmt.Errorf("ent: validator failed for field \"password\": %v", err)
		}
	}
	if len(pu.history) > 1 {
		return 0, errors.New("ent: multiple assignments on a unique edge \"history\"")
	}
	if pu.clearedHistory && pu.history == nil {
		return 0, errors.New("ent: clearing a unique edge \"history\"")
	}
	if len(pu.schedule) > 1 {
		return 0, errors.New("ent: multiple assignments on a unique edge \"schedule\"")
	}
	return pu.sqlSave(ctx)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PatientUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PatientUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PatientUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pu *PatientUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   patient.Table,
			Columns: patient.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: patient.FieldID,
			},
		},
	}
	if ps := pu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value := pu.name; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: patient.FieldName,
		})
	}
	if value := pu.phone; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: patient.FieldPhone,
		})
	}
	if value := pu.age; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  *value,
			Column: patient.FieldAge,
		})
	}
	if value := pu.addage; value != nil {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  *value,
			Column: patient.FieldAge,
		})
	}
	if pu.clearage {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: patient.FieldAge,
		})
	}
	if value := pu.email; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: patient.FieldEmail,
		})
	}
	if pu.clearemail {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: patient.FieldEmail,
		})
	}
	if value := pu.password; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: patient.FieldPassword,
		})
	}
	if pu.clearpassword {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: patient.FieldPassword,
		})
	}
	if value := pu.facebookID; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: patient.FieldFacebookID,
		})
	}
	if pu.clearfacebookID {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: patient.FieldFacebookID,
		})
	}
	if value := pu.watsonID; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: patient.FieldWatsonID,
		})
	}
	if pu.clearwatsonID {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: patient.FieldWatsonID,
		})
	}
	if value := pu.first_contact; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  *value,
			Column: patient.FieldFirstContact,
		})
	}
	if value := pu.conditions; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  *value,
			Column: patient.FieldConditions,
		})
	}
	if pu.clearconditions {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: patient.FieldConditions,
		})
	}
	if pu.clearedHistory {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   patient.HistoryTable,
			Columns: []string{patient.HistoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: history.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.history; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   patient.HistoryTable,
			Columns: []string{patient.HistoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: history.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.clearedSchedule {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   patient.ScheduleTable,
			Columns: []string{patient.ScheduleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: schedule.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.schedule; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   patient.ScheduleTable,
			Columns: []string{patient.ScheduleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: schedule.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{patient.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// PatientUpdateOne is the builder for updating a single Patient entity.
type PatientUpdateOne struct {
	config
	id              uuid.UUID
	name            *string
	phone           *string
	age             *int
	addage          *int
	clearage        bool
	email           *string
	clearemail      bool
	password        *string
	clearpassword   bool
	facebookID      *string
	clearfacebookID bool
	watsonID        *string
	clearwatsonID   bool
	first_contact   *time.Time
	conditions      *[]string
	clearconditions bool
	history         map[uuid.UUID]struct{}
	schedule        map[uuid.UUID]struct{}
	clearedHistory  bool
	clearedSchedule bool
}

// SetName sets the name field.
func (puo *PatientUpdateOne) SetName(s string) *PatientUpdateOne {
	puo.name = &s
	return puo
}

// SetPhone sets the phone field.
func (puo *PatientUpdateOne) SetPhone(s string) *PatientUpdateOne {
	puo.phone = &s
	return puo
}

// SetAge sets the age field.
func (puo *PatientUpdateOne) SetAge(i int) *PatientUpdateOne {
	puo.age = &i
	puo.addage = nil
	return puo
}

// SetNillableAge sets the age field if the given value is not nil.
func (puo *PatientUpdateOne) SetNillableAge(i *int) *PatientUpdateOne {
	if i != nil {
		puo.SetAge(*i)
	}
	return puo
}

// AddAge adds i to age.
func (puo *PatientUpdateOne) AddAge(i int) *PatientUpdateOne {
	if puo.addage == nil {
		puo.addage = &i
	} else {
		*puo.addage += i
	}
	return puo
}

// ClearAge clears the value of age.
func (puo *PatientUpdateOne) ClearAge() *PatientUpdateOne {
	puo.age = nil
	puo.clearage = true
	return puo
}

// SetEmail sets the email field.
func (puo *PatientUpdateOne) SetEmail(s string) *PatientUpdateOne {
	puo.email = &s
	return puo
}

// SetNillableEmail sets the email field if the given value is not nil.
func (puo *PatientUpdateOne) SetNillableEmail(s *string) *PatientUpdateOne {
	if s != nil {
		puo.SetEmail(*s)
	}
	return puo
}

// ClearEmail clears the value of email.
func (puo *PatientUpdateOne) ClearEmail() *PatientUpdateOne {
	puo.email = nil
	puo.clearemail = true
	return puo
}

// SetPassword sets the password field.
func (puo *PatientUpdateOne) SetPassword(s string) *PatientUpdateOne {
	puo.password = &s
	return puo
}

// SetNillablePassword sets the password field if the given value is not nil.
func (puo *PatientUpdateOne) SetNillablePassword(s *string) *PatientUpdateOne {
	if s != nil {
		puo.SetPassword(*s)
	}
	return puo
}

// ClearPassword clears the value of password.
func (puo *PatientUpdateOne) ClearPassword() *PatientUpdateOne {
	puo.password = nil
	puo.clearpassword = true
	return puo
}

// SetFacebookID sets the facebookID field.
func (puo *PatientUpdateOne) SetFacebookID(s string) *PatientUpdateOne {
	puo.facebookID = &s
	return puo
}

// SetNillableFacebookID sets the facebookID field if the given value is not nil.
func (puo *PatientUpdateOne) SetNillableFacebookID(s *string) *PatientUpdateOne {
	if s != nil {
		puo.SetFacebookID(*s)
	}
	return puo
}

// ClearFacebookID clears the value of facebookID.
func (puo *PatientUpdateOne) ClearFacebookID() *PatientUpdateOne {
	puo.facebookID = nil
	puo.clearfacebookID = true
	return puo
}

// SetWatsonID sets the watsonID field.
func (puo *PatientUpdateOne) SetWatsonID(s string) *PatientUpdateOne {
	puo.watsonID = &s
	return puo
}

// SetNillableWatsonID sets the watsonID field if the given value is not nil.
func (puo *PatientUpdateOne) SetNillableWatsonID(s *string) *PatientUpdateOne {
	if s != nil {
		puo.SetWatsonID(*s)
	}
	return puo
}

// ClearWatsonID clears the value of watsonID.
func (puo *PatientUpdateOne) ClearWatsonID() *PatientUpdateOne {
	puo.watsonID = nil
	puo.clearwatsonID = true
	return puo
}

// SetFirstContact sets the first_contact field.
func (puo *PatientUpdateOne) SetFirstContact(t time.Time) *PatientUpdateOne {
	puo.first_contact = &t
	return puo
}

// SetNillableFirstContact sets the first_contact field if the given value is not nil.
func (puo *PatientUpdateOne) SetNillableFirstContact(t *time.Time) *PatientUpdateOne {
	if t != nil {
		puo.SetFirstContact(*t)
	}
	return puo
}

// SetConditions sets the conditions field.
func (puo *PatientUpdateOne) SetConditions(s []string) *PatientUpdateOne {
	puo.conditions = &s
	return puo
}

// ClearConditions clears the value of conditions.
func (puo *PatientUpdateOne) ClearConditions() *PatientUpdateOne {
	puo.conditions = nil
	puo.clearconditions = true
	return puo
}

// SetHistoryID sets the history edge to History by id.
func (puo *PatientUpdateOne) SetHistoryID(id uuid.UUID) *PatientUpdateOne {
	if puo.history == nil {
		puo.history = make(map[uuid.UUID]struct{})
	}
	puo.history[id] = struct{}{}
	return puo
}

// SetHistory sets the history edge to History.
func (puo *PatientUpdateOne) SetHistory(h *History) *PatientUpdateOne {
	return puo.SetHistoryID(h.ID)
}

// SetScheduleID sets the schedule edge to Schedule by id.
func (puo *PatientUpdateOne) SetScheduleID(id uuid.UUID) *PatientUpdateOne {
	if puo.schedule == nil {
		puo.schedule = make(map[uuid.UUID]struct{})
	}
	puo.schedule[id] = struct{}{}
	return puo
}

// SetNillableScheduleID sets the schedule edge to Schedule by id if the given value is not nil.
func (puo *PatientUpdateOne) SetNillableScheduleID(id *uuid.UUID) *PatientUpdateOne {
	if id != nil {
		puo = puo.SetScheduleID(*id)
	}
	return puo
}

// SetSchedule sets the schedule edge to Schedule.
func (puo *PatientUpdateOne) SetSchedule(s *Schedule) *PatientUpdateOne {
	return puo.SetScheduleID(s.ID)
}

// ClearHistory clears the history edge to History.
func (puo *PatientUpdateOne) ClearHistory() *PatientUpdateOne {
	puo.clearedHistory = true
	return puo
}

// ClearSchedule clears the schedule edge to Schedule.
func (puo *PatientUpdateOne) ClearSchedule() *PatientUpdateOne {
	puo.clearedSchedule = true
	return puo
}

// Save executes the query and returns the updated entity.
func (puo *PatientUpdateOne) Save(ctx context.Context) (*Patient, error) {
	if puo.age != nil {
		if err := patient.AgeValidator(*puo.age); err != nil {
			return nil, fmt.Errorf("ent: validator failed for field \"age\": %v", err)
		}
	}
	if puo.password != nil {
		if err := patient.PasswordValidator(*puo.password); err != nil {
			return nil, fmt.Errorf("ent: validator failed for field \"password\": %v", err)
		}
	}
	if len(puo.history) > 1 {
		return nil, errors.New("ent: multiple assignments on a unique edge \"history\"")
	}
	if puo.clearedHistory && puo.history == nil {
		return nil, errors.New("ent: clearing a unique edge \"history\"")
	}
	if len(puo.schedule) > 1 {
		return nil, errors.New("ent: multiple assignments on a unique edge \"schedule\"")
	}
	return puo.sqlSave(ctx)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PatientUpdateOne) SaveX(ctx context.Context) *Patient {
	pa, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return pa
}

// Exec executes the query on the entity.
func (puo *PatientUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PatientUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (puo *PatientUpdateOne) sqlSave(ctx context.Context) (pa *Patient, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   patient.Table,
			Columns: patient.Columns,
			ID: &sqlgraph.FieldSpec{
				Value:  puo.id,
				Type:   field.TypeUUID,
				Column: patient.FieldID,
			},
		},
	}
	if value := puo.name; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: patient.FieldName,
		})
	}
	if value := puo.phone; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: patient.FieldPhone,
		})
	}
	if value := puo.age; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  *value,
			Column: patient.FieldAge,
		})
	}
	if value := puo.addage; value != nil {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  *value,
			Column: patient.FieldAge,
		})
	}
	if puo.clearage {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: patient.FieldAge,
		})
	}
	if value := puo.email; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: patient.FieldEmail,
		})
	}
	if puo.clearemail {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: patient.FieldEmail,
		})
	}
	if value := puo.password; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: patient.FieldPassword,
		})
	}
	if puo.clearpassword {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: patient.FieldPassword,
		})
	}
	if value := puo.facebookID; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: patient.FieldFacebookID,
		})
	}
	if puo.clearfacebookID {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: patient.FieldFacebookID,
		})
	}
	if value := puo.watsonID; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: patient.FieldWatsonID,
		})
	}
	if puo.clearwatsonID {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: patient.FieldWatsonID,
		})
	}
	if value := puo.first_contact; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  *value,
			Column: patient.FieldFirstContact,
		})
	}
	if value := puo.conditions; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  *value,
			Column: patient.FieldConditions,
		})
	}
	if puo.clearconditions {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: patient.FieldConditions,
		})
	}
	if puo.clearedHistory {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   patient.HistoryTable,
			Columns: []string{patient.HistoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: history.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.history; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   patient.HistoryTable,
			Columns: []string{patient.HistoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: history.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.clearedSchedule {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   patient.ScheduleTable,
			Columns: []string{patient.ScheduleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: schedule.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.schedule; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   patient.ScheduleTable,
			Columns: []string{patient.ScheduleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: schedule.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	pa = &Patient{config: puo.config}
	_spec.Assign = pa.assignValues
	_spec.ScanValues = pa.scanValues()
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{patient.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return pa, nil
}
