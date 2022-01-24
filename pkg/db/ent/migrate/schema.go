// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AppsColumns holds the columns for the "apps" table.
	AppsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_by", Type: field.TypeUUID},
		{Name: "name", Type: field.TypeString},
		{Name: "logo", Type: field.TypeString},
		{Name: "description", Type: field.TypeString},
		{Name: "create_at", Type: field.TypeUint32},
		{Name: "update_at", Type: field.TypeUint32},
		{Name: "delete_at", Type: field.TypeUint32},
	}
	// AppsTable holds the schema information for the "apps" table.
	AppsTable = &schema.Table{
		Name:       "apps",
		Columns:    AppsColumns,
		PrimaryKey: []*schema.Column{AppsColumns[0]},
	}
	// AppControlsColumns holds the columns for the "app_controls" table.
	AppControlsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "app_id", Type: field.TypeUUID, Unique: true},
		{Name: "signup_methods", Type: field.TypeJSON},
		{Name: "extern_signin_methods", Type: field.TypeJSON},
		{Name: "recaptcha_method", Type: field.TypeString},
		{Name: "kyc_enable", Type: field.TypeBool},
		{Name: "signin_verify_enable", Type: field.TypeBool},
		{Name: "create_at", Type: field.TypeUint32},
		{Name: "update_at", Type: field.TypeUint32},
		{Name: "delete_at", Type: field.TypeUint32},
	}
	// AppControlsTable holds the schema information for the "app_controls" table.
	AppControlsTable = &schema.Table{
		Name:       "app_controls",
		Columns:    AppControlsColumns,
		PrimaryKey: []*schema.Column{AppControlsColumns[0]},
	}
	// AppUsersColumns holds the columns for the "app_users" table.
	AppUsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "app_id", Type: field.TypeUUID},
		{Name: "email_address", Type: field.TypeString},
		{Name: "phone_no", Type: field.TypeString},
		{Name: "import_from_app", Type: field.TypeUUID},
		{Name: "create_at", Type: field.TypeUint32},
		{Name: "update_at", Type: field.TypeUint32},
		{Name: "delete_at", Type: field.TypeUint32},
	}
	// AppUsersTable holds the schema information for the "app_users" table.
	AppUsersTable = &schema.Table{
		Name:       "app_users",
		Columns:    AppUsersColumns,
		PrimaryKey: []*schema.Column{AppUsersColumns[0]},
	}
	// BanAppsColumns holds the columns for the "ban_apps" table.
	BanAppsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "app_id", Type: field.TypeUUID},
		{Name: "create_at", Type: field.TypeUint32},
		{Name: "update_at", Type: field.TypeUint32},
		{Name: "delete_at", Type: field.TypeUint32},
	}
	// BanAppsTable holds the schema information for the "ban_apps" table.
	BanAppsTable = &schema.Table{
		Name:       "ban_apps",
		Columns:    BanAppsColumns,
		PrimaryKey: []*schema.Column{BanAppsColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AppsTable,
		AppControlsTable,
		AppUsersTable,
		BanAppsTable,
	}
)

func init() {
}
