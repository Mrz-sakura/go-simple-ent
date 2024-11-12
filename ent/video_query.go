// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"dwbackend/ent/predicate"
	"dwbackend/ent/video"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// VideoQuery is the builder for querying Video entities.
type VideoQuery struct {
	config
	ctx        *QueryContext
	order      []video.OrderOption
	inters     []Interceptor
	predicates []predicate.Video
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the VideoQuery builder.
func (vq *VideoQuery) Where(ps ...predicate.Video) *VideoQuery {
	vq.predicates = append(vq.predicates, ps...)
	return vq
}

// Limit the number of records to be returned by this query.
func (vq *VideoQuery) Limit(limit int) *VideoQuery {
	vq.ctx.Limit = &limit
	return vq
}

// Offset to start from.
func (vq *VideoQuery) Offset(offset int) *VideoQuery {
	vq.ctx.Offset = &offset
	return vq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (vq *VideoQuery) Unique(unique bool) *VideoQuery {
	vq.ctx.Unique = &unique
	return vq
}

// Order specifies how the records should be ordered.
func (vq *VideoQuery) Order(o ...video.OrderOption) *VideoQuery {
	vq.order = append(vq.order, o...)
	return vq
}

// First returns the first Video entity from the query.
// Returns a *NotFoundError when no Video was found.
func (vq *VideoQuery) First(ctx context.Context) (*Video, error) {
	nodes, err := vq.Limit(1).All(setContextOp(ctx, vq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{video.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (vq *VideoQuery) FirstX(ctx context.Context) *Video {
	node, err := vq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Video ID from the query.
// Returns a *NotFoundError when no Video ID was found.
func (vq *VideoQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = vq.Limit(1).IDs(setContextOp(ctx, vq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{video.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (vq *VideoQuery) FirstIDX(ctx context.Context) int {
	id, err := vq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Video entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Video entity is found.
// Returns a *NotFoundError when no Video entities are found.
func (vq *VideoQuery) Only(ctx context.Context) (*Video, error) {
	nodes, err := vq.Limit(2).All(setContextOp(ctx, vq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{video.Label}
	default:
		return nil, &NotSingularError{video.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (vq *VideoQuery) OnlyX(ctx context.Context) *Video {
	node, err := vq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Video ID in the query.
// Returns a *NotSingularError when more than one Video ID is found.
// Returns a *NotFoundError when no entities are found.
func (vq *VideoQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = vq.Limit(2).IDs(setContextOp(ctx, vq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{video.Label}
	default:
		err = &NotSingularError{video.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (vq *VideoQuery) OnlyIDX(ctx context.Context) int {
	id, err := vq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Videos.
func (vq *VideoQuery) All(ctx context.Context) ([]*Video, error) {
	ctx = setContextOp(ctx, vq.ctx, ent.OpQueryAll)
	if err := vq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Video, *VideoQuery]()
	return withInterceptors[[]*Video](ctx, vq, qr, vq.inters)
}

// AllX is like All, but panics if an error occurs.
func (vq *VideoQuery) AllX(ctx context.Context) []*Video {
	nodes, err := vq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Video IDs.
func (vq *VideoQuery) IDs(ctx context.Context) (ids []int, err error) {
	if vq.ctx.Unique == nil && vq.path != nil {
		vq.Unique(true)
	}
	ctx = setContextOp(ctx, vq.ctx, ent.OpQueryIDs)
	if err = vq.Select(video.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (vq *VideoQuery) IDsX(ctx context.Context) []int {
	ids, err := vq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (vq *VideoQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, vq.ctx, ent.OpQueryCount)
	if err := vq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, vq, querierCount[*VideoQuery](), vq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (vq *VideoQuery) CountX(ctx context.Context) int {
	count, err := vq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (vq *VideoQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, vq.ctx, ent.OpQueryExist)
	switch _, err := vq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (vq *VideoQuery) ExistX(ctx context.Context) bool {
	exist, err := vq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the VideoQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (vq *VideoQuery) Clone() *VideoQuery {
	if vq == nil {
		return nil
	}
	return &VideoQuery{
		config:     vq.config,
		ctx:        vq.ctx.Clone(),
		order:      append([]video.OrderOption{}, vq.order...),
		inters:     append([]Interceptor{}, vq.inters...),
		predicates: append([]predicate.Video{}, vq.predicates...),
		// clone intermediate query.
		sql:  vq.sql.Clone(),
		path: vq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Title string `json:"title,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Video.Query().
//		GroupBy(video.FieldTitle).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (vq *VideoQuery) GroupBy(field string, fields ...string) *VideoGroupBy {
	vq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &VideoGroupBy{build: vq}
	grbuild.flds = &vq.ctx.Fields
	grbuild.label = video.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Title string `json:"title,omitempty"`
//	}
//
//	client.Video.Query().
//		Select(video.FieldTitle).
//		Scan(ctx, &v)
func (vq *VideoQuery) Select(fields ...string) *VideoSelect {
	vq.ctx.Fields = append(vq.ctx.Fields, fields...)
	sbuild := &VideoSelect{VideoQuery: vq}
	sbuild.label = video.Label
	sbuild.flds, sbuild.scan = &vq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a VideoSelect configured with the given aggregations.
func (vq *VideoQuery) Aggregate(fns ...AggregateFunc) *VideoSelect {
	return vq.Select().Aggregate(fns...)
}

func (vq *VideoQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range vq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, vq); err != nil {
				return err
			}
		}
	}
	for _, f := range vq.ctx.Fields {
		if !video.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if vq.path != nil {
		prev, err := vq.path(ctx)
		if err != nil {
			return err
		}
		vq.sql = prev
	}
	return nil
}

func (vq *VideoQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Video, error) {
	var (
		nodes = []*Video{}
		_spec = vq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Video).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Video{config: vq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, vq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (vq *VideoQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := vq.querySpec()
	_spec.Node.Columns = vq.ctx.Fields
	if len(vq.ctx.Fields) > 0 {
		_spec.Unique = vq.ctx.Unique != nil && *vq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, vq.driver, _spec)
}

func (vq *VideoQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(video.Table, video.Columns, sqlgraph.NewFieldSpec(video.FieldID, field.TypeInt))
	_spec.From = vq.sql
	if unique := vq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if vq.path != nil {
		_spec.Unique = true
	}
	if fields := vq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, video.FieldID)
		for i := range fields {
			if fields[i] != video.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := vq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := vq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := vq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := vq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (vq *VideoQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(vq.driver.Dialect())
	t1 := builder.Table(video.Table)
	columns := vq.ctx.Fields
	if len(columns) == 0 {
		columns = video.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if vq.sql != nil {
		selector = vq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if vq.ctx.Unique != nil && *vq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range vq.predicates {
		p(selector)
	}
	for _, p := range vq.order {
		p(selector)
	}
	if offset := vq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := vq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// VideoGroupBy is the group-by builder for Video entities.
type VideoGroupBy struct {
	selector
	build *VideoQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (vgb *VideoGroupBy) Aggregate(fns ...AggregateFunc) *VideoGroupBy {
	vgb.fns = append(vgb.fns, fns...)
	return vgb
}

// Scan applies the selector query and scans the result into the given value.
func (vgb *VideoGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, vgb.build.ctx, ent.OpQueryGroupBy)
	if err := vgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*VideoQuery, *VideoGroupBy](ctx, vgb.build, vgb, vgb.build.inters, v)
}

func (vgb *VideoGroupBy) sqlScan(ctx context.Context, root *VideoQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(vgb.fns))
	for _, fn := range vgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*vgb.flds)+len(vgb.fns))
		for _, f := range *vgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*vgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := vgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// VideoSelect is the builder for selecting fields of Video entities.
type VideoSelect struct {
	*VideoQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (vs *VideoSelect) Aggregate(fns ...AggregateFunc) *VideoSelect {
	vs.fns = append(vs.fns, fns...)
	return vs
}

// Scan applies the selector query and scans the result into the given value.
func (vs *VideoSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, vs.ctx, ent.OpQuerySelect)
	if err := vs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*VideoQuery, *VideoSelect](ctx, vs.VideoQuery, vs, vs.inters, v)
}

func (vs *VideoSelect) sqlScan(ctx context.Context, root *VideoQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(vs.fns))
	for _, fn := range vs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*vs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := vs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
