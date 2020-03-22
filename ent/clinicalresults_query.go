// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/google/uuid"
	"github.com/minskylab/asclepius/ent/clinicalresults"
	"github.com/minskylab/asclepius/ent/predicate"
	"github.com/minskylab/asclepius/ent/test"
)

// ClinicalResultsQuery is the builder for querying ClinicalResults entities.
type ClinicalResultsQuery struct {
	config
	limit      *int
	offset     *int
	order      []Order
	unique     []string
	predicates []predicate.ClinicalResults
	// eager-loading edges.
	withTest *TestQuery
	withFKs  bool
	// intermediate query.
	sql *sql.Selector
}

// Where adds a new predicate for the builder.
func (crq *ClinicalResultsQuery) Where(ps ...predicate.ClinicalResults) *ClinicalResultsQuery {
	crq.predicates = append(crq.predicates, ps...)
	return crq
}

// Limit adds a limit step to the query.
func (crq *ClinicalResultsQuery) Limit(limit int) *ClinicalResultsQuery {
	crq.limit = &limit
	return crq
}

// Offset adds an offset step to the query.
func (crq *ClinicalResultsQuery) Offset(offset int) *ClinicalResultsQuery {
	crq.offset = &offset
	return crq
}

// Order adds an order step to the query.
func (crq *ClinicalResultsQuery) Order(o ...Order) *ClinicalResultsQuery {
	crq.order = append(crq.order, o...)
	return crq
}

// QueryTest chains the current query on the test edge.
func (crq *ClinicalResultsQuery) QueryTest() *TestQuery {
	query := &TestQuery{config: crq.config}
	step := sqlgraph.NewStep(
		sqlgraph.From(clinicalresults.Table, clinicalresults.FieldID, crq.sqlQuery()),
		sqlgraph.To(test.Table, test.FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, clinicalresults.TestTable, clinicalresults.TestColumn),
	)
	query.sql = sqlgraph.SetNeighbors(crq.driver.Dialect(), step)
	return query
}

