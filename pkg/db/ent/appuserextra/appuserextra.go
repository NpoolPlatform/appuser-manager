// Code generated by entc, DO NOT EDIT.

package appuserextra

import (
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the appuserextra type in the database.
	Label = "app_user_extra"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldAppID holds the string denoting the app_id field in the database.
	FieldAppID = "app_id"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldUsername holds the string denoting the username field in the database.
	FieldUsername = "username"
	// FieldAddressFields holds the string denoting the address_fields field in the database.
	FieldAddressFields = "address_fields"
	// FieldGender holds the string denoting the gender field in the database.
	FieldGender = "gender"
	// FieldPostalCode holds the string denoting the postal_code field in the database.
	FieldPostalCode = "postal_code"
	// FieldAge holds the string denoting the age field in the database.
	FieldAge = "age"
	// FieldBirthday holds the string denoting the birthday field in the database.
	FieldBirthday = "birthday"
	// FieldAvatar holds the string denoting the avatar field in the database.
	FieldAvatar = "avatar"
	// FieldOrganization holds the string denoting the organization field in the database.
	FieldOrganization = "organization"
	// FieldCreateAt holds the string denoting the create_at field in the database.
	FieldCreateAt = "create_at"
	// FieldUpdateAt holds the string denoting the update_at field in the database.
	FieldUpdateAt = "update_at"
	// FieldDeleteAt holds the string denoting the delete_at field in the database.
	FieldDeleteAt = "delete_at"
	// Table holds the table name of the appuserextra in the database.
	Table = "app_user_extras"
)

// Columns holds all SQL columns for appuserextra fields.
var Columns = []string{
	FieldID,
	FieldAppID,
	FieldUserID,
	FieldUsername,
	FieldAddressFields,
	FieldGender,
	FieldPostalCode,
	FieldAge,
	FieldBirthday,
	FieldAvatar,
	FieldOrganization,
	FieldCreateAt,
	FieldUpdateAt,
	FieldDeleteAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreateAt holds the default value on creation for the "create_at" field.
	DefaultCreateAt func() uint32
	// DefaultUpdateAt holds the default value on creation for the "update_at" field.
	DefaultUpdateAt func() uint32
	// UpdateDefaultUpdateAt holds the default value on update for the "update_at" field.
	UpdateDefaultUpdateAt func() uint32
	// DefaultDeleteAt holds the default value on creation for the "delete_at" field.
	DefaultDeleteAt func() uint32
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
