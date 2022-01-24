// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent/appuser"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent/predicate"
)

// AppUserDelete is the builder for deleting a AppUser entity.
type AppUserDelete struct {
	config
	hooks    []Hook
	mutation *AppUserMutation
}

// Where appends a list predicates to the AppUserDelete builder.
func (aud *AppUserDelete) Where(ps ...predicate.AppUser) *AppUserDelete {
	aud.mutation.Where(ps...)
	return aud
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (aud *AppUserDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(aud.hooks) == 0 {
		affected, err = aud.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AppUserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			aud.mutation = mutation
			affected, err = aud.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(aud.hooks) - 1; i >= 0; i-- {
			if aud.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = aud.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, aud.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (aud *AppUserDelete) ExecX(ctx context.Context) int {
	n, err := aud.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (aud *AppUserDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: appuser.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: appuser.FieldID,
			},
		},
	}
	if ps := aud.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, aud.driver, _spec)
}

// AppUserDeleteOne is the builder for deleting a single AppUser entity.
type AppUserDeleteOne struct {
	aud *AppUserDelete
}

// Exec executes the deletion query.
func (audo *AppUserDeleteOne) Exec(ctx context.Context) error {
	n, err := audo.aud.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{appuser.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (audo *AppUserDeleteOne) ExecX(ctx context.Context) {
	audo.aud.ExecX(ctx)
}
