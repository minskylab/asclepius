// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/google/uuid"
	"github.com/minskylab/asclepius/ent/medicalnote"
	"github.com/minskylab/asclepius/ent/medicus"
)

// MedicusCreate is the builder for creating a Medicus entity.
type MedicusCreate struct {
	config
	id             *uuid.UUID
	name           *[]string
	email          *string
	phone          *string
	state          *medicus.State
	lastConnection *time.Time
	volunteer      *bool
	notes          map[uuid.UUID]struct{}
}

// SetName sets the name field.
func (mc *MedicusCreate) SetName(s []string) *MedicusCreate {
	mc.name = &s
	return mc
}

// SetEmail sets the email field.
func (mc *MedicusCreate) SetEmail(s string) *MedicusCreate {
	mc.email = &s
	return mc
}

// SetPhone sets the phone field.
func (mc *MedicusCreate) SetPhone(s string) *MedicusCreate {
	mc.phone = &s
	return mc
}

// SetState sets the state field.
func (mc *MedicusCreate) SetState(m medicus.State) *MedicusCreate {
	mc.state = &m
	return mc
}

// SetLastConnection sets the lastConnection field.
func (mc *MedicusCreate) SetLastConnection(t time.Time) *MedicusCreate {
	mc.lastConnection = &t
	return mc
}

// SetNillableLastConnection sets the lastConnection field if the given value is not nil.
func (mc *MedicusCreate) SetNillableLastConnection(t *time.Time) *MedicusCreate {
	if t != nil {
		mc.SetLastConnection(*t)
	}
	return mc
}

// SetVolunteer sets the volunteer field.
func (mc *MedicusCreate) SetVolunteer(b bool) *MedicusCreate {
	mc.volunteer = &b
	return mc
}

// SetNillableVolunteer sets the volunteer field if the given value is not nil.
func (mc *MedicusCreate) SetNillableVolunteer(b *bool) *MedicusCreate {
	if b != nil {
		mc.SetVolunteer(*b)
	}
	return mc
}

// SetID sets the id field.
func (mc *MedicusCreate) SetID(u uuid.UUID) *MedicusCreate {
	mc.id = &u
	return mc
}

// AddNoteIDs adds the notes edge to MedicalNote by ids.
func (mc *MedicusCreate) AddNoteIDs(ids ...uuid.UUID) *MedicusCreate {
	if mc.notes == nil {
		mc.notes = make(map[uuid.UUID]struct{})
	}
	for i := range ids {
		mc.notes[ids[i]] = struct{}{}
	}
	return mc
}

// AddNotes adds the notes edges to MedicalNote.
func (mc *MedicusCreate) AddNotes(m ...*MedicalNote) *MedicusCreate {
	ids := make([]uuid.UUID, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mc.AddNoteIDs(ids...)
}

// Save creates the Medicus in the database.
func (mc *MedicusCreate) Save(ctx context.Context) (*Medicus, error) {
	if mc.name == nil {
		return nil, errors.New("ent: missing required field \"name\"")
	}
	if mc.email == nil {
		return nil, errors.New("ent: missing required field \"email\"")
	}
	if mc.phone == nil {
		return nil, errors.New("ent: missing required field \"phone\"")
	}
	if mc.state == nil {
		return nil, errors.New("ent: missing required field \"state\"")
	}
	if err := medicus.StateValidator(*mc.state); err != nil {
		return nil, fmt.Errorf("ent: validator failed for field \"state\": %v", err)
	}
	if mc.lastConnection == nil {
		v := medicus.DefaultLastConnection()
		mc.lastConnection = &v
	}
	if mc.volunteer == nil {
		v := medicus.DefaultVolunteer
		mc.volunteer = &v
	}
	return mc.sqlSave(ctx)
}

// SaveX calls Save and panics if Save returns an error.
func (mc *MedicusCreate) SaveX(ctx context.Context) *Medicus {
	v, err := mc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (mc *MedicusCreate) sqlSave(ctx context.Context) (*Medicus, error) {
	var (
		m     = &Medicus{config: mc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: medicus.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: medicus.FieldID,
			},
		}
	)
	if value := mc.id; value != nil {
		m.ID = *value
		_spec.ID.Value = *value
	}
	if value := mc.name; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  *value,
			Column: medicus.FieldName,
		})
		m.Name = *value
	}
	if value := mc.email; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: medicus.FieldEmail,
		})
		m.Email = *value
	}
	if value := mc.phone; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: medicus.FieldPhone,
		})
		m.Phone = *value
	}
	if value := mc.state; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  *value,
			Column: medicus.FieldState,
		})
		m.State = *value
	}
	if value := mc.lastConnection; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  *value,
			Column: medicus.FieldLastConnection,
		})
		m.LastConnection = *value
	}
	if value := mc.volunteer; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  *value,
			Column: medicus.FieldVolunteer,
		})
		m.Volunteer = *value
	}
	if nodes := mc.notes; len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if err := sqlgraph.CreateNode(ctx, mc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return m, nil
}