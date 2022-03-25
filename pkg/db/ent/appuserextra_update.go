// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent/appuserextra"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// AppUserExtraUpdate is the builder for updating AppUserExtra entities.
type AppUserExtraUpdate struct {
	config
	hooks    []Hook
	mutation *AppUserExtraMutation
}

// Where appends a list predicates to the AppUserExtraUpdate builder.
func (aueu *AppUserExtraUpdate) Where(ps ...predicate.AppUserExtra) *AppUserExtraUpdate {
	aueu.mutation.Where(ps...)
	return aueu
}

// SetAppID sets the "app_id" field.
func (aueu *AppUserExtraUpdate) SetAppID(u uuid.UUID) *AppUserExtraUpdate {
	aueu.mutation.SetAppID(u)
	return aueu
}

// SetUserID sets the "user_id" field.
func (aueu *AppUserExtraUpdate) SetUserID(u uuid.UUID) *AppUserExtraUpdate {
	aueu.mutation.SetUserID(u)
	return aueu
}

// SetUsername sets the "username" field.
func (aueu *AppUserExtraUpdate) SetUsername(s string) *AppUserExtraUpdate {
	aueu.mutation.SetUsername(s)
	return aueu
}

// SetFirstName sets the "first_name" field.
func (aueu *AppUserExtraUpdate) SetFirstName(s string) *AppUserExtraUpdate {
	aueu.mutation.SetFirstName(s)
	return aueu
}

// SetLastName sets the "last_name" field.
func (aueu *AppUserExtraUpdate) SetLastName(s string) *AppUserExtraUpdate {
	aueu.mutation.SetLastName(s)
	return aueu
}

// SetAddressFields sets the "address_fields" field.
func (aueu *AppUserExtraUpdate) SetAddressFields(s []string) *AppUserExtraUpdate {
	aueu.mutation.SetAddressFields(s)
	return aueu
}

// SetGender sets the "gender" field.
func (aueu *AppUserExtraUpdate) SetGender(s string) *AppUserExtraUpdate {
	aueu.mutation.SetGender(s)
	return aueu
}

// SetPostalCode sets the "postal_code" field.
func (aueu *AppUserExtraUpdate) SetPostalCode(s string) *AppUserExtraUpdate {
	aueu.mutation.SetPostalCode(s)
	return aueu
}

// SetAge sets the "age" field.
func (aueu *AppUserExtraUpdate) SetAge(u uint32) *AppUserExtraUpdate {
	aueu.mutation.ResetAge()
	aueu.mutation.SetAge(u)
	return aueu
}

// AddAge adds u to the "age" field.
func (aueu *AppUserExtraUpdate) AddAge(u int32) *AppUserExtraUpdate {
	aueu.mutation.AddAge(u)
	return aueu
}

// SetBirthday sets the "birthday" field.
func (aueu *AppUserExtraUpdate) SetBirthday(u uint32) *AppUserExtraUpdate {
	aueu.mutation.ResetBirthday()
	aueu.mutation.SetBirthday(u)
	return aueu
}

// AddBirthday adds u to the "birthday" field.
func (aueu *AppUserExtraUpdate) AddBirthday(u int32) *AppUserExtraUpdate {
	aueu.mutation.AddBirthday(u)
	return aueu
}

// SetAvatar sets the "avatar" field.
func (aueu *AppUserExtraUpdate) SetAvatar(s string) *AppUserExtraUpdate {
	aueu.mutation.SetAvatar(s)
	return aueu
}

// SetOrganization sets the "organization" field.
func (aueu *AppUserExtraUpdate) SetOrganization(s string) *AppUserExtraUpdate {
	aueu.mutation.SetOrganization(s)
	return aueu
}

// SetIDNumber sets the "id_number" field.
func (aueu *AppUserExtraUpdate) SetIDNumber(s string) *AppUserExtraUpdate {
	aueu.mutation.SetIDNumber(s)
	return aueu
}

