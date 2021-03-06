// Code generated by entc, DO NOT EDIT.

package appcontrol

import (
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the appcontrol type in the database.
	Label = "app_control"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldAppID holds the string denoting the app_id field in the database.
	FieldAppID = "app_id"
	// FieldSignupMethods holds the string denoting the signup_methods field in the database.
	FieldSignupMethods = "signup_methods"
	// FieldExternSigninMethods holds the string denoting the extern_signin_methods field in the database.
	FieldExternSigninMethods = "extern_signin_methods"
	// FieldRecaptchaMethod holds the string denoting the recaptcha_method field in the database.
	FieldRecaptchaMethod = "recaptcha_method"
	// FieldKycEnable holds the string denoting the kyc_enable field in the database.
	FieldKycEnable = "kyc_enable"
	// FieldSigninVerifyEnable holds the string denoting the signin_verify_enable field in the database.
	FieldSigninVerifyEnable = "signin_verify_enable"
	// FieldInvitationCodeMust holds the string denoting the invitation_code_must field in the database.
	FieldInvitationCodeMust = "invitation_code_must"
	// FieldCreateAt holds the string denoting the create_at field in the database.
	FieldCreateAt = "create_at"
	// FieldUpdateAt holds the string denoting the update_at field in the database.
	FieldUpdateAt = "update_at"
	// FieldDeleteAt holds the string denoting the delete_at field in the database.
	FieldDeleteAt = "delete_at"
	// Table holds the table name of the appcontrol in the database.
	Table = "app_controls"
)

// Columns holds all SQL columns for appcontrol fields.
var Columns = []string{
	FieldID,
	FieldAppID,
	FieldSignupMethods,
	FieldExternSigninMethods,
	FieldRecaptchaMethod,
	FieldKycEnable,
	FieldSigninVerifyEnable,
	FieldInvitationCodeMust,
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
