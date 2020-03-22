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
	"github.com/minskylab/asclepius/ent/patient"
	"github.com/minskylab/asclepius/ent/predicate"
	"github.com/minskylab/asclepius/ent/schedule"
	"github.com/minskylab/asclepius/ent/task"
)

// ScheduleQuery is the builder for querying Schedule entities.
type ScheduleQuery struct {
	config
	limit      *int
	offset     *int
	order      []Order
	unique     []string
	predicates []predicate.Schedule
	// eager-loading edges.
	withTasks   *TaskQuery
	withPatient *PatientQuery
	withFKs     bool
	// intermediate query.
	sql *sql.Selector
}

// Where adds a new predicate for the builder.
func (sq *ScheduleQuery) Where(ps ...predicate.Schedule) *ScheduleQuery {
	sq.predicates = append(sq.predicates, ps...)
	return sq
}

// Limit adds a limit step to the query.
func (sq *ScheduleQuery) Limit(limit int) *ScheduleQuery {
	sq.limit = &limit
	return sq
}

// Offset adds an offset step to the query.
func (sq *ScheduleQuery) Offset(offset int) *ScheduleQuery {
	sq.offset = &offset
	return sq
}

// Order adds an order step to the query.
func (sq *ScheduleQuery) Order(o ...Order) *ScheduleQuery {
	sq.order = append(sq.order, o...)
	return sq
}

// QueryTasks chains the current query on the tasks edge.
func (sq *ScheduleQuery) QueryTasks() *TaskQuery {
	query := &TaskQuery{config: sq.config}
	step := sqlgraph.NewStep(
		sqlgraph.From(schedule.Table, schedule.FieldID, sq.sqlQuery()),
		sqlgraph.To(task.Table, task.FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, schedule.TasksTable, schedule.TasksColumn),
	)
	query.sql = sqlgraph.SetNeighbors(sq.driver.Dialect(), step)
	return query
}

// QueryPatient chains the current query on the patient edge.
func (sq *ScheduleQuery) QueryPatient() *PatientQuery {
	query := &PatientQuery{config: sq.config}
	step := sqlgraph.NewStep(
		sqlgraph.From(schedule.Table, schedule.FieldID, sq.sqlQuery()),
		sqlgraph.To(patient.Table, patient.FieldID),
		sqlgraph.Edge(sqlgraph.O2O, true, schedule.PatientTable, schedule.PatientColumn),
	)
	query.sql = sqlgraph.SetNeighbors(sq.driver.Dialect(), step)
	return query
}

