// Code generated by ent, DO NOT EDIT.

package skill

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the skill type in the database.
	Label = "skill"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldIsMember holds the string denoting the ismember field in the database.
	FieldIsMember = "is_member"
	// Table holds the table name of the skill in the database.
	Table = "skills"
)

// Columns holds all SQL columns for skill fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldDescription,
	FieldIsMember,
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
	// DefaultName holds the default value on creation for the "name" field.
	DefaultName string
	// DefaultDescription holds the default value on creation for the "description" field.
	DefaultDescription string
	// DefaultIsMember holds the default value on creation for the "isMember" field.
	DefaultIsMember bool
)

// OrderOption defines the ordering options for the Skill queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByIsMember orders the results by the isMember field.
func ByIsMember(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsMember, opts...).ToFunc()
}
