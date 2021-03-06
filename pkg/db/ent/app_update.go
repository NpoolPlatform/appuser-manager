// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent/app"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// AppUpdate is the builder for updating App entities.
type AppUpdate struct {
	config
	hooks    []Hook
	mutation *AppMutation
}

// Where appends a list predicates to the AppUpdate builder.
func (au *AppUpdate) Where(ps ...predicate.App) *AppUpdate {
	au.mutation.Where(ps...)
	return au
}

// SetCreatedBy sets the "created_by" field.
func (au *AppUpdate) SetCreatedBy(u uuid.UUID) *AppUpdate {
	au.mutation.SetCreatedBy(u)
	return au
}

// SetName sets the "name" field.
func (au *AppUpdate) SetName(s string) *AppUpdate {
	au.mutation.SetName(s)
	return au
}

// SetLogo sets the "logo" field.
func (au *AppUpdate) SetLogo(s string) *AppUpdate {
	au.mutation.SetLogo(s)
	return au
}

// SetDescription sets the "description" field.
func (au *AppUpdate) SetDescription(s string) *AppUpdate {
	au.mutation.SetDescription(s)
	return au
}

// SetCreateAt sets the "create_at" field.
func (au *AppUpdate) SetCreateAt(u uint32) *AppUpdate {
	au.mutation.ResetCreateAt()
	au.mutation.SetCreateAt(u)
	return au
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (au *AppUpdate) SetNillableCreateAt(u *uint32) *AppUpdate {
	if u != nil {
		au.SetCreateAt(*u)
	}
	return au
}

// AddCreateAt adds u to the "create_at" field.
func (au *AppUpdate) AddCreateAt(u int32) *AppUpdate {
	au.mutation.AddCreateAt(u)
	return au
}

// SetUpdateAt sets the "update_at" field.
func (au *AppUpdate) SetUpdateAt(u uint32) *AppUpdate {
	au.mutation.ResetUpdateAt()
	au.mutation.SetUpdateAt(u)
	return au
}

// AddUpdateAt adds u to the "update_at" field.
func (au *AppUpdate) AddUpdateAt(u int32) *AppUpdate {
	au.mutation.AddUpdateAt(u)
	return au
}

// SetDeleteAt sets the "delete_at" field.
func (au *AppUpdate) SetDeleteAt(u uint32) *AppUpdate {
	au.mutation.ResetDeleteAt()
	au.mutation.SetDeleteAt(u)
	return au
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (au *AppUpdate) SetNillableDeleteAt(u *uint32) *AppUpdate {
	if u != nil {
		au.SetDeleteAt(*u)
	}
	return au
}

// AddDeleteAt adds u to the "delete_at" field.
func (au *AppUpdate) AddDeleteAt(u int32) *AppUpdate {
	au.mutation.AddDeleteAt(u)
	return au
}

// Mutation returns the AppMutation object of the builder.
func (au *AppUpdate) Mutation() *AppMutation {
	return au.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *AppUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	au.defaults()
	if len(au.hooks) == 0 {
		affected, err = au.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AppMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			au.mutation = mutation
			affected, err = au.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(au.hooks) - 1; i >= 0; i-- {
			if au.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = au.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, au.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (au *AppUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *AppUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *AppUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (au *AppUpdate) defaults() {
	if _, ok := au.mutation.UpdateAt(); !ok {
		v := app.UpdateDefaultUpdateAt()
		au.mutation.SetUpdateAt(v)
	}
}

func (au *AppUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   app.Table,
			Columns: app.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: app.FieldID,
			},
		},
	}
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.CreatedBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: app.FieldCreatedBy,
		})
	}
	if value, ok := au.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: app.FieldName,
		})
	}
	if value, ok := au.mutation.Logo(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: app.FieldLogo,
		})
	}
	if value, ok := au.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: app.FieldDescription,
		})
	}
	if value, ok := au.mutation.CreateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: app.FieldCreateAt,
		})
	}
	if value, ok := au.mutation.AddedCreateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: app.FieldCreateAt,
		})
	}
	if value, ok := au.mutation.UpdateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: app.FieldUpdateAt,
		})
	}
	if value, ok := au.mutation.AddedUpdateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: app.FieldUpdateAt,
		})
	}
	if value, ok := au.mutation.DeleteAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: app.FieldDeleteAt,
		})
	}
	if value, ok := au.mutation.AddedDeleteAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: app.FieldDeleteAt,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{app.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// AppUpdateOne is the builder for updating a single App entity.
type AppUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AppMutation
}

// SetCreatedBy sets the "created_by" field.
func (auo *AppUpdateOne) SetCreatedBy(u uuid.UUID) *AppUpdateOne {
	auo.mutation.SetCreatedBy(u)
	return auo
}

