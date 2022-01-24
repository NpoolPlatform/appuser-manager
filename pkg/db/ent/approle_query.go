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
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent/approle"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// AppRoleQuery is the builder for querying AppRole entities.
type AppRoleQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.AppRole
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AppRoleQuery builder.
func (arq *AppRoleQuery) Where(ps ...predicate.AppRole) *AppRoleQuery {
	arq.predicates = append(arq.predicates, ps...)
	return arq
}

// Limit adds a limit step to the query.
func (arq *AppRoleQuery) Limit(limit int) *AppRoleQuery {
	arq.limit = &limit
	return arq
}

// Offset adds an offset step to the query.
func (arq *AppRoleQuery) Offset(offset int) *AppRoleQuery {
	arq.offset = &offset
	return arq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (arq *AppRoleQuery) Unique(unique bool) *AppRoleQuery {
	arq.unique = &unique
	return arq
}

// Order adds an order step to the query.
func (arq *AppRoleQuery) Order(o ...OrderFunc) *AppRoleQuery {
	arq.order = append(arq.order, o...)
	return arq
}

// First returns the first AppRole entity from the query.
// Returns a *NotFoundError when no AppRole was found.
func (arq *AppRoleQuery) First(ctx context.Context) (*AppRole, error) {
	nodes, err := arq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{approle.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (arq *AppRoleQuery) FirstX(ctx context.Context) *AppRole {
	node, err := arq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first AppRole ID from the query.
// Returns a *NotFoundError when no AppRole ID was found.
func (arq *AppRoleQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = arq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{approle.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (arq *AppRoleQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := arq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single AppRole entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one AppRole entity is not found.
// Returns a *NotFoundError when no AppRole entities are found.
func (arq *AppRoleQuery) Only(ctx context.Context) (*AppRole, error) {
	nodes, err := arq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{approle.Label}
	default:
		return nil, &NotSingularError{approle.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (arq *AppRoleQuery) OnlyX(ctx context.Context) *AppRole {
	node, err := arq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only AppRole ID in the query.
// Returns a *NotSingularError when exactly one AppRole ID is not found.
// Returns a *NotFoundError when no entities are found.
func (arq *AppRoleQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = arq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{approle.Label}
	default:
		err = &NotSingularError{approle.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (arq *AppRoleQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := arq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of AppRoles.
func (arq *AppRoleQuery) All(ctx context.Context) ([]*AppRole, error) {
	if err := arq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return arq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (arq *AppRoleQuery) AllX(ctx context.Context) []*AppRole {
	nodes, err := arq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of AppRole IDs.
func (arq *AppRoleQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := arq.Select(approle.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (arq *AppRoleQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := arq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (arq *AppRoleQuery) Count(ctx context.Context) (int, error) {
	if err := arq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return arq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (arq *AppRoleQuery) CountX(ctx context.Context) int {
	count, err := arq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (arq *AppRoleQuery) Exist(ctx context.Context) (bool, error) {
	if err := arq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return arq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (arq *AppRoleQuery) ExistX(ctx context.Context) bool {
	exist, err := arq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AppRoleQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (arq *AppRoleQuery) Clone() *AppRoleQuery {
	if arq == nil {
		return nil
	}
	return &AppRoleQuery{
		config:     arq.config,
		limit:      arq.limit,
		offset:     arq.offset,
		order:      append([]OrderFunc{}, arq.order...),
		predicates: append([]predicate.AppRole{}, arq.predicates...),
		// clone intermediate query.
		sql:  arq.sql.Clone(),
		path: arq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedBy uuid.UUID `json:"created_by,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.AppRole.Query().
//		GroupBy(approle.FieldCreatedBy).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (arq *AppRoleQuery) GroupBy(field string, fields ...string) *AppRoleGroupBy {
	group := &AppRoleGroupBy{config: arq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := arq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return arq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedBy uuid.UUID `json:"created_by,omitempty"`
//	}
//
//	client.AppRole.Query().
//		Select(approle.FieldCreatedBy).
//		Scan(ctx, &v)
//
func (arq *AppRoleQuery) Select(fields ...string) *AppRoleSelect {
	arq.fields = append(arq.fields, fields...)
	return &AppRoleSelect{AppRoleQuery: arq}
}

func (arq *AppRoleQuery) prepareQuery(ctx context.Context) error {
	for _, f := range arq.fields {
		if !approle.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if arq.path != nil {
		prev, err := arq.path(ctx)
		if err != nil {
			return err
		}
		arq.sql = prev
	}
	return nil
}

func (arq *AppRoleQuery) sqlAll(ctx context.Context) ([]*AppRole, error) {
	var (
		nodes = []*AppRole{}
		_spec = arq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &AppRole{config: arq.config}
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
	if err := sqlgraph.QueryNodes(ctx, arq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (arq *AppRoleQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := arq.querySpec()
	_spec.Node.Columns = arq.fields
	if len(arq.fields) > 0 {
		_spec.Unique = arq.unique != nil && *arq.unique
	}
	return sqlgraph.CountNodes(ctx, arq.driver, _spec)
}

func (arq *AppRoleQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := arq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (arq *AppRoleQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   approle.Table,
			Columns: approle.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: approle.FieldID,
			},
		},
		From:   arq.sql,
		Unique: true,
	}
	if unique := arq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := arq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, approle.FieldID)
		for i := range fields {
			if fields[i] != approle.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := arq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := arq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := arq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := arq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (arq *AppRoleQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(arq.driver.Dialect())
	t1 := builder.Table(approle.Table)
	columns := arq.fields
	if len(columns) == 0 {
		columns = approle.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if arq.sql != nil {
		selector = arq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if arq.unique != nil && *arq.unique {
		selector.Distinct()
	}
	for _, p := range arq.predicates {
		p(selector)
	}
	for _, p := range arq.order {
		p(selector)
	}
	if offset := arq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := arq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// AppRoleGroupBy is the group-by builder for AppRole entities.
type AppRoleGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (argb *AppRoleGroupBy) Aggregate(fns ...AggregateFunc) *AppRoleGroupBy {
	argb.fns = append(argb.fns, fns...)
	return argb
}

// Scan applies the group-by query and scans the result into the given value.
func (argb *AppRoleGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := argb.path(ctx)
	if err != nil {
		return err
	}
	argb.sql = query
	return argb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (argb *AppRoleGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := argb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (argb *AppRoleGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(argb.fields) > 1 {
		return nil, errors.New("ent: AppRoleGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := argb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (argb *AppRoleGroupBy) StringsX(ctx context.Context) []string {
	v, err := argb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (argb *AppRoleGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = argb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{approle.Label}
	default:
		err = fmt.Errorf("ent: AppRoleGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (argb *AppRoleGroupBy) StringX(ctx context.Context) string {
	v, err := argb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (argb *AppRoleGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(argb.fields) > 1 {
		return nil, errors.New("ent: AppRoleGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := argb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (argb *AppRoleGroupBy) IntsX(ctx context.Context) []int {
	v, err := argb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (argb *AppRoleGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = argb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{approle.Label}
	default:
		err = fmt.Errorf("ent: AppRoleGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (argb *AppRoleGroupBy) IntX(ctx context.Context) int {
	v, err := argb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (argb *AppRoleGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(argb.fields) > 1 {
		return nil, errors.New("ent: AppRoleGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := argb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (argb *AppRoleGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := argb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (argb *AppRoleGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = argb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{approle.Label}
	default:
		err = fmt.Errorf("ent: AppRoleGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (argb *AppRoleGroupBy) Float64X(ctx context.Context) float64 {
	v, err := argb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (argb *AppRoleGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(argb.fields) > 1 {
		return nil, errors.New("ent: AppRoleGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := argb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (argb *AppRoleGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := argb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (argb *AppRoleGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = argb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{approle.Label}
	default:
		err = fmt.Errorf("ent: AppRoleGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (argb *AppRoleGroupBy) BoolX(ctx context.Context) bool {
	v, err := argb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (argb *AppRoleGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range argb.fields {
		if !approle.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := argb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := argb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (argb *AppRoleGroupBy) sqlQuery() *sql.Selector {
	selector := argb.sql.Select()
	aggregation := make([]string, 0, len(argb.fns))
	for _, fn := range argb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(argb.fields)+len(argb.fns))
		for _, f := range argb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(argb.fields...)...)
}

// AppRoleSelect is the builder for selecting fields of AppRole entities.
type AppRoleSelect struct {
	*AppRoleQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ars *AppRoleSelect) Scan(ctx context.Context, v interface{}) error {
	if err := ars.prepareQuery(ctx); err != nil {
		return err
	}
	ars.sql = ars.AppRoleQuery.sqlQuery(ctx)
	return ars.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ars *AppRoleSelect) ScanX(ctx context.Context, v interface{}) {
	if err := ars.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (ars *AppRoleSelect) Strings(ctx context.Context) ([]string, error) {
	if len(ars.fields) > 1 {
		return nil, errors.New("ent: AppRoleSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := ars.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ars *AppRoleSelect) StringsX(ctx context.Context) []string {
	v, err := ars.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (ars *AppRoleSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ars.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{approle.Label}
	default:
		err = fmt.Errorf("ent: AppRoleSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ars *AppRoleSelect) StringX(ctx context.Context) string {
	v, err := ars.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (ars *AppRoleSelect) Ints(ctx context.Context) ([]int, error) {
	if len(ars.fields) > 1 {
		return nil, errors.New("ent: AppRoleSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := ars.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ars *AppRoleSelect) IntsX(ctx context.Context) []int {
	v, err := ars.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (ars *AppRoleSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ars.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{approle.Label}
	default:
		err = fmt.Errorf("ent: AppRoleSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ars *AppRoleSelect) IntX(ctx context.Context) int {
	v, err := ars.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (ars *AppRoleSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(ars.fields) > 1 {
		return nil, errors.New("ent: AppRoleSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := ars.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ars *AppRoleSelect) Float64sX(ctx context.Context) []float64 {
	v, err := ars.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (ars *AppRoleSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ars.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{approle.Label}
	default:
		err = fmt.Errorf("ent: AppRoleSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ars *AppRoleSelect) Float64X(ctx context.Context) float64 {
	v, err := ars.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (ars *AppRoleSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(ars.fields) > 1 {
		return nil, errors.New("ent: AppRoleSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := ars.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ars *AppRoleSelect) BoolsX(ctx context.Context) []bool {
	v, err := ars.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (ars *AppRoleSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ars.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{approle.Label}
	default:
		err = fmt.Errorf("ent: AppRoleSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ars *AppRoleSelect) BoolX(ctx context.Context) bool {
	v, err := ars.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ars *AppRoleSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ars.sql.Query()
	if err := ars.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}