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
	// AppRolesColumns holds the columns for the "app_roles" table.
	AppRolesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_by", Type: field.TypeUUID},
		{Name: "role", Type: field.TypeString},
		{Name: "description", Type: field.TypeString},
		{Name: "app_id", Type: field.TypeUUID},
		{Name: "default", Type: field.TypeBool},
		{Name: "create_at", Type: field.TypeUint32},
		{Name: "update_at", Type: field.TypeUint32},
		{Name: "delete_at", Type: field.TypeUint32},
	}
	// AppRolesTable holds the schema information for the "app_roles" table.
	AppRolesTable = &schema.Table{
		Name:       "app_roles",
		Columns:    AppRolesColumns,
		PrimaryKey: []*schema.Column{AppRolesColumns[0]},
	}
	// AppRoleUsersColumns holds the columns for the "app_role_users" table.
	AppRoleUsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "app_id", Type: field.TypeUUID},
		{Name: "role_id", Type: field.TypeUUID},
		{Name: "user_id", Type: field.TypeUUID},
		{Name: "create_at", Type: field.TypeUint32},
		{Name: "update_at", Type: field.TypeUint32},
		{Name: "delete_at", Type: field.TypeUint32},
	}
	// AppRoleUsersTable holds the schema information for the "app_role_users" table.
	AppRoleUsersTable = &schema.Table{
		Name:       "app_role_users",
		Columns:    AppRoleUsersColumns,
		PrimaryKey: []*schema.Column{AppRoleUsersColumns[0]},
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
	// AppUserExtrasColumns holds the columns for the "app_user_extras" table.
	AppUserExtrasColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "app_id", Type: field.TypeUUID},
		{Name: "user_id", Type: field.TypeUUID},
		{Name: "username", Type: field.TypeString},
		{Name: "address_fields", Type: field.TypeJSON},
		{Name: "gender", Type: field.TypeString},
		{Name: "postal_code", Type: field.TypeString},
		{Name: "age", Type: field.TypeUint32},
		{Name: "birthday", Type: field.TypeUint32},
		{Name: "avatar", Type: field.TypeString},
		{Name: "organization", Type: field.TypeString},
		{Name: "create_at", Type: field.TypeUint32},
		{Name: "update_at", Type: field.TypeUint32},
		{Name: "delete_at", Type: field.TypeUint32},
	}
	// AppUserExtrasTable holds the schema information for the "app_user_extras" table.
	AppUserExtrasTable = &schema.Table{
		Name:       "app_user_extras",
		Columns:    AppUserExtrasColumns,
		PrimaryKey: []*schema.Column{AppUserExtrasColumns[0]},
	}
	// AppUserSecretsColumns holds the columns for the "app_user_secrets" table.
	AppUserSecretsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "app_id", Type: field.TypeUUID},
		{Name: "user_id", Type: field.TypeUUID},
		{Name: "password_hash", Type: field.TypeString},
		{Name: "salt", Type: field.TypeString},
		{Name: "google_secret", Type: field.TypeString},
		{Name: "create_at", Type: field.TypeUint32},
		{Name: "update_at", Type: field.TypeUint32},
		{Name: "delete_at", Type: field.TypeUint32},
	}
	// AppUserSecretsTable holds the schema information for the "app_user_secrets" table.
	AppUserSecretsTable = &schema.Table{
		Name:       "app_user_secrets",
		Columns:    AppUserSecretsColumns,
		PrimaryKey: []*schema.Column{AppUserSecretsColumns[0]},
	}
	// BanAppsColumns holds the columns for the "ban_apps" table.
	BanAppsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "app_id", Type: field.TypeUUID},
		{Name: "message", Type: field.TypeString},
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
	// BanAppUsersColumns holds the columns for the "ban_app_users" table.
	BanAppUsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "app_id", Type: field.TypeUUID},
		{Name: "user_id", Type: field.TypeUUID},
		{Name: "create_at", Type: field.TypeUint32},
		{Name: "update_at", Type: field.TypeUint32},
		{Name: "delete_at", Type: field.TypeUint32},
	}
	// BanAppUsersTable holds the schema information for the "ban_app_users" table.
	BanAppUsersTable = &schema.Table{
		Name:       "ban_app_users",
		Columns:    BanAppUsersColumns,
		PrimaryKey: []*schema.Column{BanAppUsersColumns[0]},
	}
	// GenesisUsersColumns holds the columns for the "genesis_users" table.
	GenesisUsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "user_id", Type: field.TypeUUID},
		{Name: "create_at", Type: field.TypeUint32},
		{Name: "update_at", Type: field.TypeUint32},
		{Name: "delete_at", Type: field.TypeUint32},
	}
	// GenesisUsersTable holds the schema information for the "genesis_users" table.
	GenesisUsersTable = &schema.Table{
		Name:       "genesis_users",
		Columns:    GenesisUsersColumns,
		PrimaryKey: []*schema.Column{GenesisUsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AppsTable,
		AppControlsTable,
		AppRolesTable,
		AppRoleUsersTable,
		AppUsersTable,
		AppUserExtrasTable,
		AppUserSecretsTable,
		BanAppsTable,
		BanAppUsersTable,
		GenesisUsersTable,
	}
)

func init() {
}
