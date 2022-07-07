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

// SetCreatedAt sets the "created_at" field.
func (bauc *BanAppUserCreate) SetCreatedAt(u uint32) *BanAppUserCreate {
	bauc.mutation.SetCreatedAt(u)
	return bauc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (bauc *BanAppUserCreate) SetNillableCreatedAt(u *uint32) *BanAppUserCreate {
	if u != nil {
		bauc.SetCreatedAt(*u)
	}
	return bauc
}

// SetUpdatedAt sets the "updated_at" field.
func (bauc *BanAppUserCreate) SetUpdatedAt(u uint32) *BanAppUserCreate {
	bauc.mutation.SetUpdatedAt(u)
	return bauc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (bauc *BanAppUserCreate) SetNillableUpdatedAt(u *uint32) *BanAppUserCreate {
	if u != nil {
		bauc.SetUpdatedAt(*u)
	}
	return bauc
}

// SetDeletedAt sets the "deleted_at" field.
func (bauc *BanAppUserCreate) SetDeletedAt(u uint32) *BanAppUserCreate {
	bauc.mutation.SetDeletedAt(u)
	return bauc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (bauc *BanAppUserCreate) SetNillableDeletedAt(u *uint32) *BanAppUserCreate {
	if u != nil {
		bauc.SetDeletedAt(*u)
	}
	return bauc
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

// SetMessage sets the "message" field.
func (bauc *BanAppUserCreate) SetMessage(s string) *BanAppUserCreate {
	bauc.mutation.SetMessage(s)
	return bauc
}

// SetNillableMessage sets the "message" field if the given value is not nil.
func (bauc *BanAppUserCreate) SetNillableMessage(s *string) *BanAppUserCreate {
	if s != nil {
		bauc.SetMessage(*s)
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
	if err := bauc.defaults(); err != nil {
		return nil, err
	}
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
func (bauc *BanAppUserCreate) defaults() error {
	if _, ok := bauc.mutation.CreatedAt(); !ok {
		if banappuser.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized banappuser.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := banappuser.DefaultCreatedAt()
		bauc.mutation.SetCreatedAt(v)
	}
	if _, ok := bauc.mutation.UpdatedAt(); !ok {
		if banappuser.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized banappuser.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := banappuser.DefaultUpdatedAt()
		bauc.mutation.SetUpdatedAt(v)
	}
	if _, ok := bauc.mutation.DeletedAt(); !ok {
		if banappuser.DefaultDeletedAt == nil {
			return fmt.Errorf("ent: uninitialized banappuser.DefaultDeletedAt (forgotten import ent/runtime?)")
		}
		v := banappuser.DefaultDeletedAt()
		bauc.mutation.SetDeletedAt(v)
	}
	if _, ok := bauc.mutation.Message(); !ok {
		v := banappuser.DefaultMessage
		bauc.mutation.SetMessage(v)
	}
	if _, ok := bauc.mutation.ID(); !ok {
		if banappuser.DefaultID == nil {
			return fmt.Errorf("ent: uninitialized banappuser.DefaultID (forgotten import ent/runtime?)")
		}
		v := banappuser.DefaultID()
		bauc.mutation.SetID(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (bauc *BanAppUserCreate) check() error {
	if _, ok := bauc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "BanAppUser.created_at"`)}
	}
	if _, ok := bauc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "BanAppUser.updated_at"`)}
	}
	if _, ok := bauc.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`ent: missing required field "BanAppUser.deleted_at"`)}
	}
	if _, ok := bauc.mutation.AppID(); !ok {
		return &ValidationError{Name: "app_id", err: errors.New(`ent: missing required field "BanAppUser.app_id"`)}
	}
	if _, ok := bauc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "BanAppUser.user_id"`)}
	}
	if _, ok := bauc.mutation.Message(); !ok {
		return &ValidationError{Name: "message", err: errors.New(`ent: missing required field "BanAppUser.message"`)}
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
	if value, ok := bauc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: banappuser.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := bauc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: banappuser.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := bauc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: banappuser.FieldDeletedAt,
		})
		_node.DeletedAt = value
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
	if value, ok := bauc.mutation.Message(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: banappuser.FieldMessage,
		})
		_node.Message = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.BanAppUser.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.BanAppUserUpsert) {
//			SetCreatedAt(v+v).
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

// SetCreatedAt sets the "created_at" field.
func (u *BanAppUserUpsert) SetCreatedAt(v uint32) *BanAppUserUpsert {
	u.Set(banappuser.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *BanAppUserUpsert) UpdateCreatedAt() *BanAppUserUpsert {
	u.SetExcluded(banappuser.FieldCreatedAt)
	return u
}

// AddCreatedAt adds v to the "created_at" field.
func (u *BanAppUserUpsert) AddCreatedAt(v uint32) *BanAppUserUpsert {
	u.Add(banappuser.FieldCreatedAt, v)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *BanAppUserUpsert) SetUpdatedAt(v uint32) *BanAppUserUpsert {
	u.Set(banappuser.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *BanAppUserUpsert) UpdateUpdatedAt() *BanAppUserUpsert {
	u.SetExcluded(banappuser.FieldUpdatedAt)
	return u
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *BanAppUserUpsert) AddUpdatedAt(v uint32) *BanAppUserUpsert {
	u.Add(banappuser.FieldUpdatedAt, v)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *BanAppUserUpsert) SetDeletedAt(v uint32) *BanAppUserUpsert {
	u.Set(banappuser.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *BanAppUserUpsert) UpdateDeletedAt() *BanAppUserUpsert {
	u.SetExcluded(banappuser.FieldDeletedAt)
	return u
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *BanAppUserUpsert) AddDeletedAt(v uint32) *BanAppUserUpsert {
	u.Add(banappuser.FieldDeletedAt, v)
	return u
}

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

// SetMessage sets the "message" field.
func (u *BanAppUserUpsert) SetMessage(v string) *BanAppUserUpsert {
	u.Set(banappuser.FieldMessage, v)
	return u
}

// UpdateMessage sets the "message" field to the value that was provided on create.
func (u *BanAppUserUpsert) UpdateMessage() *BanAppUserUpsert {
	u.SetExcluded(banappuser.FieldMessage)
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

// SetCreatedAt sets the "created_at" field.
func (u *BanAppUserUpsertOne) SetCreatedAt(v uint32) *BanAppUserUpsertOne {
	return u.Update(func(s *BanAppUserUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *BanAppUserUpsertOne) AddCreatedAt(v uint32) *BanAppUserUpsertOne {
	return u.Update(func(s *BanAppUserUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *BanAppUserUpsertOne) UpdateCreatedAt() *BanAppUserUpsertOne {
	return u.Update(func(s *BanAppUserUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *BanAppUserUpsertOne) SetUpdatedAt(v uint32) *BanAppUserUpsertOne {
	return u.Update(func(s *BanAppUserUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *BanAppUserUpsertOne) AddUpdatedAt(v uint32) *BanAppUserUpsertOne {
	return u.Update(func(s *BanAppUserUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *BanAppUserUpsertOne) UpdateUpdatedAt() *BanAppUserUpsertOne {
	return u.Update(func(s *BanAppUserUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *BanAppUserUpsertOne) SetDeletedAt(v uint32) *BanAppUserUpsertOne {
	return u.Update(func(s *BanAppUserUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *BanAppUserUpsertOne) AddDeletedAt(v uint32) *BanAppUserUpsertOne {
	return u.Update(func(s *BanAppUserUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *BanAppUserUpsertOne) UpdateDeletedAt() *BanAppUserUpsertOne {
	return u.Update(func(s *BanAppUserUpsert) {
		s.UpdateDeletedAt()
	})
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

// SetMessage sets the "message" field.
func (u *BanAppUserUpsertOne) SetMessage(v string) *BanAppUserUpsertOne {
	return u.Update(func(s *BanAppUserUpsert) {
		s.SetMessage(v)
	})
}

// UpdateMessage sets the "message" field to the value that was provided on create.
func (u *BanAppUserUpsertOne) UpdateMessage() *BanAppUserUpsertOne {
	return u.Update(func(s *BanAppUserUpsert) {
		s.UpdateMessage()
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
//			SetCreatedAt(v+v).
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

// SetCreatedAt sets the "created_at" field.
func (u *BanAppUserUpsertBulk) SetCreatedAt(v uint32) *BanAppUserUpsertBulk {
	return u.Update(func(s *BanAppUserUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *BanAppUserUpsertBulk) AddCreatedAt(v uint32) *BanAppUserUpsertBulk {
	return u.Update(func(s *BanAppUserUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *BanAppUserUpsertBulk) UpdateCreatedAt() *BanAppUserUpsertBulk {
	return u.Update(func(s *BanAppUserUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *BanAppUserUpsertBulk) SetUpdatedAt(v uint32) *BanAppUserUpsertBulk {
	return u.Update(func(s *BanAppUserUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *BanAppUserUpsertBulk) AddUpdatedAt(v uint32) *BanAppUserUpsertBulk {
	return u.Update(func(s *BanAppUserUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *BanAppUserUpsertBulk) UpdateUpdatedAt() *BanAppUserUpsertBulk {
	return u.Update(func(s *BanAppUserUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *BanAppUserUpsertBulk) SetDeletedAt(v uint32) *BanAppUserUpsertBulk {
	return u.Update(func(s *BanAppUserUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *BanAppUserUpsertBulk) AddDeletedAt(v uint32) *BanAppUserUpsertBulk {
	return u.Update(func(s *BanAppUserUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *BanAppUserUpsertBulk) UpdateDeletedAt() *BanAppUserUpsertBulk {
	return u.Update(func(s *BanAppUserUpsert) {
		s.UpdateDeletedAt()
	})
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

// SetMessage sets the "message" field.
func (u *BanAppUserUpsertBulk) SetMessage(v string) *BanAppUserUpsertBulk {
	return u.Update(func(s *BanAppUserUpsert) {
		s.SetMessage(v)
	})
}

// UpdateMessage sets the "message" field to the value that was provided on create.
func (u *BanAppUserUpsertBulk) UpdateMessage() *BanAppUserUpsertBulk {
	return u.Update(func(s *BanAppUserUpsert) {
		s.UpdateMessage()
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
