// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent/banapp"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// BanAppUpdate is the builder for updating BanApp entities.
type BanAppUpdate struct {
	config
	hooks    []Hook
	mutation *BanAppMutation
}

// Where appends a list predicates to the BanAppUpdate builder.
func (bau *BanAppUpdate) Where(ps ...predicate.BanApp) *BanAppUpdate {
	bau.mutation.Where(ps...)
	return bau
}

// SetAppID sets the "app_id" field.
func (bau *BanAppUpdate) SetAppID(u uuid.UUID) *BanAppUpdate {
	bau.mutation.SetAppID(u)
	return bau
}

// SetCreateAt sets the "create_at" field.
func (bau *BanAppUpdate) SetCreateAt(u uint32) *BanAppUpdate {
	bau.mutation.ResetCreateAt()
	bau.mutation.SetCreateAt(u)
	return bau
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (bau *BanAppUpdate) SetNillableCreateAt(u *uint32) *BanAppUpdate {
	if u != nil {
		bau.SetCreateAt(*u)
	}
	return bau
}

// AddCreateAt adds u to the "create_at" field.
func (bau *BanAppUpdate) AddCreateAt(u int32) *BanAppUpdate {
	bau.mutation.AddCreateAt(u)
	return bau
}

// SetUpdateAt sets the "update_at" field.
func (bau *BanAppUpdate) SetUpdateAt(u uint32) *BanAppUpdate {
	bau.mutation.ResetUpdateAt()
	bau.mutation.SetUpdateAt(u)
	return bau
}

// AddUpdateAt adds u to the "update_at" field.
func (bau *BanAppUpdate) AddUpdateAt(u int32) *BanAppUpdate {
	bau.mutation.AddUpdateAt(u)
	return bau
}

// SetDeleteAt sets the "delete_at" field.
func (bau *BanAppUpdate) SetDeleteAt(u uint32) *BanAppUpdate {
	bau.mutation.ResetDeleteAt()
	bau.mutation.SetDeleteAt(u)
	return bau
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (bau *BanAppUpdate) SetNillableDeleteAt(u *uint32) *BanAppUpdate {
	if u != nil {
		bau.SetDeleteAt(*u)
	}
	return bau
}

// AddDeleteAt adds u to the "delete_at" field.
func (bau *BanAppUpdate) AddDeleteAt(u int32) *BanAppUpdate {
	bau.mutation.AddDeleteAt(u)
	return bau
}

// Mutation returns the BanAppMutation object of the builder.
func (bau *BanAppUpdate) Mutation() *BanAppMutation {
	return bau.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (bau *BanAppUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	bau.defaults()
	if len(bau.hooks) == 0 {
		affected, err = bau.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BanAppMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			bau.mutation = mutation
			affected, err = bau.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(bau.hooks) - 1; i >= 0; i-- {
			if bau.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = bau.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bau.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (bau *BanAppUpdate) SaveX(ctx context.Context) int {
	affected, err := bau.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (bau *BanAppUpdate) Exec(ctx context.Context) error {
	_, err := bau.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bau *BanAppUpdate) ExecX(ctx context.Context) {
	if err := bau.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (bau *BanAppUpdate) defaults() {
	if _, ok := bau.mutation.UpdateAt(); !ok {
		v := banapp.UpdateDefaultUpdateAt()
		bau.mutation.SetUpdateAt(v)
	}
}

func (bau *BanAppUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   banapp.Table,
			Columns: banapp.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: banapp.FieldID,
			},
		},
	}
	if ps := bau.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bau.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: banapp.FieldAppID,
		})
	}
	if value, ok := bau.mutation.CreateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: banapp.FieldCreateAt,
		})
	}
	if value, ok := bau.mutation.AddedCreateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: banapp.FieldCreateAt,
		})
	}
	if value, ok := bau.mutation.UpdateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: banapp.FieldUpdateAt,
		})
	}
	if value, ok := bau.mutation.AddedUpdateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: banapp.FieldUpdateAt,
		})
	}
	if value, ok := bau.mutation.DeleteAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: banapp.FieldDeleteAt,
		})
	}
	if value, ok := bau.mutation.AddedDeleteAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: banapp.FieldDeleteAt,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, bau.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{banapp.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// BanAppUpdateOne is the builder for updating a single BanApp entity.
type BanAppUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *BanAppMutation
}

// SetAppID sets the "app_id" field.
func (bauo *BanAppUpdateOne) SetAppID(u uuid.UUID) *BanAppUpdateOne {
	bauo.mutation.SetAppID(u)
	return bauo
}

