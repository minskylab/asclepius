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
	"github.com/minskylab/asclepius/ent/doctor"
	"github.com/minskylab/asclepius/ent/medicalnote"
	"github.com/minskylab/asclepius/ent/predicate"
	"github.com/minskylab/asclepius/ent/task"
	"github.com/minskylab/asclepius/ent/taskresponse"
)

// DoctorQuery is the builder for querying Doctor entities.
type DoctorQuery struct {
	config
	limit      *int
	offset     *int
	order      []Order
	unique     []string
	predicates []predicate.Doctor
	// eager-loading edges.
	withNotes     *MedicalNoteQuery
	withResponses *TaskResponseQuery
	withTasks     *TaskQuery
	// intermediate query.
	sql *sql.Selector
}

// Where adds a new predicate for the builder.
func (dq *DoctorQuery) Where(ps ...predicate.Doctor) *DoctorQuery {
	dq.predicates = append(dq.predicates, ps...)
	return dq
}

// Limit adds a limit step to the query.
func (dq *DoctorQuery) Limit(limit int) *DoctorQuery {
	dq.limit = &limit
	return dq
}

// Offset adds an offset step to the query.
func (dq *DoctorQuery) Offset(offset int) *DoctorQuery {
	dq.offset = &offset
	return dq
}

// Order adds an order step to the query.
func (dq *DoctorQuery) Order(o ...Order) *DoctorQuery {
	dq.order = append(dq.order, o...)
	return dq
}

// QueryNotes chains the current query on the notes edge.
func (dq *DoctorQuery) QueryNotes() *MedicalNoteQuery {
	query := &MedicalNoteQuery{config: dq.config}
	step := sqlgraph.NewStep(
		sqlgraph.From(doctor.Table, doctor.FieldID, dq.sqlQuery()),
		sqlgraph.To(medicalnote.Table, medicalnote.FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, doctor.NotesTable, doctor.NotesColumn),
	)
	query.sql = sqlgraph.SetNeighbors(dq.driver.Dialect(), step)
	return query
}

// QueryResponses chains the current query on the responses edge.
func (dq *DoctorQuery) QueryResponses() *TaskResponseQuery {
	query := &TaskResponseQuery{config: dq.config}
	step := sqlgraph.NewStep(
		sqlgraph.From(doctor.Table, doctor.FieldID, dq.sqlQuery()),
		sqlgraph.To(taskresponse.Table, taskresponse.FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, doctor.ResponsesTable, doctor.ResponsesColumn),
	)
	query.sql = sqlgraph.SetNeighbors(dq.driver.Dialect(), step)
	return query
}

// QueryTasks chains the current query on the tasks edge.
func (dq *DoctorQuery) QueryTasks() *TaskQuery {
	query := &TaskQuery{config: dq.config}
	step := sqlgraph.NewStep(
		sqlgraph.From(doctor.Table, doctor.FieldID, dq.sqlQuery()),
		sqlgraph.To(task.Table, task.FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, doctor.TasksTable, doctor.TasksPrimaryKey...),
	)
	query.sql = sqlgraph.SetNeighbors(dq.driver.Dialect(), step)
	return query
}

