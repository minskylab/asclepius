// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"time"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/google/uuid"
	"github.com/minskylab/asclepius/ent/doctor"
	"github.com/minskylab/asclepius/ent/schedule"
	"github.com/minskylab/asclepius/ent/task"
)

// TaskCreate is the builder for creating a Task entity.
type TaskCreate struct {
	config
	id          *uuid.UUID
	title       *[]string
	startAt     *time.Time
	endsAt      *time.Time
	description *[]string
	responsible map[uuid.UUID]struct{}
	schedule    map[uuid.UUID]struct{}
}

// SetTitle sets the title field.
func (tc *TaskCreate) SetTitle(s []string) *TaskCreate {
	tc.title = &s
	return tc
}

// SetStartAt sets the startAt field.
func (tc *TaskCreate) SetStartAt(t time.Time) *TaskCreate {
	tc.startAt = &t
	return tc
}

// SetNillableStartAt sets the startAt field if the given value is not nil.
func (tc *TaskCreate) SetNillableStartAt(t *time.Time) *TaskCreate {
	if t != nil {
		tc.SetStartAt(*t)
	}
	return tc
}

// SetEndsAt sets the endsAt field.
func (tc *TaskCreate) SetEndsAt(t time.Time) *TaskCreate {
	tc.endsAt = &t
	return tc
}

// SetNillableEndsAt sets the endsAt field if the given value is not nil.
func (tc *TaskCreate) SetNillableEndsAt(t *time.Time) *TaskCreate {
	if t != nil {
		tc.SetEndsAt(*t)
	}
	return tc
}

// SetDescription sets the description field.
func (tc *TaskCreate) SetDescription(s []string) *TaskCreate {
	tc.description = &s
	return tc
}

// SetID sets the id field.
func (tc *TaskCreate) SetID(u uuid.UUID) *TaskCreate {
	tc.id = &u
	return tc
}

// AddResponsibleIDs adds the responsible edge to Doctor by ids.
func (tc *TaskCreate) AddResponsibleIDs(ids ...uuid.UUID) *TaskCreate {
	if tc.responsible == nil {
		tc.responsible = make(map[uuid.UUID]struct{})
	}
	for i := range ids {
		tc.responsible[ids[i]] = struct{}{}
	}
	return tc
}

// AddResponsible adds the responsible edges to Doctor.
func (tc *TaskCreate) AddResponsible(d ...*Doctor) *TaskCreate {
	ids := make([]uuid.UUID, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return tc.AddResponsibleIDs(ids...)
}

// SetScheduleID sets the schedule edge to Schedule by id.
func (tc *TaskCreate) SetScheduleID(id uuid.UUID) *TaskCreate {
	if tc.schedule == nil {
		tc.schedule = make(map[uuid.UUID]struct{})
	}
	tc.schedule[id] = struct{}{}
	return tc
}

// SetNillableScheduleID sets the schedule edge to Schedule by id if the given value is not nil.
func (tc *TaskCreate) SetNillableScheduleID(id *uuid.UUID) *TaskCreate {
	if id != nil {
		tc = tc.SetScheduleID(*id)
	}
	return tc
}

// SetSchedule sets the schedule edge to Schedule.
func (tc *TaskCreate) SetSchedule(s *Schedule) *TaskCreate {
	return tc.SetScheduleID(s.ID)
}

// Save creates the Task in the database.
func (tc *TaskCreate) Save(ctx context.Context) (*Task, error) {
	if tc.title == nil {
		return nil, errors.New("ent: missing required field \"title\"")
	}
	if tc.startAt == nil {
		v := task.DefaultStartAt()
		tc.startAt = &v
	}
	if tc.endsAt == nil {
		v := task.DefaultEndsAt()
		tc.endsAt = &v
	}
	if len(tc.schedule) > 1 {
		return nil, errors.New("ent: multiple assignments on a unique edge \"schedule\"")
	}
	return tc.sqlSave(ctx)
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TaskCreate) SaveX(ctx context.Context) *Task {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (tc *TaskCreate) sqlSave(ctx context.Context) (*Task, error) {
	var (
		t     = &Task{config: tc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: task.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: task.FieldID,
			},
		}
	)
	if value := tc.id; value != nil {
		t.ID = *value
		_spec.ID.Value = *value
	}
	if value := tc.title; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  *value,
			Column: task.FieldTitle,
		})
		t.Title = *value
	}
	if value := tc.startAt; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  *value,
			Column: task.FieldStartAt,
		})
		t.StartAt = *value
	}
	if value := tc.endsAt; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  *value,
			Column: task.FieldEndsAt,
		})
		t.EndsAt = *value
	}
	if value := tc.description; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  *value,
			Column: task.FieldDescription,
		})
		t.Description = *value
	}
	if nodes := tc.responsible; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   task.ResponsibleTable,
			Columns: task.ResponsiblePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: doctor.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.schedule; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   task.ScheduleTable,
			Columns: []string{task.ScheduleColumn},
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return t, nil
}
