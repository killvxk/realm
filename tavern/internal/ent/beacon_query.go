// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/kcarretto/realm/tavern/internal/ent/beacon"
	"github.com/kcarretto/realm/tavern/internal/ent/host"
	"github.com/kcarretto/realm/tavern/internal/ent/predicate"
	"github.com/kcarretto/realm/tavern/internal/ent/task"
)

// BeaconQuery is the builder for querying Beacon entities.
type BeaconQuery struct {
	config
	ctx            *QueryContext
	order          []beacon.OrderOption
	inters         []Interceptor
	predicates     []predicate.Beacon
	withHost       *HostQuery
	withTasks      *TaskQuery
	withFKs        bool
	modifiers      []func(*sql.Selector)
	loadTotal      []func(context.Context, []*Beacon) error
	withNamedTasks map[string]*TaskQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the BeaconQuery builder.
func (bq *BeaconQuery) Where(ps ...predicate.Beacon) *BeaconQuery {
	bq.predicates = append(bq.predicates, ps...)
	return bq
}

// Limit the number of records to be returned by this query.
func (bq *BeaconQuery) Limit(limit int) *BeaconQuery {
	bq.ctx.Limit = &limit
	return bq
}

// Offset to start from.
func (bq *BeaconQuery) Offset(offset int) *BeaconQuery {
	bq.ctx.Offset = &offset
	return bq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (bq *BeaconQuery) Unique(unique bool) *BeaconQuery {
	bq.ctx.Unique = &unique
	return bq
}

// Order specifies how the records should be ordered.
func (bq *BeaconQuery) Order(o ...beacon.OrderOption) *BeaconQuery {
	bq.order = append(bq.order, o...)
	return bq
}

// QueryHost chains the current query on the "host" edge.
func (bq *BeaconQuery) QueryHost() *HostQuery {
	query := (&HostClient{config: bq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := bq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := bq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(beacon.Table, beacon.FieldID, selector),
			sqlgraph.To(host.Table, host.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, beacon.HostTable, beacon.HostColumn),
		)
		fromU = sqlgraph.SetNeighbors(bq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryTasks chains the current query on the "tasks" edge.
func (bq *BeaconQuery) QueryTasks() *TaskQuery {
	query := (&TaskClient{config: bq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := bq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := bq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(beacon.Table, beacon.FieldID, selector),
			sqlgraph.To(task.Table, task.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, beacon.TasksTable, beacon.TasksColumn),
		)
		fromU = sqlgraph.SetNeighbors(bq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Beacon entity from the query.
// Returns a *NotFoundError when no Beacon was found.
func (bq *BeaconQuery) First(ctx context.Context) (*Beacon, error) {
	nodes, err := bq.Limit(1).All(setContextOp(ctx, bq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{beacon.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (bq *BeaconQuery) FirstX(ctx context.Context) *Beacon {
	node, err := bq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Beacon ID from the query.
// Returns a *NotFoundError when no Beacon ID was found.
func (bq *BeaconQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = bq.Limit(1).IDs(setContextOp(ctx, bq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{beacon.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (bq *BeaconQuery) FirstIDX(ctx context.Context) int {
	id, err := bq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Beacon entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Beacon entity is found.
// Returns a *NotFoundError when no Beacon entities are found.
func (bq *BeaconQuery) Only(ctx context.Context) (*Beacon, error) {
	nodes, err := bq.Limit(2).All(setContextOp(ctx, bq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{beacon.Label}
	default:
		return nil, &NotSingularError{beacon.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (bq *BeaconQuery) OnlyX(ctx context.Context) *Beacon {
	node, err := bq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Beacon ID in the query.
// Returns a *NotSingularError when more than one Beacon ID is found.
// Returns a *NotFoundError when no entities are found.
func (bq *BeaconQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = bq.Limit(2).IDs(setContextOp(ctx, bq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{beacon.Label}
	default:
		err = &NotSingularError{beacon.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (bq *BeaconQuery) OnlyIDX(ctx context.Context) int {
	id, err := bq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Beacons.
func (bq *BeaconQuery) All(ctx context.Context) ([]*Beacon, error) {
	ctx = setContextOp(ctx, bq.ctx, "All")
	if err := bq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Beacon, *BeaconQuery]()
	return withInterceptors[[]*Beacon](ctx, bq, qr, bq.inters)
}

// AllX is like All, but panics if an error occurs.
func (bq *BeaconQuery) AllX(ctx context.Context) []*Beacon {
	nodes, err := bq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Beacon IDs.
func (bq *BeaconQuery) IDs(ctx context.Context) (ids []int, err error) {
	if bq.ctx.Unique == nil && bq.path != nil {
		bq.Unique(true)
	}
	ctx = setContextOp(ctx, bq.ctx, "IDs")
	if err = bq.Select(beacon.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (bq *BeaconQuery) IDsX(ctx context.Context) []int {
	ids, err := bq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (bq *BeaconQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, bq.ctx, "Count")
	if err := bq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, bq, querierCount[*BeaconQuery](), bq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (bq *BeaconQuery) CountX(ctx context.Context) int {
	count, err := bq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (bq *BeaconQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, bq.ctx, "Exist")
	switch _, err := bq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (bq *BeaconQuery) ExistX(ctx context.Context) bool {
	exist, err := bq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the BeaconQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (bq *BeaconQuery) Clone() *BeaconQuery {
	if bq == nil {
		return nil
	}
	return &BeaconQuery{
		config:     bq.config,
		ctx:        bq.ctx.Clone(),
		order:      append([]beacon.OrderOption{}, bq.order...),
		inters:     append([]Interceptor{}, bq.inters...),
		predicates: append([]predicate.Beacon{}, bq.predicates...),
		withHost:   bq.withHost.Clone(),
		withTasks:  bq.withTasks.Clone(),
		// clone intermediate query.
		sql:  bq.sql.Clone(),
		path: bq.path,
	}
}

// WithHost tells the query-builder to eager-load the nodes that are connected to
// the "host" edge. The optional arguments are used to configure the query builder of the edge.
func (bq *BeaconQuery) WithHost(opts ...func(*HostQuery)) *BeaconQuery {
	query := (&HostClient{config: bq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	bq.withHost = query
	return bq
}

// WithTasks tells the query-builder to eager-load the nodes that are connected to
// the "tasks" edge. The optional arguments are used to configure the query builder of the edge.
func (bq *BeaconQuery) WithTasks(opts ...func(*TaskQuery)) *BeaconQuery {
	query := (&TaskClient{config: bq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	bq.withTasks = query
	return bq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Beacon.Query().
//		GroupBy(beacon.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (bq *BeaconQuery) GroupBy(field string, fields ...string) *BeaconGroupBy {
	bq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &BeaconGroupBy{build: bq}
	grbuild.flds = &bq.ctx.Fields
	grbuild.label = beacon.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.Beacon.Query().
//		Select(beacon.FieldName).
//		Scan(ctx, &v)
func (bq *BeaconQuery) Select(fields ...string) *BeaconSelect {
	bq.ctx.Fields = append(bq.ctx.Fields, fields...)
	sbuild := &BeaconSelect{BeaconQuery: bq}
	sbuild.label = beacon.Label
	sbuild.flds, sbuild.scan = &bq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a BeaconSelect configured with the given aggregations.
func (bq *BeaconQuery) Aggregate(fns ...AggregateFunc) *BeaconSelect {
	return bq.Select().Aggregate(fns...)
}

func (bq *BeaconQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range bq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, bq); err != nil {
				return err
			}
		}
	}
	for _, f := range bq.ctx.Fields {
		if !beacon.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if bq.path != nil {
		prev, err := bq.path(ctx)
		if err != nil {
			return err
		}
		bq.sql = prev
	}
	return nil
}

func (bq *BeaconQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Beacon, error) {
	var (
		nodes       = []*Beacon{}
		withFKs     = bq.withFKs
		_spec       = bq.querySpec()
		loadedTypes = [2]bool{
			bq.withHost != nil,
			bq.withTasks != nil,
		}
	)
	if bq.withHost != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, beacon.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Beacon).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Beacon{config: bq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(bq.modifiers) > 0 {
		_spec.Modifiers = bq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, bq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := bq.withHost; query != nil {
		if err := bq.loadHost(ctx, query, nodes, nil,
			func(n *Beacon, e *Host) { n.Edges.Host = e }); err != nil {
			return nil, err
		}
	}
	if query := bq.withTasks; query != nil {
		if err := bq.loadTasks(ctx, query, nodes,
			func(n *Beacon) { n.Edges.Tasks = []*Task{} },
			func(n *Beacon, e *Task) { n.Edges.Tasks = append(n.Edges.Tasks, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range bq.withNamedTasks {
		if err := bq.loadTasks(ctx, query, nodes,
			func(n *Beacon) { n.appendNamedTasks(name) },
			func(n *Beacon, e *Task) { n.appendNamedTasks(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range bq.loadTotal {
		if err := bq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (bq *BeaconQuery) loadHost(ctx context.Context, query *HostQuery, nodes []*Beacon, init func(*Beacon), assign func(*Beacon, *Host)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Beacon)
	for i := range nodes {
		if nodes[i].beacon_host == nil {
			continue
		}
		fk := *nodes[i].beacon_host
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(host.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "beacon_host" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (bq *BeaconQuery) loadTasks(ctx context.Context, query *TaskQuery, nodes []*Beacon, init func(*Beacon), assign func(*Beacon, *Task)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Beacon)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Task(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(beacon.TasksColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.task_beacon
		if fk == nil {
			return fmt.Errorf(`foreign-key "task_beacon" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "task_beacon" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (bq *BeaconQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := bq.querySpec()
	if len(bq.modifiers) > 0 {
		_spec.Modifiers = bq.modifiers
	}
	_spec.Node.Columns = bq.ctx.Fields
	if len(bq.ctx.Fields) > 0 {
		_spec.Unique = bq.ctx.Unique != nil && *bq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, bq.driver, _spec)
}

func (bq *BeaconQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(beacon.Table, beacon.Columns, sqlgraph.NewFieldSpec(beacon.FieldID, field.TypeInt))
	_spec.From = bq.sql
	if unique := bq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if bq.path != nil {
		_spec.Unique = true
	}
	if fields := bq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, beacon.FieldID)
		for i := range fields {
			if fields[i] != beacon.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := bq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := bq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := bq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := bq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (bq *BeaconQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(bq.driver.Dialect())
	t1 := builder.Table(beacon.Table)
	columns := bq.ctx.Fields
	if len(columns) == 0 {
		columns = beacon.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if bq.sql != nil {
		selector = bq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if bq.ctx.Unique != nil && *bq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range bq.predicates {
		p(selector)
	}
	for _, p := range bq.order {
		p(selector)
	}
	if offset := bq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := bq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WithNamedTasks tells the query-builder to eager-load the nodes that are connected to the "tasks"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (bq *BeaconQuery) WithNamedTasks(name string, opts ...func(*TaskQuery)) *BeaconQuery {
	query := (&TaskClient{config: bq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if bq.withNamedTasks == nil {
		bq.withNamedTasks = make(map[string]*TaskQuery)
	}
	bq.withNamedTasks[name] = query
	return bq
}

// BeaconGroupBy is the group-by builder for Beacon entities.
type BeaconGroupBy struct {
	selector
	build *BeaconQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (bgb *BeaconGroupBy) Aggregate(fns ...AggregateFunc) *BeaconGroupBy {
	bgb.fns = append(bgb.fns, fns...)
	return bgb
}

// Scan applies the selector query and scans the result into the given value.
func (bgb *BeaconGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, bgb.build.ctx, "GroupBy")
	if err := bgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*BeaconQuery, *BeaconGroupBy](ctx, bgb.build, bgb, bgb.build.inters, v)
}

func (bgb *BeaconGroupBy) sqlScan(ctx context.Context, root *BeaconQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(bgb.fns))
	for _, fn := range bgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*bgb.flds)+len(bgb.fns))
		for _, f := range *bgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*bgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := bgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// BeaconSelect is the builder for selecting fields of Beacon entities.
type BeaconSelect struct {
	*BeaconQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (bs *BeaconSelect) Aggregate(fns ...AggregateFunc) *BeaconSelect {
	bs.fns = append(bs.fns, fns...)
	return bs
}

// Scan applies the selector query and scans the result into the given value.
func (bs *BeaconSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, bs.ctx, "Select")
	if err := bs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*BeaconQuery, *BeaconSelect](ctx, bs.BeaconQuery, bs, bs.inters, v)
}

func (bs *BeaconSelect) sqlScan(ctx context.Context, root *BeaconQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(bs.fns))
	for _, fn := range bs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*bs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := bs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}