// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/minskylab/asclepius/ent/epidemiologicresults"
	"github.com/minskylab/asclepius/ent/predicate"
)

// EpidemiologicResultsDelete is the builder for deleting a EpidemiologicResults entity.
type EpidemiologicResultsDelete struct {
	config
	predicates []predicate.EpidemiologicResults
}

// Where adds a new predicate to the delete builder.
func (erd *EpidemiologicResultsDelete) Where(ps ...predicate.EpidemiologicResults) *EpidemiologicResultsDelete {
	erd.predicates = append(erd.predicates, ps...)
	return erd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (erd *EpidemiologicResultsDelete) Exec(ctx context.Context) (int, error) {
	return erd.sqlExec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (erd *EpidemiologicResultsDelete) ExecX(ctx context.Context) int {
	n, err := erd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (erd *EpidemiologicResultsDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: epidemiologicresults.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: epidemiologicresults.FieldID,
			},
		},
	}
	if ps := erd.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, erd.driver, _spec)
}

// EpidemiologicResultsDeleteOne is the builder for deleting a single EpidemiologicResults entity.
type EpidemiologicResultsDeleteOne struct {
	erd *EpidemiologicResultsDelete
}

// Exec executes the deletion query.
func (erdo *EpidemiologicResultsDeleteOne) Exec(ctx context.Context) error {
	n, err := erdo.erd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{epidemiologicresults.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (erdo *EpidemiologicResultsDeleteOne) ExecX(ctx context.Context) {
	erdo.erd.ExecX(ctx)
}
