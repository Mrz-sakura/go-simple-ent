// Code generated by ent, DO NOT EDIT.

package video

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the video type in the database.
	Label = "video"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldURL holds the string denoting the url field in the database.
	FieldURL = "url"
	// FieldCover holds the string denoting the cover field in the database.
	FieldCover = "cover"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// Table holds the table name of the video in the database.
	Table = "videos"
)

// Columns holds all SQL columns for video fields.
var Columns = []string{
	FieldID,
	FieldTitle,
	FieldURL,
	FieldCover,
	FieldDescription,
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
	// TitleValidator is a validator for the "title" field. It is called by the builders before save.
	TitleValidator func(string) error
	// URLValidator is a validator for the "url" field. It is called by the builders before save.
	URLValidator func(string) error
)

// OrderOption defines the ordering options for the Video queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByTitle orders the results by the title field.
func ByTitle(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTitle, opts...).ToFunc()
}

// ByURL orders the results by the url field.
func ByURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldURL, opts...).ToFunc()
}

// ByCover orders the results by the cover field.
func ByCover(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCover, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}