// First returns the first Doctor entity in the query. Returns *NotFoundError when no doctor was found.
func (dq *DoctorQuery) First(ctx context.Context) (*Doctor, error) {
	ds, err := dq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(ds) == 0 {
		return nil, &NotFoundError{doctor.Label}
	}
	return ds[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (dq *DoctorQuery) FirstX(ctx context.Context) *Doctor {
	d, err := dq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return d
}

// FirstID returns the first Doctor id in the query. Returns *NotFoundError when no id was found.
func (dq *DoctorQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = dq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{doctor.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (dq *DoctorQuery) FirstXID(ctx context.Context) uuid.UUID {
	id, err := dq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only Doctor entity in the query, returns an error if not exactly one entity was returned.
func (dq *DoctorQuery) Only(ctx context.Context) (*Doctor, error) {
	ds, err := dq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(ds) {
	case 1:
		return ds[0], nil
	case 0:
		return nil, &NotFoundError{doctor.Label}
	default:
		return nil, &NotSingularError{doctor.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (dq *DoctorQuery) OnlyX(ctx context.Context) *Doctor {
	d, err := dq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return d
}

// OnlyID returns the only Doctor id in the query, returns an error if not exactly one id was returned.
func (dq *DoctorQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = dq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{doctor.Label}
	default:
		err = &NotSingularError{doctor.Label}
	}
	return
}

// OnlyXID is like OnlyID, but panics if an error occurs.
func (dq *DoctorQuery) OnlyXID(ctx context.Context) uuid.UUID {
	id, err := dq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Doctors.
func (dq *DoctorQuery) All(ctx context.Context) ([]*Doctor, error) {
	return dq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (dq *DoctorQuery) AllX(ctx context.Context) []*Doctor {
	ds, err := dq.All(ctx)
	if err != nil {
		panic(err)
	}
	return ds
}

// IDs executes the query and returns a list of Doctor ids.
func (dq *DoctorQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := dq.Select(doctor.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (dq *DoctorQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := dq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (dq *DoctorQuery) Count(ctx context.Context) (int, error) {
	return dq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (dq *DoctorQuery) CountX(ctx context.Context) int {
	count, err := dq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (dq *DoctorQuery) Exist(ctx context.Context) (bool, error) {
	return dq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (dq *DoctorQuery) ExistX(ctx context.Context) bool {
	exist, err := dq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (dq *DoctorQuery) Clone() *DoctorQuery {
	return &DoctorQuery{
		config:     dq.config,
		limit:      dq.limit,
		offset:     dq.offset,
		order:      append([]Order{}, dq.order...),
		unique:     append([]string{}, dq.unique...),
		predicates: append([]predicate.Doctor{}, dq.predicates...),
		// clone intermediate query.
		sql: dq.sql.Clone(),
	}
}

//  WithNotes tells the query-builder to eager-loads the nodes that are connected to
// the "notes" edge. The optional arguments used to configure the query builder of the edge.
func (dq *DoctorQuery) WithNotes(opts ...func(*MedicalNoteQuery)) *DoctorQuery {
	query := &MedicalNoteQuery{config: dq.config}
	for _, opt := range opts {
		opt(query)
	}
	dq.withNotes = query
	return dq
}

//  WithResponses tells the query-builder to eager-loads the nodes that are connected to
// the "responses" edge. The optional arguments used to configure the query builder of the edge.
func (dq *DoctorQuery) WithResponses(opts ...func(*TaskResponseQuery)) *DoctorQuery {
	query := &TaskResponseQuery{config: dq.config}
	for _, opt := range opts {
		opt(query)
	}
	dq.withResponses = query
	return dq
}

//  WithTasks tells the query-builder to eager-loads the nodes that are connected to
// the "tasks" edge. The optional arguments used to configure the query builder of the edge.
func (dq *DoctorQuery) WithTasks(opts ...func(*TaskQuery)) *DoctorQuery {
	query := &TaskQuery{config: dq.config}
	for _, opt := range opts {
		opt(query)
	}
	dq.withTasks = query
	return dq
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name []string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Doctor.Query().
//		GroupBy(doctor.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (dq *DoctorQuery) GroupBy(field string, fields ...string) *DoctorGroupBy {
	group := &DoctorGroupBy{config: dq.config}
	group.fields = append([]string{field}, fields...)
	group.sql = dq.sqlQuery()
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		Name []string `json:"name,omitempty"`
//	}
//
//	client.Doctor.Query().
//		Select(doctor.FieldName).
//		Scan(ctx, &v)
//
func (dq *DoctorQuery) Select(field string, fields ...string) *DoctorSelect {
	selector := &DoctorSelect{config: dq.config}
	selector.fields = append([]string{field}, fields...)
	selector.sql = dq.sqlQuery()
	return selector
}

func (dq *DoctorQuery) sqlAll(ctx context.Context) ([]*Doctor, error) {
	var (
		nodes       = []*Doctor{}
		_spec       = dq.querySpec()
		loadedTypes = [3]bool{
			dq.withNotes != nil,
			dq.withResponses != nil,
			dq.withTasks != nil,
		}
	)
	_spec.ScanValues = func() []interface{} {
		node := &Doctor{config: dq.config}
		nodes = append(nodes, node)
		values := node.scanValues()
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
	if err := sqlgraph.QueryNodes(ctx, dq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := dq.withNotes; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[uuid.UUID]*Doctor)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
		}
		query.withFKs = true
		query.Where(predicate.MedicalNote(func(s *sql.Selector) {
			s.Where(sql.InValues(doctor.NotesColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.doctor_notes
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "doctor_notes" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "doctor_notes" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Notes = append(node.Edges.Notes, n)
		}
	}

	if query := dq.withResponses; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[uuid.UUID]*Doctor)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
		}
		query.withFKs = true
		query.Where(predicate.TaskResponse(func(s *sql.Selector) {
			s.Where(sql.InValues(doctor.ResponsesColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.doctor_responses
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "doctor_responses" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "doctor_responses" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Responses = append(node.Edges.Responses, n)
		}
	}

	if query := dq.withTasks; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		ids := make(map[uuid.UUID]*Doctor, len(nodes))
		for _, node := range nodes {
			ids[node.ID] = node
			fks = append(fks, node.ID)
		}
		var (
			edgeids []uuid.UUID
			edges   = make(map[uuid.UUID][]*Doctor)
		)
		_spec := &sqlgraph.EdgeQuerySpec{
			Edge: &sqlgraph.EdgeSpec{
				Inverse: true,
				Table:   doctor.TasksTable,
				Columns: doctor.TasksPrimaryKey,
			},
			Predicate: func(s *sql.Selector) {
				s.Where(sql.InValues(doctor.TasksPrimaryKey[1], fks...))
			},

			ScanValues: func() [2]interface{} {
				return [2]interface{}{&uuid.UUID{}, &uuid.UUID{}}
			},
			Assign: func(out, in interface{}) error {
				eout, ok := out.(*uuid.UUID)
				if !ok || eout == nil {
					return fmt.Errorf("unexpected id value for edge-out")
				}
				ein, ok := in.(*uuid.UUID)
				if !ok || ein == nil {
					return fmt.Errorf("unexpected id value for edge-in")
				}
				outValue := *eout
				inValue := *ein
				node, ok := ids[outValue]
				if !ok {
					return fmt.Errorf("unexpected node id in edges: %v", outValue)
				}
				edgeids = append(edgeids, inValue)
				edges[inValue] = append(edges[inValue], node)
				return nil
			},
		}
		if err := sqlgraph.QueryEdges(ctx, dq.driver, _spec); err != nil {
			return nil, fmt.Errorf(`query edges "tasks": %v`, err)
		}
		query.Where(task.IDIn(edgeids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := edges[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected "tasks" node returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Tasks = append(nodes[i].Edges.Tasks, n)
			}
		}
	}

	return nodes, nil
}

func (dq *DoctorQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := dq.querySpec()
	return sqlgraph.CountNodes(ctx, dq.driver, _spec)
}

func (dq *DoctorQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := dq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (dq *DoctorQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   doctor.Table,
			Columns: doctor.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: doctor.FieldID,
			},
		},
		From:   dq.sql,
		Unique: true,
	}
	if ps := dq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := dq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := dq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := dq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (dq *DoctorQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(dq.driver.Dialect())
	t1 := builder.Table(doctor.Table)
	selector := builder.Select(t1.Columns(doctor.Columns...)...).From(t1)
	if dq.sql != nil {
		selector = dq.sql
		selector.Select(selector.Columns(doctor.Columns...)...)
	}
	for _, p := range dq.predicates {
		p(selector)
	}
	for _, p := range dq.order {
		p(selector)
	}
	if offset := dq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := dq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// DoctorGroupBy is the builder for group-by Doctor entities.
type DoctorGroupBy struct {
	config
	fields []string
	fns    []Aggregate
	// intermediate query.
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the group-by query.
func (dgb *DoctorGroupBy) Aggregate(fns ...Aggregate) *DoctorGroupBy {
	dgb.fns = append(dgb.fns, fns...)
	return dgb
}

// Scan applies the group-by query and scan the result into the given value.
func (dgb *DoctorGroupBy) Scan(ctx context.Context, v interface{}) error {
	return dgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (dgb *DoctorGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := dgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (dgb *DoctorGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(dgb.fields) > 1 {
		return nil, errors.New("ent: DoctorGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := dgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (dgb *DoctorGroupBy) StringsX(ctx context.Context) []string {
	v, err := dgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (dgb *DoctorGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(dgb.fields) > 1 {
		return nil, errors.New("ent: DoctorGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := dgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (dgb *DoctorGroupBy) IntsX(ctx context.Context) []int {
	v, err := dgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (dgb *DoctorGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(dgb.fields) > 1 {
		return nil, errors.New("ent: DoctorGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := dgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (dgb *DoctorGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := dgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (dgb *DoctorGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(dgb.fields) > 1 {
		return nil, errors.New("ent: DoctorGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := dgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (dgb *DoctorGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := dgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (dgb *DoctorGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := dgb.sqlQuery().Query()
	if err := dgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (dgb *DoctorGroupBy) sqlQuery() *sql.Selector {
	selector := dgb.sql
	columns := make([]string, 0, len(dgb.fields)+len(dgb.fns))
	columns = append(columns, dgb.fields...)
	for _, fn := range dgb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(dgb.fields...)
}

// DoctorSelect is the builder for select fields of Doctor entities.
type DoctorSelect struct {
	config
	fields []string
	// intermediate queries.
	sql *sql.Selector
}

// Scan applies the selector query and scan the result into the given value.
func (ds *DoctorSelect) Scan(ctx context.Context, v interface{}) error {
	return ds.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ds *DoctorSelect) ScanX(ctx context.Context, v interface{}) {
	if err := ds.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (ds *DoctorSelect) Strings(ctx context.Context) ([]string, error) {
	if len(ds.fields) > 1 {
		return nil, errors.New("ent: DoctorSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := ds.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ds *DoctorSelect) StringsX(ctx context.Context) []string {
	v, err := ds.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (ds *DoctorSelect) Ints(ctx context.Context) ([]int, error) {
	if len(ds.fields) > 1 {
		return nil, errors.New("ent: DoctorSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := ds.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ds *DoctorSelect) IntsX(ctx context.Context) []int {
	v, err := ds.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (ds *DoctorSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(ds.fields) > 1 {
		return nil, errors.New("ent: DoctorSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := ds.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ds *DoctorSelect) Float64sX(ctx context.Context) []float64 {
	v, err := ds.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (ds *DoctorSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(ds.fields) > 1 {
		return nil, errors.New("ent: DoctorSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := ds.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ds *DoctorSelect) BoolsX(ctx context.Context) []bool {
	v, err := ds.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ds *DoctorSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ds.sqlQuery().Query()
	if err := ds.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ds *DoctorSelect) sqlQuery() sql.Querier {
	selector := ds.sql
	selector.Select(selector.Columns(ds.fields...)...)
	return selector
}
