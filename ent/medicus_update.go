// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/google/uuid"
	"github.com/minskylab/asclepius/ent/medicalnote"
	"github.com/minskylab/asclepius/ent/medicus"
	"github.com/minskylab/asclepius/ent/predicate"
)

// MedicusUpdate is the builder for updating Medicus entities.
type MedicusUpdate struct {
	config
	name                *[]string
	email               *string
	phone               *string
	state               *medicus.State
	lastConnection      *time.Time
	clearlastConnection bool
	volunteer           *bool
	notes               map[uuid.UUID]struct{}
	removedNotes        map[uuid.UUID]struct{}
	predicates          []predicate.Medicus
}

// Where adds a new predicate for the builder.
func (mu *MedicusUpdate) Where(ps ...predicate.Medicus) *MedicusUpdate {
	mu.predicates = append(mu.predicates, ps...)
	return mu
}

// SetName sets the name field.
func (mu *MedicusUpdate) SetName(s []string) *MedicusUpdate {
	mu.name = &s
	return mu
}

// SetEmail sets the email field.
func (mu *MedicusUpdate) SetEmail(s string) *MedicusUpdate {
	mu.email = &s
	return mu
}

// SetPhone sets the phone field.
func (mu *MedicusUpdate) SetPhone(s string) *MedicusUpdate {
	mu.phone = &s
	return mu
}

// SetState sets the state field.
func (mu *MedicusUpdate) SetState(m medicus.State) *MedicusUpdate {
	mu.state = &m
	return mu
}

// SetLastConnection sets the lastConnection field.
func (mu *MedicusUpdate) SetLastConnection(t time.Time) *MedicusUpdate {
	mu.lastConnection = &t
	return mu
}

// SetNillableLastConnection sets the lastConnection field if the given value is not nil.
func (mu *MedicusUpdate) SetNillableLastConnection(t *time.Time) *MedicusUpdate {
	if t != nil {
		mu.SetLastConnection(*t)
	}
	return mu
}

// ClearLastConnection clears the value of lastConnection.
func (mu *MedicusUpdate) ClearLastConnection() *MedicusUpdate {
	mu.lastConnection = nil
	mu.clearlastConnection = true
	return mu
}

// SetVolunteer sets the volunteer field.
func (mu *MedicusUpdate) SetVolunteer(b bool) *MedicusUpdate {
	mu.volunteer = &b
	return mu
}

// SetNillableVolunteer sets the volunteer field if the given value is not nil.
func (mu *MedicusUpdate) SetNillableVolunteer(b *bool) *MedicusUpdate {
	if b != nil {
		mu.SetVolunteer(*b)
	}
	return mu
}

// AddNoteIDs adds the notes edge to MedicalNote by ids.
func (mu *MedicusUpdate) AddNoteIDs(ids ...uuid.UUID) *MedicusUpdate {
	if mu.notes == nil {
		mu.notes = make(map[uuid.UUID]struct{})
	}
	for i := range ids {
		mu.notes[ids[i]] = struct{}{}
	}
	return mu
}