// SetCreateAt sets the "create_at" field.
func (aueu *AppUserExtraUpdate) SetCreateAt(u uint32) *AppUserExtraUpdate {
	aueu.mutation.ResetCreateAt()
	aueu.mutation.SetCreateAt(u)
	return aueu
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (aueu *AppUserExtraUpdate) SetNillableCreateAt(u *uint32) *AppUserExtraUpdate {
	if u != nil {
		aueu.SetCreateAt(*u)
	}
	return aueu
}

// AddCreateAt adds u to the "create_at" field.
func (aueu *AppUserExtraUpdate) AddCreateAt(u int32) *AppUserExtraUpdate {
	aueu.mutation.AddCreateAt(u)
	return aueu
}

// SetUpdateAt sets the "update_at" field.
func (aueu *AppUserExtraUpdate) SetUpdateAt(u uint32) *AppUserExtraUpdate {
	aueu.mutation.ResetUpdateAt()
	aueu.mutation.SetUpdateAt(u)
	return aueu
}

// AddUpdateAt adds u to the "update_at" field.
func (aueu *AppUserExtraUpdate) AddUpdateAt(u int32) *AppUserExtraUpdate {
	aueu.mutation.AddUpdateAt(u)
	return aueu
}

// SetDeleteAt sets the "delete_at" field.
func (aueu *AppUserExtraUpdate) SetDeleteAt(u uint32) *AppUserExtraUpdate {
	aueu.mutation.ResetDeleteAt()
	aueu.mutation.SetDeleteAt(u)
	return aueu
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (aueu *AppUserExtraUpdate) SetNillableDeleteAt(u *uint32) *AppUserExtraUpdate {
	if u != nil {
		aueu.SetDeleteAt(*u)
	}
	return aueu
}

// AddDeleteAt adds u to the "delete_at" field.
func (aueu *AppUserExtraUpdate) AddDeleteAt(u int32) *AppUserExtraUpdate {
	aueu.mutation.AddDeleteAt(u)
	return aueu
}

// Mutation returns the AppUserExtraMutation object of the builder.
func (aueu *AppUserExtraUpdate) Mutation() *AppUserExtraMutation {
	return aueu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (aueu *AppUserExtraUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	aueu.defaults()
	if len(aueu.hooks) == 0 {
		affected, err = aueu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AppUserExtraMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			aueu.mutation = mutation
			affected, err = aueu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(aueu.hooks) - 1; i >= 0; i-- {
			if aueu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = aueu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, aueu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (aueu *AppUserExtraUpdate) SaveX(ctx context.Context) int {
	affected, err := aueu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (aueu *AppUserExtraUpdate) Exec(ctx context.Context) error {
	_, err := aueu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aueu *AppUserExtraUpdate) ExecX(ctx context.Context) {
	if err := aueu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (aueu *AppUserExtraUpdate) defaults() {
	if _, ok := aueu.mutation.UpdateAt(); !ok {
		v := appuserextra.UpdateDefaultUpdateAt()
		aueu.mutation.SetUpdateAt(v)
	}
}

func (aueu *AppUserExtraUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   appuserextra.Table,
			Columns: appuserextra.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: appuserextra.FieldID,
			},
		},
	}
	if ps := aueu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := aueu.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: appuserextra.FieldAppID,
		})
	}
	if value, ok := aueu.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: appuserextra.FieldUserID,
		})
	}
	if value, ok := aueu.mutation.Username(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: appuserextra.FieldUsername,
		})
	}
	if value, ok := aueu.mutation.FirstName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: appuserextra.FieldFirstName,
		})
	}
	if value, ok := aueu.mutation.LastName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: appuserextra.FieldLastName,
		})
	}
	if value, ok := aueu.mutation.AddressFields(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: appuserextra.FieldAddressFields,
		})
	}
	if value, ok := aueu.mutation.Gender(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: appuserextra.FieldGender,
		})
	}
	if value, ok := aueu.mutation.PostalCode(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: appuserextra.FieldPostalCode,
		})
	}
	if value, ok := aueu.mutation.Age(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appuserextra.FieldAge,
		})
	}
	if value, ok := aueu.mutation.AddedAge(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appuserextra.FieldAge,
		})
	}
	if value, ok := aueu.mutation.Birthday(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appuserextra.FieldBirthday,
		})
	}
	if value, ok := aueu.mutation.AddedBirthday(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appuserextra.FieldBirthday,
		})
	}
	if value, ok := aueu.mutation.Avatar(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: appuserextra.FieldAvatar,
		})
	}
	if value, ok := aueu.mutation.Organization(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: appuserextra.FieldOrganization,
		})
	}
	if value, ok := aueu.mutation.IDNumber(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: appuserextra.FieldIDNumber,
		})
	}
	if value, ok := aueu.mutation.CreateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appuserextra.FieldCreateAt,
		})
	}
	if value, ok := aueu.mutation.AddedCreateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appuserextra.FieldCreateAt,
		})
	}
	if value, ok := aueu.mutation.UpdateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appuserextra.FieldUpdateAt,
		})
	}
	if value, ok := aueu.mutation.AddedUpdateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appuserextra.FieldUpdateAt,
		})
	}
	if value, ok := aueu.mutation.DeleteAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appuserextra.FieldDeleteAt,
		})
	}
	if value, ok := aueu.mutation.AddedDeleteAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appuserextra.FieldDeleteAt,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, aueu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{appuserextra.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// AppUserExtraUpdateOne is the builder for updating a single AppUserExtra entity.
type AppUserExtraUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AppUserExtraMutation
}