// First returns the first ClinicalResults entity in the query. Returns *NotFoundError when no clinicalresults was found.
func (crq *ClinicalResultsQuery) First(ctx context.Context) (*ClinicalResults, error) {
	crs, err := crq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(crs) == 0 {
		return nil, &NotFoundError{clinicalresults.Label}
	}
	return crs[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (crq *ClinicalResultsQuery) FirstX(ctx context.Context) *ClinicalResults {
	cr, err := crq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return cr
}

// FirstID returns the first ClinicalResults id in the query. Returns *NotFoundError when no id was found.
func (crq *ClinicalResultsQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = crq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{clinicalresults.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (crq *ClinicalResultsQuery) FirstXID(ctx context.Context) uuid.UUID {
	id, err := crq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only ClinicalResults entity in the query, returns an error if not exactly one entity was returned.
func (crq *ClinicalResultsQuery) Only(ctx context.Context) (*ClinicalResults, error) {
	crs, err := crq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(crs) {
	case 1:
		return crs[0], nil
	case 0:
		return nil, &NotFoundError{clinicalresults.Label}
	default:
		return nil, &NotSingularError{clinicalresults.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (crq *ClinicalResultsQuery) OnlyX(ctx context.Context) *ClinicalResults {
	cr, err := crq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return cr
}

// OnlyID returns the only ClinicalResults id in the query, returns an error if not exactly one id was returned.
func (crq *ClinicalResultsQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = crq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{clinicalresults.Label}
	default:
		err = &NotSingularError{clinicalresults.Label}
	}
	return
}

// OnlyXID is like OnlyID, but panics if an error occurs.
func (crq *ClinicalResultsQuery) OnlyXID(ctx context.Context) uuid.UUID {
	id, err := crq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ClinicalResultsSlice.
func (crq *ClinicalResultsQuery) All(ctx context.Context) ([]*ClinicalResults, error) {
	return crq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (crq *ClinicalResultsQuery) AllX(ctx context.Context) []*ClinicalResults {
	crs, err := crq.All(ctx)
	if err != nil {
		panic(err)
	}
	return crs
}

// IDs executes the query and returns a list of ClinicalResults ids.
func (crq *ClinicalResultsQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := crq.Select(clinicalresults.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (crq *ClinicalResultsQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := crq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (crq *ClinicalResultsQuery) Count(ctx context.Context) (int, error) {
	return crq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (crq *ClinicalResultsQuery) CountX(ctx context.Context) int {
	count, err := crq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (crq *ClinicalResultsQuery) Exist(ctx context.Context) (bool, error) {
	return crq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (crq *ClinicalResultsQuery) ExistX(ctx context.Context) bool {
	exist, err := crq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (crq *ClinicalResultsQuery) Clone() *ClinicalResultsQuery {
	return &ClinicalResultsQuery{
		config:     crq.config,
		limit:      crq.limit,
		offset:     crq.offset,
		order:      append([]Order{}, crq.order...),
		unique:     append([]string{}, crq.unique...),
		predicates: append([]predicate.ClinicalResults{}, crq.predicates...),
		// clone intermediate query.
		sql: crq.sql.Clone(),
	}
}

//  WithTest tells the query-builder to eager-loads the nodes that are connected to
// the "test" edge. The optional arguments used to configure the query builder of the edge.
func (crq *ClinicalResultsQuery) WithTest(opts ...func(*TestQuery)) *ClinicalResultsQuery {
	query := &TestQuery{config: crq.config}
	for _, opt := range opts {
		opt(query)
	}
	crq.withTest = query
	return crq
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		GeneralDiscomfort bool `json:"generalDiscomfort,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.ClinicalResults.Query().
//		GroupBy(clinicalresults.FieldGeneralDiscomfort).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (crq *ClinicalResultsQuery) GroupBy(field string, fields ...string) *ClinicalResultsGroupBy {
	group := &ClinicalResultsGroupBy{config: crq.config}
	group.fields = append([]string{field}, fields...)
	group.sql = crq.sqlQuery()
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		GeneralDiscomfort bool `json:"generalDiscomfort,omitempty"`
//	}
//
//	client.ClinicalResults.Query().
//		Select(clinicalresults.FieldGeneralDiscomfort).
//		Scan(ctx, &v)
//
func (crq *ClinicalResultsQuery) Select(field string, fields ...string) *ClinicalResultsSelect {
	selector := &ClinicalResultsSelect{config: crq.config}
	selector.fields = append([]string{field}, fields...)
	selector.sql = crq.sqlQuery()
	return selector
}

func (crq *ClinicalResultsQuery) sqlAll(ctx context.Context) ([]*ClinicalResults, error) {
	var (
		nodes       = []*ClinicalResults{}
		withFKs     = crq.withFKs
		_spec       = crq.querySpec()
		loadedTypes = [1]bool{
			crq.withTest != nil,
		}
	)
	if crq.withTest != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, clinicalresults.ForeignKeys...)
	}
	_spec.ScanValues = func() []interface{} {
		node := &ClinicalResults{config: crq.config}
		nodes = append(nodes, node)
		values := node.scanValues()
		if withFKs {
			values = append(values, node.fkValues()...)
		}
		return values
	}
	_spec.Assign = func(values ...interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(values...)
	}
	if err := sqlgraph.QueryNodes(ctx, crq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := crq.withTest; query != nil {
		ids := make([]uuid.UUID, 0, len(nodes))
		nodeids := make(map[uuid.UUID][]*ClinicalResults)
		for i := range nodes {
			if fk := nodes[i].test_clinical; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(test.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "test_clinical" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Test = n
			}
		}
	}

	return nodes, nil
}

func (crq *ClinicalResultsQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := crq.querySpec()
	return sqlgraph.CountNodes(ctx, crq.driver, _spec)
}

func (crq *ClinicalResultsQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := crq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (crq *ClinicalResultsQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   clinicalresults.Table,
			Columns: clinicalresults.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: clinicalresults.FieldID,
			},
		},
		From:   crq.sql,
		Unique: true,
	}
	if ps := crq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := crq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := crq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := crq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (crq *ClinicalResultsQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(crq.driver.Dialect())
	t1 := builder.Table(clinicalresults.Table)
	selector := builder.Select(t1.Columns(clinicalresults.Columns...)...).From(t1)
	if crq.sql != nil {
		selector = crq.sql
		selector.Select(selector.Columns(clinicalresults.Columns...)...)
	}
	for _, p := range crq.predicates {
		p(selector)
	}
	for _, p := range crq.order {
		p(selector)
	}
	if offset := crq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := crq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ClinicalResultsGroupBy is the builder for group-by ClinicalResults entities.
type ClinicalResultsGroupBy struct {
	config
	fields []string
	fns    []Aggregate
	// intermediate query.
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the group-by query.
func (crgb *ClinicalResultsGroupBy) Aggregate(fns ...Aggregate) *ClinicalResultsGroupBy {
	crgb.fns = append(crgb.fns, fns...)
	return crgb
}

// Scan applies the group-by query and scan the result into the given value.
func (crgb *ClinicalResultsGroupBy) Scan(ctx context.Context, v interface{}) error {
	return crgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (crgb *ClinicalResultsGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := crgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (crgb *ClinicalResultsGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(crgb.fields) > 1 {
		return nil, errors.New("ent: ClinicalResultsGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := crgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (crgb *ClinicalResultsGroupBy) StringsX(ctx context.Context) []string {
	v, err := crgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (crgb *ClinicalResultsGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(crgb.fields) > 1 {
		return nil, errors.New("ent: ClinicalResultsGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := crgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (crgb *ClinicalResultsGroupBy) IntsX(ctx context.Context) []int {
	v, err := crgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (crgb *ClinicalResultsGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(crgb.fields) > 1 {
		return nil, errors.New("ent: ClinicalResultsGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := crgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (crgb *ClinicalResultsGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := crgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (crgb *ClinicalResultsGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(crgb.fields) > 1 {
		return nil, errors.New("ent: ClinicalResultsGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := crgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (crgb *ClinicalResultsGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := crgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (crgb *ClinicalResultsGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := crgb.sqlQuery().Query()
	if err := crgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (crgb *ClinicalResultsGroupBy) sqlQuery() *sql.Selector {
	selector := crgb.sql
	columns := make([]string, 0, len(crgb.fields)+len(crgb.fns))
	columns = append(columns, crgb.fields...)
	for _, fn := range crgb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(crgb.fields...)
}

// ClinicalResultsSelect is the builder for select fields of ClinicalResults entities.
type ClinicalResultsSelect struct {
	config
	fields []string
	// intermediate queries.
	sql *sql.Selector
}

// Scan applies the selector query and scan the result into the given value.
func (crs *ClinicalResultsSelect) Scan(ctx context.Context, v interface{}) error {
	return crs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (crs *ClinicalResultsSelect) ScanX(ctx context.Context, v interface{}) {
	if err := crs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (crs *ClinicalResultsSelect) Strings(ctx context.Context) ([]string, error) {
	if len(crs.fields) > 1 {
		return nil, errors.New("ent: ClinicalResultsSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := crs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (crs *ClinicalResultsSelect) StringsX(ctx context.Context) []string {
	v, err := crs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (crs *ClinicalResultsSelect) Ints(ctx context.Context) ([]int, error) {
	if len(crs.fields) > 1 {
		return nil, errors.New("ent: ClinicalResultsSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := crs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (crs *ClinicalResultsSelect) IntsX(ctx context.Context) []int {
	v, err := crs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (crs *ClinicalResultsSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(crs.fields) > 1 {
		return nil, errors.New("ent: ClinicalResultsSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := crs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (crs *ClinicalResultsSelect) Float64sX(ctx context.Context) []float64 {
	v, err := crs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (crs *ClinicalResultsSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(crs.fields) > 1 {
		return nil, errors.New("ent: ClinicalResultsSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := crs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (crs *ClinicalResultsSelect) BoolsX(ctx context.Context) []bool {
	v, err := crs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (crs *ClinicalResultsSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := crs.sqlQuery().Query()
	if err := crs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (crs *ClinicalResultsSelect) sqlQuery() sql.Querier {
	selector := crs.sql
	selector.Select(selector.Columns(crs.fields...)...)
	return selector
}