// AddNotes adds the notes edges to MedicalNote.
func (mu *MedicusUpdate) AddNotes(m ...*MedicalNote) *MedicusUpdate {
	ids := make([]uuid.UUID, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mu.AddNoteIDs(ids...)
}

// RemoveNoteIDs removes the notes edge to MedicalNote by ids.
func (mu *MedicusUpdate) RemoveNoteIDs(ids ...uuid.UUID) *MedicusUpdate {
	if mu.removedNotes == nil {
		mu.removedNotes = make(map[uuid.UUID]struct{})
	}
	for i := range ids {
		mu.removedNotes[ids[i]] = struct{}{}
	}
	return mu
}

// RemoveNotes removes notes edges to MedicalNote.
func (mu *MedicusUpdate) RemoveNotes(m ...*MedicalNote) *MedicusUpdate {
	ids := make([]uuid.UUID, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mu.RemoveNoteIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (mu *MedicusUpdate) Save(ctx context.Context) (int, error) {
	if mu.state != nil {
		if err := medicus.StateValidator(*mu.state); err != nil {
			return 0, fmt.Errorf("ent: validator failed for field \"state\": %v", err)
		}
	}
	return mu.sqlSave(ctx)
}

// SaveX is like Save, but panics if an error occurs.
func (mu *MedicusUpdate) SaveX(ctx context.Context) int {
	affected, err := mu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mu *MedicusUpdate) Exec(ctx context.Context) error {
	_, err := mu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mu *MedicusUpdate) ExecX(ctx context.Context) {
	if err := mu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (mu *MedicusUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   medicus.Table,
			Columns: medicus.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: medicus.FieldID,
			},
		},
	}
	if ps := mu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value := mu.name; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  *value,
			Column: medicus.FieldName,
		})
	}
	if value := mu.email; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: medicus.FieldEmail,
		})
	}
	if value := mu.phone; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: medicus.FieldPhone,
		})
	}
	if value := mu.state; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  *value,
			Column: medicus.FieldState,
		})
	}
	if value := mu.lastConnection; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  *value,
			Column: medicus.FieldLastConnection,
		})
	}
	if mu.clearlastConnection {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: medicus.FieldLastConnection,
		})
	}
	if value := mu.volunteer; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  *value,
			Column: medicus.FieldVolunteer,
		})
	}
	if nodes := mu.removedNotes; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   medicus.NotesTable,
			Columns: []string{medicus.NotesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: medicalnote.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.notes; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   medicus.NotesTable,
			Columns: []string{medicus.NotesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: medicalnote.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, mu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{medicus.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// MedicusUpdateOne is the builder for updating a single Medicus entity.
type MedicusUpdateOne struct {
	config
	id                  uuid.UUID
	name                *[]string
	email               *string
	phone               *string
	state               *medicus.State
	lastConnection      *time.Time
	clearlastConnection bool
	volunteer           *bool
	notes               map[uuid.UUID]struct{}
	removedNotes        map[uuid.UUID]struct{}
}

// SetName sets the name field.
func (muo *MedicusUpdateOne) SetName(s []string) *MedicusUpdateOne {
	muo.name = &s
	return muo
}

// SetEmail sets the email field.
func (muo *MedicusUpdateOne) SetEmail(s string) *MedicusUpdateOne {
	muo.email = &s
	return muo
}

// SetPhone sets the phone field.
func (muo *MedicusUpdateOne) SetPhone(s string) *MedicusUpdateOne {
	muo.phone = &s
	return muo
}

// SetState sets the state field.
func (muo *MedicusUpdateOne) SetState(m medicus.State) *MedicusUpdateOne {
	muo.state = &m
	return muo
}

// SetLastConnection sets the lastConnection field.
func (muo *MedicusUpdateOne) SetLastConnection(t time.Time) *MedicusUpdateOne {
	muo.lastConnection = &t
	return muo
}

// SetNillableLastConnection sets the lastConnection field if the given value is not nil.
func (muo *MedicusUpdateOne) SetNillableLastConnection(t *time.Time) *MedicusUpdateOne {
	if t != nil {
		muo.SetLastConnection(*t)
	}
	return muo
}

// ClearLastConnection clears the value of lastConnection.
func (muo *MedicusUpdateOne) ClearLastConnection() *MedicusUpdateOne {
	muo.lastConnection = nil
	muo.clearlastConnection = true
	return muo
}

// SetVolunteer sets the volunteer field.
func (muo *MedicusUpdateOne) SetVolunteer(b bool) *MedicusUpdateOne {
	muo.volunteer = &b
	return muo
}

// SetNillableVolunteer sets the volunteer field if the given value is not nil.
func (muo *MedicusUpdateOne) SetNillableVolunteer(b *bool) *MedicusUpdateOne {
	if b != nil {
		muo.SetVolunteer(*b)
	}
	return muo
}

// AddNoteIDs adds the notes edge to MedicalNote by ids.
func (muo *MedicusUpdateOne) AddNoteIDs(ids ...uuid.UUID) *MedicusUpdateOne {
	if muo.notes == nil {
		muo.notes = make(map[uuid.UUID]struct{})
	}
	for i := range ids {
		muo.notes[ids[i]] = struct{}{}
	}
	return muo
}

// AddNotes adds the notes edges to MedicalNote.
func (muo *MedicusUpdateOne) AddNotes(m ...*MedicalNote) *MedicusUpdateOne {
	ids := make([]uuid.UUID, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return muo.AddNoteIDs(ids...)
}

// RemoveNoteIDs removes the notes edge to MedicalNote by ids.
func (muo *MedicusUpdateOne) RemoveNoteIDs(ids ...uuid.UUID) *MedicusUpdateOne {
	if muo.removedNotes == nil {
		muo.removedNotes = make(map[uuid.UUID]struct{})
	}
	for i := range ids {
		muo.removedNotes[ids[i]] = struct{}{}
	}
	return muo
}

// RemoveNotes removes notes edges to MedicalNote.
func (muo *MedicusUpdateOne) RemoveNotes(m ...*MedicalNote) *MedicusUpdateOne {
	ids := make([]uuid.UUID, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return muo.RemoveNoteIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (muo *MedicusUpdateOne) Save(ctx context.Context) (*Medicus, error) {
	if muo.state != nil {
		if err := medicus.StateValidator(*muo.state); err != nil {
			return nil, fmt.Errorf("ent: validator failed for field \"state\": %v", err)
		}
	}
	return muo.sqlSave(ctx)
}

// SaveX is like Save, but panics if an error occurs.
func (muo *MedicusUpdateOne) SaveX(ctx context.Context) *Medicus {
	m, err := muo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return m
}

// Exec executes the query on the entity.
func (muo *MedicusUpdateOne) Exec(ctx context.Context) error {
	_, err := muo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (muo *MedicusUpdateOne) ExecX(ctx context.Context) {
	if err := muo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (muo *MedicusUpdateOne) sqlSave(ctx context.Context) (m *Medicus, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   medicus.Table,
			Columns: medicus.Columns,
			ID: &sqlgraph.FieldSpec{
				Value:  muo.id,
				Type:   field.TypeUUID,
				Column: medicus.FieldID,
			},
		},
	}
	if value := muo.name; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  *value,
			Column: medicus.FieldName,
		})
	}
	if value := muo.email; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: medicus.FieldEmail,
		})
	}
	if value := muo.phone; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: medicus.FieldPhone,
		})
	}
	if value := muo.state; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  *value,
			Column: medicus.FieldState,
		})
	}
	if value := muo.lastConnection; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  *value,
			Column: medicus.FieldLastConnection,
		})
	}
	if muo.clearlastConnection {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: medicus.FieldLastConnection,
		})
	}
	if value := muo.volunteer; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  *value,
			Column: medicus.FieldVolunteer,
		})
	}
	if nodes := muo.removedNotes; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   medicus.NotesTable,
			Columns: []string{medicus.NotesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: medicalnote.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.notes; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   medicus.NotesTable,
			Columns: []string{medicus.NotesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: medicalnote.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	m = &Medicus{config: muo.config}
	_spec.Assign = m.assignValues
	_spec.ScanValues = m.scanValues()
	if err = sqlgraph.UpdateNode(ctx, muo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{medicus.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return m, nil
}