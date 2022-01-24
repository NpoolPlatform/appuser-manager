// Code generated by entc, DO NOT EDIT.

package ent

import (
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent/app"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent/appcontrol"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent/approle"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent/appuser"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent/appuserextra"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent/appusersecret"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent/banapp"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent/banappuser"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent/schema"
	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	appFields := schema.App{}.Fields()
	_ = appFields
	// appDescCreateAt is the schema descriptor for create_at field.
	appDescCreateAt := appFields[5].Descriptor()
	// app.DefaultCreateAt holds the default value on creation for the create_at field.
	app.DefaultCreateAt = appDescCreateAt.Default.(func() uint32)
	// appDescUpdateAt is the schema descriptor for update_at field.
	appDescUpdateAt := appFields[6].Descriptor()
	// app.DefaultUpdateAt holds the default value on creation for the update_at field.
	app.DefaultUpdateAt = appDescUpdateAt.Default.(func() uint32)
	// app.UpdateDefaultUpdateAt holds the default value on update for the update_at field.
	app.UpdateDefaultUpdateAt = appDescUpdateAt.UpdateDefault.(func() uint32)
	// appDescDeleteAt is the schema descriptor for delete_at field.
	appDescDeleteAt := appFields[7].Descriptor()
	// app.DefaultDeleteAt holds the default value on creation for the delete_at field.
	app.DefaultDeleteAt = appDescDeleteAt.Default.(func() uint32)
	// appDescID is the schema descriptor for id field.
	appDescID := appFields[0].Descriptor()
	// app.DefaultID holds the default value on creation for the id field.
	app.DefaultID = appDescID.Default.(func() uuid.UUID)
	appcontrolFields := schema.AppControl{}.Fields()
	_ = appcontrolFields
	// appcontrolDescCreateAt is the schema descriptor for create_at field.
	appcontrolDescCreateAt := appcontrolFields[7].Descriptor()
	// appcontrol.DefaultCreateAt holds the default value on creation for the create_at field.
	appcontrol.DefaultCreateAt = appcontrolDescCreateAt.Default.(func() uint32)
	// appcontrolDescUpdateAt is the schema descriptor for update_at field.
	appcontrolDescUpdateAt := appcontrolFields[8].Descriptor()
	// appcontrol.DefaultUpdateAt holds the default value on creation for the update_at field.
	appcontrol.DefaultUpdateAt = appcontrolDescUpdateAt.Default.(func() uint32)
	// appcontrol.UpdateDefaultUpdateAt holds the default value on update for the update_at field.
	appcontrol.UpdateDefaultUpdateAt = appcontrolDescUpdateAt.UpdateDefault.(func() uint32)
	// appcontrolDescDeleteAt is the schema descriptor for delete_at field.
	appcontrolDescDeleteAt := appcontrolFields[9].Descriptor()
	// appcontrol.DefaultDeleteAt holds the default value on creation for the delete_at field.
	appcontrol.DefaultDeleteAt = appcontrolDescDeleteAt.Default.(func() uint32)
	// appcontrolDescID is the schema descriptor for id field.
	appcontrolDescID := appcontrolFields[0].Descriptor()
	// appcontrol.DefaultID holds the default value on creation for the id field.
	appcontrol.DefaultID = appcontrolDescID.Default.(func() uuid.UUID)
	approleFields := schema.AppRole{}.Fields()
	_ = approleFields
	// approleDescCreateAt is the schema descriptor for create_at field.
	approleDescCreateAt := approleFields[6].Descriptor()
	// approle.DefaultCreateAt holds the default value on creation for the create_at field.
	approle.DefaultCreateAt = approleDescCreateAt.Default.(func() uint32)
	// approleDescUpdateAt is the schema descriptor for update_at field.
	approleDescUpdateAt := approleFields[7].Descriptor()
	// approle.DefaultUpdateAt holds the default value on creation for the update_at field.
	approle.DefaultUpdateAt = approleDescUpdateAt.Default.(func() uint32)
	// approle.UpdateDefaultUpdateAt holds the default value on update for the update_at field.
	approle.UpdateDefaultUpdateAt = approleDescUpdateAt.UpdateDefault.(func() uint32)
	// approleDescDeleteAt is the schema descriptor for delete_at field.
	approleDescDeleteAt := approleFields[8].Descriptor()
	// approle.DefaultDeleteAt holds the default value on creation for the delete_at field.
	approle.DefaultDeleteAt = approleDescDeleteAt.Default.(func() uint32)
	// approleDescID is the schema descriptor for id field.
	approleDescID := approleFields[0].Descriptor()
	// approle.DefaultID holds the default value on creation for the id field.
	approle.DefaultID = approleDescID.Default.(func() uuid.UUID)
	appuserFields := schema.AppUser{}.Fields()
	_ = appuserFields
	// appuserDescCreateAt is the schema descriptor for create_at field.
	appuserDescCreateAt := appuserFields[5].Descriptor()
	// appuser.DefaultCreateAt holds the default value on creation for the create_at field.
	appuser.DefaultCreateAt = appuserDescCreateAt.Default.(func() uint32)
	// appuserDescUpdateAt is the schema descriptor for update_at field.
	appuserDescUpdateAt := appuserFields[6].Descriptor()
	// appuser.DefaultUpdateAt holds the default value on creation for the update_at field.
	appuser.DefaultUpdateAt = appuserDescUpdateAt.Default.(func() uint32)
	// appuser.UpdateDefaultUpdateAt holds the default value on update for the update_at field.
	appuser.UpdateDefaultUpdateAt = appuserDescUpdateAt.UpdateDefault.(func() uint32)
	// appuserDescDeleteAt is the schema descriptor for delete_at field.
	appuserDescDeleteAt := appuserFields[7].Descriptor()
	// appuser.DefaultDeleteAt holds the default value on creation for the delete_at field.
	appuser.DefaultDeleteAt = appuserDescDeleteAt.Default.(func() uint32)
	// appuserDescID is the schema descriptor for id field.
	appuserDescID := appuserFields[0].Descriptor()
	// appuser.DefaultID holds the default value on creation for the id field.
	appuser.DefaultID = appuserDescID.Default.(func() uuid.UUID)
	appuserextraFields := schema.AppUserExtra{}.Fields()
	_ = appuserextraFields
	// appuserextraDescCreateAt is the schema descriptor for create_at field.
	appuserextraDescCreateAt := appuserextraFields[11].Descriptor()
	// appuserextra.DefaultCreateAt holds the default value on creation for the create_at field.
	appuserextra.DefaultCreateAt = appuserextraDescCreateAt.Default.(func() uint32)
	// appuserextraDescUpdateAt is the schema descriptor for update_at field.
	appuserextraDescUpdateAt := appuserextraFields[12].Descriptor()
	// appuserextra.DefaultUpdateAt holds the default value on creation for the update_at field.
	appuserextra.DefaultUpdateAt = appuserextraDescUpdateAt.Default.(func() uint32)
	// appuserextra.UpdateDefaultUpdateAt holds the default value on update for the update_at field.
	appuserextra.UpdateDefaultUpdateAt = appuserextraDescUpdateAt.UpdateDefault.(func() uint32)
	// appuserextraDescDeleteAt is the schema descriptor for delete_at field.
	appuserextraDescDeleteAt := appuserextraFields[13].Descriptor()
	// appuserextra.DefaultDeleteAt holds the default value on creation for the delete_at field.
	appuserextra.DefaultDeleteAt = appuserextraDescDeleteAt.Default.(func() uint32)
	// appuserextraDescID is the schema descriptor for id field.
	appuserextraDescID := appuserextraFields[0].Descriptor()
	// appuserextra.DefaultID holds the default value on creation for the id field.
	appuserextra.DefaultID = appuserextraDescID.Default.(func() uuid.UUID)
	appusersecretFields := schema.AppUserSecret{}.Fields()
	_ = appusersecretFields
	// appusersecretDescCreateAt is the schema descriptor for create_at field.
	appusersecretDescCreateAt := appusersecretFields[6].Descriptor()
	// appusersecret.DefaultCreateAt holds the default value on creation for the create_at field.
	appusersecret.DefaultCreateAt = appusersecretDescCreateAt.Default.(func() uint32)
	// appusersecretDescUpdateAt is the schema descriptor for update_at field.
	appusersecretDescUpdateAt := appusersecretFields[7].Descriptor()
	// appusersecret.DefaultUpdateAt holds the default value on creation for the update_at field.
	appusersecret.DefaultUpdateAt = appusersecretDescUpdateAt.Default.(func() uint32)
	// appusersecret.UpdateDefaultUpdateAt holds the default value on update for the update_at field.
	appusersecret.UpdateDefaultUpdateAt = appusersecretDescUpdateAt.UpdateDefault.(func() uint32)
	// appusersecretDescDeleteAt is the schema descriptor for delete_at field.
	appusersecretDescDeleteAt := appusersecretFields[8].Descriptor()
	// appusersecret.DefaultDeleteAt holds the default value on creation for the delete_at field.
	appusersecret.DefaultDeleteAt = appusersecretDescDeleteAt.Default.(func() uint32)
	// appusersecretDescID is the schema descriptor for id field.
	appusersecretDescID := appusersecretFields[0].Descriptor()
	// appusersecret.DefaultID holds the default value on creation for the id field.
	appusersecret.DefaultID = appusersecretDescID.Default.(func() uuid.UUID)
	banappFields := schema.BanApp{}.Fields()
	_ = banappFields
	// banappDescCreateAt is the schema descriptor for create_at field.
	banappDescCreateAt := banappFields[2].Descriptor()
	// banapp.DefaultCreateAt holds the default value on creation for the create_at field.
	banapp.DefaultCreateAt = banappDescCreateAt.Default.(func() uint32)
	// banappDescUpdateAt is the schema descriptor for update_at field.
	banappDescUpdateAt := banappFields[3].Descriptor()
	// banapp.DefaultUpdateAt holds the default value on creation for the update_at field.
	banapp.DefaultUpdateAt = banappDescUpdateAt.Default.(func() uint32)
	// banapp.UpdateDefaultUpdateAt holds the default value on update for the update_at field.
	banapp.UpdateDefaultUpdateAt = banappDescUpdateAt.UpdateDefault.(func() uint32)
	// banappDescDeleteAt is the schema descriptor for delete_at field.
	banappDescDeleteAt := banappFields[4].Descriptor()
	// banapp.DefaultDeleteAt holds the default value on creation for the delete_at field.
	banapp.DefaultDeleteAt = banappDescDeleteAt.Default.(func() uint32)
	// banappDescID is the schema descriptor for id field.
	banappDescID := banappFields[0].Descriptor()
	// banapp.DefaultID holds the default value on creation for the id field.
	banapp.DefaultID = banappDescID.Default.(func() uuid.UUID)
	banappuserFields := schema.BanAppUser{}.Fields()
	_ = banappuserFields
	// banappuserDescCreateAt is the schema descriptor for create_at field.
	banappuserDescCreateAt := banappuserFields[3].Descriptor()
	// banappuser.DefaultCreateAt holds the default value on creation for the create_at field.
	banappuser.DefaultCreateAt = banappuserDescCreateAt.Default.(func() uint32)
	// banappuserDescUpdateAt is the schema descriptor for update_at field.
	banappuserDescUpdateAt := banappuserFields[4].Descriptor()
	// banappuser.DefaultUpdateAt holds the default value on creation for the update_at field.
	banappuser.DefaultUpdateAt = banappuserDescUpdateAt.Default.(func() uint32)
	// banappuser.UpdateDefaultUpdateAt holds the default value on update for the update_at field.
	banappuser.UpdateDefaultUpdateAt = banappuserDescUpdateAt.UpdateDefault.(func() uint32)
	// banappuserDescDeleteAt is the schema descriptor for delete_at field.
	banappuserDescDeleteAt := banappuserFields[5].Descriptor()
	// banappuser.DefaultDeleteAt holds the default value on creation for the delete_at field.
	banappuser.DefaultDeleteAt = banappuserDescDeleteAt.Default.(func() uint32)
	// banappuserDescID is the schema descriptor for id field.
	banappuserDescID := banappuserFields[0].Descriptor()
	// banappuser.DefaultID holds the default value on creation for the id field.
	banappuser.DefaultID = banappuserDescID.Default.(func() uuid.UUID)
}
