// Code generated by entc, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent/appuserextra"
	"github.com/google/uuid"
)

// AppUserExtra is the model entity for the AppUserExtra schema.
type AppUserExtra struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// AppID holds the value of the "app_id" field.
	AppID uuid.UUID `json:"app_id,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID uuid.UUID `json:"user_id,omitempty"`
	// Username holds the value of the "username" field.
	Username string `json:"username,omitempty"`
	// AddressFields holds the value of the "address_fields" field.
	AddressFields []string `json:"address_fields,omitempty"`
	// Gender holds the value of the "gender" field.
	Gender string `json:"gender,omitempty"`
	// PostalCode holds the value of the "postal_code" field.
	PostalCode string `json:"postal_code,omitempty"`
	// Age holds the value of the "age" field.
	Age uint32 `json:"age,omitempty"`
	// Birthday holds the value of the "birthday" field.
	Birthday uint32 `json:"birthday,omitempty"`
	// Avatar holds the value of the "avatar" field.
	Avatar string `json:"avatar,omitempty"`
	// Organization holds the value of the "Organization" field.
	Organization string `json:"Organization,omitempty"`
	// CreateAt holds the value of the "create_at" field.
	CreateAt uint32 `json:"create_at,omitempty"`
	// UpdateAt holds the value of the "update_at" field.
	UpdateAt uint32 `json:"update_at,omitempty"`
	// DeleteAt holds the value of the "delete_at" field.
	DeleteAt uint32 `json:"delete_at,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*AppUserExtra) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case appuserextra.FieldAddressFields:
			values[i] = new([]byte)
		case appuserextra.FieldAge, appuserextra.FieldBirthday, appuserextra.FieldCreateAt, appuserextra.FieldUpdateAt, appuserextra.FieldDeleteAt:
			values[i] = new(sql.NullInt64)
		case appuserextra.FieldUsername, appuserextra.FieldGender, appuserextra.FieldPostalCode, appuserextra.FieldAvatar, appuserextra.FieldOrganization:
			values[i] = new(sql.NullString)
		case appuserextra.FieldID, appuserextra.FieldAppID, appuserextra.FieldUserID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type AppUserExtra", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the AppUserExtra fields.
func (aue *AppUserExtra) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case appuserextra.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				aue.ID = *value
			}
		case appuserextra.FieldAppID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field app_id", values[i])
			} else if value != nil {
				aue.AppID = *value
			}
		case appuserextra.FieldUserID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value != nil {
				aue.UserID = *value
			}
		case appuserextra.FieldUsername:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field username", values[i])
			} else if value.Valid {
				aue.Username = value.String
			}
		case appuserextra.FieldAddressFields:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field address_fields", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &aue.AddressFields); err != nil {
					return fmt.Errorf("unmarshal field address_fields: %w", err)
				}
			}
		case appuserextra.FieldGender:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field gender", values[i])
			} else if value.Valid {
				aue.Gender = value.String
			}
		case appuserextra.FieldPostalCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field postal_code", values[i])
			} else if value.Valid {
				aue.PostalCode = value.String
			}
		case appuserextra.FieldAge:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field age", values[i])
			} else if value.Valid {
				aue.Age = uint32(value.Int64)
			}
		case appuserextra.FieldBirthday:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field birthday", values[i])
			} else if value.Valid {
				aue.Birthday = uint32(value.Int64)
			}
		case appuserextra.FieldAvatar:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field avatar", values[i])
			} else if value.Valid {
				aue.Avatar = value.String
			}
		case appuserextra.FieldOrganization:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Organization", values[i])
			} else if value.Valid {
				aue.Organization = value.String
			}
		case appuserextra.FieldCreateAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field create_at", values[i])
			} else if value.Valid {
				aue.CreateAt = uint32(value.Int64)
			}
		case appuserextra.FieldUpdateAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field update_at", values[i])
			} else if value.Valid {
				aue.UpdateAt = uint32(value.Int64)
			}
		case appuserextra.FieldDeleteAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field delete_at", values[i])
			} else if value.Valid {
				aue.DeleteAt = uint32(value.Int64)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this AppUserExtra.
// Note that you need to call AppUserExtra.Unwrap() before calling this method if this AppUserExtra
// was returned from a transaction, and the transaction was committed or rolled back.
func (aue *AppUserExtra) Update() *AppUserExtraUpdateOne {
	return (&AppUserExtraClient{config: aue.config}).UpdateOne(aue)
}

// Unwrap unwraps the AppUserExtra entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (aue *AppUserExtra) Unwrap() *AppUserExtra {
	tx, ok := aue.config.driver.(*txDriver)
	if !ok {
		panic("ent: AppUserExtra is not a transactional entity")
	}
	aue.config.driver = tx.drv
	return aue
}

// String implements the fmt.Stringer.
func (aue *AppUserExtra) String() string {
	var builder strings.Builder
	builder.WriteString("AppUserExtra(")
	builder.WriteString(fmt.Sprintf("id=%v", aue.ID))
	builder.WriteString(", app_id=")
	builder.WriteString(fmt.Sprintf("%v", aue.AppID))
	builder.WriteString(", user_id=")
	builder.WriteString(fmt.Sprintf("%v", aue.UserID))
	builder.WriteString(", username=")
	builder.WriteString(aue.Username)
	builder.WriteString(", address_fields=")
	builder.WriteString(fmt.Sprintf("%v", aue.AddressFields))
	builder.WriteString(", gender=")
	builder.WriteString(aue.Gender)
	builder.WriteString(", postal_code=")
	builder.WriteString(aue.PostalCode)
	builder.WriteString(", age=")
	builder.WriteString(fmt.Sprintf("%v", aue.Age))
	builder.WriteString(", birthday=")
	builder.WriteString(fmt.Sprintf("%v", aue.Birthday))
	builder.WriteString(", avatar=")
	builder.WriteString(aue.Avatar)
	builder.WriteString(", Organization=")
	builder.WriteString(aue.Organization)
	builder.WriteString(", create_at=")
	builder.WriteString(fmt.Sprintf("%v", aue.CreateAt))
	builder.WriteString(", update_at=")
	builder.WriteString(fmt.Sprintf("%v", aue.UpdateAt))
	builder.WriteString(", delete_at=")
	builder.WriteString(fmt.Sprintf("%v", aue.DeleteAt))
	builder.WriteByte(')')
	return builder.String()
}

// AppUserExtras is a parsable slice of AppUserExtra.
type AppUserExtras []*AppUserExtra

func (aue AppUserExtras) config(cfg config) {
	for _i := range aue {
		aue[_i].config = cfg
	}
}
