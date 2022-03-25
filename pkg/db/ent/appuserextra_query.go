// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent/appuserextra"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// AppUserExtraQuery is the builder for querying AppUserExtra entities.
type AppUserExtraQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.AppUserExtra
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AppUserExtraQuery builder.
func (aueq *AppUserExtraQuery) Where(ps ...predicate.AppUserExtra) *AppUserExtraQuery {
	aueq.predicates = append(aueq.predicates, ps...)
	return aueq
}

// Limit adds a limit step to the query.
func (aueq *AppUserExtraQuery) Limit(limit int) *AppUserExtraQuery {
	aueq.limit = &limit
	return aueq
}

// Offset adds an offset step to the query.
func (aueq *AppUserExtraQuery) Offset(offset int) *AppUserExtraQuery {
	aueq.offset = &offset
	return aueq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (aueq *AppUserExtraQuery) Unique(unique bool) *AppUserExtraQuery {
	aueq.unique = &unique
	return aueq
}

// Order adds an order step to the query.
func (aueq *AppUserExtraQuery) Order(o ...OrderFunc) *AppUserExtraQuery {
	aueq.order = append(aueq.order, o...)
	return aueq
}

// First returns the first AppUserExtra entity from the query.
// Returns a *NotFoundError when no AppUserExtra was found.
func (aueq *AppUserExtraQuery) First(ctx context.Context) (*AppUserExtra, error) {
	nodes, err := aueq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{appuserextra.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (aueq *AppUserExtraQuery) FirstX(ctx context.Context) *AppUserExtra {
	node, err := aueq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first AppUserExtra ID from the query.
// Returns a *NotFoundError when no AppUserExtra ID was found.
func (aueq *AppUserExtraQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = aueq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{appuserextra.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (aueq *AppUserExtraQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := aueq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single AppUserExtra entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one AppUserExtra entity is found.
// Returns a *NotFoundError when no AppUserExtra entities are found.
func (aueq *AppUserExtraQuery) Only(ctx context.Context) (*AppUserExtra, error) {
	nodes, err := aueq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{appuserextra.Label}
	default:
		return nil, &NotSingularError{appuserextra.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (aueq *AppUserExtraQuery) OnlyX(ctx context.Context) *AppUserExtra {
	node, err := aueq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only AppUserExtra ID in the query.
// Returns a *NotSingularError when more than one AppUserExtra ID is found.
// Returns a *NotFoundError when no entities are found.
func (aueq *AppUserExtraQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = aueq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{appuserextra.Label}
	default:
		err = &NotSingularError{appuserextra.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (aueq *AppUserExtraQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := aueq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of AppUserExtras.
func (aueq *AppUserExtraQuery) All(ctx context.Context) ([]*AppUserExtra, error) {
	if err := aueq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return aueq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (aueq *AppUserExtraQuery) AllX(ctx context.Context) []*AppUserExtra {
	nodes, err := aueq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of AppUserExtra IDs.
func (aueq *AppUserExtraQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := aueq.Select(appuserextra.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (aueq *AppUserExtraQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := aueq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (aueq *AppUserExtraQuery) Count(ctx context.Context) (int, error) {
	if err := aueq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return aueq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (aueq *AppUserExtraQuery) CountX(ctx context.Context) int {
	count, err := aueq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (aueq *AppUserExtraQuery) Exist(ctx context.Context) (bool, error) {
	if err := aueq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return aueq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (aueq *AppUserExtraQuery) ExistX(ctx context.Context) bool {
	exist, err := aueq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AppUserExtraQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (aueq *AppUserExtraQuery) Clone() *AppUserExtraQuery {
	if aueq == nil {
		return nil
	}
	return &AppUserExtraQuery{
		config:     aueq.config,
		limit:      aueq.limit,
		offset:     aueq.offset,
		order:      append([]OrderFunc{}, aueq.order...),
		predicates: append([]predicate.AppUserExtra{}, aueq.predicates...),
		// clone intermediate query.
		sql:    aueq.sql.Clone(),
		path:   aueq.path,
		unique: aueq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		AppID uuid.UUID `json:"app_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.AppUserExtra.Query().
//		GroupBy(appuserextra.FieldAppID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (aueq *AppUserExtraQuery) GroupBy(field string, fields ...string) *AppUserExtraGroupBy {
	group := &AppUserExtraGroupBy{config: aueq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := aueq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return aueq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		AppID uuid.UUID `json:"app_id,omitempty"`
//	}
//
//	client.AppUserExtra.Query().
//		Select(appuserextra.FieldAppID).
//		Scan(ctx, &v)
//
func (aueq *AppUserExtraQuery) Select(fields ...string) *AppUserExtraSelect {
	aueq.fields = append(aueq.fields, fields...)
	return &AppUserExtraSelect{AppUserExtraQuery: aueq}
}

func (aueq *AppUserExtraQuery) prepareQuery(ctx context.Context) error {
	for _, f := range aueq.fields {
		if !appuserextra.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if aueq.path != nil {
		prev, err := aueq.path(ctx)
		if err != nil {
			return err
		}
		aueq.sql = prev
	}
	return nil
}

func (aueq *AppUserExtraQuery) sqlAll(ctx context.Context) ([]*AppUserExtra, error) {
	var (
		nodes = []*AppUserExtra{}
		_spec = aueq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &AppUserExtra{config: aueq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, aueq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (aueq *AppUserExtraQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := aueq.querySpec()
	_spec.Node.Columns = aueq.fields
	if len(aueq.fields) > 0 {
		_spec.Unique = aueq.unique != nil && *aueq.unique
	}
	return sqlgraph.CountNodes(ctx, aueq.driver, _spec)
}

func (aueq *AppUserExtraQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := aueq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (aueq *AppUserExtraQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   appuserextra.Table,
			Columns: appuserextra.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: appuserextra.FieldID,
			},
		},
		From:   aueq.sql,
		Unique: true,
	}
	if unique := aueq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := aueq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, appuserextra.FieldID)
		for i := range fields {
			if fields[i] != appuserextra.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := aueq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := aueq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := aueq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := aueq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (aueq *AppUserExtraQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(aueq.driver.Dialect())
	t1 := builder.Table(appuserextra.Table)
	columns := aueq.fields
	if len(columns) == 0 {
		columns = appuserextra.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if aueq.sql != nil {
		selector = aueq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if aueq.unique != nil && *aueq.unique {
		selector.Distinct()
	}
	for _, p := range aueq.predicates {
		p(selector)
	}
	for _, p := range aueq.order {
		p(selector)
	}
	if offset := aueq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := aueq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// AppUserExtraGroupBy is the group-by builder for AppUserExtra entities.
type AppUserExtraGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (auegb *AppUserExtraGroupBy) Aggregate(fns ...AggregateFunc) *AppUserExtraGroupBy {
	auegb.fns = append(auegb.fns, fns...)
	return auegb
}

// Scan applies the group-by query and scans the result into the given value.
func (auegb *AppUserExtraGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := auegb.path(ctx)
	if err != nil {
		return err
	}
	auegb.sql = query
	return auegb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (auegb *AppUserExtraGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := auegb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (auegb *AppUserExtraGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(auegb.fields) > 1 {
		return nil, errors.New("ent: AppUserExtraGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := auegb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (auegb *AppUserExtraGroupBy) StringsX(ctx context.Context) []string {
	v, err := auegb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (auegb *AppUserExtraGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = auegb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{appuserextra.Label}
	default:
		err = fmt.Errorf("ent: AppUserExtraGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (auegb *AppUserExtraGroupBy) StringX(ctx context.Context) string {
	v, err := auegb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (auegb *AppUserExtraGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(auegb.fields) > 1 {
		return nil, errors.New("ent: AppUserExtraGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := auegb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (auegb *AppUserExtraGroupBy) IntsX(ctx context.Context) []int {
	v, err := auegb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (auegb *AppUserExtraGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = auegb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{appuserextra.Label}
	default:
		err = fmt.Errorf("ent: AppUserExtraGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (auegb *AppUserExtraGroupBy) IntX(ctx context.Context) int {
	v, err := auegb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (auegb *AppUserExtraGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(auegb.fields) > 1 {
		return nil, errors.New("ent: AppUserExtraGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := auegb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (auegb *AppUserExtraGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := auegb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (auegb *AppUserExtraGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = auegb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{appuserextra.Label}
	default:
		err = fmt.Errorf("ent: AppUserExtraGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (auegb *AppUserExtraGroupBy) Float64X(ctx context.Context) float64 {
	v, err := auegb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (auegb *AppUserExtraGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(auegb.fields) > 1 {
		return nil, errors.New("ent: AppUserExtraGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := auegb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (auegb *AppUserExtraGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := auegb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (auegb *AppUserExtraGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = auegb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{appuserextra.Label}
	default:
		err = fmt.Errorf("ent: AppUserExtraGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (auegb *AppUserExtraGroupBy) BoolX(ctx context.Context) bool {
	v, err := auegb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (auegb *AppUserExtraGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range auegb.fields {
		if !appuserextra.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := auegb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := auegb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (auegb *AppUserExtraGroupBy) sqlQuery() *sql.Selector {
	selector := auegb.sql.Select()
	aggregation := make([]string, 0, len(auegb.fns))
	for _, fn := range auegb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(auegb.fields)+len(auegb.fns))
		for _, f := range auegb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(auegb.fields...)...)
}

// AppUserExtraSelect is the builder for selecting fields of AppUserExtra entities.
type AppUserExtraSelect struct {
	*AppUserExtraQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (aues *AppUserExtraSelect) Scan(ctx context.Context, v interface{}) error {
	if err := aues.prepareQuery(ctx); err != nil {
		return err
	}
	aues.sql = aues.AppUserExtraQuery.sqlQuery(ctx)
	return aues.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (aues *AppUserExtraSelect) ScanX(ctx context.Context, v interface{}) {
	if err := aues.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (aues *AppUserExtraSelect) Strings(ctx context.Context) ([]string, error) {
	if len(aues.fields) > 1 {
		return nil, errors.New("ent: AppUserExtraSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := aues.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (aues *AppUserExtraSelect) StringsX(ctx context.Context) []string {
	v, err := aues.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (aues *AppUserExtraSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = aues.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{appuserextra.Label}
	default:
		err = fmt.Errorf("ent: AppUserExtraSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (aues *AppUserExtraSelect) StringX(ctx context.Context) string {
	v, err := aues.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (aues *AppUserExtraSelect) Ints(ctx context.Context) ([]int, error) {
	if len(aues.fields) > 1 {
		return nil, errors.New("ent: AppUserExtraSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := aues.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (aues *AppUserExtraSelect) IntsX(ctx context.Context) []int {
	v, err := aues.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (aues *AppUserExtraSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = aues.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{appuserextra.Label}
	default:
		err = fmt.Errorf("ent: AppUserExtraSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (aues *AppUserExtraSelect) IntX(ctx context.Context) int {
	v, err := aues.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (aues *AppUserExtraSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(aues.fields) > 1 {
		return nil, errors.New("ent: AppUserExtraSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := aues.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (aues *AppUserExtraSelect) Float64sX(ctx context.Context) []float64 {
	v, err := aues.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (aues *AppUserExtraSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = aues.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{appuserextra.Label}
	default:
		err = fmt.Errorf("ent: AppUserExtraSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (aues *AppUserExtraSelect) Float64X(ctx context.Context) float64 {
	v, err := aues.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (aues *AppUserExtraSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(aues.fields) > 1 {
		return nil, errors.New("ent: AppUserExtraSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := aues.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (aues *AppUserExtraSelect) BoolsX(ctx context.Context) []bool {
	v, err := aues.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (aues *AppUserExtraSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = aues.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{appuserextra.Label}
	default:
		err = fmt.Errorf("ent: AppUserExtraSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (aues *AppUserExtraSelect) BoolX(ctx context.Context) bool {
	v, err := aues.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (aues *AppUserExtraSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := aues.sql.Query()
	if err := aues.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
