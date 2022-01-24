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
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent/appuser"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// AppUserQuery is the builder for querying AppUser entities.
type AppUserQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.AppUser
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AppUserQuery builder.
func (auq *AppUserQuery) Where(ps ...predicate.AppUser) *AppUserQuery {
	auq.predicates = append(auq.predicates, ps...)
	return auq
}

// Limit adds a limit step to the query.
func (auq *AppUserQuery) Limit(limit int) *AppUserQuery {
	auq.limit = &limit
	return auq
}

// Offset adds an offset step to the query.
func (auq *AppUserQuery) Offset(offset int) *AppUserQuery {
	auq.offset = &offset
	return auq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (auq *AppUserQuery) Unique(unique bool) *AppUserQuery {
	auq.unique = &unique
	return auq
}

// Order adds an order step to the query.
func (auq *AppUserQuery) Order(o ...OrderFunc) *AppUserQuery {
	auq.order = append(auq.order, o...)
	return auq
}

// First returns the first AppUser entity from the query.
// Returns a *NotFoundError when no AppUser was found.
func (auq *AppUserQuery) First(ctx context.Context) (*AppUser, error) {
	nodes, err := auq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{appuser.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (auq *AppUserQuery) FirstX(ctx context.Context) *AppUser {
	node, err := auq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first AppUser ID from the query.
// Returns a *NotFoundError when no AppUser ID was found.
func (auq *AppUserQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = auq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{appuser.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (auq *AppUserQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := auq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single AppUser entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one AppUser entity is not found.
// Returns a *NotFoundError when no AppUser entities are found.
func (auq *AppUserQuery) Only(ctx context.Context) (*AppUser, error) {
	nodes, err := auq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{appuser.Label}
	default:
		return nil, &NotSingularError{appuser.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (auq *AppUserQuery) OnlyX(ctx context.Context) *AppUser {
	node, err := auq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only AppUser ID in the query.
// Returns a *NotSingularError when exactly one AppUser ID is not found.
// Returns a *NotFoundError when no entities are found.
func (auq *AppUserQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = auq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{appuser.Label}
	default:
		err = &NotSingularError{appuser.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (auq *AppUserQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := auq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of AppUsers.
func (auq *AppUserQuery) All(ctx context.Context) ([]*AppUser, error) {
	if err := auq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return auq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (auq *AppUserQuery) AllX(ctx context.Context) []*AppUser {
	nodes, err := auq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of AppUser IDs.
func (auq *AppUserQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := auq.Select(appuser.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (auq *AppUserQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := auq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (auq *AppUserQuery) Count(ctx context.Context) (int, error) {
	if err := auq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return auq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (auq *AppUserQuery) CountX(ctx context.Context) int {
	count, err := auq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (auq *AppUserQuery) Exist(ctx context.Context) (bool, error) {
	if err := auq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return auq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (auq *AppUserQuery) ExistX(ctx context.Context) bool {
	exist, err := auq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AppUserQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (auq *AppUserQuery) Clone() *AppUserQuery {
	if auq == nil {
		return nil
	}
	return &AppUserQuery{
		config:     auq.config,
		limit:      auq.limit,
		offset:     auq.offset,
		order:      append([]OrderFunc{}, auq.order...),
		predicates: append([]predicate.AppUser{}, auq.predicates...),
		// clone intermediate query.
		sql:  auq.sql.Clone(),
		path: auq.path,
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
//	client.AppUser.Query().
//		GroupBy(appuser.FieldAppID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (auq *AppUserQuery) GroupBy(field string, fields ...string) *AppUserGroupBy {
	group := &AppUserGroupBy{config: auq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := auq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return auq.sqlQuery(ctx), nil
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
//	client.AppUser.Query().
//		Select(appuser.FieldAppID).
//		Scan(ctx, &v)
//
func (auq *AppUserQuery) Select(fields ...string) *AppUserSelect {
	auq.fields = append(auq.fields, fields...)
	return &AppUserSelect{AppUserQuery: auq}
}

func (auq *AppUserQuery) prepareQuery(ctx context.Context) error {
	for _, f := range auq.fields {
		if !appuser.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if auq.path != nil {
		prev, err := auq.path(ctx)
		if err != nil {
			return err
		}
		auq.sql = prev
	}
	return nil
}

func (auq *AppUserQuery) sqlAll(ctx context.Context) ([]*AppUser, error) {
	var (
		nodes = []*AppUser{}
		_spec = auq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &AppUser{config: auq.config}
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
	if err := sqlgraph.QueryNodes(ctx, auq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (auq *AppUserQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := auq.querySpec()
	_spec.Node.Columns = auq.fields
	if len(auq.fields) > 0 {
		_spec.Unique = auq.unique != nil && *auq.unique
	}
	return sqlgraph.CountNodes(ctx, auq.driver, _spec)
}

func (auq *AppUserQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := auq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (auq *AppUserQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   appuser.Table,
			Columns: appuser.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: appuser.FieldID,
			},
		},
		From:   auq.sql,
		Unique: true,
	}
	if unique := auq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := auq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, appuser.FieldID)
		for i := range fields {
			if fields[i] != appuser.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := auq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := auq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := auq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := auq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (auq *AppUserQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(auq.driver.Dialect())
	t1 := builder.Table(appuser.Table)
	columns := auq.fields
	if len(columns) == 0 {
		columns = appuser.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if auq.sql != nil {
		selector = auq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if auq.unique != nil && *auq.unique {
		selector.Distinct()
	}
	for _, p := range auq.predicates {
		p(selector)
	}
	for _, p := range auq.order {
		p(selector)
	}
	if offset := auq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := auq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// AppUserGroupBy is the group-by builder for AppUser entities.
type AppUserGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (augb *AppUserGroupBy) Aggregate(fns ...AggregateFunc) *AppUserGroupBy {
	augb.fns = append(augb.fns, fns...)
	return augb
}

// Scan applies the group-by query and scans the result into the given value.
func (augb *AppUserGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := augb.path(ctx)
	if err != nil {
		return err
	}
	augb.sql = query
	return augb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (augb *AppUserGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := augb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (augb *AppUserGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(augb.fields) > 1 {
		return nil, errors.New("ent: AppUserGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := augb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (augb *AppUserGroupBy) StringsX(ctx context.Context) []string {
	v, err := augb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (augb *AppUserGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = augb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{appuser.Label}
	default:
		err = fmt.Errorf("ent: AppUserGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (augb *AppUserGroupBy) StringX(ctx context.Context) string {
	v, err := augb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (augb *AppUserGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(augb.fields) > 1 {
		return nil, errors.New("ent: AppUserGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := augb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (augb *AppUserGroupBy) IntsX(ctx context.Context) []int {
	v, err := augb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (augb *AppUserGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = augb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{appuser.Label}
	default:
		err = fmt.Errorf("ent: AppUserGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (augb *AppUserGroupBy) IntX(ctx context.Context) int {
	v, err := augb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (augb *AppUserGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(augb.fields) > 1 {
		return nil, errors.New("ent: AppUserGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := augb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (augb *AppUserGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := augb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (augb *AppUserGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = augb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{appuser.Label}
	default:
		err = fmt.Errorf("ent: AppUserGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (augb *AppUserGroupBy) Float64X(ctx context.Context) float64 {
	v, err := augb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (augb *AppUserGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(augb.fields) > 1 {
		return nil, errors.New("ent: AppUserGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := augb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (augb *AppUserGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := augb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (augb *AppUserGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = augb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{appuser.Label}
	default:
		err = fmt.Errorf("ent: AppUserGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (augb *AppUserGroupBy) BoolX(ctx context.Context) bool {
	v, err := augb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (augb *AppUserGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range augb.fields {
		if !appuser.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := augb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := augb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (augb *AppUserGroupBy) sqlQuery() *sql.Selector {
	selector := augb.sql.Select()
	aggregation := make([]string, 0, len(augb.fns))
	for _, fn := range augb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(augb.fields)+len(augb.fns))
		for _, f := range augb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(augb.fields...)...)
}

// AppUserSelect is the builder for selecting fields of AppUser entities.
type AppUserSelect struct {
	*AppUserQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (aus *AppUserSelect) Scan(ctx context.Context, v interface{}) error {
	if err := aus.prepareQuery(ctx); err != nil {
		return err
	}
	aus.sql = aus.AppUserQuery.sqlQuery(ctx)
	return aus.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (aus *AppUserSelect) ScanX(ctx context.Context, v interface{}) {
	if err := aus.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (aus *AppUserSelect) Strings(ctx context.Context) ([]string, error) {
	if len(aus.fields) > 1 {
		return nil, errors.New("ent: AppUserSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := aus.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (aus *AppUserSelect) StringsX(ctx context.Context) []string {
	v, err := aus.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (aus *AppUserSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = aus.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{appuser.Label}
	default:
		err = fmt.Errorf("ent: AppUserSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (aus *AppUserSelect) StringX(ctx context.Context) string {
	v, err := aus.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (aus *AppUserSelect) Ints(ctx context.Context) ([]int, error) {
	if len(aus.fields) > 1 {
		return nil, errors.New("ent: AppUserSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := aus.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (aus *AppUserSelect) IntsX(ctx context.Context) []int {
	v, err := aus.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (aus *AppUserSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = aus.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{appuser.Label}
	default:
		err = fmt.Errorf("ent: AppUserSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (aus *AppUserSelect) IntX(ctx context.Context) int {
	v, err := aus.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (aus *AppUserSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(aus.fields) > 1 {
		return nil, errors.New("ent: AppUserSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := aus.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (aus *AppUserSelect) Float64sX(ctx context.Context) []float64 {
	v, err := aus.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (aus *AppUserSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = aus.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{appuser.Label}
	default:
		err = fmt.Errorf("ent: AppUserSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (aus *AppUserSelect) Float64X(ctx context.Context) float64 {
	v, err := aus.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (aus *AppUserSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(aus.fields) > 1 {
		return nil, errors.New("ent: AppUserSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := aus.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (aus *AppUserSelect) BoolsX(ctx context.Context) []bool {
	v, err := aus.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (aus *AppUserSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = aus.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{appuser.Label}
	default:
		err = fmt.Errorf("ent: AppUserSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (aus *AppUserSelect) BoolX(ctx context.Context) bool {
	v, err := aus.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (aus *AppUserSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := aus.sql.Query()
	if err := aus.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
