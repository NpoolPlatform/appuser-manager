// Code generated by entc, DO NOT EDIT.

package appuser

import (
	"entgo.io/ent"
)

const (
	// Label holds the string label denoting the appuser type in the database.
	Label = "app_user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateAt holds the string denoting the create_at field in the database.
	FieldCreateAt = "create_at"
	// FieldUpdateAt holds the string denoting the update_at field in the database.
	FieldUpdateAt = "update_at"
	// FieldDeleteAt holds the string denoting the delete_at field in the database.
	FieldDeleteAt = "delete_at"
	// FieldAppID holds the string denoting the app_id field in the database.
	FieldAppID = "app_id"
	// FieldEmailAddress holds the string denoting the email_address field in the database.
	FieldEmailAddress = "email_address"
	// FieldPhoneNo holds the string denoting the phone_no field in the database.
	FieldPhoneNo = "phone_no"
	// FieldImportFromApp holds the string denoting the import_from_app field in the database.
	FieldImportFromApp = "import_from_app"
	// Table holds the table name of the appuser in the database.
	Table = "app_users"
)

// Columns holds all SQL columns for appuser fields.
var Columns = []string{
	FieldID,
	FieldCreateAt,
	FieldUpdateAt,
	FieldDeleteAt,
	FieldAppID,
	FieldEmailAddress,
	FieldPhoneNo,
	FieldImportFromApp,
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

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "github.com/NpoolPlatform/appuser-manager/pkg/db/ent/runtime"
//
var (
	Hooks  [1]ent.Hook
	Policy ent.Policy
	// DefaultCreateAt holds the default value on creation for the "create_at" field.
	DefaultCreateAt func() uint32
	// DefaultUpdateAt holds the default value on creation for the "update_at" field.
	DefaultUpdateAt func() uint32
	// UpdateDefaultUpdateAt holds the default value on update for the "update_at" field.
	UpdateDefaultUpdateAt func() uint32
	// DefaultDeleteAt holds the default value on creation for the "delete_at" field.
	DefaultDeleteAt func() uint32
)
