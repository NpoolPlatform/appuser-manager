// Code generated by entc, DO NOT EDIT.

package appuserthirdparty

import (
	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v uint32) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v uint32) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v uint32) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// AppID applies equality check predicate on the "app_id" field. It's identical to AppIDEQ.
func AppID(v uuid.UUID) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAppID), v))
	})
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v uuid.UUID) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserID), v))
	})
}

// ThirdPartyUserID applies equality check predicate on the "third_party_user_id" field. It's identical to ThirdPartyUserIDEQ.
func ThirdPartyUserID(v string) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldThirdPartyUserID), v))
	})
}

// ThirdPartyID applies equality check predicate on the "third_party_id" field. It's identical to ThirdPartyIDEQ.
func ThirdPartyID(v string) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldThirdPartyID), v))
	})
}

// ThirdPartyUsername applies equality check predicate on the "third_party_username" field. It's identical to ThirdPartyUsernameEQ.
func ThirdPartyUsername(v string) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldThirdPartyUsername), v))
	})
}

// ThirdPartyUserAvatar applies equality check predicate on the "third_party_user_avatar" field. It's identical to ThirdPartyUserAvatarEQ.
func ThirdPartyUserAvatar(v string) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldThirdPartyUserAvatar), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v uint32) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v uint32) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...uint32) predicate.AppUserThirdParty {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...uint32) predicate.AppUserThirdParty {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v uint32) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v uint32) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v uint32) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v uint32) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v uint32) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v uint32) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...uint32) predicate.AppUserThirdParty {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...uint32) predicate.AppUserThirdParty {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v uint32) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v uint32) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v uint32) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v uint32) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v uint32) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v uint32) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...uint32) predicate.AppUserThirdParty {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...uint32) predicate.AppUserThirdParty {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v uint32) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v uint32) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v uint32) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v uint32) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeletedAt), v))
	})
}

// AppIDEQ applies the EQ predicate on the "app_id" field.
func AppIDEQ(v uuid.UUID) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAppID), v))
	})
}

// AppIDNEQ applies the NEQ predicate on the "app_id" field.
func AppIDNEQ(v uuid.UUID) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAppID), v))
	})
}

// AppIDIn applies the In predicate on the "app_id" field.
func AppIDIn(vs ...uuid.UUID) predicate.AppUserThirdParty {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAppID), v...))
	})
}

// AppIDNotIn applies the NotIn predicate on the "app_id" field.
func AppIDNotIn(vs ...uuid.UUID) predicate.AppUserThirdParty {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAppID), v...))
	})
}

// AppIDGT applies the GT predicate on the "app_id" field.
func AppIDGT(v uuid.UUID) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAppID), v))
	})
}

// AppIDGTE applies the GTE predicate on the "app_id" field.
func AppIDGTE(v uuid.UUID) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAppID), v))
	})
}

// AppIDLT applies the LT predicate on the "app_id" field.
func AppIDLT(v uuid.UUID) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAppID), v))
	})
}

// AppIDLTE applies the LTE predicate on the "app_id" field.
func AppIDLTE(v uuid.UUID) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAppID), v))
	})
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v uuid.UUID) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserID), v))
	})
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v uuid.UUID) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUserID), v))
	})
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...uuid.UUID) predicate.AppUserThirdParty {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUserID), v...))
	})
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...uuid.UUID) predicate.AppUserThirdParty {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUserID), v...))
	})
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v uuid.UUID) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUserID), v))
	})
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v uuid.UUID) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUserID), v))
	})
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v uuid.UUID) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUserID), v))
	})
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v uuid.UUID) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUserID), v))
	})
}

// ThirdPartyUserIDEQ applies the EQ predicate on the "third_party_user_id" field.
func ThirdPartyUserIDEQ(v string) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldThirdPartyUserID), v))
	})
}

// ThirdPartyUserIDNEQ applies the NEQ predicate on the "third_party_user_id" field.
func ThirdPartyUserIDNEQ(v string) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldThirdPartyUserID), v))
	})
}

// ThirdPartyUserIDIn applies the In predicate on the "third_party_user_id" field.
func ThirdPartyUserIDIn(vs ...string) predicate.AppUserThirdParty {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldThirdPartyUserID), v...))
	})
}

// ThirdPartyUserIDNotIn applies the NotIn predicate on the "third_party_user_id" field.
func ThirdPartyUserIDNotIn(vs ...string) predicate.AppUserThirdParty {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldThirdPartyUserID), v...))
	})
}

// ThirdPartyUserIDGT applies the GT predicate on the "third_party_user_id" field.
func ThirdPartyUserIDGT(v string) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldThirdPartyUserID), v))
	})
}

// ThirdPartyUserIDGTE applies the GTE predicate on the "third_party_user_id" field.
func ThirdPartyUserIDGTE(v string) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldThirdPartyUserID), v))
	})
}

