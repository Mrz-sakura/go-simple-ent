// Code generated by ent, DO NOT EDIT.

package ent

import (
	"dwbackend/ent/video"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Video is the model entity for the Video schema.
type Video struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// URL holds the value of the "url" field.
	URL string `json:"url,omitempty"`
	// Cover holds the value of the "cover" field.
	Cover string `json:"cover,omitempty"`
	// Description holds the value of the "description" field.
	Description  string `json:"description,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Video) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case video.FieldID:
			values[i] = new(sql.NullInt64)
		case video.FieldTitle, video.FieldURL, video.FieldCover, video.FieldDescription:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Video fields.
func (v *Video) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case video.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			v.ID = int(value.Int64)
		case video.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				v.Title = value.String
			}
		case video.FieldURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field url", values[i])
			} else if value.Valid {
				v.URL = value.String
			}
		case video.FieldCover:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field cover", values[i])
			} else if value.Valid {
				v.Cover = value.String
			}
		case video.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				v.Description = value.String
			}
		default:
			v.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Video.
// This includes values selected through modifiers, order, etc.
func (v *Video) Value(name string) (ent.Value, error) {
	return v.selectValues.Get(name)
}

// Update returns a builder for updating this Video.
// Note that you need to call Video.Unwrap() before calling this method if this Video
// was returned from a transaction, and the transaction was committed or rolled back.
func (v *Video) Update() *VideoUpdateOne {
	return NewVideoClient(v.config).UpdateOne(v)
}

// Unwrap unwraps the Video entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (v *Video) Unwrap() *Video {
	_tx, ok := v.config.driver.(*txDriver)
	if !ok {
		panic("ent: Video is not a transactional entity")
	}
	v.config.driver = _tx.drv
	return v
}

// String implements the fmt.Stringer.
func (v *Video) String() string {
	var builder strings.Builder
	builder.WriteString("Video(")
	builder.WriteString(fmt.Sprintf("id=%v, ", v.ID))
	builder.WriteString("title=")
	builder.WriteString(v.Title)
	builder.WriteString(", ")
	builder.WriteString("url=")
	builder.WriteString(v.URL)
	builder.WriteString(", ")
	builder.WriteString("cover=")
	builder.WriteString(v.Cover)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(v.Description)
	builder.WriteByte(')')
	return builder.String()
}

// Videos is a parsable slice of Video.
type Videos []*Video