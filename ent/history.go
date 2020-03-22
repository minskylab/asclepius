// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/minskylab/asclepius/ent/history"
	"github.com/minskylab/asclepius/ent/patient"
)

// History is the model entity for the History schema.
type History struct {
	config
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the HistoryQuery when eager-loading is set.
	Edges           HistoryEdges `json:"edges"`
	patient_history *uuid.UUID
}

// HistoryEdges holds the relations/edges for other nodes in the graph.
type HistoryEdges struct {
	// Patient holds the value of the patient edge.
	Patient *Patient
	// Tests holds the value of the tests edge.
	Tests []*Test
	// Notes holds the value of the notes edge.
	Notes []*MedicalNote
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// PatientOrErr returns the Patient value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e HistoryEdges) PatientOrErr() (*Patient, error) {
	if e.loadedTypes[0] {
		if e.Patient == nil {
			// The edge patient was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: patient.Label}
		}
		return e.Patient, nil
	}
	return nil, &NotLoadedError{edge: "patient"}
}

// TestsOrErr returns the Tests value or an error if the edge
// was not loaded in eager-loading.
func (e HistoryEdges) TestsOrErr() ([]*Test, error) {
	if e.loadedTypes[1] {
		return e.Tests, nil
	}
	return nil, &NotLoadedError{edge: "tests"}
}

// NotesOrErr returns the Notes value or an error if the edge
// was not loaded in eager-loading.
func (e HistoryEdges) NotesOrErr() ([]*MedicalNote, error) {
	if e.loadedTypes[2] {
		return e.Notes, nil
	}
	return nil, &NotLoadedError{edge: "notes"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*History) scanValues() []interface{} {
	return []interface{}{
		&uuid.UUID{}, // id
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*History) fkValues() []interface{} {
	return []interface{}{
		&uuid.UUID{}, // patient_history
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the History fields.
func (h *History) assignValues(values ...interface{}) error {
	if m, n := len(values), len(history.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	if value, ok := values[0].(*uuid.UUID); !ok {
		return fmt.Errorf("unexpected type %T for field id", values[0])
	} else if value != nil {
		h.ID = *value
	}
	values = values[1:]
	values = values[0:]
	if len(values) == len(history.ForeignKeys) {
		if value, ok := values[0].(*uuid.UUID); !ok {
			return fmt.Errorf("unexpected type %T for field patient_history", values[0])
		} else if value != nil {
			h.patient_history = value
		}
	}
	return nil
}

// QueryPatient queries the patient edge of the History.
func (h *History) QueryPatient() *PatientQuery {
	return (&HistoryClient{config: h.config}).QueryPatient(h)
}

// QueryTests queries the tests edge of the History.
func (h *History) QueryTests() *TestQuery {
	return (&HistoryClient{config: h.config}).QueryTests(h)
}

// QueryNotes queries the notes edge of the History.
func (h *History) QueryNotes() *MedicalNoteQuery {
	return (&HistoryClient{config: h.config}).QueryNotes(h)
}

// Update returns a builder for updating this History.
// Note that, you need to call History.Unwrap() before calling this method, if this History
// was returned from a transaction, and the transaction was committed or rolled back.
func (h *History) Update() *HistoryUpdateOne {
	return (&HistoryClient{config: h.config}).UpdateOne(h)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (h *History) Unwrap() *History {
	tx, ok := h.config.driver.(*txDriver)
	if !ok {
		panic("ent: History is not a transactional entity")
	}
	h.config.driver = tx.drv
	return h
}

// String implements the fmt.Stringer.
func (h *History) String() string {
	var builder strings.Builder
	builder.WriteString("History(")
	builder.WriteString(fmt.Sprintf("id=%v", h.ID))
	builder.WriteByte(')')
	return builder.String()
}

// Histories is a parsable slice of History.
type Histories []*History

func (h Histories) config(cfg config) {
	for _i := range h {
		h[_i].config = cfg
	}
}