// First returns the first Schedule entity in the query. Returns *NotFoundError when no schedule was found.
func (sq *ScheduleQuery) First(ctx context.Context) (*Schedule, error) {
	sSlice, err := sq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(sSlice) == 0 {
		return nil, &NotFoundError{schedule.Label}
	}
	return sSlice[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (sq *ScheduleQuery) FirstX(ctx context.Context) *Schedule {
	s, err := sq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return s
}

// FirstID returns the first Schedule id in the query. Returns *NotFoundError when no id was found.
func (sq *ScheduleQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = sq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{schedule.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (sq *ScheduleQuery) FirstXID(ctx context.Context) uuid.UUID {
	id, err := sq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only Schedule entity in the query, returns an error if not exactly one entity was returned.
func (sq *ScheduleQuery) Only(ctx context.Context) (*Schedule, error) {
	sSlice, err := sq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(sSlice) {
	case 1:
		return sSlice[0], nil
	case 0:
		return nil, &NotFoundError{schedule.Label}
	default:
		return nil, &NotSingularError{schedule.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (sq *ScheduleQuery) OnlyX(ctx context.Context) *Schedule {
	s, err := sq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return s
}

// OnlyID returns the only Schedule id in the query, returns an error if not exactly one id was returned.
func (sq *ScheduleQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = sq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{schedule.Label}
	default:
		err = &NotSingularError{schedule.Label}
	}
	return
}

// OnlyXID is like OnlyID, but panics if an error occurs.
func (sq *ScheduleQuery) OnlyXID(ctx context.Context) uuid.UUID {
	id, err := sq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Schedules.
func (sq *ScheduleQuery) All(ctx context.Context) ([]*Schedule, error) {
	return sq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (sq *ScheduleQuery) AllX(ctx context.Context) []*Schedule {
	sSlice, err := sq.All(ctx)
	if err != nil {
		panic(err)
	}
	return sSlice
}

// IDs executes the query and returns a list of Schedule ids.
func (sq *ScheduleQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := sq.Select(schedule.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (sq *ScheduleQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := sq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (sq *ScheduleQuery) Count(ctx context.Context) (int, error) {
	return sq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (sq *ScheduleQuery) CountX(ctx context.Context) int {
	count, err := sq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (sq *ScheduleQuery) Exist(ctx context.Context) (bool, error) {
	return sq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (sq *ScheduleQuery) ExistX(ctx context.Context) bool {
	exist, err := sq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (sq *ScheduleQuery) Clone() *ScheduleQuery {
	return &ScheduleQuery{
		config:     sq.config,
		limit:      sq.limit,
		offset:     sq.offset,
		order:      append([]Order{}, sq.order...),
		unique:     append([]string{}, sq.unique...),
		predicates: append([]predicate.Schedule{}, sq.predicates...),
		// clone intermediate query.
		sql: sq.sql.Clone(),
	}
}

//  WithTasks tells the query-builder to eager-loads the nodes that are connected to
// the "tasks" edge. The optional arguments used to configure the query builder of the edge.
func (sq *ScheduleQuery) WithTasks(opts ...func(*TaskQuery)) *ScheduleQuery {
	query := &TaskQuery{config: sq.config}
	for _, opt := range opts {
		opt(query)
	}
	sq.withTasks = query
	return sq
}

//  WithPatient tells the query-builder to eager-loads the nodes that are connected to
// the "patient" edge. The optional arguments used to configure the query builder of the edge.
func (sq *ScheduleQuery) WithPatient(opts ...func(*PatientQuery)) *ScheduleQuery {
	query := &PatientQuery{config: sq.config}
	for _, opt := range opts {
		opt(query)
	}
	sq.withPatient = query
	return sq
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
func (sq *ScheduleQuery) GroupBy(field string, fields ...string) *ScheduleGroupBy {
	group := &ScheduleGroupBy{config: sq.config}
	group.fields = append([]string{field}, fields...)
	group.sql = sq.sqlQuery()
	return group
}

// Select one or more fields from the given query.
func (sq *ScheduleQuery) Select(field string, fields ...string) *ScheduleSelect {
	selector := &ScheduleSelect{config: sq.config}
	selector.fields = append([]string{field}, fields...)
	selector.sql = sq.sqlQuery()
	return selector
}

func (sq *ScheduleQuery) sqlAll(ctx context.Context) ([]*Schedule, error) {
	var (
		nodes       = []*Schedule{}
		withFKs     = sq.withFKs
		_spec       = sq.querySpec()
		loadedTypes = [2]bool{
			sq.withTasks != nil,
			sq.withPatient != nil,
		}
	)
	if sq.withPatient != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, schedule.ForeignKeys...)
	}
	_spec.ScanValues = func() []interface{} {
		node := &Schedule{config: sq.config}
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
	if err := sqlgraph.QueryNodes(ctx, sq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := sq.withTasks; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[uuid.UUID]*Schedule)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
		}
		query.withFKs = true
		query.Where(predicate.Task(func(s *sql.Selector) {
			s.Where(sql.InValues(schedule.TasksColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.schedule_tasks
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "schedule_tasks" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "schedule_tasks" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Tasks = append(node.Edges.Tasks, n)
		}
	}

	if query := sq.withPatient; query != nil {
		ids := make([]uuid.UUID, 0, len(nodes))
		nodeids := make(map[uuid.UUID][]*Schedule)
		for i := range nodes {
			if fk := nodes[i].patient_schedule; fk != nil {
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
				return nil, fmt.Errorf(`unexpected foreign-key "patient_schedule" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Patient = n
			}
		}
	}

	return nodes, nil
}

func (sq *ScheduleQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := sq.querySpec()
	return sqlgraph.CountNodes(ctx, sq.driver, _spec)
}

func (sq *ScheduleQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := sq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (sq *ScheduleQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   schedule.Table,
			Columns: schedule.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: schedule.FieldID,
			},
		},
		From:   sq.sql,
		Unique: true,
	}
	if ps := sq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := sq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := sq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := sq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (sq *ScheduleQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(sq.driver.Dialect())
	t1 := builder.Table(schedule.Table)
	selector := builder.Select(t1.Columns(schedule.Columns...)...).From(t1)
	if sq.sql != nil {
		selector = sq.sql
		selector.Select(selector.Columns(schedule.Columns...)...)
	}
	for _, p := range sq.predicates {
		p(selector)
	}
	for _, p := range sq.order {
		p(selector)
	}
	if offset := sq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := sq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ScheduleGroupBy is the builder for group-by Schedule entities.
type ScheduleGroupBy struct {
	config
	fields []string
	fns    []Aggregate
	// intermediate query.
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the group-by query.
func (sgb *ScheduleGroupBy) Aggregate(fns ...Aggregate) *ScheduleGroupBy {
	sgb.fns = append(sgb.fns, fns...)
	return sgb
}

// Scan applies the group-by query and scan the result into the given value.
func (sgb *ScheduleGroupBy) Scan(ctx context.Context, v interface{}) error {
	return sgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (sgb *ScheduleGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := sgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (sgb *ScheduleGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(sgb.fields) > 1 {
		return nil, errors.New("ent: ScheduleGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := sgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (sgb *ScheduleGroupBy) StringsX(ctx context.Context) []string {
	v, err := sgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (sgb *ScheduleGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(sgb.fields) > 1 {
		return nil, errors.New("ent: ScheduleGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := sgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (sgb *ScheduleGroupBy) IntsX(ctx context.Context) []int {
	v, err := sgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (sgb *ScheduleGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(sgb.fields) > 1 {
		return nil, errors.New("ent: ScheduleGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := sgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (sgb *ScheduleGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := sgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (sgb *ScheduleGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(sgb.fields) > 1 {
		return nil, errors.New("ent: ScheduleGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := sgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (sgb *ScheduleGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := sgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (sgb *ScheduleGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := sgb.sqlQuery().Query()
	if err := sgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (sgb *ScheduleGroupBy) sqlQuery() *sql.Selector {
	selector := sgb.sql
	columns := make([]string, 0, len(sgb.fields)+len(sgb.fns))
	columns = append(columns, sgb.fields...)
	for _, fn := range sgb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(sgb.fields...)
}

// ScheduleSelect is the builder for select fields of Schedule entities.
type ScheduleSelect struct {
	config
	fields []string
	// intermediate queries.
	sql *sql.Selector
}

// Scan applies the selector query and scan the result into the given value.
func (ss *ScheduleSelect) Scan(ctx context.Context, v interface{}) error {
	return ss.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ss *ScheduleSelect) ScanX(ctx context.Context, v interface{}) {
	if err := ss.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (ss *ScheduleSelect) Strings(ctx context.Context) ([]string, error) {
	if len(ss.fields) > 1 {
		return nil, errors.New("ent: ScheduleSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := ss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ss *ScheduleSelect) StringsX(ctx context.Context) []string {
	v, err := ss.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (ss *ScheduleSelect) Ints(ctx context.Context) ([]int, error) {
	if len(ss.fields) > 1 {
		return nil, errors.New("ent: ScheduleSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := ss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ss *ScheduleSelect) IntsX(ctx context.Context) []int {
	v, err := ss.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (ss *ScheduleSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(ss.fields) > 1 {
		return nil, errors.New("ent: ScheduleSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := ss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ss *ScheduleSelect) Float64sX(ctx context.Context) []float64 {
	v, err := ss.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (ss *ScheduleSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(ss.fields) > 1 {
		return nil, errors.New("ent: ScheduleSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := ss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ss *ScheduleSelect) BoolsX(ctx context.Context) []bool {
	v, err := ss.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ss *ScheduleSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ss.sqlQuery().Query()
	if err := ss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ss *ScheduleSelect) sqlQuery() sql.Querier {
	selector := ss.sql
	selector.Select(selector.Columns(ss.fields...)...)
	return selector
}