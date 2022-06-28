// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent/appusercontrol"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// AppUserControlQuery is the builder for querying AppUserControl entities.
type AppUserControlQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.AppUserControl
	modifiers  []func(s *sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AppUserControlQuery builder.
func (aucq *AppUserControlQuery) Where(ps ...predicate.AppUserControl) *AppUserControlQuery {
	aucq.predicates = append(aucq.predicates, ps...)
	return aucq
}

// Limit adds a limit step to the query.
func (aucq *AppUserControlQuery) Limit(limit int) *AppUserControlQuery {
	aucq.limit = &limit
	return aucq
}

// Offset adds an offset step to the query.
func (aucq *AppUserControlQuery) Offset(offset int) *AppUserControlQuery {
	aucq.offset = &offset
	return aucq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (aucq *AppUserControlQuery) Unique(unique bool) *AppUserControlQuery {
	aucq.unique = &unique
	return aucq
}

// Order adds an order step to the query.
func (aucq *AppUserControlQuery) Order(o ...OrderFunc) *AppUserControlQuery {
	aucq.order = append(aucq.order, o...)
	return aucq
}

// First returns the first AppUserControl entity from the query.
// Returns a *NotFoundError when no AppUserControl was found.
func (aucq *AppUserControlQuery) First(ctx context.Context) (*AppUserControl, error) {
	nodes, err := aucq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{appusercontrol.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (aucq *AppUserControlQuery) FirstX(ctx context.Context) *AppUserControl {
	node, err := aucq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first AppUserControl ID from the query.
// Returns a *NotFoundError when no AppUserControl ID was found.
func (aucq *AppUserControlQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = aucq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{appusercontrol.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (aucq *AppUserControlQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := aucq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single AppUserControl entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one AppUserControl entity is found.
// Returns a *NotFoundError when no AppUserControl entities are found.
func (aucq *AppUserControlQuery) Only(ctx context.Context) (*AppUserControl, error) {
	nodes, err := aucq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{appusercontrol.Label}
	default:
		return nil, &NotSingularError{appusercontrol.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (aucq *AppUserControlQuery) OnlyX(ctx context.Context) *AppUserControl {
	node, err := aucq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only AppUserControl ID in the query.
// Returns a *NotSingularError when more than one AppUserControl ID is found.
// Returns a *NotFoundError when no entities are found.
func (aucq *AppUserControlQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = aucq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{appusercontrol.Label}
	default:
		err = &NotSingularError{appusercontrol.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (aucq *AppUserControlQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := aucq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of AppUserControls.
func (aucq *AppUserControlQuery) All(ctx context.Context) ([]*AppUserControl, error) {
	if err := aucq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return aucq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (aucq *AppUserControlQuery) AllX(ctx context.Context) []*AppUserControl {
	nodes, err := aucq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of AppUserControl IDs.
func (aucq *AppUserControlQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := aucq.Select(appusercontrol.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (aucq *AppUserControlQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := aucq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (aucq *AppUserControlQuery) Count(ctx context.Context) (int, error) {
	if err := aucq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return aucq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (aucq *AppUserControlQuery) CountX(ctx context.Context) int {
	count, err := aucq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (aucq *AppUserControlQuery) Exist(ctx context.Context) (bool, error) {
	if err := aucq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return aucq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (aucq *AppUserControlQuery) ExistX(ctx context.Context) bool {
	exist, err := aucq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AppUserControlQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (aucq *AppUserControlQuery) Clone() *AppUserControlQuery {
	if aucq == nil {
		return nil
	}
	return &AppUserControlQuery{
		config:     aucq.config,
		limit:      aucq.limit,
		offset:     aucq.offset,
		order:      append([]OrderFunc{}, aucq.order...),
		predicates: append([]predicate.AppUserControl{}, aucq.predicates...),
		// clone intermediate query.
		sql:    aucq.sql.Clone(),
		path:   aucq.path,
		unique: aucq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt uint32 `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.AppUserControl.Query().
//		GroupBy(appusercontrol.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (aucq *AppUserControlQuery) GroupBy(field string, fields ...string) *AppUserControlGroupBy {
	group := &AppUserControlGroupBy{config: aucq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := aucq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return aucq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt uint32 `json:"created_at,omitempty"`
//	}
//
//	client.AppUserControl.Query().
//		Select(appusercontrol.FieldCreatedAt).
//		Scan(ctx, &v)
//
func (aucq *AppUserControlQuery) Select(fields ...string) *AppUserControlSelect {
	aucq.fields = append(aucq.fields, fields...)
	return &AppUserControlSelect{AppUserControlQuery: aucq}
}

func (aucq *AppUserControlQuery) prepareQuery(ctx context.Context) error {
	for _, f := range aucq.fields {
		if !appusercontrol.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if aucq.path != nil {
		prev, err := aucq.path(ctx)
		if err != nil {
			return err
		}
		aucq.sql = prev
	}
	if appusercontrol.Policy == nil {
		return errors.New("ent: uninitialized appusercontrol.Policy (forgotten import ent/runtime?)")
	}
	if err := appusercontrol.Policy.EvalQuery(ctx, aucq); err != nil {
		return err
	}
	return nil
}

func (aucq *AppUserControlQuery) sqlAll(ctx context.Context) ([]*AppUserControl, error) {
	var (
		nodes = []*AppUserControl{}
		_spec = aucq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &AppUserControl{config: aucq.config}
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
	if len(aucq.modifiers) > 0 {
		_spec.Modifiers = aucq.modifiers
	}
	if err := sqlgraph.QueryNodes(ctx, aucq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (aucq *AppUserControlQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := aucq.querySpec()
	if len(aucq.modifiers) > 0 {
		_spec.Modifiers = aucq.modifiers
	}
	_spec.Node.Columns = aucq.fields
	if len(aucq.fields) > 0 {
		_spec.Unique = aucq.unique != nil && *aucq.unique
	}
	return sqlgraph.CountNodes(ctx, aucq.driver, _spec)
}

func (aucq *AppUserControlQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := aucq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (aucq *AppUserControlQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   appusercontrol.Table,
			Columns: appusercontrol.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: appusercontrol.FieldID,
			},
		},
		From:   aucq.sql,
		Unique: true,
	}
	if unique := aucq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := aucq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, appusercontrol.FieldID)
		for i := range fields {
			if fields[i] != appusercontrol.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := aucq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := aucq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := aucq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := aucq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (aucq *AppUserControlQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(aucq.driver.Dialect())
	t1 := builder.Table(appusercontrol.Table)
	columns := aucq.fields
	if len(columns) == 0 {
		columns = appusercontrol.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if aucq.sql != nil {
		selector = aucq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if aucq.unique != nil && *aucq.unique {
		selector.Distinct()
	}
	for _, m := range aucq.modifiers {
		m(selector)
	}
	for _, p := range aucq.predicates {
		p(selector)
	}
	for _, p := range aucq.order {
		p(selector)
	}
	if offset := aucq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := aucq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (aucq *AppUserControlQuery) ForUpdate(opts ...sql.LockOption) *AppUserControlQuery {
	if aucq.driver.Dialect() == dialect.Postgres {
		aucq.Unique(false)
	}
	aucq.modifiers = append(aucq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return aucq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (aucq *AppUserControlQuery) ForShare(opts ...sql.LockOption) *AppUserControlQuery {
	if aucq.driver.Dialect() == dialect.Postgres {
		aucq.Unique(false)
	}
	aucq.modifiers = append(aucq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return aucq
}

// AppUserControlGroupBy is the group-by builder for AppUserControl entities.
type AppUserControlGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (aucgb *AppUserControlGroupBy) Aggregate(fns ...AggregateFunc) *AppUserControlGroupBy {
	aucgb.fns = append(aucgb.fns, fns...)
	return aucgb
}

// Scan applies the group-by query and scans the result into the given value.
func (aucgb *AppUserControlGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := aucgb.path(ctx)
	if err != nil {
		return err
	}
	aucgb.sql = query
	return aucgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (aucgb *AppUserControlGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := aucgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (aucgb *AppUserControlGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(aucgb.fields) > 1 {
		return nil, errors.New("ent: AppUserControlGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := aucgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (aucgb *AppUserControlGroupBy) StringsX(ctx context.Context) []string {
	v, err := aucgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (aucgb *AppUserControlGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = aucgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{appusercontrol.Label}
	default:
		err = fmt.Errorf("ent: AppUserControlGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (aucgb *AppUserControlGroupBy) StringX(ctx context.Context) string {
	v, err := aucgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (aucgb *AppUserControlGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(aucgb.fields) > 1 {
		return nil, errors.New("ent: AppUserControlGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := aucgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (aucgb *AppUserControlGroupBy) IntsX(ctx context.Context) []int {
	v, err := aucgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (aucgb *AppUserControlGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = aucgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{appusercontrol.Label}
	default:
		err = fmt.Errorf("ent: AppUserControlGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (aucgb *AppUserControlGroupBy) IntX(ctx context.Context) int {
	v, err := aucgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (aucgb *AppUserControlGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(aucgb.fields) > 1 {
		return nil, errors.New("ent: AppUserControlGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := aucgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (aucgb *AppUserControlGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := aucgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (aucgb *AppUserControlGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = aucgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{appusercontrol.Label}
	default:
		err = fmt.Errorf("ent: AppUserControlGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (aucgb *AppUserControlGroupBy) Float64X(ctx context.Context) float64 {
	v, err := aucgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (aucgb *AppUserControlGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(aucgb.fields) > 1 {
		return nil, errors.New("ent: AppUserControlGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := aucgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (aucgb *AppUserControlGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := aucgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (aucgb *AppUserControlGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = aucgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{appusercontrol.Label}
	default:
		err = fmt.Errorf("ent: AppUserControlGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (aucgb *AppUserControlGroupBy) BoolX(ctx context.Context) bool {
	v, err := aucgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (aucgb *AppUserControlGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range aucgb.fields {
		if !appusercontrol.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := aucgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := aucgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (aucgb *AppUserControlGroupBy) sqlQuery() *sql.Selector {
	selector := aucgb.sql.Select()
	aggregation := make([]string, 0, len(aucgb.fns))
	for _, fn := range aucgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(aucgb.fields)+len(aucgb.fns))
		for _, f := range aucgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(aucgb.fields...)...)
}

// AppUserControlSelect is the builder for selecting fields of AppUserControl entities.
type AppUserControlSelect struct {
	*AppUserControlQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (aucs *AppUserControlSelect) Scan(ctx context.Context, v interface{}) error {
	if err := aucs.prepareQuery(ctx); err != nil {
		return err
	}
	aucs.sql = aucs.AppUserControlQuery.sqlQuery(ctx)
	return aucs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (aucs *AppUserControlSelect) ScanX(ctx context.Context, v interface{}) {
	if err := aucs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (aucs *AppUserControlSelect) Strings(ctx context.Context) ([]string, error) {
	if len(aucs.fields) > 1 {
		return nil, errors.New("ent: AppUserControlSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := aucs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (aucs *AppUserControlSelect) StringsX(ctx context.Context) []string {
	v, err := aucs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (aucs *AppUserControlSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = aucs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{appusercontrol.Label}
	default:
		err = fmt.Errorf("ent: AppUserControlSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (aucs *AppUserControlSelect) StringX(ctx context.Context) string {
	v, err := aucs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (aucs *AppUserControlSelect) Ints(ctx context.Context) ([]int, error) {
	if len(aucs.fields) > 1 {
		return nil, errors.New("ent: AppUserControlSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := aucs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (aucs *AppUserControlSelect) IntsX(ctx context.Context) []int {
	v, err := aucs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (aucs *AppUserControlSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = aucs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{appusercontrol.Label}
	default:
		err = fmt.Errorf("ent: AppUserControlSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (aucs *AppUserControlSelect) IntX(ctx context.Context) int {
	v, err := aucs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (aucs *AppUserControlSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(aucs.fields) > 1 {
		return nil, errors.New("ent: AppUserControlSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := aucs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (aucs *AppUserControlSelect) Float64sX(ctx context.Context) []float64 {
	v, err := aucs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (aucs *AppUserControlSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = aucs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{appusercontrol.Label}
	default:
		err = fmt.Errorf("ent: AppUserControlSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (aucs *AppUserControlSelect) Float64X(ctx context.Context) float64 {
	v, err := aucs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (aucs *AppUserControlSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(aucs.fields) > 1 {
		return nil, errors.New("ent: AppUserControlSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := aucs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (aucs *AppUserControlSelect) BoolsX(ctx context.Context) []bool {
	v, err := aucs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (aucs *AppUserControlSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = aucs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{appusercontrol.Label}
	default:
		err = fmt.Errorf("ent: AppUserControlSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (aucs *AppUserControlSelect) BoolX(ctx context.Context) bool {
	v, err := aucs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (aucs *AppUserControlSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := aucs.sql.Query()
	if err := aucs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
