package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// AppUserThird holds the schema definition for the AppUserThird entity.
type AppUserThird struct {
	ent.Schema
}

func (AppUserThird) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the AppUserThird.
func (AppUserThird) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.UUID("app_id", uuid.UUID{}),
		field.UUID("user_id", uuid.UUID{}),
		field.String("third_user_id"),
		field.String("third"),
		field.String("third_id"),
		field.String("third_user_name"),
		field.String("third_user_avatar").MaxLen(1024),
	}
}

// Edges of the AppUserThird.
func (AppUserThird) Edges() []ent.Edge {
	return nil
}

func (AppUserThird) Indexes() []ent.Index {
	return []ent.Index{}
}