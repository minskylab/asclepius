// Code generated by entc, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/minskylab/asclepius/ent/schedule"
	"github.com/minskylab/asclepius/ent/task"
)

// Task is the model entity for the Task schema.
type Task struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// StartAt holds the value of the "startAt" field.
	StartAt time.Time `json:"startAt,omitempty"`
	// EndsAt holds the value of the "endsAt" field.
	EndsAt time.Time `json:"endsAt,omitempty"`
	// Description holds the value of the "description" field.
	Description []string `json:"description,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TaskQuery when eager-loading is set.
	Edges          TaskEdges `json:"edges"`
	schedule_tasks *uuid.UUID
}

// TaskEdges holds the relations/edges for other nodes in the graph.
type TaskEdges struct {
	// Responsible holds the value of the responsible edge.
	Responsible []*Doctor
	// Responses holds the value of the responses edge.
	Responses []*TaskResponse
	// Schedule holds the value of the schedule edge.
	Schedule *Schedule
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// ResponsibleOrErr returns the Responsible value or an error if the edge
// was not loaded in eager-loading.
func (e TaskEdges) ResponsibleOrErr() ([]*Doctor, error) {
	if e.loadedTypes[0] {
		return e.Responsible, nil
	}
	return nil, &NotLoadedError{edge: "responsible"}
}

// ResponsesOrErr returns the Responses value or an error if the edge
// was not loaded in eager-loading.
func (e TaskEdges) ResponsesOrErr() ([]*TaskResponse, error) {
	if e.loadedTypes[1] {
		return e.Responses, nil
	}
	return nil, &NotLoadedError{edge: "responses"}
}

// ScheduleOrErr returns the Schedule value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TaskEdges) ScheduleOrErr() (*Schedule, error) {
	if e.loadedTypes[2] {
		if e.Schedule == nil {
			// The edge schedule was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: schedule.Label}
		}
		return e.Schedule, nil
	}
	return nil, &NotLoadedError{edge: "schedule"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Task) scanValues() []interface{} {
	return []interface{}{
		&uuid.UUID{},      // id
		&sql.NullString{}, // title
		&sql.NullTime{},   // startAt
		&sql.NullTime{},   // endsAt
		&[]byte{},         // description
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*Task) fkValues() []interface{} {
	return []interface{}{
		&uuid.UUID{}, // schedule_tasks
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Task fields.
func (t *Task) assignValues(values ...interface{}) error {
	if m, n := len(values), len(task.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	if value, ok := values[0].(*uuid.UUID); !ok {
		return fmt.Errorf("unexpected type %T for field id", values[0])
	} else if value != nil {
		t.ID = *value
	}
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field title", values[0])
	} else if value.Valid {
		t.Title = value.String
	}
	if value, ok := values[1].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field startAt", values[1])
	} else if value.Valid {
		t.StartAt = value.Time
	}
	if value, ok := values[2].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field endsAt", values[2])
	} else if value.Valid {
		t.EndsAt = value.Time
	}

	if value, ok := values[3].(*[]byte); !ok {
		return fmt.Errorf("unexpected type %T for field description", values[3])
	} else if value != nil && len(*value) > 0 {
		if err := json.Unmarshal(*value, &t.Description); err != nil {
			return fmt.Errorf("unmarshal field description: %v", err)
		}
	}
	values = values[4:]
	if len(values) == len(task.ForeignKeys) {
		if value, ok := values[0].(*uuid.UUID); !ok {
			return fmt.Errorf("unexpected type %T for field schedule_tasks", values[0])
		} else if value != nil {
			t.schedule_tasks = value
		}
	}
	return nil
}

// QueryResponsible queries the responsible edge of the Task.
func (t *Task) QueryResponsible() *DoctorQuery {
	return (&TaskClient{config: t.config}).QueryResponsible(t)
}

// QueryResponses queries the responses edge of the Task.
func (t *Task) QueryResponses() *TaskResponseQuery {
	return (&TaskClient{config: t.config}).QueryResponses(t)
}

// QuerySchedule queries the schedule edge of the Task.
func (t *Task) QuerySchedule() *ScheduleQuery {
	return (&TaskClient{config: t.config}).QuerySchedule(t)
}

// Update returns a builder for updating this Task.
// Note that, you need to call Task.Unwrap() before calling this method, if this Task
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Task) Update() *TaskUpdateOne {
	return (&TaskClient{config: t.config}).UpdateOne(t)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (t *Task) Unwrap() *Task {
	tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Task is not a transactional entity")
	}
	t.config.driver = tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Task) String() string {
	var builder strings.Builder
	builder.WriteString("Task(")
	builder.WriteString(fmt.Sprintf("id=%v", t.ID))
	builder.WriteString(", title=")
	builder.WriteString(t.Title)
	builder.WriteString(", startAt=")
	builder.WriteString(t.StartAt.Format(time.ANSIC))
	builder.WriteString(", endsAt=")
	builder.WriteString(t.EndsAt.Format(time.ANSIC))
	builder.WriteString(", description=")
	builder.WriteString(fmt.Sprintf("%v", t.Description))
	builder.WriteByte(')')
	return builder.String()
}

// Tasks is a parsable slice of Task.
type Tasks []*Task

func (t Tasks) config(cfg config) {
	for _i := range t {
		t[_i].config = cfg
	}
}