// ThirdPartyUserIDLT applies the LT predicate on the "third_party_user_id" field.
func ThirdPartyUserIDLT(v string) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldThirdPartyUserID), v))
	})
}

// ThirdPartyUserIDLTE applies the LTE predicate on the "third_party_user_id" field.
func ThirdPartyUserIDLTE(v string) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldThirdPartyUserID), v))
	})
}

// ThirdPartyUserIDContains applies the Contains predicate on the "third_party_user_id" field.
func ThirdPartyUserIDContains(v string) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldThirdPartyUserID), v))
	})
}

// ThirdPartyUserIDHasPrefix applies the HasPrefix predicate on the "third_party_user_id" field.
func ThirdPartyUserIDHasPrefix(v string) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldThirdPartyUserID), v))
	})
}

// ThirdPartyUserIDHasSuffix applies the HasSuffix predicate on the "third_party_user_id" field.
func ThirdPartyUserIDHasSuffix(v string) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldThirdPartyUserID), v))
	})
}

// ThirdPartyUserIDEqualFold applies the EqualFold predicate on the "third_party_user_id" field.
func ThirdPartyUserIDEqualFold(v string) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldThirdPartyUserID), v))
	})
}

// ThirdPartyUserIDContainsFold applies the ContainsFold predicate on the "third_party_user_id" field.
func ThirdPartyUserIDContainsFold(v string) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldThirdPartyUserID), v))
	})
}

// ThirdPartyIDEQ applies the EQ predicate on the "third_party_id" field.
func ThirdPartyIDEQ(v string) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldThirdPartyID), v))
	})
}

// ThirdPartyIDNEQ applies the NEQ predicate on the "third_party_id" field.
func ThirdPartyIDNEQ(v string) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldThirdPartyID), v))
	})
}

// ThirdPartyIDIn applies the In predicate on the "third_party_id" field.
func ThirdPartyIDIn(vs ...string) predicate.AppUserThirdParty {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldThirdPartyID), v...))
	})
}

// ThirdPartyIDNotIn applies the NotIn predicate on the "third_party_id" field.
func ThirdPartyIDNotIn(vs ...string) predicate.AppUserThirdParty {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldThirdPartyID), v...))
	})
}

// ThirdPartyIDGT applies the GT predicate on the "third_party_id" field.
func ThirdPartyIDGT(v string) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldThirdPartyID), v))
	})
}

// ThirdPartyIDGTE applies the GTE predicate on the "third_party_id" field.
func ThirdPartyIDGTE(v string) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldThirdPartyID), v))
	})
}

// ThirdPartyIDLT applies the LT predicate on the "third_party_id" field.
func ThirdPartyIDLT(v string) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldThirdPartyID), v))
	})
}

// ThirdPartyIDLTE applies the LTE predicate on the "third_party_id" field.
func ThirdPartyIDLTE(v string) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldThirdPartyID), v))
	})
}

// ThirdPartyIDContains applies the Contains predicate on the "third_party_id" field.
func ThirdPartyIDContains(v string) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldThirdPartyID), v))
	})
}

// ThirdPartyIDHasPrefix applies the HasPrefix predicate on the "third_party_id" field.
func ThirdPartyIDHasPrefix(v string) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldThirdPartyID), v))
	})
}

// ThirdPartyIDHasSuffix applies the HasSuffix predicate on the "third_party_id" field.
func ThirdPartyIDHasSuffix(v string) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldThirdPartyID), v))
	})
}

// ThirdPartyIDEqualFold applies the EqualFold predicate on the "third_party_id" field.
func ThirdPartyIDEqualFold(v string) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldThirdPartyID), v))
	})
}

// ThirdPartyIDContainsFold applies the ContainsFold predicate on the "third_party_id" field.
func ThirdPartyIDContainsFold(v string) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldThirdPartyID), v))
	})
}

// ThirdPartyUsernameEQ applies the EQ predicate on the "third_party_username" field.
func ThirdPartyUsernameEQ(v string) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldThirdPartyUsername), v))
	})
}

// ThirdPartyUsernameNEQ applies the NEQ predicate on the "third_party_username" field.
func ThirdPartyUsernameNEQ(v string) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldThirdPartyUsername), v))
	})
}

// ThirdPartyUsernameIn applies the In predicate on the "third_party_username" field.
func ThirdPartyUsernameIn(vs ...string) predicate.AppUserThirdParty {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldThirdPartyUsername), v...))
	})
}

// ThirdPartyUsernameNotIn applies the NotIn predicate on the "third_party_username" field.
func ThirdPartyUsernameNotIn(vs ...string) predicate.AppUserThirdParty {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldThirdPartyUsername), v...))
	})
}

// ThirdPartyUsernameGT applies the GT predicate on the "third_party_username" field.
func ThirdPartyUsernameGT(v string) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldThirdPartyUsername), v))
	})
}