// SetCreateAt sets the "create_at" field.
func (bauo *BanAppUpdateOne) SetCreateAt(u uint32) *BanAppUpdateOne {
	bauo.mutation.ResetCreateAt()
	bauo.mutation.SetCreateAt(u)
	return bauo
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (bauo *BanAppUpdateOne) SetNillableCreateAt(u *uint32) *BanAppUpdateOne {
	if u != nil {
		bauo.SetCreateAt(*u)
	}
	return bauo
}

// AddCreateAt adds u to the "create_at" field.
func (bauo *BanAppUpdateOne) AddCreateAt(u int32) *BanAppUpdateOne {
	bauo.mutation.AddCreateAt(u)
	return bauo
}

// SetUpdateAt sets the "update_at" field.
func (bauo *BanAppUpdateOne) SetUpdateAt(u uint32) *BanAppUpdateOne {
	bauo.mutation.ResetUpdateAt()
	bauo.mutation.SetUpdateAt(u)
	return bauo
}

// AddUpdateAt adds u to the "update_at" field.
func (bauo *BanAppUpdateOne) AddUpdateAt(u int32) *BanAppUpdateOne {
	bauo.mutation.AddUpdateAt(u)
	return bauo
}

// SetDeleteAt sets the "delete_at" field.
func (bauo *BanAppUpdateOne) SetDeleteAt(u uint32) *BanAppUpdateOne {
	bauo.mutation.ResetDeleteAt()
	bauo.mutation.SetDeleteAt(u)
	return bauo
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (bauo *BanAppUpdateOne) SetNillableDeleteAt(u *uint32) *BanAppUpdateOne {
	if u != nil {
		bauo.SetDeleteAt(*u)
	}
	return bauo
}

// AddDeleteAt adds u to the "delete_at" field.
func (bauo *BanAppUpdateOne) AddDeleteAt(u int32) *BanAppUpdateOne {
	bauo.mutation.AddDeleteAt(u)
	return bauo
}

// Mutation returns the BanAppMutation object of the builder.
func (bauo *BanAppUpdateOne) Mutation() *BanAppMutation {
	return bauo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (bauo *BanAppUpdateOne) Select(field string, fields ...string) *BanAppUpdateOne {
	bauo.fields = append([]string{field}, fields...)
	return bauo
}

// Save executes the query and returns the updated BanApp entity.
func (bauo *BanAppUpdateOne) Save(ctx context.Context) (*BanApp, error) {
	var (
		err  error
		node *BanApp
	)
	bauo.defaults()
	if len(bauo.hooks) == 0 {
		node, err = bauo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BanAppMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			bauo.mutation = mutation
			node, err = bauo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(bauo.hooks) - 1; i >= 0; i-- {
			if bauo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = bauo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bauo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (bauo *BanAppUpdateOne) SaveX(ctx context.Context) *BanApp {
	node, err := bauo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (bauo *BanAppUpdateOne) Exec(ctx context.Context) error {
	_, err := bauo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bauo *BanAppUpdateOne) ExecX(ctx context.Context) {
	if err := bauo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (bauo *BanAppUpdateOne) defaults() {
	if _, ok := bauo.mutation.UpdateAt(); !ok {
		v := banapp.UpdateDefaultUpdateAt()
		bauo.mutation.SetUpdateAt(v)
	}
}

func (bauo *BanAppUpdateOne) sqlSave(ctx context.Context) (_node *BanApp, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   banapp.Table,
			Columns: banapp.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: banapp.FieldID,
			},
		},
	}
	id, ok := bauo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "BanApp.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := bauo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, banapp.FieldID)
		for _, f := range fields {
			if !banapp.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != banapp.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := bauo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bauo.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: banapp.FieldAppID,
		})
	}
	if value, ok := bauo.mutation.CreateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: banapp.FieldCreateAt,
		})
	}
	if value, ok := bauo.mutation.AddedCreateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: banapp.FieldCreateAt,
		})
	}
	if value, ok := bauo.mutation.UpdateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: banapp.FieldUpdateAt,
		})
	}
	if value, ok := bauo.mutation.AddedUpdateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: banapp.FieldUpdateAt,
		})
	}
	if value, ok := bauo.mutation.DeleteAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: banapp.FieldDeleteAt,
		})
	}
	if value, ok := bauo.mutation.AddedDeleteAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: banapp.FieldDeleteAt,
		})
	}
	_node = &BanApp{config: bauo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, bauo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{banapp.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