// SetAppID sets the "app_id" field.
func (aueuo *AppUserExtraUpdateOne) SetAppID(u uuid.UUID) *AppUserExtraUpdateOne {
	aueuo.mutation.SetAppID(u)
	return aueuo
}

// SetUserID sets the "user_id" field.
func (aueuo *AppUserExtraUpdateOne) SetUserID(u uuid.UUID) *AppUserExtraUpdateOne {
	aueuo.mutation.SetUserID(u)
	return aueuo
}

// SetUsername sets the "username" field.
func (aueuo *AppUserExtraUpdateOne) SetUsername(s string) *AppUserExtraUpdateOne {
	aueuo.mutation.SetUsername(s)
	return aueuo
}

// SetFirstName sets the "first_name" field.
func (aueuo *AppUserExtraUpdateOne) SetFirstName(s string) *AppUserExtraUpdateOne {
	aueuo.mutation.SetFirstName(s)
	return aueuo
}

// SetLastName sets the "last_name" field.
func (aueuo *AppUserExtraUpdateOne) SetLastName(s string) *AppUserExtraUpdateOne {
	aueuo.mutation.SetLastName(s)
	return aueuo
}

// SetAddressFields sets the "address_fields" field.
func (aueuo *AppUserExtraUpdateOne) SetAddressFields(s []string) *AppUserExtraUpdateOne {
	aueuo.mutation.SetAddressFields(s)
	return aueuo
}

// SetGender sets the "gender" field.
func (aueuo *AppUserExtraUpdateOne) SetGender(s string) *AppUserExtraUpdateOne {
	aueuo.mutation.SetGender(s)
	return aueuo
}

// SetPostalCode sets the "postal_code" field.
func (aueuo *AppUserExtraUpdateOne) SetPostalCode(s string) *AppUserExtraUpdateOne {
	aueuo.mutation.SetPostalCode(s)
	return aueuo
}

// SetAge sets the "age" field.
func (aueuo *AppUserExtraUpdateOne) SetAge(u uint32) *AppUserExtraUpdateOne {
	aueuo.mutation.ResetAge()
	aueuo.mutation.SetAge(u)
	return aueuo
}

