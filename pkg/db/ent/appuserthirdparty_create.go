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
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent/appuserthirdparty"
	"github.com/google/uuid"
)

// AppUserThirdPartyCreate is the builder for creating a AppUserThirdParty entity.
type AppUserThirdPartyCreate struct {
	config
	mutation *AppUserThirdPartyMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (autpc *AppUserThirdPartyCreate) SetCreatedAt(u uint32) *AppUserThirdPartyCreate {
	autpc.mutation.SetCreatedAt(u)
	return autpc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (autpc *AppUserThirdPartyCreate) SetNillableCreatedAt(u *uint32) *AppUserThirdPartyCreate {
	if u != nil {
		autpc.SetCreatedAt(*u)
	}
	return autpc
}

// SetUpdatedAt sets the "updated_at" field.
func (autpc *AppUserThirdPartyCreate) SetUpdatedAt(u uint32) *AppUserThirdPartyCreate {
	autpc.mutation.SetUpdatedAt(u)
	return autpc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (autpc *AppUserThirdPartyCreate) SetNillableUpdatedAt(u *uint32) *AppUserThirdPartyCreate {
	if u != nil {
		autpc.SetUpdatedAt(*u)
	}
	return autpc
}

// SetDeletedAt sets the "deleted_at" field.
func (autpc *AppUserThirdPartyCreate) SetDeletedAt(u uint32) *AppUserThirdPartyCreate {
	autpc.mutation.SetDeletedAt(u)
	return autpc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (autpc *AppUserThirdPartyCreate) SetNillableDeletedAt(u *uint32) *AppUserThirdPartyCreate {
	if u != nil {
		autpc.SetDeletedAt(*u)
	}
	return autpc
}

// SetAppID sets the "app_id" field.
func (autpc *AppUserThirdPartyCreate) SetAppID(u uuid.UUID) *AppUserThirdPartyCreate {
	autpc.mutation.SetAppID(u)
	return autpc
}

// SetUserID sets the "user_id" field.
func (autpc *AppUserThirdPartyCreate) SetUserID(u uuid.UUID) *AppUserThirdPartyCreate {
	autpc.mutation.SetUserID(u)
	return autpc
}

// SetThirdPartyUserID sets the "third_party_user_id" field.
func (autpc *AppUserThirdPartyCreate) SetThirdPartyUserID(s string) *AppUserThirdPartyCreate {
	autpc.mutation.SetThirdPartyUserID(s)
	return autpc
}

// SetThirdPartyID sets the "third_party_id" field.
func (autpc *AppUserThirdPartyCreate) SetThirdPartyID(s string) *AppUserThirdPartyCreate {
	autpc.mutation.SetThirdPartyID(s)
	return autpc
}

// SetThirdPartyUsername sets the "third_party_username" field.
func (autpc *AppUserThirdPartyCreate) SetThirdPartyUsername(s string) *AppUserThirdPartyCreate {
	autpc.mutation.SetThirdPartyUsername(s)
	return autpc
}

// SetNillableThirdPartyUsername sets the "third_party_username" field if the given value is not nil.
func (autpc *AppUserThirdPartyCreate) SetNillableThirdPartyUsername(s *string) *AppUserThirdPartyCreate {
	if s != nil {
		autpc.SetThirdPartyUsername(*s)
	}
	return autpc
}

// SetThirdPartyUserAvatar sets the "third_party_user_avatar" field.
func (autpc *AppUserThirdPartyCreate) SetThirdPartyUserAvatar(s string) *AppUserThirdPartyCreate {
	autpc.mutation.SetThirdPartyUserAvatar(s)
	return autpc
}

// SetNillableThirdPartyUserAvatar sets the "third_party_user_avatar" field if the given value is not nil.
func (autpc *AppUserThirdPartyCreate) SetNillableThirdPartyUserAvatar(s *string) *AppUserThirdPartyCreate {
	if s != nil {
		autpc.SetThirdPartyUserAvatar(*s)
	}
	return autpc
}

// SetID sets the "id" field.
func (autpc *AppUserThirdPartyCreate) SetID(u uuid.UUID) *AppUserThirdPartyCreate {
	autpc.mutation.SetID(u)
	return autpc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (autpc *AppUserThirdPartyCreate) SetNillableID(u *uuid.UUID) *AppUserThirdPartyCreate {
	if u != nil {
		autpc.SetID(*u)
	}
	return autpc
}

// Mutation returns the AppUserThirdPartyMutation object of the builder.
func (autpc *AppUserThirdPartyCreate) Mutation() *AppUserThirdPartyMutation {
	return autpc.mutation
}

// Save creates the AppUserThirdParty in the database.
func (autpc *AppUserThirdPartyCreate) Save(ctx context.Context) (*AppUserThirdParty, error) {
	var (
		err  error
		node *AppUserThirdParty
	)
	if err := autpc.defaults(); err != nil {
		return nil, err
	}
	if len(autpc.hooks) == 0 {
		if err = autpc.check(); err != nil {
			return nil, err
		}
		node, err = autpc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AppUserThirdPartyMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = autpc.check(); err != nil {
				return nil, err
			}
			autpc.mutation = mutation
			if node, err = autpc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(autpc.hooks) - 1; i >= 0; i-- {
			if autpc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = autpc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, autpc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (autpc *AppUserThirdPartyCreate) SaveX(ctx context.Context) *AppUserThirdParty {
	v, err := autpc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (autpc *AppUserThirdPartyCreate) Exec(ctx context.Context) error {
	_, err := autpc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (autpc *AppUserThirdPartyCreate) ExecX(ctx context.Context) {
	if err := autpc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (autpc *AppUserThirdPartyCreate) defaults() error {
	if _, ok := autpc.mutation.CreatedAt(); !ok {
		if appuserthirdparty.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized appuserthirdparty.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := appuserthirdparty.DefaultCreatedAt()
		autpc.mutation.SetCreatedAt(v)
	}
	if _, ok := autpc.mutation.UpdatedAt(); !ok {
		if appuserthirdparty.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized appuserthirdparty.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := appuserthirdparty.DefaultUpdatedAt()
		autpc.mutation.SetUpdatedAt(v)
	}
	if _, ok := autpc.mutation.DeletedAt(); !ok {
		if appuserthirdparty.DefaultDeletedAt == nil {
			return fmt.Errorf("ent: uninitialized appuserthirdparty.DefaultDeletedAt (forgotten import ent/runtime?)")
		}
		v := appuserthirdparty.DefaultDeletedAt()
		autpc.mutation.SetDeletedAt(v)
	}
	if _, ok := autpc.mutation.ThirdPartyUsername(); !ok {
		v := appuserthirdparty.DefaultThirdPartyUsername
		autpc.mutation.SetThirdPartyUsername(v)
	}
	if _, ok := autpc.mutation.ThirdPartyUserAvatar(); !ok {
		v := appuserthirdparty.DefaultThirdPartyUserAvatar
		autpc.mutation.SetThirdPartyUserAvatar(v)
	}
	if _, ok := autpc.mutation.ID(); !ok {
		if appuserthirdparty.DefaultID == nil {
			return fmt.Errorf("ent: uninitialized appuserthirdparty.DefaultID (forgotten import ent/runtime?)")
		}
		v := appuserthirdparty.DefaultID()
		autpc.mutation.SetID(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (autpc *AppUserThirdPartyCreate) check() error {
	if _, ok := autpc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "AppUserThirdParty.created_at"`)}
	}
	if _, ok := autpc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "AppUserThirdParty.updated_at"`)}
	}
	if _, ok := autpc.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`ent: missing required field "AppUserThirdParty.deleted_at"`)}
	}
	if _, ok := autpc.mutation.AppID(); !ok {
		return &ValidationError{Name: "app_id", err: errors.New(`ent: missing required field "AppUserThirdParty.app_id"`)}
	}
	if _, ok := autpc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "AppUserThirdParty.user_id"`)}
	}
	if _, ok := autpc.mutation.ThirdPartyUserID(); !ok {
		return &ValidationError{Name: "third_party_user_id", err: errors.New(`ent: missing required field "AppUserThirdParty.third_party_user_id"`)}
	}
	if _, ok := autpc.mutation.ThirdPartyID(); !ok {
		return &ValidationError{Name: "third_party_id", err: errors.New(`ent: missing required field "AppUserThirdParty.third_party_id"`)}
	}
	if _, ok := autpc.mutation.ThirdPartyUsername(); !ok {
		return &ValidationError{Name: "third_party_username", err: errors.New(`ent: missing required field "AppUserThirdParty.third_party_username"`)}
	}
	if _, ok := autpc.mutation.ThirdPartyUserAvatar(); !ok {
		return &ValidationError{Name: "third_party_user_avatar", err: errors.New(`ent: missing required field "AppUserThirdParty.third_party_user_avatar"`)}
	}
	if v, ok := autpc.mutation.ThirdPartyUserAvatar(); ok {
		if err := appuserthirdparty.ThirdPartyUserAvatarValidator(v); err != nil {
			return &ValidationError{Name: "third_party_user_avatar", err: fmt.Errorf(`ent: validator failed for field "AppUserThirdParty.third_party_user_avatar": %w`, err)}
		}
	}
	return nil
}

func (autpc *AppUserThirdPartyCreate) sqlSave(ctx context.Context) (*AppUserThirdParty, error) {
	_node, _spec := autpc.createSpec()
	if err := sqlgraph.CreateNode(ctx, autpc.driver, _spec); err != nil {
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

func (autpc *AppUserThirdPartyCreate) createSpec() (*AppUserThirdParty, *sqlgraph.CreateSpec) {
	var (
		_node = &AppUserThirdParty{config: autpc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: appuserthirdparty.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: appuserthirdparty.FieldID,
			},
		}
	)
	_spec.OnConflict = autpc.conflict
	if id, ok := autpc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := autpc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appuserthirdparty.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := autpc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appuserthirdparty.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := autpc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appuserthirdparty.FieldDeletedAt,
		})
		_node.DeletedAt = value
	}
	if value, ok := autpc.mutation.AppID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: appuserthirdparty.FieldAppID,
		})
		_node.AppID = value
	}
	if value, ok := autpc.mutation.UserID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: appuserthirdparty.FieldUserID,
		})
		_node.UserID = value
	}
	if value, ok := autpc.mutation.ThirdPartyUserID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: appuserthirdparty.FieldThirdPartyUserID,
		})
		_node.ThirdPartyUserID = value
	}
	if value, ok := autpc.mutation.ThirdPartyID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: appuserthirdparty.FieldThirdPartyID,
		})
		_node.ThirdPartyID = value
	}
	if value, ok := autpc.mutation.ThirdPartyUsername(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: appuserthirdparty.FieldThirdPartyUsername,
		})
		_node.ThirdPartyUsername = value
	}
	if value, ok := autpc.mutation.ThirdPartyUserAvatar(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: appuserthirdparty.FieldThirdPartyUserAvatar,
		})
		_node.ThirdPartyUserAvatar = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.AppUserThirdParty.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AppUserThirdPartyUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (autpc *AppUserThirdPartyCreate) OnConflict(opts ...sql.ConflictOption) *AppUserThirdPartyUpsertOne {
	autpc.conflict = opts
	return &AppUserThirdPartyUpsertOne{
		create: autpc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.AppUserThirdParty.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (autpc *AppUserThirdPartyCreate) OnConflictColumns(columns ...string) *AppUserThirdPartyUpsertOne {
	autpc.conflict = append(autpc.conflict, sql.ConflictColumns(columns...))
	return &AppUserThirdPartyUpsertOne{
		create: autpc,
	}
}

type (
	// AppUserThirdPartyUpsertOne is the builder for "upsert"-ing
	//  one AppUserThirdParty node.
	AppUserThirdPartyUpsertOne struct {
		create *AppUserThirdPartyCreate
	}

	// AppUserThirdPartyUpsert is the "OnConflict" setter.
	AppUserThirdPartyUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedAt sets the "created_at" field.
func (u *AppUserThirdPartyUpsert) SetCreatedAt(v uint32) *AppUserThirdPartyUpsert {
	u.Set(appuserthirdparty.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *AppUserThirdPartyUpsert) UpdateCreatedAt() *AppUserThirdPartyUpsert {
	u.SetExcluded(appuserthirdparty.FieldCreatedAt)
	return u
}

// AddCreatedAt adds v to the "created_at" field.
func (u *AppUserThirdPartyUpsert) AddCreatedAt(v uint32) *AppUserThirdPartyUpsert {
	u.Add(appuserthirdparty.FieldCreatedAt, v)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *AppUserThirdPartyUpsert) SetUpdatedAt(v uint32) *AppUserThirdPartyUpsert {
	u.Set(appuserthirdparty.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *AppUserThirdPartyUpsert) UpdateUpdatedAt() *AppUserThirdPartyUpsert {
	u.SetExcluded(appuserthirdparty.FieldUpdatedAt)
	return u
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *AppUserThirdPartyUpsert) AddUpdatedAt(v uint32) *AppUserThirdPartyUpsert {
	u.Add(appuserthirdparty.FieldUpdatedAt, v)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *AppUserThirdPartyUpsert) SetDeletedAt(v uint32) *AppUserThirdPartyUpsert {
	u.Set(appuserthirdparty.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *AppUserThirdPartyUpsert) UpdateDeletedAt() *AppUserThirdPartyUpsert {
	u.SetExcluded(appuserthirdparty.FieldDeletedAt)
	return u
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *AppUserThirdPartyUpsert) AddDeletedAt(v uint32) *AppUserThirdPartyUpsert {
	u.Add(appuserthirdparty.FieldDeletedAt, v)
	return u
}

// SetAppID sets the "app_id" field.
func (u *AppUserThirdPartyUpsert) SetAppID(v uuid.UUID) *AppUserThirdPartyUpsert {
	u.Set(appuserthirdparty.FieldAppID, v)
	return u
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *AppUserThirdPartyUpsert) UpdateAppID() *AppUserThirdPartyUpsert {
	u.SetExcluded(appuserthirdparty.FieldAppID)
	return u
}

// SetUserID sets the "user_id" field.
func (u *AppUserThirdPartyUpsert) SetUserID(v uuid.UUID) *AppUserThirdPartyUpsert {
	u.Set(appuserthirdparty.FieldUserID, v)
	return u
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *AppUserThirdPartyUpsert) UpdateUserID() *AppUserThirdPartyUpsert {
	u.SetExcluded(appuserthirdparty.FieldUserID)
	return u
}

// SetThirdPartyUserID sets the "third_party_user_id" field.
func (u *AppUserThirdPartyUpsert) SetThirdPartyUserID(v string) *AppUserThirdPartyUpsert {
	u.Set(appuserthirdparty.FieldThirdPartyUserID, v)
	return u
}

// UpdateThirdPartyUserID sets the "third_party_user_id" field to the value that was provided on create.
func (u *AppUserThirdPartyUpsert) UpdateThirdPartyUserID() *AppUserThirdPartyUpsert {
	u.SetExcluded(appuserthirdparty.FieldThirdPartyUserID)
	return u
}

// SetThirdPartyID sets the "third_party_id" field.
func (u *AppUserThirdPartyUpsert) SetThirdPartyID(v string) *AppUserThirdPartyUpsert {
	u.Set(appuserthirdparty.FieldThirdPartyID, v)
	return u
}

// UpdateThirdPartyID sets the "third_party_id" field to the value that was provided on create.
func (u *AppUserThirdPartyUpsert) UpdateThirdPartyID() *AppUserThirdPartyUpsert {
	u.SetExcluded(appuserthirdparty.FieldThirdPartyID)
	return u
}

// SetThirdPartyUsername sets the "third_party_username" field.
func (u *AppUserThirdPartyUpsert) SetThirdPartyUsername(v string) *AppUserThirdPartyUpsert {
	u.Set(appuserthirdparty.FieldThirdPartyUsername, v)
	return u
}

// UpdateThirdPartyUsername sets the "third_party_username" field to the value that was provided on create.
func (u *AppUserThirdPartyUpsert) UpdateThirdPartyUsername() *AppUserThirdPartyUpsert {
	u.SetExcluded(appuserthirdparty.FieldThirdPartyUsername)
	return u
}

// SetThirdPartyUserAvatar sets the "third_party_user_avatar" field.
func (u *AppUserThirdPartyUpsert) SetThirdPartyUserAvatar(v string) *AppUserThirdPartyUpsert {
	u.Set(appuserthirdparty.FieldThirdPartyUserAvatar, v)
	return u
}

// UpdateThirdPartyUserAvatar sets the "third_party_user_avatar" field to the value that was provided on create.
func (u *AppUserThirdPartyUpsert) UpdateThirdPartyUserAvatar() *AppUserThirdPartyUpsert {
	u.SetExcluded(appuserthirdparty.FieldThirdPartyUserAvatar)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.AppUserThirdParty.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(appuserthirdparty.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *AppUserThirdPartyUpsertOne) UpdateNewValues() *AppUserThirdPartyUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(appuserthirdparty.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.AppUserThirdParty.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *AppUserThirdPartyUpsertOne) Ignore() *AppUserThirdPartyUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AppUserThirdPartyUpsertOne) DoNothing() *AppUserThirdPartyUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AppUserThirdPartyCreate.OnConflict
// documentation for more info.
func (u *AppUserThirdPartyUpsertOne) Update(set func(*AppUserThirdPartyUpsert)) *AppUserThirdPartyUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AppUserThirdPartyUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *AppUserThirdPartyUpsertOne) SetCreatedAt(v uint32) *AppUserThirdPartyUpsertOne {
	return u.Update(func(s *AppUserThirdPartyUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *AppUserThirdPartyUpsertOne) AddCreatedAt(v uint32) *AppUserThirdPartyUpsertOne {
	return u.Update(func(s *AppUserThirdPartyUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *AppUserThirdPartyUpsertOne) UpdateCreatedAt() *AppUserThirdPartyUpsertOne {
	return u.Update(func(s *AppUserThirdPartyUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *AppUserThirdPartyUpsertOne) SetUpdatedAt(v uint32) *AppUserThirdPartyUpsertOne {
	return u.Update(func(s *AppUserThirdPartyUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *AppUserThirdPartyUpsertOne) AddUpdatedAt(v uint32) *AppUserThirdPartyUpsertOne {
	return u.Update(func(s *AppUserThirdPartyUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *AppUserThirdPartyUpsertOne) UpdateUpdatedAt() *AppUserThirdPartyUpsertOne {
	return u.Update(func(s *AppUserThirdPartyUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *AppUserThirdPartyUpsertOne) SetDeletedAt(v uint32) *AppUserThirdPartyUpsertOne {
	return u.Update(func(s *AppUserThirdPartyUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *AppUserThirdPartyUpsertOne) AddDeletedAt(v uint32) *AppUserThirdPartyUpsertOne {
	return u.Update(func(s *AppUserThirdPartyUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *AppUserThirdPartyUpsertOne) UpdateDeletedAt() *AppUserThirdPartyUpsertOne {
	return u.Update(func(s *AppUserThirdPartyUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetAppID sets the "app_id" field.
func (u *AppUserThirdPartyUpsertOne) SetAppID(v uuid.UUID) *AppUserThirdPartyUpsertOne {
	return u.Update(func(s *AppUserThirdPartyUpsert) {
		s.SetAppID(v)
	})
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *AppUserThirdPartyUpsertOne) UpdateAppID() *AppUserThirdPartyUpsertOne {
	return u.Update(func(s *AppUserThirdPartyUpsert) {
		s.UpdateAppID()
	})
}

// SetUserID sets the "user_id" field.
func (u *AppUserThirdPartyUpsertOne) SetUserID(v uuid.UUID) *AppUserThirdPartyUpsertOne {
	return u.Update(func(s *AppUserThirdPartyUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *AppUserThirdPartyUpsertOne) UpdateUserID() *AppUserThirdPartyUpsertOne {
	return u.Update(func(s *AppUserThirdPartyUpsert) {
		s.UpdateUserID()
	})
}

// SetThirdPartyUserID sets the "third_party_user_id" field.
func (u *AppUserThirdPartyUpsertOne) SetThirdPartyUserID(v string) *AppUserThirdPartyUpsertOne {
	return u.Update(func(s *AppUserThirdPartyUpsert) {
		s.SetThirdPartyUserID(v)
	})
}

// UpdateThirdPartyUserID sets the "third_party_user_id" field to the value that was provided on create.
func (u *AppUserThirdPartyUpsertOne) UpdateThirdPartyUserID() *AppUserThirdPartyUpsertOne {
	return u.Update(func(s *AppUserThirdPartyUpsert) {
		s.UpdateThirdPartyUserID()
	})
}

// SetThirdPartyID sets the "third_party_id" field.
func (u *AppUserThirdPartyUpsertOne) SetThirdPartyID(v string) *AppUserThirdPartyUpsertOne {
	return u.Update(func(s *AppUserThirdPartyUpsert) {
		s.SetThirdPartyID(v)
	})
}

// UpdateThirdPartyID sets the "third_party_id" field to the value that was provided on create.
func (u *AppUserThirdPartyUpsertOne) UpdateThirdPartyID() *AppUserThirdPartyUpsertOne {
	return u.Update(func(s *AppUserThirdPartyUpsert) {
		s.UpdateThirdPartyID()
	})
}

// SetThirdPartyUsername sets the "third_party_username" field.
func (u *AppUserThirdPartyUpsertOne) SetThirdPartyUsername(v string) *AppUserThirdPartyUpsertOne {
	return u.Update(func(s *AppUserThirdPartyUpsert) {
		s.SetThirdPartyUsername(v)
	})
}

// UpdateThirdPartyUsername sets the "third_party_username" field to the value that was provided on create.
func (u *AppUserThirdPartyUpsertOne) UpdateThirdPartyUsername() *AppUserThirdPartyUpsertOne {
	return u.Update(func(s *AppUserThirdPartyUpsert) {
		s.UpdateThirdPartyUsername()
	})
}

// SetThirdPartyUserAvatar sets the "third_party_user_avatar" field.
func (u *AppUserThirdPartyUpsertOne) SetThirdPartyUserAvatar(v string) *AppUserThirdPartyUpsertOne {
	return u.Update(func(s *AppUserThirdPartyUpsert) {
		s.SetThirdPartyUserAvatar(v)
	})
}

// UpdateThirdPartyUserAvatar sets the "third_party_user_avatar" field to the value that was provided on create.
func (u *AppUserThirdPartyUpsertOne) UpdateThirdPartyUserAvatar() *AppUserThirdPartyUpsertOne {
	return u.Update(func(s *AppUserThirdPartyUpsert) {
		s.UpdateThirdPartyUserAvatar()
	})
}

// Exec executes the query.
func (u *AppUserThirdPartyUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AppUserThirdPartyCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AppUserThirdPartyUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *AppUserThirdPartyUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: AppUserThirdPartyUpsertOne.ID is not supported by MySQL driver. Use AppUserThirdPartyUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *AppUserThirdPartyUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// AppUserThirdPartyCreateBulk is the builder for creating many AppUserThirdParty entities in bulk.
type AppUserThirdPartyCreateBulk struct {
	config
	builders []*AppUserThirdPartyCreate
	conflict []sql.ConflictOption
}

// Save creates the AppUserThirdParty entities in the database.
func (autpcb *AppUserThirdPartyCreateBulk) Save(ctx context.Context) ([]*AppUserThirdParty, error) {
	specs := make([]*sqlgraph.CreateSpec, len(autpcb.builders))
	nodes := make([]*AppUserThirdParty, len(autpcb.builders))
	mutators := make([]Mutator, len(autpcb.builders))
	for i := range autpcb.builders {
		func(i int, root context.Context) {
			builder := autpcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AppUserThirdPartyMutation)
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
					_, err = mutators[i+1].Mutate(root, autpcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = autpcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, autpcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, autpcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (autpcb *AppUserThirdPartyCreateBulk) SaveX(ctx context.Context) []*AppUserThirdParty {
	v, err := autpcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (autpcb *AppUserThirdPartyCreateBulk) Exec(ctx context.Context) error {
	_, err := autpcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (autpcb *AppUserThirdPartyCreateBulk) ExecX(ctx context.Context) {
	if err := autpcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.AppUserThirdParty.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AppUserThirdPartyUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (autpcb *AppUserThirdPartyCreateBulk) OnConflict(opts ...sql.ConflictOption) *AppUserThirdPartyUpsertBulk {
	autpcb.conflict = opts
	return &AppUserThirdPartyUpsertBulk{
		create: autpcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.AppUserThirdParty.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (autpcb *AppUserThirdPartyCreateBulk) OnConflictColumns(columns ...string) *AppUserThirdPartyUpsertBulk {
	autpcb.conflict = append(autpcb.conflict, sql.ConflictColumns(columns...))
	return &AppUserThirdPartyUpsertBulk{
		create: autpcb,
	}
}

// AppUserThirdPartyUpsertBulk is the builder for "upsert"-ing
// a bulk of AppUserThirdParty nodes.
type AppUserThirdPartyUpsertBulk struct {
	create *AppUserThirdPartyCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.AppUserThirdParty.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(appuserthirdparty.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *AppUserThirdPartyUpsertBulk) UpdateNewValues() *AppUserThirdPartyUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(appuserthirdparty.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.AppUserThirdParty.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *AppUserThirdPartyUpsertBulk) Ignore() *AppUserThirdPartyUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AppUserThirdPartyUpsertBulk) DoNothing() *AppUserThirdPartyUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AppUserThirdPartyCreateBulk.OnConflict
// documentation for more info.
func (u *AppUserThirdPartyUpsertBulk) Update(set func(*AppUserThirdPartyUpsert)) *AppUserThirdPartyUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AppUserThirdPartyUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *AppUserThirdPartyUpsertBulk) SetCreatedAt(v uint32) *AppUserThirdPartyUpsertBulk {
	return u.Update(func(s *AppUserThirdPartyUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *AppUserThirdPartyUpsertBulk) AddCreatedAt(v uint32) *AppUserThirdPartyUpsertBulk {
	return u.Update(func(s *AppUserThirdPartyUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *AppUserThirdPartyUpsertBulk) UpdateCreatedAt() *AppUserThirdPartyUpsertBulk {
	return u.Update(func(s *AppUserThirdPartyUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *AppUserThirdPartyUpsertBulk) SetUpdatedAt(v uint32) *AppUserThirdPartyUpsertBulk {
	return u.Update(func(s *AppUserThirdPartyUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *AppUserThirdPartyUpsertBulk) AddUpdatedAt(v uint32) *AppUserThirdPartyUpsertBulk {
	return u.Update(func(s *AppUserThirdPartyUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *AppUserThirdPartyUpsertBulk) UpdateUpdatedAt() *AppUserThirdPartyUpsertBulk {
	return u.Update(func(s *AppUserThirdPartyUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *AppUserThirdPartyUpsertBulk) SetDeletedAt(v uint32) *AppUserThirdPartyUpsertBulk {
	return u.Update(func(s *AppUserThirdPartyUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *AppUserThirdPartyUpsertBulk) AddDeletedAt(v uint32) *AppUserThirdPartyUpsertBulk {
	return u.Update(func(s *AppUserThirdPartyUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *AppUserThirdPartyUpsertBulk) UpdateDeletedAt() *AppUserThirdPartyUpsertBulk {
	return u.Update(func(s *AppUserThirdPartyUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetAppID sets the "app_id" field.
func (u *AppUserThirdPartyUpsertBulk) SetAppID(v uuid.UUID) *AppUserThirdPartyUpsertBulk {
	return u.Update(func(s *AppUserThirdPartyUpsert) {
		s.SetAppID(v)
	})
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *AppUserThirdPartyUpsertBulk) UpdateAppID() *AppUserThirdPartyUpsertBulk {
	return u.Update(func(s *AppUserThirdPartyUpsert) {
		s.UpdateAppID()
	})
}

// SetUserID sets the "user_id" field.
func (u *AppUserThirdPartyUpsertBulk) SetUserID(v uuid.UUID) *AppUserThirdPartyUpsertBulk {
	return u.Update(func(s *AppUserThirdPartyUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *AppUserThirdPartyUpsertBulk) UpdateUserID() *AppUserThirdPartyUpsertBulk {
	return u.Update(func(s *AppUserThirdPartyUpsert) {
		s.UpdateUserID()
	})
}

// SetThirdPartyUserID sets the "third_party_user_id" field.
func (u *AppUserThirdPartyUpsertBulk) SetThirdPartyUserID(v string) *AppUserThirdPartyUpsertBulk {
	return u.Update(func(s *AppUserThirdPartyUpsert) {
		s.SetThirdPartyUserID(v)
	})
}

// UpdateThirdPartyUserID sets the "third_party_user_id" field to the value that was provided on create.
func (u *AppUserThirdPartyUpsertBulk) UpdateThirdPartyUserID() *AppUserThirdPartyUpsertBulk {
	return u.Update(func(s *AppUserThirdPartyUpsert) {
		s.UpdateThirdPartyUserID()
	})
}

// SetThirdPartyID sets the "third_party_id" field.
func (u *AppUserThirdPartyUpsertBulk) SetThirdPartyID(v string) *AppUserThirdPartyUpsertBulk {
	return u.Update(func(s *AppUserThirdPartyUpsert) {
		s.SetThirdPartyID(v)
	})
}

// UpdateThirdPartyID sets the "third_party_id" field to the value that was provided on create.
func (u *AppUserThirdPartyUpsertBulk) UpdateThirdPartyID() *AppUserThirdPartyUpsertBulk {
	return u.Update(func(s *AppUserThirdPartyUpsert) {
		s.UpdateThirdPartyID()
	})
}

// SetThirdPartyUsername sets the "third_party_username" field.
func (u *AppUserThirdPartyUpsertBulk) SetThirdPartyUsername(v string) *AppUserThirdPartyUpsertBulk {
	return u.Update(func(s *AppUserThirdPartyUpsert) {
		s.SetThirdPartyUsername(v)
	})
}

// UpdateThirdPartyUsername sets the "third_party_username" field to the value that was provided on create.
func (u *AppUserThirdPartyUpsertBulk) UpdateThirdPartyUsername() *AppUserThirdPartyUpsertBulk {
	return u.Update(func(s *AppUserThirdPartyUpsert) {
		s.UpdateThirdPartyUsername()
	})
}

// SetThirdPartyUserAvatar sets the "third_party_user_avatar" field.
func (u *AppUserThirdPartyUpsertBulk) SetThirdPartyUserAvatar(v string) *AppUserThirdPartyUpsertBulk {
	return u.Update(func(s *AppUserThirdPartyUpsert) {
		s.SetThirdPartyUserAvatar(v)
	})
}

// UpdateThirdPartyUserAvatar sets the "third_party_user_avatar" field to the value that was provided on create.
func (u *AppUserThirdPartyUpsertBulk) UpdateThirdPartyUserAvatar() *AppUserThirdPartyUpsertBulk {
	return u.Update(func(s *AppUserThirdPartyUpsert) {
		s.UpdateThirdPartyUserAvatar()
	})
}

// Exec executes the query.
func (u *AppUserThirdPartyUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the AppUserThirdPartyCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AppUserThirdPartyCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AppUserThirdPartyUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
