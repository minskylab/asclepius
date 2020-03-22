// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/google/uuid"
	"github.com/minskylab/asclepius/ent/history"
	"github.com/minskylab/asclepius/ent/medicalnote"
	"github.com/minskylab/asclepius/ent/patient"
	"github.com/minskylab/asclepius/ent/predicate"
	"github.com/minskylab/asclepius/ent/test"
)

// HistoryQuery is the builder for querying History entities.
type HistoryQuery struct {
	config
	limit      *int
	offset     *int
	order      []Order
	unique     []string
	predicates []predicate.History
	// eager-loading edges.
	withPatient *PatientQuery
	withTests   *TestQuery
	withNotes   *MedicalNoteQuery
	withFKs     bool
	// intermediate query.
	sql *sql.Selector
}

// Where adds a new predicate for the builder.
func (hq *HistoryQuery) Where(ps ...predicate.History) *HistoryQuery {
	hq.predicates = append(hq.predicates, ps...)
	return hq
}

// Limit adds a limit step to the query.
func (hq *HistoryQuery) Limit(limit int) *HistoryQuery {
	hq.limit = &limit
	return hq
}

// Offset adds an offset step to the query.
func (hq *HistoryQuery) Offset(offset int) *HistoryQuery {
	hq.offset = &offset
	return hq
}

// Order adds an order step to the query.
func (hq *HistoryQuery) Order(o ...Order) *HistoryQuery {
	hq.order = append(hq.order, o...)
	return hq
}

// QueryPatient chains the current query on the patient edge.
func (hq *HistoryQuery) QueryPatient() *PatientQuery {
	query := &PatientQuery{config: hq.config}
	step := sqlgraph.NewStep(
		sqlgraph.From(history.Table, history.FieldID, hq.sqlQuery()),
		sqlgraph.To(patient.Table, patient.FieldID),
		sqlgraph.Edge(sqlgraph.O2O, true, history.PatientTable, history.PatientColumn),
	)
	query.sql = sqlgraph.SetNeighbors(hq.driver.Dialect(), step)
	return query
}

// QueryTests chains the current query on the tests edge.
func (hq *HistoryQuery) QueryTests() *TestQuery {
	query := &TestQuery{config: hq.config}
	step := sqlgraph.NewStep(
		sqlgraph.From(history.Table, history.FieldID, hq.sqlQuery()),
		sqlgraph.To(test.Table, test.FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, history.TestsTable, history.TestsColumn),
	)
	query.sql = sqlgraph.SetNeighbors(hq.driver.Dialect(), step)
	return query
}

// QueryNotes chains the current query on the notes edge.
func (hq *HistoryQuery) QueryNotes() *MedicalNoteQuery {
	query := &MedicalNoteQuery{config: hq.config}
	step := sqlgraph.NewStep(
		sqlgraph.From(history.Table, history.FieldID, hq.sqlQuery()),
		sqlgraph.To(medicalnote.Table, medicalnote.FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, history.NotesTable, history.NotesColumn),
	)
	query.sql = sqlgraph.SetNeighbors(hq.driver.Dialect(), step)
	return query
}

