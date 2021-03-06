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
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent/banapp"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// BanAppQuery is the builder for querying BanApp entities.
type BanAppQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.BanApp
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the BanAppQuery builder.
func (baq *BanAppQuery) Where(ps ...predicate.BanApp) *BanAppQuery {
	baq.predicates = append(baq.predicates, ps...)
	return baq
}

// Limit adds a limit step to the query.
func (baq *BanAppQuery) Limit(limit int) *BanAppQuery {
	baq.limit = &limit
	return baq
}

// Offset adds an offset step to the query.
func (baq *BanAppQuery) Offset(offset int) *BanAppQuery {
	baq.offset = &offset
	return baq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (baq *BanAppQuery) Unique(unique bool) *BanAppQuery {
	baq.unique = &unique
	return baq
}

// Order adds an order step to the query.
func (baq *BanAppQuery) Order(o ...OrderFunc) *BanAppQuery {
	baq.order = append(baq.order, o...)
	return baq
}

// First returns the first BanApp entity from the query.
// Returns a *NotFoundError when no BanApp was found.
func (baq *BanAppQuery) First(ctx context.Context) (*BanApp, error) {
	nodes, err := baq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{banapp.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (baq *BanAppQuery) FirstX(ctx context.Context) *BanApp {
	node, err := baq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first BanApp ID from the query.
// Returns a *NotFoundError when no BanApp ID was found.
func (baq *BanAppQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = baq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{banapp.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (baq *BanAppQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := baq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single BanApp entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one BanApp entity is found.
// Returns a *NotFoundError when no BanApp entities are found.
func (baq *BanAppQuery) Only(ctx context.Context) (*BanApp, error) {
	nodes, err := baq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{banapp.Label}
	default:
		return nil, &NotSingularError{banapp.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (baq *BanAppQuery) OnlyX(ctx context.Context) *BanApp {
	node, err := baq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only BanApp ID in the query.
// Returns a *NotSingularError when more than one BanApp ID is found.
// Returns a *NotFoundError when no entities are found.
func (baq *BanAppQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = baq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{banapp.Label}
	default:
		err = &NotSingularError{banapp.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (baq *BanAppQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := baq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of BanApps.
func (baq *BanAppQuery) All(ctx context.Context) ([]*BanApp, error) {
	if err := baq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return baq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (baq *BanAppQuery) AllX(ctx context.Context) []*BanApp {
	nodes, err := baq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of BanApp IDs.
func (baq *BanAppQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := baq.Select(banapp.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (baq *BanAppQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := baq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (baq *BanAppQuery) Count(ctx context.Context) (int, error) {
	if err := baq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return baq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (baq *BanAppQuery) CountX(ctx context.Context) int {
	count, err := baq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (baq *BanAppQuery) Exist(ctx context.Context) (bool, error) {
	if err := baq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return baq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (baq *BanAppQuery) ExistX(ctx context.Context) bool {
	exist, err := baq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the BanAppQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (baq *BanAppQuery) Clone() *BanAppQuery {
	if baq == nil {
		return nil
	}
	return &BanAppQuery{
		config:     baq.config,
		limit:      baq.limit,
		offset:     baq.offset,
		order:      append([]OrderFunc{}, baq.order...),
		predicates: append([]predicate.BanApp{}, baq.predicates...),
		// clone intermediate query.
		sql:    baq.sql.Clone(),
		path:   baq.path,
		unique: baq.unique,
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
//	client.BanApp.Query().
//		GroupBy(banapp.FieldAppID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (baq *BanAppQuery) GroupBy(field string, fields ...string) *BanAppGroupBy {
	group := &BanAppGroupBy{config: baq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := baq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return baq.sqlQuery(ctx), nil
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
//	client.BanApp.Query().
//		Select(banapp.FieldAppID).
//		Scan(ctx, &v)
//
func (baq *BanAppQuery) Select(fields ...string) *BanAppSelect {
	baq.fields = append(baq.fields, fields...)
	return &BanAppSelect{BanAppQuery: baq}
}

func (baq *BanAppQuery) prepareQuery(ctx context.Context) error {
	for _, f := range baq.fields {
		if !banapp.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if baq.path != nil {
		prev, err := baq.path(ctx)
		if err != nil {
			return err
		}
		baq.sql = prev
	}
	return nil
}

func (baq *BanAppQuery) sqlAll(ctx context.Context) ([]*BanApp, error) {
	var (
		nodes = []*BanApp{}
		_spec = baq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &BanApp{config: baq.config}
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
	if err := sqlgraph.QueryNodes(ctx, baq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (baq *BanAppQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := baq.querySpec()
	_spec.Node.Columns = baq.fields
	if len(baq.fields) > 0 {
		_spec.Unique = baq.unique != nil && *baq.unique
	}
	return sqlgraph.CountNodes(ctx, baq.driver, _spec)
}

func (baq *BanAppQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := baq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (baq *BanAppQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   banapp.Table,
			Columns: banapp.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: banapp.FieldID,
			},
		},
		From:   baq.sql,
		Unique: true,
	}
	if unique := baq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := baq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, banapp.FieldID)
		for i := range fields {
			if fields[i] != banapp.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := baq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := baq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := baq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := baq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (baq *BanAppQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(baq.driver.Dialect())
	t1 := builder.Table(banapp.Table)
	columns := baq.fields
	if len(columns) == 0 {
		columns = banapp.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if baq.sql != nil {
		selector = baq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if baq.unique != nil && *baq.unique {
		selector.Distinct()
	}
	for _, p := range baq.predicates {
		p(selector)
	}
	for _, p := range baq.order {
		p(selector)
	}
	if offset := baq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := baq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// BanAppGroupBy is the group-by builder for BanApp entities.
type BanAppGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (bagb *BanAppGroupBy) Aggregate(fns ...AggregateFunc) *BanAppGroupBy {
	bagb.fns = append(bagb.fns, fns...)
	return bagb
}

// Scan applies the group-by query and scans the result into the given value.
func (bagb *BanAppGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := bagb.path(ctx)
	if err != nil {
		return err
	}
	bagb.sql = query
	return bagb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (bagb *BanAppGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := bagb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (bagb *BanAppGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(bagb.fields) > 1 {
		return nil, errors.New("ent: BanAppGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := bagb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (bagb *BanAppGroupBy) StringsX(ctx context.Context) []string {
	v, err := bagb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (bagb *BanAppGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = bagb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{banapp.Label}
	default:
		err = fmt.Errorf("ent: BanAppGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (bagb *BanAppGroupBy) StringX(ctx context.Context) string {
	v, err := bagb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (bagb *BanAppGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(bagb.fields) > 1 {
		return nil, errors.New("ent: BanAppGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := bagb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (bagb *BanAppGroupBy) IntsX(ctx context.Context) []int {
	v, err := bagb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (bagb *BanAppGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = bagb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{banapp.Label}
	default:
		err = fmt.Errorf("ent: BanAppGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (bagb *BanAppGroupBy) IntX(ctx context.Context) int {
	v, err := bagb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (bagb *BanAppGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(bagb.fields) > 1 {
		return nil, errors.New("ent: BanAppGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := bagb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (bagb *BanAppGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := bagb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (bagb *BanAppGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = bagb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{banapp.Label}
	default:
		err = fmt.Errorf("ent: BanAppGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (bagb *BanAppGroupBy) Float64X(ctx context.Context) float64 {
	v, err := bagb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (bagb *BanAppGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(bagb.fields) > 1 {
		return nil, errors.New("ent: BanAppGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := bagb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (bagb *BanAppGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := bagb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (bagb *BanAppGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = bagb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{banapp.Label}
	default:
		err = fmt.Errorf("ent: BanAppGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (bagb *BanAppGroupBy) BoolX(ctx context.Context) bool {
	v, err := bagb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (bagb *BanAppGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range bagb.fields {
		if !banapp.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := bagb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := bagb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (bagb *BanAppGroupBy) sqlQuery() *sql.Selector {
	selector := bagb.sql.Select()
	aggregation := make([]string, 0, len(bagb.fns))
	for _, fn := range bagb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(bagb.fields)+len(bagb.fns))
		for _, f := range bagb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(bagb.fields...)...)
}

// BanAppSelect is the builder for selecting fields of BanApp entities.
type BanAppSelect struct {
	*BanAppQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (bas *BanAppSelect) Scan(ctx context.Context, v interface{}) error {
	if err := bas.prepareQuery(ctx); err != nil {
		return err
	}
	bas.sql = bas.BanAppQuery.sqlQuery(ctx)
	return bas.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (bas *BanAppSelect) ScanX(ctx context.Context, v interface{}) {
	if err := bas.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (bas *BanAppSelect) Strings(ctx context.Context) ([]string, error) {
	if len(bas.fields) > 1 {
		return nil, errors.New("ent: BanAppSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := bas.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (bas *BanAppSelect) StringsX(ctx context.Context) []string {
	v, err := bas.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (bas *BanAppSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = bas.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{banapp.Label}
	default:
		err = fmt.Errorf("ent: BanAppSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (bas *BanAppSelect) StringX(ctx context.Context) string {
	v, err := bas.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (bas *BanAppSelect) Ints(ctx context.Context) ([]int, error) {
	if len(bas.fields) > 1 {
		return nil, errors.New("ent: BanAppSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := bas.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (bas *BanAppSelect) IntsX(ctx context.Context) []int {
	v, err := bas.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (bas *BanAppSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = bas.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{banapp.Label}
	default:
		err = fmt.Errorf("ent: BanAppSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (bas *BanAppSelect) IntX(ctx context.Context) int {
	v, err := bas.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (bas *BanAppSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(bas.fields) > 1 {
		return nil, errors.New("ent: BanAppSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := bas.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (bas *BanAppSelect) Float64sX(ctx context.Context) []float64 {
	v, err := bas.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (bas *BanAppSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = bas.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{banapp.Label}
	default:
		err = fmt.Errorf("ent: BanAppSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (bas *BanAppSelect) Float64X(ctx context.Context) float64 {
	v, err := bas.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (bas *BanAppSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(bas.fields) > 1 {
		return nil, errors.New("ent: BanAppSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := bas.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (bas *BanAppSelect) BoolsX(ctx context.Context) []bool {
	v, err := bas.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (bas *BanAppSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = bas.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{banapp.Label}
	default:
		err = fmt.Errorf("ent: BanAppSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (bas *BanAppSelect) BoolX(ctx context.Context) bool {
	v, err := bas.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (bas *BanAppSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := bas.sql.Query()
	if err := bas.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
