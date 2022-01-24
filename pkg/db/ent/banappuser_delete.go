// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent/banappuser"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent/predicate"
)

// BanAppUserDelete is the builder for deleting a BanAppUser entity.
type BanAppUserDelete struct {
	config
	hooks    []Hook
	mutation *BanAppUserMutation
}

// Where appends a list predicates to the BanAppUserDelete builder.
func (baud *BanAppUserDelete) Where(ps ...predicate.BanAppUser) *BanAppUserDelete {
	baud.mutation.Where(ps...)
	return baud
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (baud *BanAppUserDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(baud.hooks) == 0 {
		affected, err = baud.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BanAppUserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			baud.mutation = mutation
			affected, err = baud.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(baud.hooks) - 1; i >= 0; i-- {
			if baud.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = baud.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, baud.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (baud *BanAppUserDelete) ExecX(ctx context.Context) int {
	n, err := baud.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (baud *BanAppUserDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: banappuser.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: banappuser.FieldID,
			},
		},
	}
	if ps := baud.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, baud.driver, _spec)
}

// BanAppUserDeleteOne is the builder for deleting a single BanAppUser entity.
type BanAppUserDeleteOne struct {
	baud *BanAppUserDelete
}

// Exec executes the deletion query.
func (baudo *BanAppUserDeleteOne) Exec(ctx context.Context) error {
	n, err := baudo.baud.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{banappuser.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (baudo *BanAppUserDeleteOne) ExecX(ctx context.Context) {
	baudo.baud.ExecX(ctx)
}
