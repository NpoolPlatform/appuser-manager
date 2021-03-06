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
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent/appusersecret"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// AppUserSecretQuery is the builder for querying AppUserSecret entities.
type AppUserSecretQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.AppUserSecret
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AppUserSecretQuery builder.
func (ausq *AppUserSecretQuery) Where(ps ...predicate.AppUserSecret) *AppUserSecretQuery {
	ausq.predicates = append(ausq.predicates, ps...)
	return ausq
}

// Limit adds a limit step to the query.
func (ausq *AppUserSecretQuery) Limit(limit int) *AppUserSecretQuery {
	ausq.limit = &limit
	return ausq
}

// Offset adds an offset step to the query.
func (ausq *AppUserSecretQuery) Offset(offset int) *AppUserSecretQuery {
	ausq.offset = &offset
	return ausq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ausq *AppUserSecretQuery) Unique(unique bool) *AppUserSecretQuery {
	ausq.unique = &unique
	return ausq
}

// Order adds an order step to the query.
func (ausq *AppUserSecretQuery) Order(o ...OrderFunc) *AppUserSecretQuery {
	ausq.order = append(ausq.order, o...)
	return ausq
}

// First returns the first AppUserSecret entity from the query.
// Returns a *NotFoundError when no AppUserSecret was found.
func (ausq *AppUserSecretQuery) First(ctx context.Context) (*AppUserSecret, error) {
	nodes, err := ausq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{appusersecret.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ausq *AppUserSecretQuery) FirstX(ctx context.Context) *AppUserSecret {
	node, err := ausq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first AppUserSecret ID from the query.
// Returns a *NotFoundError when no AppUserSecret ID was found.
func (ausq *AppUserSecretQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = ausq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{appusersecret.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ausq *AppUserSecretQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := ausq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single AppUserSecret entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one AppUserSecret entity is found.
// Returns a *NotFoundError when no AppUserSecret entities are found.
func (ausq *AppUserSecretQuery) Only(ctx context.Context) (*AppUserSecret, error) {
	nodes, err := ausq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{appusersecret.Label}
	default:
		return nil, &NotSingularError{appusersecret.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ausq *AppUserSecretQuery) OnlyX(ctx context.Context) *AppUserSecret {
	node, err := ausq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only AppUserSecret ID in the query.
// Returns a *NotSingularError when more than one AppUserSecret ID is found.
// Returns a *NotFoundError when no entities are found.
func (ausq *AppUserSecretQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = ausq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{appusersecret.Label}
	default:
		err = &NotSingularError{appusersecret.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ausq *AppUserSecretQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := ausq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of AppUserSecrets.
func (ausq *AppUserSecretQuery) All(ctx context.Context) ([]*AppUserSecret, error) {
	if err := ausq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return ausq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (ausq *AppUserSecretQuery) AllX(ctx context.Context) []*AppUserSecret {
	nodes, err := ausq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of AppUserSecret IDs.
func (ausq *AppUserSecretQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := ausq.Select(appusersecret.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ausq *AppUserSecretQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := ausq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ausq *AppUserSecretQuery) Count(ctx context.Context) (int, error) {
	if err := ausq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return ausq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (ausq *AppUserSecretQuery) CountX(ctx context.Context) int {
	count, err := ausq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ausq *AppUserSecretQuery) Exist(ctx context.Context) (bool, error) {
	if err := ausq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return ausq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (ausq *AppUserSecretQuery) ExistX(ctx context.Context) bool {
	exist, err := ausq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AppUserSecretQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ausq *AppUserSecretQuery) Clone() *AppUserSecretQuery {
	if ausq == nil {
		return nil
	}
	return &AppUserSecretQuery{
		config:     ausq.config,
		limit:      ausq.limit,
		offset:     ausq.offset,
		order:      append([]OrderFunc{}, ausq.order...),
		predicates: append([]predicate.AppUserSecret{}, ausq.predicates...),
		// clone intermediate query.
		sql:    ausq.sql.Clone(),
		path:   ausq.path,
		unique: ausq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreateAt uint32 `json:"create_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.AppUserSecret.Query().
//		GroupBy(appusersecret.FieldCreateAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (ausq *AppUserSecretQuery) GroupBy(field string, fields ...string) *AppUserSecretGroupBy {
	group := &AppUserSecretGroupBy{config: ausq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := ausq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return ausq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreateAt uint32 `json:"create_at,omitempty"`
//	}
//
//	client.AppUserSecret.Query().
//		Select(appusersecret.FieldCreateAt).
//		Scan(ctx, &v)
//
func (ausq *AppUserSecretQuery) Select(fields ...string) *AppUserSecretSelect {
	ausq.fields = append(ausq.fields, fields...)
	return &AppUserSecretSelect{AppUserSecretQuery: ausq}
}

func (ausq *AppUserSecretQuery) prepareQuery(ctx context.Context) error {
	for _, f := range ausq.fields {
		if !appusersecret.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ausq.path != nil {
		prev, err := ausq.path(ctx)
		if err != nil {
			return err
		}
		ausq.sql = prev
	}
	if appusersecret.Policy == nil {
		return errors.New("ent: uninitialized appusersecret.Policy (forgotten import ent/runtime?)")
	}
	if err := appusersecret.Policy.EvalQuery(ctx, ausq); err != nil {
		return err
	}
	return nil
}

func (ausq *AppUserSecretQuery) sqlAll(ctx context.Context) ([]*AppUserSecret, error) {
	var (
		nodes = []*AppUserSecret{}
		_spec = ausq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &AppUserSecret{config: ausq.config}
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
	if err := sqlgraph.QueryNodes(ctx, ausq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (ausq *AppUserSecretQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ausq.querySpec()
	_spec.Node.Columns = ausq.fields
	if len(ausq.fields) > 0 {
		_spec.Unique = ausq.unique != nil && *ausq.unique
	}
	return sqlgraph.CountNodes(ctx, ausq.driver, _spec)
}

func (ausq *AppUserSecretQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := ausq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (ausq *AppUserSecretQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   appusersecret.Table,
			Columns: appusersecret.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: appusersecret.FieldID,
			},
		},
		From:   ausq.sql,
		Unique: true,
	}
	if unique := ausq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := ausq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, appusersecret.FieldID)
		for i := range fields {
			if fields[i] != appusersecret.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ausq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ausq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ausq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ausq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ausq *AppUserSecretQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ausq.driver.Dialect())
	t1 := builder.Table(appusersecret.Table)
	columns := ausq.fields
	if len(columns) == 0 {
		columns = appusersecret.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ausq.sql != nil {
		selector = ausq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ausq.unique != nil && *ausq.unique {
		selector.Distinct()
	}
	for _, p := range ausq.predicates {
		p(selector)
	}
	for _, p := range ausq.order {
		p(selector)
	}
	if offset := ausq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ausq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// AppUserSecretGroupBy is the group-by builder for AppUserSecret entities.
type AppUserSecretGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ausgb *AppUserSecretGroupBy) Aggregate(fns ...AggregateFunc) *AppUserSecretGroupBy {
	ausgb.fns = append(ausgb.fns, fns...)
	return ausgb
}

// Scan applies the group-by query and scans the result into the given value.
func (ausgb *AppUserSecretGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := ausgb.path(ctx)
	if err != nil {
		return err
	}
	ausgb.sql = query
	return ausgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ausgb *AppUserSecretGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := ausgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (ausgb *AppUserSecretGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(ausgb.fields) > 1 {
		return nil, errors.New("ent: AppUserSecretGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := ausgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ausgb *AppUserSecretGroupBy) StringsX(ctx context.Context) []string {
	v, err := ausgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ausgb *AppUserSecretGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ausgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{appusersecret.Label}
	default:
		err = fmt.Errorf("ent: AppUserSecretGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ausgb *AppUserSecretGroupBy) StringX(ctx context.Context) string {
	v, err := ausgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (ausgb *AppUserSecretGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(ausgb.fields) > 1 {
		return nil, errors.New("ent: AppUserSecretGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := ausgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ausgb *AppUserSecretGroupBy) IntsX(ctx context.Context) []int {
	v, err := ausgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ausgb *AppUserSecretGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ausgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{appusersecret.Label}
	default:
		err = fmt.Errorf("ent: AppUserSecretGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ausgb *AppUserSecretGroupBy) IntX(ctx context.Context) int {
	v, err := ausgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (ausgb *AppUserSecretGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(ausgb.fields) > 1 {
		return nil, errors.New("ent: AppUserSecretGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := ausgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ausgb *AppUserSecretGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := ausgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ausgb *AppUserSecretGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ausgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{appusersecret.Label}
	default:
		err = fmt.Errorf("ent: AppUserSecretGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ausgb *AppUserSecretGroupBy) Float64X(ctx context.Context) float64 {
	v, err := ausgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (ausgb *AppUserSecretGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(ausgb.fields) > 1 {
		return nil, errors.New("ent: AppUserSecretGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := ausgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ausgb *AppUserSecretGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := ausgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ausgb *AppUserSecretGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ausgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{appusersecret.Label}
	default:
		err = fmt.Errorf("ent: AppUserSecretGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ausgb *AppUserSecretGroupBy) BoolX(ctx context.Context) bool {
	v, err := ausgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ausgb *AppUserSecretGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range ausgb.fields {
		if !appusersecret.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := ausgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ausgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ausgb *AppUserSecretGroupBy) sqlQuery() *sql.Selector {
	selector := ausgb.sql.Select()
	aggregation := make([]string, 0, len(ausgb.fns))
	for _, fn := range ausgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(ausgb.fields)+len(ausgb.fns))
		for _, f := range ausgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(ausgb.fields...)...)
}

// AppUserSecretSelect is the builder for selecting fields of AppUserSecret entities.
type AppUserSecretSelect struct {
	*AppUserSecretQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (auss *AppUserSecretSelect) Scan(ctx context.Context, v interface{}) error {
	if err := auss.prepareQuery(ctx); err != nil {
		return err
	}
	auss.sql = auss.AppUserSecretQuery.sqlQuery(ctx)
	return auss.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (auss *AppUserSecretSelect) ScanX(ctx context.Context, v interface{}) {
	if err := auss.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (auss *AppUserSecretSelect) Strings(ctx context.Context) ([]string, error) {
	if len(auss.fields) > 1 {
		return nil, errors.New("ent: AppUserSecretSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := auss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (auss *AppUserSecretSelect) StringsX(ctx context.Context) []string {
	v, err := auss.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (auss *AppUserSecretSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = auss.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{appusersecret.Label}
	default:
		err = fmt.Errorf("ent: AppUserSecretSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (auss *AppUserSecretSelect) StringX(ctx context.Context) string {
	v, err := auss.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (auss *AppUserSecretSelect) Ints(ctx context.Context) ([]int, error) {
	if len(auss.fields) > 1 {
		return nil, errors.New("ent: AppUserSecretSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := auss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (auss *AppUserSecretSelect) IntsX(ctx context.Context) []int {
	v, err := auss.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (auss *AppUserSecretSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = auss.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{appusersecret.Label}
	default:
		err = fmt.Errorf("ent: AppUserSecretSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (auss *AppUserSecretSelect) IntX(ctx context.Context) int {
	v, err := auss.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (auss *AppUserSecretSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(auss.fields) > 1 {
		return nil, errors.New("ent: AppUserSecretSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := auss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (auss *AppUserSecretSelect) Float64sX(ctx context.Context) []float64 {
	v, err := auss.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (auss *AppUserSecretSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = auss.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{appusersecret.Label}
	default:
		err = fmt.Errorf("ent: AppUserSecretSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (auss *AppUserSecretSelect) Float64X(ctx context.Context) float64 {
	v, err := auss.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (auss *AppUserSecretSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(auss.fields) > 1 {
		return nil, errors.New("ent: AppUserSecretSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := auss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (auss *AppUserSecretSelect) BoolsX(ctx context.Context) []bool {
	v, err := auss.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (auss *AppUserSecretSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = auss.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{appusersecret.Label}
	default:
		err = fmt.Errorf("ent: AppUserSecretSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (auss *AppUserSecretSelect) BoolX(ctx context.Context) bool {
	v, err := auss.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (auss *AppUserSecretSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := auss.sql.Query()
	if err := auss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
