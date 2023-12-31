// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"{{ cookiecutter.project_name }}/data/ent/operatelog"
	"{{ cookiecutter.project_name }}/data/ent/predicate"
	"{{ cookiecutter.project_name }}/data/ent/user"
	"math"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// OperateLogQuery is the builder for querying OperateLog entities.
type OperateLogQuery struct {
	config
	ctx        *QueryContext
	order      []OrderFunc
	inters     []Interceptor
	predicates []predicate.OperateLog
	withUser   *UserQuery
	withFKs    bool
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the OperateLogQuery builder.
func (olq *OperateLogQuery) Where(ps ...predicate.OperateLog) *OperateLogQuery {
	olq.predicates = append(olq.predicates, ps...)
	return olq
}

// Limit the number of records to be returned by this query.
func (olq *OperateLogQuery) Limit(limit int) *OperateLogQuery {
	olq.ctx.Limit = &limit
	return olq
}

// Offset to start from.
func (olq *OperateLogQuery) Offset(offset int) *OperateLogQuery {
	olq.ctx.Offset = &offset
	return olq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (olq *OperateLogQuery) Unique(unique bool) *OperateLogQuery {
	olq.ctx.Unique = &unique
	return olq
}

// Order specifies how the records should be ordered.
func (olq *OperateLogQuery) Order(o ...OrderFunc) *OperateLogQuery {
	olq.order = append(olq.order, o...)
	return olq
}

// QueryUser chains the current query on the "user" edge.
func (olq *OperateLogQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: olq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := olq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := olq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(operatelog.Table, operatelog.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, operatelog.UserTable, operatelog.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(olq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first OperateLog entity from the query.
// Returns a *NotFoundError when no OperateLog was found.
func (olq *OperateLogQuery) First(ctx context.Context) (*OperateLog, error) {
	nodes, err := olq.Limit(1).All(setContextOp(ctx, olq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{operatelog.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (olq *OperateLogQuery) FirstX(ctx context.Context) *OperateLog {
	node, err := olq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first OperateLog ID from the query.
// Returns a *NotFoundError when no OperateLog ID was found.
func (olq *OperateLogQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = olq.Limit(1).IDs(setContextOp(ctx, olq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{operatelog.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (olq *OperateLogQuery) FirstIDX(ctx context.Context) int {
	id, err := olq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single OperateLog entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one OperateLog entity is found.
// Returns a *NotFoundError when no OperateLog entities are found.
func (olq *OperateLogQuery) Only(ctx context.Context) (*OperateLog, error) {
	nodes, err := olq.Limit(2).All(setContextOp(ctx, olq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{operatelog.Label}
	default:
		return nil, &NotSingularError{operatelog.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (olq *OperateLogQuery) OnlyX(ctx context.Context) *OperateLog {
	node, err := olq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only OperateLog ID in the query.
// Returns a *NotSingularError when more than one OperateLog ID is found.
// Returns a *NotFoundError when no entities are found.
func (olq *OperateLogQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = olq.Limit(2).IDs(setContextOp(ctx, olq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{operatelog.Label}
	default:
		err = &NotSingularError{operatelog.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (olq *OperateLogQuery) OnlyIDX(ctx context.Context) int {
	id, err := olq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of OperateLogs.
func (olq *OperateLogQuery) All(ctx context.Context) ([]*OperateLog, error) {
	ctx = setContextOp(ctx, olq.ctx, "All")
	if err := olq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*OperateLog, *OperateLogQuery]()
	return withInterceptors[[]*OperateLog](ctx, olq, qr, olq.inters)
}

// AllX is like All, but panics if an error occurs.
func (olq *OperateLogQuery) AllX(ctx context.Context) []*OperateLog {
	nodes, err := olq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of OperateLog IDs.
func (olq *OperateLogQuery) IDs(ctx context.Context) (ids []int, err error) {
	if olq.ctx.Unique == nil && olq.path != nil {
		olq.Unique(true)
	}
	ctx = setContextOp(ctx, olq.ctx, "IDs")
	if err = olq.Select(operatelog.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (olq *OperateLogQuery) IDsX(ctx context.Context) []int {
	ids, err := olq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (olq *OperateLogQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, olq.ctx, "Count")
	if err := olq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, olq, querierCount[*OperateLogQuery](), olq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (olq *OperateLogQuery) CountX(ctx context.Context) int {
	count, err := olq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (olq *OperateLogQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, olq.ctx, "Exist")
	switch _, err := olq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (olq *OperateLogQuery) ExistX(ctx context.Context) bool {
	exist, err := olq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the OperateLogQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (olq *OperateLogQuery) Clone() *OperateLogQuery {
	if olq == nil {
		return nil
	}
	return &OperateLogQuery{
		config:     olq.config,
		ctx:        olq.ctx.Clone(),
		order:      append([]OrderFunc{}, olq.order...),
		inters:     append([]Interceptor{}, olq.inters...),
		predicates: append([]predicate.OperateLog{}, olq.predicates...),
		withUser:   olq.withUser.Clone(),
		// clone intermediate query.
		sql:  olq.sql.Clone(),
		path: olq.path,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (olq *OperateLogQuery) WithUser(opts ...func(*UserQuery)) *OperateLogQuery {
	query := (&UserClient{config: olq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	olq.withUser = query
	return olq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.OperateLog.Query().
//		GroupBy(operatelog.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (olq *OperateLogQuery) GroupBy(field string, fields ...string) *OperateLogGroupBy {
	olq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &OperateLogGroupBy{build: olq}
	grbuild.flds = &olq.ctx.Fields
	grbuild.label = operatelog.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.OperateLog.Query().
//		Select(operatelog.FieldCreatedAt).
//		Scan(ctx, &v)
func (olq *OperateLogQuery) Select(fields ...string) *OperateLogSelect {
	olq.ctx.Fields = append(olq.ctx.Fields, fields...)
	sbuild := &OperateLogSelect{OperateLogQuery: olq}
	sbuild.label = operatelog.Label
	sbuild.flds, sbuild.scan = &olq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a OperateLogSelect configured with the given aggregations.
func (olq *OperateLogQuery) Aggregate(fns ...AggregateFunc) *OperateLogSelect {
	return olq.Select().Aggregate(fns...)
}

func (olq *OperateLogQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range olq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, olq); err != nil {
				return err
			}
		}
	}
	for _, f := range olq.ctx.Fields {
		if !operatelog.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if olq.path != nil {
		prev, err := olq.path(ctx)
		if err != nil {
			return err
		}
		olq.sql = prev
	}
	return nil
}

func (olq *OperateLogQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*OperateLog, error) {
	var (
		nodes       = []*OperateLog{}
		withFKs     = olq.withFKs
		_spec       = olq.querySpec()
		loadedTypes = [1]bool{
			olq.withUser != nil,
		}
	)
	if olq.withUser != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, operatelog.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*OperateLog).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &OperateLog{config: olq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(olq.modifiers) > 0 {
		_spec.Modifiers = olq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, olq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := olq.withUser; query != nil {
		if err := olq.loadUser(ctx, query, nodes, nil,
			func(n *OperateLog, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (olq *OperateLogQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*OperateLog, init func(*OperateLog), assign func(*OperateLog, *User)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*OperateLog)
	for i := range nodes {
		if nodes[i].user_logs == nil {
			continue
		}
		fk := *nodes[i].user_logs
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "user_logs" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (olq *OperateLogQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := olq.querySpec()
	if len(olq.modifiers) > 0 {
		_spec.Modifiers = olq.modifiers
	}
	_spec.Node.Columns = olq.ctx.Fields
	if len(olq.ctx.Fields) > 0 {
		_spec.Unique = olq.ctx.Unique != nil && *olq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, olq.driver, _spec)
}

func (olq *OperateLogQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(operatelog.Table, operatelog.Columns, sqlgraph.NewFieldSpec(operatelog.FieldID, field.TypeInt))
	_spec.From = olq.sql
	if unique := olq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if olq.path != nil {
		_spec.Unique = true
	}
	if fields := olq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, operatelog.FieldID)
		for i := range fields {
			if fields[i] != operatelog.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := olq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := olq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := olq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := olq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (olq *OperateLogQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(olq.driver.Dialect())
	t1 := builder.Table(operatelog.Table)
	columns := olq.ctx.Fields
	if len(columns) == 0 {
		columns = operatelog.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if olq.sql != nil {
		selector = olq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if olq.ctx.Unique != nil && *olq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range olq.modifiers {
		m(selector)
	}
	for _, p := range olq.predicates {
		p(selector)
	}
	for _, p := range olq.order {
		p(selector)
	}
	if offset := olq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := olq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (olq *OperateLogQuery) ForUpdate(opts ...sql.LockOption) *OperateLogQuery {
	if olq.driver.Dialect() == dialect.Postgres {
		olq.Unique(false)
	}
	olq.modifiers = append(olq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return olq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (olq *OperateLogQuery) ForShare(opts ...sql.LockOption) *OperateLogQuery {
	if olq.driver.Dialect() == dialect.Postgres {
		olq.Unique(false)
	}
	olq.modifiers = append(olq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return olq
}

// OperateLogGroupBy is the group-by builder for OperateLog entities.
type OperateLogGroupBy struct {
	selector
	build *OperateLogQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (olgb *OperateLogGroupBy) Aggregate(fns ...AggregateFunc) *OperateLogGroupBy {
	olgb.fns = append(olgb.fns, fns...)
	return olgb
}

// Scan applies the selector query and scans the result into the given value.
func (olgb *OperateLogGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, olgb.build.ctx, "GroupBy")
	if err := olgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*OperateLogQuery, *OperateLogGroupBy](ctx, olgb.build, olgb, olgb.build.inters, v)
}

func (olgb *OperateLogGroupBy) sqlScan(ctx context.Context, root *OperateLogQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(olgb.fns))
	for _, fn := range olgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*olgb.flds)+len(olgb.fns))
		for _, f := range *olgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*olgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := olgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// OperateLogSelect is the builder for selecting fields of OperateLog entities.
type OperateLogSelect struct {
	*OperateLogQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ols *OperateLogSelect) Aggregate(fns ...AggregateFunc) *OperateLogSelect {
	ols.fns = append(ols.fns, fns...)
	return ols
}

// Scan applies the selector query and scans the result into the given value.
func (ols *OperateLogSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ols.ctx, "Select")
	if err := ols.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*OperateLogQuery, *OperateLogSelect](ctx, ols.OperateLogQuery, ols, ols.inters, v)
}

func (ols *OperateLogSelect) sqlScan(ctx context.Context, root *OperateLogQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ols.fns))
	for _, fn := range ols.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ols.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ols.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
