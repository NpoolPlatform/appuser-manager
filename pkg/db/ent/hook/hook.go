// Code generated by entc, DO NOT EDIT.

package hook

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent"
)

// The AppFunc type is an adapter to allow the use of ordinary
// function as App mutator.
type AppFunc func(context.Context, *ent.AppMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f AppFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.AppMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.AppMutation", m)
	}
	return f(ctx, mv)
}

// The AppControlFunc type is an adapter to allow the use of ordinary
// function as AppControl mutator.
type AppControlFunc func(context.Context, *ent.AppControlMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f AppControlFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.AppControlMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.AppControlMutation", m)
	}
	return f(ctx, mv)
}

// The AppRoleFunc type is an adapter to allow the use of ordinary
// function as AppRole mutator.
type AppRoleFunc func(context.Context, *ent.AppRoleMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f AppRoleFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.AppRoleMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.AppRoleMutation", m)
	}
	return f(ctx, mv)
}

// The AppRoleUserFunc type is an adapter to allow the use of ordinary
// function as AppRoleUser mutator.
type AppRoleUserFunc func(context.Context, *ent.AppRoleUserMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f AppRoleUserFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.AppRoleUserMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.AppRoleUserMutation", m)
	}
	return f(ctx, mv)
}

// The AppUserFunc type is an adapter to allow the use of ordinary
// function as AppUser mutator.
type AppUserFunc func(context.Context, *ent.AppUserMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f AppUserFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.AppUserMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.AppUserMutation", m)
	}
	return f(ctx, mv)
}

// The AppUserControlFunc type is an adapter to allow the use of ordinary
// function as AppUserControl mutator.
type AppUserControlFunc func(context.Context, *ent.AppUserControlMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f AppUserControlFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.AppUserControlMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.AppUserControlMutation", m)
	}
	return f(ctx, mv)
}

// The AppUserExtraFunc type is an adapter to allow the use of ordinary
// function as AppUserExtra mutator.
type AppUserExtraFunc func(context.Context, *ent.AppUserExtraMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f AppUserExtraFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.AppUserExtraMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.AppUserExtraMutation", m)
	}
	return f(ctx, mv)
}

// The AppUserSecretFunc type is an adapter to allow the use of ordinary
// function as AppUserSecret mutator.
type AppUserSecretFunc func(context.Context, *ent.AppUserSecretMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f AppUserSecretFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.AppUserSecretMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.AppUserSecretMutation", m)
	}
	return f(ctx, mv)
}

// The BanAppFunc type is an adapter to allow the use of ordinary
// function as BanApp mutator.
type BanAppFunc func(context.Context, *ent.BanAppMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f BanAppFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.BanAppMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.BanAppMutation", m)
	}
	return f(ctx, mv)
}

// The BanAppUserFunc type is an adapter to allow the use of ordinary
// function as BanAppUser mutator.
type BanAppUserFunc func(context.Context, *ent.BanAppUserMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f BanAppUserFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.BanAppUserMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.BanAppUserMutation", m)
	}
	return f(ctx, mv)
}

// Condition is a hook condition function.
type Condition func(context.Context, ent.Mutation) bool

// And groups conditions with the AND operator.
func And(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		if !first(ctx, m) || !second(ctx, m) {
			return false
		}
		for _, cond := range rest {
			if !cond(ctx, m) {
				return false
			}
		}
		return true
	}
}

// Or groups conditions with the OR operator.
func Or(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		if first(ctx, m) || second(ctx, m) {
			return true
		}
		for _, cond := range rest {
			if cond(ctx, m) {
				return true
			}
		}
		return false
	}
}

// Not negates a given condition.
func Not(cond Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		return !cond(ctx, m)
	}
}

// HasOp is a condition testing mutation operation.
func HasOp(op ent.Op) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		return m.Op().Is(op)
	}
}

// HasAddedFields is a condition validating `.AddedField` on fields.
func HasAddedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if _, exists := m.AddedField(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.AddedField(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasClearedFields is a condition validating `.FieldCleared` on fields.
func HasClearedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if exists := m.FieldCleared(field); !exists {
			return false
		}
		for _, field := range fields {
			if exists := m.FieldCleared(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasFields is a condition validating `.Field` on fields.
func HasFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if _, exists := m.Field(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.Field(field); !exists {
				return false
			}
		}
		return true
	}
}

// If executes the given hook under condition.
//
//	hook.If(ComputeAverage, And(HasFields(...), HasAddedFields(...)))
//
func If(hk ent.Hook, cond Condition) ent.Hook {
	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if cond(ctx, m) {
				return hk(next).Mutate(ctx, m)
			}
			return next.Mutate(ctx, m)
		})
	}
}

// On executes the given hook only for the given operation.
//
//	hook.On(Log, ent.Delete|ent.Create)
//
func On(hk ent.Hook, op ent.Op) ent.Hook {
	return If(hk, HasOp(op))
}

// Unless skips the given hook only for the given operation.
//
//	hook.Unless(Log, ent.Update|ent.UpdateOne)
//
func Unless(hk ent.Hook, op ent.Op) ent.Hook {
	return If(hk, Not(HasOp(op)))
}

// FixedError is a hook returning a fixed error.
func FixedError(err error) ent.Hook {
	return func(ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(context.Context, ent.Mutation) (ent.Value, error) {
			return nil, err
		})
	}
}

// Reject returns a hook that rejects all operations that match op.
//
//	func (T) Hooks() []ent.Hook {
//		return []ent.Hook{
//			Reject(ent.Delete|ent.Update),
//		}
//	}
//
func Reject(op ent.Op) ent.Hook {
	hk := FixedError(fmt.Errorf("%s operation is not allowed", op))
	return On(hk, op)
}

// Chain acts as a list of hooks and is effectively immutable.
// Once created, it will always hold the same set of hooks in the same order.
type Chain struct {
	hooks []ent.Hook
}

// NewChain creates a new chain of hooks.
func NewChain(hooks ...ent.Hook) Chain {
	return Chain{append([]ent.Hook(nil), hooks...)}
}

// Hook chains the list of hooks and returns the final hook.
func (c Chain) Hook() ent.Hook {
	return func(mutator ent.Mutator) ent.Mutator {
		for i := len(c.hooks) - 1; i >= 0; i-- {
			mutator = c.hooks[i](mutator)
		}
		return mutator
	}
}

// Append extends a chain, adding the specified hook
// as the last ones in the mutation flow.
func (c Chain) Append(hooks ...ent.Hook) Chain {
	newHooks := make([]ent.Hook, 0, len(c.hooks)+len(hooks))
	newHooks = append(newHooks, c.hooks...)
	newHooks = append(newHooks, hooks...)
	return Chain{newHooks}
}

// Extend extends a chain, adding the specified chain
// as the last ones in the mutation flow.
func (c Chain) Extend(chain Chain) Chain {
	return c.Append(chain.hooks...)
}
