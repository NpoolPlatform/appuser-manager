// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent/banappuser"
	"github.com/google/uuid"
)

// BanAppUserCreate is the builder for creating a BanAppUser entity.
type BanAppUserCreate struct {
	config
	mutation *BanAppUserMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetAppID sets the "app_id" field.
func (bauc *BanAppUserCreate) SetAppID(u uuid.UUID) *BanAppUserCreate {
	bauc.mutation.SetAppID(u)
	return bauc
}

// SetUserID sets the "user_id" field.
func (bauc *BanAppUserCreate) SetUserID(u uuid.UUID) *BanAppUserCreate {
	bauc.mutation.SetUserID(u)
	return bauc
}

// SetCreateAt sets the "create_at" field.
func (bauc *BanAppUserCreate) SetCreateAt(u uint32) *BanAppUserCreate {
	bauc.mutation.SetCreateAt(u)
	return bauc
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (bauc *BanAppUserCreate) SetNillableCreateAt(u *uint32) *BanAppUserCreate {
	if u != nil {
		bauc.SetCreateAt(*u)
	}
	return bauc
}

// SetUpdateAt sets the "update_at" field.
func (bauc *BanAppUserCreate) SetUpdateAt(u uint32) *BanAppUserCreate {
	bauc.mutation.SetUpdateAt(u)
	return bauc
}

// SetNillableUpdateAt sets the "update_at" field if the given value is not nil.
func (bauc *BanAppUserCreate) SetNillableUpdateAt(u *uint32) *BanAppUserCreate {
	if u != nil {
		bauc.SetUpdateAt(*u)
	}
	return bauc
}

// SetDeleteAt sets the "delete_at" field.
func (bauc *BanAppUserCreate) SetDeleteAt(u uint32) *BanAppUserCreate {
	bauc.mutation.SetDeleteAt(u)
	return bauc
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (bauc *BanAppUserCreate) SetNillableDeleteAt(u *uint32) *BanAppUserCreate {
	if u != nil {
		bauc.SetDeleteAt(*u)
	}
	return bauc
}

// SetID sets the "id" field.
func (bauc *BanAppUserCreate) SetID(u uuid.UUID) *BanAppUserCreate {
	bauc.mutation.SetID(u)
	return bauc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (bauc *BanAppUserCreate) SetNillableID(u *uuid.UUID) *BanAppUserCreate {
	if u != nil {
		bauc.SetID(*u)
	}
	return bauc
}

// Mutation returns the BanAppUserMutation object of the builder.
func (bauc *BanAppUserCreate) Mutation() *BanAppUserMutation {
	return bauc.mutation
}

// Save creates the BanAppUser in the database.
func (bauc *BanAppUserCreate) Save(ctx context.Context) (*BanAppUser, error) {
	var (
		err  error
		node *BanAppUser
	)
	bauc.defaults()
	if len(bauc.hooks) == 0 {
		if err = bauc.check(); err != nil {
			return nil, err
		}
		node, err = bauc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BanAppUserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = bauc.check(); err != nil {
				return nil, err
			}
			bauc.mutation = mutation
			if node, err = bauc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(bauc.hooks) - 1; i >= 0; i-- {
			if bauc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = bauc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bauc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (bauc *BanAppUserCreate) SaveX(ctx context.Context) *BanAppUser {
	v, err := bauc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bauc *BanAppUserCreate) Exec(ctx context.Context) error {
	_, err := bauc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bauc *BanAppUserCreate) ExecX(ctx context.Context) {
	if err := bauc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (bauc *BanAppUserCreate) defaults() {
	if _, ok := bauc.mutation.CreateAt(); !ok {
		v := banappuser.DefaultCreateAt()
		bauc.mutation.SetCreateAt(v)
	}
	if _, ok := bauc.mutation.UpdateAt(); !ok {
		v := banappuser.DefaultUpdateAt()
		bauc.mutation.SetUpdateAt(v)
	}
	if _, ok := bauc.mutation.DeleteAt(); !ok {
		v := banappuser.DefaultDeleteAt()
		bauc.mutation.SetDeleteAt(v)
	}
	if _, ok := bauc.mutation.ID(); !ok {
		v := banappuser.DefaultID()
		bauc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bauc *BanAppUserCreate) check() error {
	if _, ok := bauc.mutation.AppID(); !ok {
		return &ValidationError{Name: "app_id", err: errors.New(`ent: missing required field "BanAppUser.app_id"`)}
	}
	if _, ok := bauc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "BanAppUser.user_id"`)}
	}
	if _, ok := bauc.mutation.CreateAt(); !ok {
		return &ValidationError{Name: "create_at", err: errors.New(`ent: missing required field "BanAppUser.create_at"`)}
	}
	if _, ok := bauc.mutation.UpdateAt(); !ok {
		return &ValidationError{Name: "update_at", err: errors.New(`ent: missing required field "BanAppUser.update_at"`)}
	}
	if _, ok := bauc.mutation.DeleteAt(); !ok {
		return &ValidationError{Name: "delete_at", err: errors.New(`ent: missing required field "BanAppUser.delete_at"`)}
	}
	return nil
}

func (bauc *BanAppUserCreate) sqlSave(ctx context.Context) (*BanAppUser, error) {
	_node, _spec := bauc.createSpec()
	if err := sqlgraph.CreateNode(ctx, bauc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	return _node, nil
}

func (bauc *BanAppUserCreate) createSpec() (*BanAppUser, *sqlgraph.CreateSpec) {
	var (
		_node = &BanAppUser{config: bauc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: banappuser.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: banappuser.FieldID,
			},
		}
	)
	_spec.OnConflict = bauc.conflict
	if id, ok := bauc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := bauc.mutation.AppID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: banappuser.FieldAppID,
		})
		_node.AppID = value
	}
	if value, ok := bauc.mutation.UserID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: banappuser.FieldUserID,
		})
		_node.UserID = value
	}
	if value, ok := bauc.mutation.CreateAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: banappuser.FieldCreateAt,
		})
		_node.CreateAt = value
	}
	if value, ok := bauc.mutation.UpdateAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: banappuser.FieldUpdateAt,
		})
		_node.UpdateAt = value
	}
	if value, ok := bauc.mutation.DeleteAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: banappuser.FieldDeleteAt,
		})
		_node.DeleteAt = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.BanAppUser.Create().
//		SetAppID(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.BanAppUserUpsert) {
//			SetAppID(v+v).
//		}).
//		Exec(ctx)
//
func (bauc *BanAppUserCreate) OnConflict(opts ...sql.ConflictOption) *BanAppUserUpsertOne {
	bauc.conflict = opts
	return &BanAppUserUpsertOne{
		create: bauc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.BanAppUser.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (bauc *BanAppUserCreate) OnConflictColumns(columns ...string) *BanAppUserUpsertOne {
	bauc.conflict = append(bauc.conflict, sql.ConflictColumns(columns...))
	return &BanAppUserUpsertOne{
		create: bauc,
	}
}

type (
	// BanAppUserUpsertOne is the builder for "upsert"-ing
	//  one BanAppUser node.
	BanAppUserUpsertOne struct {
		create *BanAppUserCreate
	}

	// BanAppUserUpsert is the "OnConflict" setter.
	BanAppUserUpsert struct {
		*sql.UpdateSet
	}
)

// SetAppID sets the "app_id" field.
func (u *BanAppUserUpsert) SetAppID(v uuid.UUID) *BanAppUserUpsert {
	u.Set(banappuser.FieldAppID, v)
	return u
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *BanAppUserUpsert) UpdateAppID() *BanAppUserUpsert {
	u.SetExcluded(banappuser.FieldAppID)
	return u
}

// SetUserID sets the "user_id" field.
func (u *BanAppUserUpsert) SetUserID(v uuid.UUID) *BanAppUserUpsert {
	u.Set(banappuser.FieldUserID, v)
	return u
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *BanAppUserUpsert) UpdateUserID() *BanAppUserUpsert {
	u.SetExcluded(banappuser.FieldUserID)
	return u
}

// SetCreateAt sets the "create_at" field.
func (u *BanAppUserUpsert) SetCreateAt(v uint32) *BanAppUserUpsert {
	u.Set(banappuser.FieldCreateAt, v)
	return u
}

// UpdateCreateAt sets the "create_at" field to the value that was provided on create.
func (u *BanAppUserUpsert) UpdateCreateAt() *BanAppUserUpsert {
	u.SetExcluded(banappuser.FieldCreateAt)
	return u
}

// AddCreateAt adds v to the "create_at" field.
func (u *BanAppUserUpsert) AddCreateAt(v uint32) *BanAppUserUpsert {
	u.Add(banappuser.FieldCreateAt, v)
	return u
}

// SetUpdateAt sets the "update_at" field.
func (u *BanAppUserUpsert) SetUpdateAt(v uint32) *BanAppUserUpsert {
	u.Set(banappuser.FieldUpdateAt, v)
	return u
}

// UpdateUpdateAt sets the "update_at" field to the value that was provided on create.
func (u *BanAppUserUpsert) UpdateUpdateAt() *BanAppUserUpsert {
	u.SetExcluded(banappuser.FieldUpdateAt)
	return u
}

// AddUpdateAt adds v to the "update_at" field.
func (u *BanAppUserUpsert) AddUpdateAt(v uint32) *BanAppUserUpsert {
	u.Add(banappuser.FieldUpdateAt, v)
	return u
}

// SetDeleteAt sets the "delete_at" field.
func (u *BanAppUserUpsert) SetDeleteAt(v uint32) *BanAppUserUpsert {
	u.Set(banappuser.FieldDeleteAt, v)
	return u
}

// UpdateDeleteAt sets the "delete_at" field to the value that was provided on create.
func (u *BanAppUserUpsert) UpdateDeleteAt() *BanAppUserUpsert {
	u.SetExcluded(banappuser.FieldDeleteAt)
	return u
}

// AddDeleteAt adds v to the "delete_at" field.
func (u *BanAppUserUpsert) AddDeleteAt(v uint32) *BanAppUserUpsert {
	u.Add(banappuser.FieldDeleteAt, v)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.BanAppUser.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(banappuser.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *BanAppUserUpsertOne) UpdateNewValues() *BanAppUserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(banappuser.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.BanAppUser.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *BanAppUserUpsertOne) Ignore() *BanAppUserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *BanAppUserUpsertOne) DoNothing() *BanAppUserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the BanAppUserCreate.OnConflict
// documentation for more info.
func (u *BanAppUserUpsertOne) Update(set func(*BanAppUserUpsert)) *BanAppUserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&BanAppUserUpsert{UpdateSet: update})
	}))
	return u
}

// SetAppID sets the "app_id" field.
func (u *BanAppUserUpsertOne) SetAppID(v uuid.UUID) *BanAppUserUpsertOne {
	return u.Update(func(s *BanAppUserUpsert) {
		s.SetAppID(v)
	})
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *BanAppUserUpsertOne) UpdateAppID() *BanAppUserUpsertOne {
	return u.Update(func(s *BanAppUserUpsert) {
		s.UpdateAppID()
	})
}

// SetUserID sets the "user_id" field.
func (u *BanAppUserUpsertOne) SetUserID(v uuid.UUID) *BanAppUserUpsertOne {
	return u.Update(func(s *BanAppUserUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *BanAppUserUpsertOne) UpdateUserID() *BanAppUserUpsertOne {
	return u.Update(func(s *BanAppUserUpsert) {
		s.UpdateUserID()
	})
}

// SetCreateAt sets the "create_at" field.
func (u *BanAppUserUpsertOne) SetCreateAt(v uint32) *BanAppUserUpsertOne {
	return u.Update(func(s *BanAppUserUpsert) {
		s.SetCreateAt(v)
	})
}

// AddCreateAt adds v to the "create_at" field.
func (u *BanAppUserUpsertOne) AddCreateAt(v uint32) *BanAppUserUpsertOne {
	return u.Update(func(s *BanAppUserUpsert) {
		s.AddCreateAt(v)
	})
}

// UpdateCreateAt sets the "create_at" field to the value that was provided on create.
func (u *BanAppUserUpsertOne) UpdateCreateAt() *BanAppUserUpsertOne {
	return u.Update(func(s *BanAppUserUpsert) {
		s.UpdateCreateAt()
	})
}

// SetUpdateAt sets the "update_at" field.
func (u *BanAppUserUpsertOne) SetUpdateAt(v uint32) *BanAppUserUpsertOne {
	return u.Update(func(s *BanAppUserUpsert) {
		s.SetUpdateAt(v)
	})
}

// AddUpdateAt adds v to the "update_at" field.
func (u *BanAppUserUpsertOne) AddUpdateAt(v uint32) *BanAppUserUpsertOne {
	return u.Update(func(s *BanAppUserUpsert) {
		s.AddUpdateAt(v)
	})
}

// UpdateUpdateAt sets the "update_at" field to the value that was provided on create.
func (u *BanAppUserUpsertOne) UpdateUpdateAt() *BanAppUserUpsertOne {
	return u.Update(func(s *BanAppUserUpsert) {
		s.UpdateUpdateAt()
	})
}

// SetDeleteAt sets the "delete_at" field.
func (u *BanAppUserUpsertOne) SetDeleteAt(v uint32) *BanAppUserUpsertOne {
	return u.Update(func(s *BanAppUserUpsert) {
		s.SetDeleteAt(v)
	})
}

// AddDeleteAt adds v to the "delete_at" field.
func (u *BanAppUserUpsertOne) AddDeleteAt(v uint32) *BanAppUserUpsertOne {
	return u.Update(func(s *BanAppUserUpsert) {
		s.AddDeleteAt(v)
	})
}

// UpdateDeleteAt sets the "delete_at" field to the value that was provided on create.
func (u *BanAppUserUpsertOne) UpdateDeleteAt() *BanAppUserUpsertOne {
	return u.Update(func(s *BanAppUserUpsert) {
		s.UpdateDeleteAt()
	})
}

// Exec executes the query.
func (u *BanAppUserUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for BanAppUserCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *BanAppUserUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *BanAppUserUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: BanAppUserUpsertOne.ID is not supported by MySQL driver. Use BanAppUserUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *BanAppUserUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// BanAppUserCreateBulk is the builder for creating many BanAppUser entities in bulk.
type BanAppUserCreateBulk struct {
	config
	builders []*BanAppUserCreate
	conflict []sql.ConflictOption
}

// Save creates the BanAppUser entities in the database.
func (baucb *BanAppUserCreateBulk) Save(ctx context.Context) ([]*BanAppUser, error) {
	specs := make([]*sqlgraph.CreateSpec, len(baucb.builders))
	nodes := make([]*BanAppUser, len(baucb.builders))
	mutators := make([]Mutator, len(baucb.builders))
	for i := range baucb.builders {
		func(i int, root context.Context) {
			builder := baucb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*BanAppUserMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, baucb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = baucb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, baucb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, baucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (baucb *BanAppUserCreateBulk) SaveX(ctx context.Context) []*BanAppUser {
	v, err := baucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (baucb *BanAppUserCreateBulk) Exec(ctx context.Context) error {
	_, err := baucb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (baucb *BanAppUserCreateBulk) ExecX(ctx context.Context) {
	if err := baucb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.BanAppUser.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.BanAppUserUpsert) {
//			SetAppID(v+v).
//		}).
//		Exec(ctx)
//
func (baucb *BanAppUserCreateBulk) OnConflict(opts ...sql.ConflictOption) *BanAppUserUpsertBulk {
	baucb.conflict = opts
	return &BanAppUserUpsertBulk{
		create: baucb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.BanAppUser.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (baucb *BanAppUserCreateBulk) OnConflictColumns(columns ...string) *BanAppUserUpsertBulk {
	baucb.conflict = append(baucb.conflict, sql.ConflictColumns(columns...))
	return &BanAppUserUpsertBulk{
		create: baucb,
	}
}

// BanAppUserUpsertBulk is the builder for "upsert"-ing
// a bulk of BanAppUser nodes.
type BanAppUserUpsertBulk struct {
	create *BanAppUserCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.BanAppUser.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(banappuser.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *BanAppUserUpsertBulk) UpdateNewValues() *BanAppUserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(banappuser.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.BanAppUser.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *BanAppUserUpsertBulk) Ignore() *BanAppUserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *BanAppUserUpsertBulk) DoNothing() *BanAppUserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the BanAppUserCreateBulk.OnConflict
// documentation for more info.
func (u *BanAppUserUpsertBulk) Update(set func(*BanAppUserUpsert)) *BanAppUserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&BanAppUserUpsert{UpdateSet: update})
	}))
	return u
}

// SetAppID sets the "app_id" field.
func (u *BanAppUserUpsertBulk) SetAppID(v uuid.UUID) *BanAppUserUpsertBulk {
	return u.Update(func(s *BanAppUserUpsert) {
		s.SetAppID(v)
	})
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *BanAppUserUpsertBulk) UpdateAppID() *BanAppUserUpsertBulk {
	return u.Update(func(s *BanAppUserUpsert) {
		s.UpdateAppID()
	})
}

// SetUserID sets the "user_id" field.
func (u *BanAppUserUpsertBulk) SetUserID(v uuid.UUID) *BanAppUserUpsertBulk {
	return u.Update(func(s *BanAppUserUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *BanAppUserUpsertBulk) UpdateUserID() *BanAppUserUpsertBulk {
	return u.Update(func(s *BanAppUserUpsert) {
		s.UpdateUserID()
	})
}

// SetCreateAt sets the "create_at" field.
func (u *BanAppUserUpsertBulk) SetCreateAt(v uint32) *BanAppUserUpsertBulk {
	return u.Update(func(s *BanAppUserUpsert) {
		s.SetCreateAt(v)
	})
}

// AddCreateAt adds v to the "create_at" field.
func (u *BanAppUserUpsertBulk) AddCreateAt(v uint32) *BanAppUserUpsertBulk {
	return u.Update(func(s *BanAppUserUpsert) {
		s.AddCreateAt(v)
	})
}

// UpdateCreateAt sets the "create_at" field to the value that was provided on create.
func (u *BanAppUserUpsertBulk) UpdateCreateAt() *BanAppUserUpsertBulk {
	return u.Update(func(s *BanAppUserUpsert) {
		s.UpdateCreateAt()
	})
}

// SetUpdateAt sets the "update_at" field.
func (u *BanAppUserUpsertBulk) SetUpdateAt(v uint32) *BanAppUserUpsertBulk {
	return u.Update(func(s *BanAppUserUpsert) {
		s.SetUpdateAt(v)
	})
}

// AddUpdateAt adds v to the "update_at" field.
func (u *BanAppUserUpsertBulk) AddUpdateAt(v uint32) *BanAppUserUpsertBulk {
	return u.Update(func(s *BanAppUserUpsert) {
		s.AddUpdateAt(v)
	})
}

// UpdateUpdateAt sets the "update_at" field to the value that was provided on create.
func (u *BanAppUserUpsertBulk) UpdateUpdateAt() *BanAppUserUpsertBulk {
	return u.Update(func(s *BanAppUserUpsert) {
		s.UpdateUpdateAt()
	})
}

// SetDeleteAt sets the "delete_at" field.
func (u *BanAppUserUpsertBulk) SetDeleteAt(v uint32) *BanAppUserUpsertBulk {
	return u.Update(func(s *BanAppUserUpsert) {
		s.SetDeleteAt(v)
	})
}

// AddDeleteAt adds v to the "delete_at" field.
func (u *BanAppUserUpsertBulk) AddDeleteAt(v uint32) *BanAppUserUpsertBulk {
	return u.Update(func(s *BanAppUserUpsert) {
		s.AddDeleteAt(v)
	})
}

// UpdateDeleteAt sets the "delete_at" field to the value that was provided on create.
func (u *BanAppUserUpsertBulk) UpdateDeleteAt() *BanAppUserUpsertBulk {
	return u.Update(func(s *BanAppUserUpsert) {
		s.UpdateDeleteAt()
	})
}

// Exec executes the query.
func (u *BanAppUserUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the BanAppUserCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for BanAppUserCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *BanAppUserUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
