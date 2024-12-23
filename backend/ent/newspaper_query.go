// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/jijimama/newspaper-app/ent/column"
	"github.com/jijimama/newspaper-app/ent/newspaper"
	"github.com/jijimama/newspaper-app/ent/predicate"
)

// NewspaperQuery is the builder for querying Newspaper entities.
type NewspaperQuery struct {
	config
	ctx         *QueryContext
	order       []newspaper.OrderOption
	inters      []Interceptor
	predicates  []predicate.Newspaper
	withColumns *ColumnQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the NewspaperQuery builder.
func (nq *NewspaperQuery) Where(ps ...predicate.Newspaper) *NewspaperQuery {
	nq.predicates = append(nq.predicates, ps...)
	return nq
}

// Limit the number of records to be returned by this query.
func (nq *NewspaperQuery) Limit(limit int) *NewspaperQuery {
	nq.ctx.Limit = &limit
	return nq
}

// Offset to start from.
func (nq *NewspaperQuery) Offset(offset int) *NewspaperQuery {
	nq.ctx.Offset = &offset
	return nq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (nq *NewspaperQuery) Unique(unique bool) *NewspaperQuery {
	nq.ctx.Unique = &unique
	return nq
}

// Order specifies how the records should be ordered.
func (nq *NewspaperQuery) Order(o ...newspaper.OrderOption) *NewspaperQuery {
	nq.order = append(nq.order, o...)
	return nq
}

// QueryColumns chains the current query on the "columns" edge.
func (nq *NewspaperQuery) QueryColumns() *ColumnQuery {
	query := (&ColumnClient{config: nq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := nq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := nq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(newspaper.Table, newspaper.FieldID, selector),
			sqlgraph.To(column.Table, column.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, newspaper.ColumnsTable, newspaper.ColumnsColumn),
		)
		fromU = sqlgraph.SetNeighbors(nq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Newspaper entity from the query.
// Returns a *NotFoundError when no Newspaper was found.
func (nq *NewspaperQuery) First(ctx context.Context) (*Newspaper, error) {
	nodes, err := nq.Limit(1).All(setContextOp(ctx, nq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{newspaper.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (nq *NewspaperQuery) FirstX(ctx context.Context) *Newspaper {
	node, err := nq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Newspaper ID from the query.
// Returns a *NotFoundError when no Newspaper ID was found.
func (nq *NewspaperQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = nq.Limit(1).IDs(setContextOp(ctx, nq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{newspaper.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (nq *NewspaperQuery) FirstIDX(ctx context.Context) int {
	id, err := nq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Newspaper entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Newspaper entity is found.
// Returns a *NotFoundError when no Newspaper entities are found.
func (nq *NewspaperQuery) Only(ctx context.Context) (*Newspaper, error) {
	nodes, err := nq.Limit(2).All(setContextOp(ctx, nq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{newspaper.Label}
	default:
		return nil, &NotSingularError{newspaper.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (nq *NewspaperQuery) OnlyX(ctx context.Context) *Newspaper {
	node, err := nq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Newspaper ID in the query.
// Returns a *NotSingularError when more than one Newspaper ID is found.
// Returns a *NotFoundError when no entities are found.
func (nq *NewspaperQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = nq.Limit(2).IDs(setContextOp(ctx, nq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{newspaper.Label}
	default:
		err = &NotSingularError{newspaper.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (nq *NewspaperQuery) OnlyIDX(ctx context.Context) int {
	id, err := nq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Newspapers.
func (nq *NewspaperQuery) All(ctx context.Context) ([]*Newspaper, error) {
	ctx = setContextOp(ctx, nq.ctx, ent.OpQueryAll)
	if err := nq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Newspaper, *NewspaperQuery]()
	return withInterceptors[[]*Newspaper](ctx, nq, qr, nq.inters)
}

// AllX is like All, but panics if an error occurs.
func (nq *NewspaperQuery) AllX(ctx context.Context) []*Newspaper {
	nodes, err := nq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Newspaper IDs.
func (nq *NewspaperQuery) IDs(ctx context.Context) (ids []int, err error) {
	if nq.ctx.Unique == nil && nq.path != nil {
		nq.Unique(true)
	}
	ctx = setContextOp(ctx, nq.ctx, ent.OpQueryIDs)
	if err = nq.Select(newspaper.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (nq *NewspaperQuery) IDsX(ctx context.Context) []int {
	ids, err := nq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (nq *NewspaperQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, nq.ctx, ent.OpQueryCount)
	if err := nq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, nq, querierCount[*NewspaperQuery](), nq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (nq *NewspaperQuery) CountX(ctx context.Context) int {
	count, err := nq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (nq *NewspaperQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, nq.ctx, ent.OpQueryExist)
	switch _, err := nq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (nq *NewspaperQuery) ExistX(ctx context.Context) bool {
	exist, err := nq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the NewspaperQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (nq *NewspaperQuery) Clone() *NewspaperQuery {
	if nq == nil {
		return nil
	}
	return &NewspaperQuery{
		config:      nq.config,
		ctx:         nq.ctx.Clone(),
		order:       append([]newspaper.OrderOption{}, nq.order...),
		inters:      append([]Interceptor{}, nq.inters...),
		predicates:  append([]predicate.Newspaper{}, nq.predicates...),
		withColumns: nq.withColumns.Clone(),
		// clone intermediate query.
		sql:  nq.sql.Clone(),
		path: nq.path,
	}
}

// WithColumns tells the query-builder to eager-load the nodes that are connected to
// the "columns" edge. The optional arguments are used to configure the query builder of the edge.
func (nq *NewspaperQuery) WithColumns(opts ...func(*ColumnQuery)) *NewspaperQuery {
	query := (&ColumnClient{config: nq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	nq.withColumns = query
	return nq
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
//	client.Newspaper.Query().
//		GroupBy(newspaper.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (nq *NewspaperQuery) GroupBy(field string, fields ...string) *NewspaperGroupBy {
	nq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &NewspaperGroupBy{build: nq}
	grbuild.flds = &nq.ctx.Fields
	grbuild.label = newspaper.Label
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
//	client.Newspaper.Query().
//		Select(newspaper.FieldName).
//		Scan(ctx, &v)
func (nq *NewspaperQuery) Select(fields ...string) *NewspaperSelect {
	nq.ctx.Fields = append(nq.ctx.Fields, fields...)
	sbuild := &NewspaperSelect{NewspaperQuery: nq}
	sbuild.label = newspaper.Label
	sbuild.flds, sbuild.scan = &nq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a NewspaperSelect configured with the given aggregations.
func (nq *NewspaperQuery) Aggregate(fns ...AggregateFunc) *NewspaperSelect {
	return nq.Select().Aggregate(fns...)
}

func (nq *NewspaperQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range nq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, nq); err != nil {
				return err
			}
		}
	}
	for _, f := range nq.ctx.Fields {
		if !newspaper.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if nq.path != nil {
		prev, err := nq.path(ctx)
		if err != nil {
			return err
		}
		nq.sql = prev
	}
	return nil
}

func (nq *NewspaperQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Newspaper, error) {
	var (
		nodes       = []*Newspaper{}
		_spec       = nq.querySpec()
		loadedTypes = [1]bool{
			nq.withColumns != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Newspaper).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Newspaper{config: nq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, nq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := nq.withColumns; query != nil {
		if err := nq.loadColumns(ctx, query, nodes,
			func(n *Newspaper) { n.Edges.Columns = []*Column{} },
			func(n *Newspaper, e *Column) { n.Edges.Columns = append(n.Edges.Columns, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (nq *NewspaperQuery) loadColumns(ctx context.Context, query *ColumnQuery, nodes []*Newspaper, init func(*Newspaper), assign func(*Newspaper, *Column)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Newspaper)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Column(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(newspaper.ColumnsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.newspaper_columns
		if fk == nil {
			return fmt.Errorf(`foreign-key "newspaper_columns" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "newspaper_columns" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (nq *NewspaperQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := nq.querySpec()
	_spec.Node.Columns = nq.ctx.Fields
	if len(nq.ctx.Fields) > 0 {
		_spec.Unique = nq.ctx.Unique != nil && *nq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, nq.driver, _spec)
}

func (nq *NewspaperQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(newspaper.Table, newspaper.Columns, sqlgraph.NewFieldSpec(newspaper.FieldID, field.TypeInt))
	_spec.From = nq.sql
	if unique := nq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if nq.path != nil {
		_spec.Unique = true
	}
	if fields := nq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, newspaper.FieldID)
		for i := range fields {
			if fields[i] != newspaper.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := nq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := nq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := nq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := nq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (nq *NewspaperQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(nq.driver.Dialect())
	t1 := builder.Table(newspaper.Table)
	columns := nq.ctx.Fields
	if len(columns) == 0 {
		columns = newspaper.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if nq.sql != nil {
		selector = nq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if nq.ctx.Unique != nil && *nq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range nq.predicates {
		p(selector)
	}
	for _, p := range nq.order {
		p(selector)
	}
	if offset := nq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := nq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// NewspaperGroupBy is the group-by builder for Newspaper entities.
type NewspaperGroupBy struct {
	selector
	build *NewspaperQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ngb *NewspaperGroupBy) Aggregate(fns ...AggregateFunc) *NewspaperGroupBy {
	ngb.fns = append(ngb.fns, fns...)
	return ngb
}

// Scan applies the selector query and scans the result into the given value.
func (ngb *NewspaperGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ngb.build.ctx, ent.OpQueryGroupBy)
	if err := ngb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*NewspaperQuery, *NewspaperGroupBy](ctx, ngb.build, ngb, ngb.build.inters, v)
}

func (ngb *NewspaperGroupBy) sqlScan(ctx context.Context, root *NewspaperQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ngb.fns))
	for _, fn := range ngb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ngb.flds)+len(ngb.fns))
		for _, f := range *ngb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ngb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ngb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// NewspaperSelect is the builder for selecting fields of Newspaper entities.
type NewspaperSelect struct {
	*NewspaperQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ns *NewspaperSelect) Aggregate(fns ...AggregateFunc) *NewspaperSelect {
	ns.fns = append(ns.fns, fns...)
	return ns
}

// Scan applies the selector query and scans the result into the given value.
func (ns *NewspaperSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ns.ctx, ent.OpQuerySelect)
	if err := ns.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*NewspaperQuery, *NewspaperSelect](ctx, ns.NewspaperQuery, ns, ns.inters, v)
}

func (ns *NewspaperSelect) sqlScan(ctx context.Context, root *NewspaperQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ns.fns))
	for _, fn := range ns.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ns.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ns.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}