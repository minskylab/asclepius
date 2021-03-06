// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/google/uuid"
	"github.com/minskylab/asclepius/ent/clinicalresults"
	"github.com/minskylab/asclepius/ent/test"
)

// ClinicalResultsCreate is the builder for creating a ClinicalResults entity.
type ClinicalResultsCreate struct {
	config
	id                *uuid.UUID
	generalDiscomfort *bool
	fever             *bool
	thirdAge          *bool
	dyspnea           *bool
	test              map[uuid.UUID]struct{}
}

// SetGeneralDiscomfort sets the generalDiscomfort field.
func (crc *ClinicalResultsCreate) SetGeneralDiscomfort(b bool) *ClinicalResultsCreate {
	crc.generalDiscomfort = &b
	return crc
}

// SetFever sets the fever field.
func (crc *ClinicalResultsCreate) SetFever(b bool) *ClinicalResultsCreate {
	crc.fever = &b
	return crc
}

// SetThirdAge sets the thirdAge field.
func (crc *ClinicalResultsCreate) SetThirdAge(b bool) *ClinicalResultsCreate {
	crc.thirdAge = &b
	return crc
}

// SetDyspnea sets the dyspnea field.
func (crc *ClinicalResultsCreate) SetDyspnea(b bool) *ClinicalResultsCreate {
	crc.dyspnea = &b
	return crc
}

// SetID sets the id field.
func (crc *ClinicalResultsCreate) SetID(u uuid.UUID) *ClinicalResultsCreate {
	crc.id = &u
	return crc
}

// SetTestID sets the test edge to Test by id.
func (crc *ClinicalResultsCreate) SetTestID(id uuid.UUID) *ClinicalResultsCreate {
	if crc.test == nil {
		crc.test = make(map[uuid.UUID]struct{})
	}
	crc.test[id] = struct{}{}
	return crc
}

// SetNillableTestID sets the test edge to Test by id if the given value is not nil.
func (crc *ClinicalResultsCreate) SetNillableTestID(id *uuid.UUID) *ClinicalResultsCreate {
	if id != nil {
		crc = crc.SetTestID(*id)
	}
	return crc
}

// SetTest sets the test edge to Test.
func (crc *ClinicalResultsCreate) SetTest(t *Test) *ClinicalResultsCreate {
	return crc.SetTestID(t.ID)
}

// Save creates the ClinicalResults in the database.
func (crc *ClinicalResultsCreate) Save(ctx context.Context) (*ClinicalResults, error) {
	if crc.generalDiscomfort == nil {
		return nil, errors.New("ent: missing required field \"generalDiscomfort\"")
	}
	if crc.fever == nil {
		return nil, errors.New("ent: missing required field \"fever\"")
	}
	if crc.thirdAge == nil {
		return nil, errors.New("ent: missing required field \"thirdAge\"")
	}
	if crc.dyspnea == nil {
		return nil, errors.New("ent: missing required field \"dyspnea\"")
	}
	if len(crc.test) > 1 {
		return nil, errors.New("ent: multiple assignments on a unique edge \"test\"")
	}
	return crc.sqlSave(ctx)
}

// SaveX calls Save and panics if Save returns an error.
func (crc *ClinicalResultsCreate) SaveX(ctx context.Context) *ClinicalResults {
	v, err := crc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (crc *ClinicalResultsCreate) sqlSave(ctx context.Context) (*ClinicalResults, error) {
	var (
		cr    = &ClinicalResults{config: crc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: clinicalresults.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: clinicalresults.FieldID,
			},
		}
	)
	if value := crc.id; value != nil {
		cr.ID = *value
		_spec.ID.Value = *value
	}
	if value := crc.generalDiscomfort; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  *value,
			Column: clinicalresults.FieldGeneralDiscomfort,
		})
		cr.GeneralDiscomfort = *value
	}
	if value := crc.fever; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  *value,
			Column: clinicalresults.FieldFever,
		})
		cr.Fever = *value
	}
	if value := crc.thirdAge; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  *value,
			Column: clinicalresults.FieldThirdAge,
		})
		cr.ThirdAge = *value
	}
	if value := crc.dyspnea; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  *value,
			Column: clinicalresults.FieldDyspnea,
		})
		cr.Dyspnea = *value
	}
	if nodes := crc.test; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   clinicalresults.TestTable,
			Columns: []string{clinicalresults.TestColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: test.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if err := sqlgraph.CreateNode(ctx, crc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return cr, nil
}