// ThirdPartyUsernameGTE applies the GTE predicate on the "third_party_username" field.
func ThirdPartyUsernameGTE(v string) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldThirdPartyUsername), v))
	})
}

// ThirdPartyUsernameLT applies the LT predicate on the "third_party_username" field.
func ThirdPartyUsernameLT(v string) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldThirdPartyUsername), v))
	})
}

// ThirdPartyUsernameLTE applies the LTE predicate on the "third_party_username" field.
func ThirdPartyUsernameLTE(v string) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldThirdPartyUsername), v))
	})
}

// ThirdPartyUsernameContains applies the Contains predicate on the "third_party_username" field.
func ThirdPartyUsernameContains(v string) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldThirdPartyUsername), v))
	})
}

// ThirdPartyUsernameHasPrefix applies the HasPrefix predicate on the "third_party_username" field.
func ThirdPartyUsernameHasPrefix(v string) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldThirdPartyUsername), v))
	})
}

// ThirdPartyUsernameHasSuffix applies the HasSuffix predicate on the "third_party_username" field.
func ThirdPartyUsernameHasSuffix(v string) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldThirdPartyUsername), v))
	})
}

// ThirdPartyUsernameEqualFold applies the EqualFold predicate on the "third_party_username" field.
func ThirdPartyUsernameEqualFold(v string) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldThirdPartyUsername), v))
	})
}

// ThirdPartyUsernameContainsFold applies the ContainsFold predicate on the "third_party_username" field.
func ThirdPartyUsernameContainsFold(v string) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldThirdPartyUsername), v))
	})
}

// ThirdPartyUserAvatarEQ applies the EQ predicate on the "third_party_user_avatar" field.
func ThirdPartyUserAvatarEQ(v string) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldThirdPartyUserAvatar), v))
	})
}

// ThirdPartyUserAvatarNEQ applies the NEQ predicate on the "third_party_user_avatar" field.
func ThirdPartyUserAvatarNEQ(v string) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldThirdPartyUserAvatar), v))
	})
}

// ThirdPartyUserAvatarIn applies the In predicate on the "third_party_user_avatar" field.
func ThirdPartyUserAvatarIn(vs ...string) predicate.AppUserThirdParty {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldThirdPartyUserAvatar), v...))
	})
}

// ThirdPartyUserAvatarNotIn applies the NotIn predicate on the "third_party_user_avatar" field.
func ThirdPartyUserAvatarNotIn(vs ...string) predicate.AppUserThirdParty {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldThirdPartyUserAvatar), v...))
	})
}

// ThirdPartyUserAvatarGT applies the GT predicate on the "third_party_user_avatar" field.
func ThirdPartyUserAvatarGT(v string) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldThirdPartyUserAvatar), v))
	})
}

// ThirdPartyUserAvatarGTE applies the GTE predicate on the "third_party_user_avatar" field.
func ThirdPartyUserAvatarGTE(v string) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldThirdPartyUserAvatar), v))
	})
}

// ThirdPartyUserAvatarLT applies the LT predicate on the "third_party_user_avatar" field.
func ThirdPartyUserAvatarLT(v string) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldThirdPartyUserAvatar), v))
	})
}

// ThirdPartyUserAvatarLTE applies the LTE predicate on the "third_party_user_avatar" field.
func ThirdPartyUserAvatarLTE(v string) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldThirdPartyUserAvatar), v))
	})
}

// ThirdPartyUserAvatarContains applies the Contains predicate on the "third_party_user_avatar" field.
func ThirdPartyUserAvatarContains(v string) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldThirdPartyUserAvatar), v))
	})
}

// ThirdPartyUserAvatarHasPrefix applies the HasPrefix predicate on the "third_party_user_avatar" field.
func ThirdPartyUserAvatarHasPrefix(v string) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldThirdPartyUserAvatar), v))
	})
}

// ThirdPartyUserAvatarHasSuffix applies the HasSuffix predicate on the "third_party_user_avatar" field.
func ThirdPartyUserAvatarHasSuffix(v string) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldThirdPartyUserAvatar), v))
	})
}

// ThirdPartyUserAvatarEqualFold applies the EqualFold predicate on the "third_party_user_avatar" field.
func ThirdPartyUserAvatarEqualFold(v string) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldThirdPartyUserAvatar), v))
	})
}

// ThirdPartyUserAvatarContainsFold applies the ContainsFold predicate on the "third_party_user_avatar" field.
func ThirdPartyUserAvatarContainsFold(v string) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldThirdPartyUserAvatar), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.AppUserThirdParty) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.AppUserThirdParty) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.AppUserThirdParty) predicate.AppUserThirdParty {
	return predicate.AppUserThirdParty(func(s *sql.Selector) {
		p(s.Not())
	})
}