// First returns the first History entity in the query. Returns *NotFoundError when no history was found.
func (hq *HistoryQuery) First(ctx context.Context) (*History, error) {
	hs, err := hq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(hs) == 0 {
		return nil, &NotFoundError{history.Label}
	}
	return hs[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (hq *HistoryQuery) FirstX(ctx context.Context) *History {
	h, err := hq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return h
}

// FirstID returns the first History id in the query. Returns *NotFoundError when no id was found.
func (hq *HistoryQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = hq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{history.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (hq *HistoryQuery) FirstXID(ctx context.Context) uuid.UUID {
	id, err := hq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only History entity in the query, returns an error if not exactly one entity was returned.
func (hq *HistoryQuery) Only(ctx context.Context) (*History, error) {
	hs, err := hq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(hs) {
	case 1:
		return hs[0], nil
	case 0:
		return nil, &NotFoundError{history.Label}
	default:
		return nil, &NotSingularError{history.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (hq *HistoryQuery) OnlyX(ctx context.Context) *History {
	h, err := hq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return h
}

// OnlyID returns the only History id in the query, returns an error if not exactly one id was returned.
func (hq *HistoryQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = hq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{history.Label}
	default:
		err = &NotSingularError{history.Label}
	}
	return
}

// OnlyXID is like OnlyID, but panics if an error occurs.
func (hq *HistoryQuery) OnlyXID(ctx context.Context) uuid.UUID {
	id, err := hq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Histories.
func (hq *HistoryQuery) All(ctx context.Context) ([]*History, error) {
	return hq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (hq *HistoryQuery) AllX(ctx context.Context) []*History {
	hs, err := hq.All(ctx)
	if err != nil {
		panic(err)
	}
	return hs
}

// IDs executes the query and returns a list of History ids.
func (hq *HistoryQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := hq.Select(history.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (hq *HistoryQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := hq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (hq *HistoryQuery) Count(ctx context.Context) (int, error) {
	return hq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (hq *HistoryQuery) CountX(ctx context.Context) int {
	count, err := hq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (hq *HistoryQuery) Exist(ctx context.Context) (bool, error) {
	return hq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (hq *HistoryQuery) ExistX(ctx context.Context) bool {
	exist, err := hq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (hq *HistoryQuery) Clone() *HistoryQuery {
	return &HistoryQuery{
		config:     hq.config,
		limit:      hq.limit,
		offset:     hq.offset,
		order:      append([]Order{}, hq.order...),
		unique:     append([]string{}, hq.unique...),
		predicates: append([]predicate.History{}, hq.predicates...),
		// clone intermediate query.
		sql: hq.sql.Clone(),
	}
}

//  WithPatient tells the query-builder to eager-loads the nodes that are connected to
// the "patient" edge. The optional arguments used to configure the query builder of the edge.
func (hq *HistoryQuery) WithPatient(opts ...func(*PatientQuery)) *HistoryQuery {
	query := &PatientQuery{config: hq.config}
	for _, opt := range opts {
		opt(query)
	}
	hq.withPatient = query
	return hq
}

//  WithTests tells the query-builder to eager-loads the nodes that are connected to
// the "tests" edge. The optional arguments used to configure the query builder of the edge.
func (hq *HistoryQuery) WithTests(opts ...func(*TestQuery)) *HistoryQuery {
	query := &TestQuery{config: hq.config}
	for _, opt := range opts {
		opt(query)
	}
	hq.withTests = query
	return hq
}

//  WithNotes tells the query-builder to eager-loads the nodes that are connected to
// the "notes" edge. The optional arguments used to configure the query builder of the edge.
func (hq *HistoryQuery) WithNotes(opts ...func(*MedicalNoteQuery)) *HistoryQuery {
	query := &MedicalNoteQuery{config: hq.config}
	for _, opt := range opts {
		opt(query)
	}
	hq.withNotes = query
	return hq
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
func (hq *HistoryQuery) GroupBy(field string, fields ...string) *HistoryGroupBy {
	group := &HistoryGroupBy{config: hq.config}
	group.fields = append([]string{field}, fields...)
	group.sql = hq.sqlQuery()
	return group
}

// Select one or more fields from the given query.
func (hq *HistoryQuery) Select(field string, fields ...string) *HistorySelect {
	selector := &HistorySelect{config: hq.config}
	selector.fields = append([]string{field}, fields...)
	selector.sql = hq.sqlQuery()
	return selector
}

func (hq *HistoryQuery) sqlAll(ctx context.Context) ([]*History, error) {
	var (
		nodes       = []*History{}
		withFKs     = hq.withFKs
		_spec       = hq.querySpec()
		loadedTypes = [3]bool{
			hq.withPatient != nil,
			hq.withTests != nil,
			hq.withNotes != nil,
		}
	)
	if hq.withPatient != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, history.ForeignKeys...)
	}
	_spec.ScanValues = func() []interface{} {
		node := &History{config: hq.config}
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
	if err := sqlgraph.QueryNodes(ctx, hq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := hq.withPatient; query != nil {
		ids := make([]uuid.UUID, 0, len(nodes))
		nodeids := make(map[uuid.UUID][]*History)
		for i := range nodes {
			if fk := nodes[i].patient_history; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(patient.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "patient_history" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Patient = n
			}
		}
	}

	if query := hq.withTests; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[uuid.UUID]*History)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
		}
		query.withFKs = true
		query.Where(predicate.Test(func(s *sql.Selector) {
			s.Where(sql.InValues(history.TestsColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.history_tests
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "history_tests" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "history_tests" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Tests = append(node.Edges.Tests, n)
		}
	}

	if query := hq.withNotes; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[uuid.UUID]*History)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
		}
		query.withFKs = true
		query.Where(predicate.MedicalNote(func(s *sql.Selector) {
			s.Where(sql.InValues(history.NotesColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.history_notes
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "history_notes" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "history_notes" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Notes = append(node.Edges.Notes, n)
		}
	}

	return nodes, nil
}

func (hq *HistoryQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := hq.querySpec()
	return sqlgraph.CountNodes(ctx, hq.driver, _spec)
}

func (hq *HistoryQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := hq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (hq *HistoryQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   history.Table,
			Columns: history.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: history.FieldID,
			},
		},
		From:   hq.sql,
		Unique: true,
	}
	if ps := hq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := hq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := hq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := hq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (hq *HistoryQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(hq.driver.Dialect())
	t1 := builder.Table(history.Table)
	selector := builder.Select(t1.Columns(history.Columns...)...).From(t1)
	if hq.sql != nil {
		selector = hq.sql
		selector.Select(selector.Columns(history.Columns...)...)
	}
	for _, p := range hq.predicates {
		p(selector)
	}
	for _, p := range hq.order {
		p(selector)
	}
	if offset := hq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := hq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// HistoryGroupBy is the builder for group-by History entities.
type HistoryGroupBy struct {
	config
	fields []string
	fns    []Aggregate
	// intermediate query.
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the group-by query.
func (hgb *HistoryGroupBy) Aggregate(fns ...Aggregate) *HistoryGroupBy {
	hgb.fns = append(hgb.fns, fns...)
	return hgb
}

// Scan applies the group-by query and scan the result into the given value.
func (hgb *HistoryGroupBy) Scan(ctx context.Context, v interface{}) error {
	return hgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (hgb *HistoryGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := hgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (hgb *HistoryGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(hgb.fields) > 1 {
		return nil, errors.New("ent: HistoryGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := hgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (hgb *HistoryGroupBy) StringsX(ctx context.Context) []string {
	v, err := hgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (hgb *HistoryGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(hgb.fields) > 1 {
		return nil, errors.New("ent: HistoryGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := hgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (hgb *HistoryGroupBy) IntsX(ctx context.Context) []int {
	v, err := hgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (hgb *HistoryGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(hgb.fields) > 1 {
		return nil, errors.New("ent: HistoryGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := hgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (hgb *HistoryGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := hgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (hgb *HistoryGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(hgb.fields) > 1 {
		return nil, errors.New("ent: HistoryGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := hgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (hgb *HistoryGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := hgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (hgb *HistoryGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := hgb.sqlQuery().Query()
	if err := hgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (hgb *HistoryGroupBy) sqlQuery() *sql.Selector {
	selector := hgb.sql
	columns := make([]string, 0, len(hgb.fields)+len(hgb.fns))
	columns = append(columns, hgb.fields...)
	for _, fn := range hgb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(hgb.fields...)
}

// HistorySelect is the builder for select fields of History entities.
type HistorySelect struct {
	config
	fields []string
	// intermediate queries.
	sql *sql.Selector
}

// Scan applies the selector query and scan the result into the given value.
func (hs *HistorySelect) Scan(ctx context.Context, v interface{}) error {
	return hs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (hs *HistorySelect) ScanX(ctx context.Context, v interface{}) {
	if err := hs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (hs *HistorySelect) Strings(ctx context.Context) ([]string, error) {
	if len(hs.fields) > 1 {
		return nil, errors.New("ent: HistorySelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := hs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (hs *HistorySelect) StringsX(ctx context.Context) []string {
	v, err := hs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (hs *HistorySelect) Ints(ctx context.Context) ([]int, error) {
	if len(hs.fields) > 1 {
		return nil, errors.New("ent: HistorySelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := hs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (hs *HistorySelect) IntsX(ctx context.Context) []int {
	v, err := hs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (hs *HistorySelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(hs.fields) > 1 {
		return nil, errors.New("ent: HistorySelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := hs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (hs *HistorySelect) Float64sX(ctx context.Context) []float64 {
	v, err := hs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (hs *HistorySelect) Bools(ctx context.Context) ([]bool, error) {
	if len(hs.fields) > 1 {
		return nil, errors.New("ent: HistorySelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := hs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (hs *HistorySelect) BoolsX(ctx context.Context) []bool {
	v, err := hs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (hs *HistorySelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := hs.sqlQuery().Query()
	if err := hs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (hs *HistorySelect) sqlQuery() sql.Querier {
	selector := hs.sql
	selector.Select(selector.Columns(hs.fields...)...)
	return selector
}
