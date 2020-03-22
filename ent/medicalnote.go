// Code generated by entc, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/minskylab/asclepius/ent/doctor"
	"github.com/minskylab/asclepius/ent/history"
	"github.com/minskylab/asclepius/ent/medicalnote"
)

// MedicalNote is the model entity for the MedicalNote schema.
type MedicalNote struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// At holds the value of the "at" field.
	At time.Time `json:"at,omitempty"`
	// LastChange holds the value of the "lastChange" field.
	LastChange time.Time `json:"lastChange,omitempty"`
	// Observations holds the value of the "observations" field.
	Observations []string `json:"observations,omitempty"`
	// Meta holds the value of the "meta" field.
	Meta []string `json:"meta,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the MedicalNoteQuery when eager-loading is set.
	Edges         MedicalNoteEdges `json:"edges"`
	doctor_notes  *uuid.UUID
	history_notes *uuid.UUID
}

// MedicalNoteEdges holds the relations/edges for other nodes in the graph.
type MedicalNoteEdges struct {
	// History holds the value of the history edge.
	History *History
	// Owner holds the value of the owner edge.
	Owner *Doctor
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// HistoryOrErr returns the History value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MedicalNoteEdges) HistoryOrErr() (*History, error) {
	if e.loadedTypes[0] {
		if e.History == nil {
			// The edge history was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: history.Label}
		}
		return e.History, nil
	}
	return nil, &NotLoadedError{edge: "history"}
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MedicalNoteEdges) OwnerOrErr() (*Doctor, error) {
	if e.loadedTypes[1] {
		if e.Owner == nil {
			// The edge owner was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: doctor.Label}
		}
		return e.Owner, nil
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*MedicalNote) scanValues() []interface{} {
	return []interface{}{
		&uuid.UUID{},    // id
		&sql.NullTime{}, // at
		&sql.NullTime{}, // lastChange
		&[]byte{},       // observations
		&[]byte{},       // meta
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*MedicalNote) fkValues() []interface{} {
	return []interface{}{
		&uuid.UUID{}, // doctor_notes
		&uuid.UUID{}, // history_notes
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the MedicalNote fields.
func (mn *MedicalNote) assignValues(values ...interface{}) error {
	if m, n := len(values), len(medicalnote.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	if value, ok := values[0].(*uuid.UUID); !ok {
		return fmt.Errorf("unexpected type %T for field id", values[0])
	} else if value != nil {
		mn.ID = *value
	}
	values = values[1:]
	if value, ok := values[0].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field at", values[0])
	} else if value.Valid {
		mn.At = value.Time
	}
	if value, ok := values[1].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field lastChange", values[1])
	} else if value.Valid {
		mn.LastChange = value.Time
	}

	if value, ok := values[2].(*[]byte); !ok {
		return fmt.Errorf("unexpected type %T for field observations", values[2])
	} else if value != nil && len(*value) > 0 {
		if err := json.Unmarshal(*value, &mn.Observations); err != nil {
			return fmt.Errorf("unmarshal field observations: %v", err)
		}
	}

	if value, ok := values[3].(*[]byte); !ok {
		return fmt.Errorf("unexpected type %T for field meta", values[3])
	} else if value != nil && len(*value) > 0 {
		if err := json.Unmarshal(*value, &mn.Meta); err != nil {
			return fmt.Errorf("unmarshal field meta: %v", err)
		}
	}
	values = values[4:]
	if len(values) == len(medicalnote.ForeignKeys) {
		if value, ok := values[0].(*uuid.UUID); !ok {
			return fmt.Errorf("unexpected type %T for field doctor_notes", values[0])
		} else if value != nil {
			mn.doctor_notes = value
		}
		if value, ok := values[0].(*uuid.UUID); !ok {
			return fmt.Errorf("unexpected type %T for field history_notes", values[0])
		} else if value != nil {
			mn.history_notes = value
		}
	}
	return nil
}

// QueryHistory queries the history edge of the MedicalNote.
func (mn *MedicalNote) QueryHistory() *HistoryQuery {
	return (&MedicalNoteClient{config: mn.config}).QueryHistory(mn)
}

// QueryOwner queries the owner edge of the MedicalNote.
func (mn *MedicalNote) QueryOwner() *DoctorQuery {
	return (&MedicalNoteClient{config: mn.config}).QueryOwner(mn)
}

// Update returns a builder for updating this MedicalNote.
// Note that, you need to call MedicalNote.Unwrap() before calling this method, if this MedicalNote
// was returned from a transaction, and the transaction was committed or rolled back.
func (mn *MedicalNote) Update() *MedicalNoteUpdateOne {
	return (&MedicalNoteClient{config: mn.config}).UpdateOne(mn)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (mn *MedicalNote) Unwrap() *MedicalNote {
	tx, ok := mn.config.driver.(*txDriver)
	if !ok {
		panic("ent: MedicalNote is not a transactional entity")
	}
	mn.config.driver = tx.drv
	return mn
}

// String implements the fmt.Stringer.
func (mn *MedicalNote) String() string {
	var builder strings.Builder
	builder.WriteString("MedicalNote(")
	builder.WriteString(fmt.Sprintf("id=%v", mn.ID))
	builder.WriteString(", at=")
	builder.WriteString(mn.At.Format(time.ANSIC))
	builder.WriteString(", lastChange=")
	builder.WriteString(mn.LastChange.Format(time.ANSIC))
	builder.WriteString(", observations=")
	builder.WriteString(fmt.Sprintf("%v", mn.Observations))
	builder.WriteString(", meta=")
	builder.WriteString(fmt.Sprintf("%v", mn.Meta))
	builder.WriteByte(')')
	return builder.String()
}

// MedicalNotes is a parsable slice of MedicalNote.
type MedicalNotes []*MedicalNote

func (mn MedicalNotes) config(cfg config) {
	for _i := range mn {
		mn[_i].config = cfg
	}
}