// AddAge adds u to the "age" field.
func (aueuo *AppUserExtraUpdateOne) AddAge(u int32) *AppUserExtraUpdateOne {
	aueuo.mutation.AddAge(u)
	return aueuo
}

// SetBirthday sets the "birthday" field.
func (aueuo *AppUserExtraUpdateOne) SetBirthday(u uint32) *AppUserExtraUpdateOne {
	aueuo.mutation.ResetBirthday()
	aueuo.mutation.SetBirthday(u)
	return aueuo
}

// AddBirthday adds u to the "birthday" field.
func (aueuo *AppUserExtraUpdateOne) AddBirthday(u int32) *AppUserExtraUpdateOne {
	aueuo.mutation.AddBirthday(u)
	return aueuo
}

// SetAvatar sets the "avatar" field.
func (aueuo *AppUserExtraUpdateOne) SetAvatar(s string) *AppUserExtraUpdateOne {
	aueuo.mutation.SetAvatar(s)
	return aueuo
}

// SetOrganization sets the "organization" field.
func (aueuo *AppUserExtraUpdateOne) SetOrganization(s string) *AppUserExtraUpdateOne {
	aueuo.mutation.SetOrganization(s)
	return aueuo
}

// SetIDNumber sets the "id_number" field.
func (aueuo *AppUserExtraUpdateOne) SetIDNumber(s string) *AppUserExtraUpdateOne {
	aueuo.mutation.SetIDNumber(s)
	return aueuo
}