// SetName sets the "name" field.
func (auo *AppUpdateOne) SetName(s string) *AppUpdateOne {
	auo.mutation.SetName(s)
	return auo
}

// SetLogo sets the "logo" field.
func (auo *AppUpdateOne) SetLogo(s string) *AppUpdateOne {
	auo.mutation.SetLogo(s)
	return auo
}

// SetDescription sets the "description" field.
func (auo *AppUpdateOne) SetDescription(s string) *AppUpdateOne {
	auo.mutation.SetDescription(s)
	return auo
}

// SetCreateAt sets the "create_at" field.
func (auo *AppUpdateOne) SetCreateAt(u uint32) *AppUpdateOne {
	auo.mutation.ResetCreateAt()
	auo.mutation.SetCreateAt(u)
	return auo
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (auo *AppUpdateOne) SetNillableCreateAt(u *uint32) *AppUpdateOne {
	if u != nil {
		auo.SetCreateAt(*u)
	}
	return auo
}

// AddCreateAt adds u to the "create_at" field.
func (auo *AppUpdateOne) AddCreateAt(u int32) *AppUpdateOne {
	auo.mutation.AddCreateAt(u)
	return auo
}

// SetUpdateAt sets the "update_at" field.
func (auo *AppUpdateOne) SetUpdateAt(u uint32) *AppUpdateOne {
	auo.mutation.ResetUpdateAt()
	auo.mutation.SetUpdateAt(u)
	return auo
}

// AddUpdateAt adds u to the "update_at" field.
func (auo *AppUpdateOne) AddUpdateAt(u int32) *AppUpdateOne {
	auo.mutation.AddUpdateAt(u)
	return auo
}

// SetDeleteAt sets the "delete_at" field.
func (auo *AppUpdateOne) SetDeleteAt(u uint32) *AppUpdateOne {
	auo.mutation.ResetDeleteAt()
	auo.mutation.SetDeleteAt(u)
	return auo
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (auo *AppUpdateOne) SetNillableDeleteAt(u *uint32) *AppUpdateOne {
	if u != nil {
		auo.SetDeleteAt(*u)
	}
	return auo
}

// AddDeleteAt adds u to the "delete_at" field.
func (auo *AppUpdateOne) AddDeleteAt(u int32) *AppUpdateOne {
	auo.mutation.AddDeleteAt(u)
	return auo
}

// Mutation returns the AppMutation object of the builder.
func (auo *AppUpdateOne) Mutation() *AppMutation {
	return auo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auo *AppUpdateOne) Select(field string, fields ...string) *AppUpdateOne {
	auo.fields = append([]string{field}, fields...)
	return auo
}

// Save executes the query and returns the updated App entity.
func (auo *AppUpdateOne) Save(ctx context.Context) (*App, error) {
	var (
		err  error
		node *App
	)
	auo.defaults()
	if len(auo.hooks) == 0 {
		node, err = auo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AppMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			auo.mutation = mutation
			node, err = auo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(auo.hooks) - 1; i >= 0; i-- {
			if auo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = auo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, auo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (auo *AppUpdateOne) SaveX(ctx context.Context) *App {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *AppUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *AppUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (auo *AppUpdateOne) defaults() {
	if _, ok := auo.mutation.UpdateAt(); !ok {
		v := app.UpdateDefaultUpdateAt()
		auo.mutation.SetUpdateAt(v)
	}
}

func (auo *AppUpdateOne) sqlSave(ctx context.Context) (_node *App, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   app.Table,
			Columns: app.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: app.FieldID,
			},
		},
	}
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "App.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := auo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, app.FieldID)
		for _, f := range fields {
			if !app.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != app.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := auo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := auo.mutation.CreatedBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: app.FieldCreatedBy,
		})
	}
	if value, ok := auo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: app.FieldName,
		})
	}
	if value, ok := auo.mutation.Logo(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: app.FieldLogo,
		})
	}
	if value, ok := auo.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: app.FieldDescription,
		})
	}
	if value, ok := auo.mutation.CreateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: app.FieldCreateAt,
		})
	}
	if value, ok := auo.mutation.AddedCreateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: app.FieldCreateAt,
		})
	}
	if value, ok := auo.mutation.UpdateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: app.FieldUpdateAt,
		})
	}
	if value, ok := auo.mutation.AddedUpdateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: app.FieldUpdateAt,
		})
	}
	if value, ok := auo.mutation.DeleteAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: app.FieldDeleteAt,
		})
	}
	if value, ok := auo.mutation.AddedDeleteAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: app.FieldDeleteAt,
		})
	}
	_node = &App{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{app.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