// SetCreateAt sets the "create_at" field.
func (aueuo *AppUserExtraUpdateOne) SetCreateAt(u uint32) *AppUserExtraUpdateOne {
	aueuo.mutation.ResetCreateAt()
	aueuo.mutation.SetCreateAt(u)
	return aueuo
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (aueuo *AppUserExtraUpdateOne) SetNillableCreateAt(u *uint32) *AppUserExtraUpdateOne {
	if u != nil {
		aueuo.SetCreateAt(*u)
	}
	return aueuo
}

// AddCreateAt adds u to the "create_at" field.
func (aueuo *AppUserExtraUpdateOne) AddCreateAt(u int32) *AppUserExtraUpdateOne {
	aueuo.mutation.AddCreateAt(u)
	return aueuo
}

// SetUpdateAt sets the "update_at" field.
func (aueuo *AppUserExtraUpdateOne) SetUpdateAt(u uint32) *AppUserExtraUpdateOne {
	aueuo.mutation.ResetUpdateAt()
	aueuo.mutation.SetUpdateAt(u)
	return aueuo
}

// AddUpdateAt adds u to the "update_at" field.
func (aueuo *AppUserExtraUpdateOne) AddUpdateAt(u int32) *AppUserExtraUpdateOne {
	aueuo.mutation.AddUpdateAt(u)
	return aueuo
}

// SetDeleteAt sets the "delete_at" field.
func (aueuo *AppUserExtraUpdateOne) SetDeleteAt(u uint32) *AppUserExtraUpdateOne {
	aueuo.mutation.ResetDeleteAt()
	aueuo.mutation.SetDeleteAt(u)
	return aueuo
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (aueuo *AppUserExtraUpdateOne) SetNillableDeleteAt(u *uint32) *AppUserExtraUpdateOne {
	if u != nil {
		aueuo.SetDeleteAt(*u)
	}
	return aueuo
}

// AddDeleteAt adds u to the "delete_at" field.
func (aueuo *AppUserExtraUpdateOne) AddDeleteAt(u int32) *AppUserExtraUpdateOne {
	aueuo.mutation.AddDeleteAt(u)
	return aueuo
}

// Mutation returns the AppUserExtraMutation object of the builder.
func (aueuo *AppUserExtraUpdateOne) Mutation() *AppUserExtraMutation {
	return aueuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (aueuo *AppUserExtraUpdateOne) Select(field string, fields ...string) *AppUserExtraUpdateOne {
	aueuo.fields = append([]string{field}, fields...)
	return aueuo
}

// Save executes the query and returns the updated AppUserExtra entity.
func (aueuo *AppUserExtraUpdateOne) Save(ctx context.Context) (*AppUserExtra, error) {
	var (
		err  error
		node *AppUserExtra
	)
	aueuo.defaults()
	if len(aueuo.hooks) == 0 {
		node, err = aueuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AppUserExtraMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			aueuo.mutation = mutation
			node, err = aueuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(aueuo.hooks) - 1; i >= 0; i-- {
			if aueuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = aueuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, aueuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (aueuo *AppUserExtraUpdateOne) SaveX(ctx context.Context) *AppUserExtra {
	node, err := aueuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (aueuo *AppUserExtraUpdateOne) Exec(ctx context.Context) error {
	_, err := aueuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aueuo *AppUserExtraUpdateOne) ExecX(ctx context.Context) {
	if err := aueuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (aueuo *AppUserExtraUpdateOne) defaults() {
	if _, ok := aueuo.mutation.UpdateAt(); !ok {
		v := appuserextra.UpdateDefaultUpdateAt()
		aueuo.mutation.SetUpdateAt(v)
	}
}

func (aueuo *AppUserExtraUpdateOne) sqlSave(ctx context.Context) (_node *AppUserExtra, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   appuserextra.Table,
			Columns: appuserextra.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: appuserextra.FieldID,
			},
		},
	}
	id, ok := aueuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "AppUserExtra.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := aueuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, appuserextra.FieldID)
		for _, f := range fields {
			if !appuserextra.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != appuserextra.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := aueuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := aueuo.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: appuserextra.FieldAppID,
		})
	}
	if value, ok := aueuo.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: appuserextra.FieldUserID,
		})
	}
	if value, ok := aueuo.mutation.Username(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: appuserextra.FieldUsername,
		})
	}
	if value, ok := aueuo.mutation.FirstName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: appuserextra.FieldFirstName,
		})
	}
	if value, ok := aueuo.mutation.LastName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: appuserextra.FieldLastName,
		})
	}
	if value, ok := aueuo.mutation.AddressFields(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: appuserextra.FieldAddressFields,
		})
	}
	if value, ok := aueuo.mutation.Gender(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: appuserextra.FieldGender,
		})
	}
	if value, ok := aueuo.mutation.PostalCode(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: appuserextra.FieldPostalCode,
		})
	}
	if value, ok := aueuo.mutation.Age(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appuserextra.FieldAge,
		})
	}
	if value, ok := aueuo.mutation.AddedAge(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appuserextra.FieldAge,
		})
	}
	if value, ok := aueuo.mutation.Birthday(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appuserextra.FieldBirthday,
		})
	}
	if value, ok := aueuo.mutation.AddedBirthday(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appuserextra.FieldBirthday,
		})
	}
	if value, ok := aueuo.mutation.Avatar(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: appuserextra.FieldAvatar,
		})
	}
	if value, ok := aueuo.mutation.Organization(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: appuserextra.FieldOrganization,
		})
	}
	if value, ok := aueuo.mutation.IDNumber(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: appuserextra.FieldIDNumber,
		})
	}
	if value, ok := aueuo.mutation.CreateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appuserextra.FieldCreateAt,
		})
	}
	if value, ok := aueuo.mutation.AddedCreateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appuserextra.FieldCreateAt,
		})
	}
	if value, ok := aueuo.mutation.UpdateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appuserextra.FieldUpdateAt,
		})
	}
	if value, ok := aueuo.mutation.AddedUpdateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appuserextra.FieldUpdateAt,
		})
	}
	if value, ok := aueuo.mutation.DeleteAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appuserextra.FieldDeleteAt,
		})
	}
	if value, ok := aueuo.mutation.AddedDeleteAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appuserextra.FieldDeleteAt,
		})
	}
	_node = &AppUserExtra{config: aueuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, aueuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{appuserextra.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
