// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"runescape_http_server/ent/skill"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Skill is the model entity for the Skill schema.
type Skill struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// IsMember holds the value of the "isMember" field.
	IsMember     bool `json:"isMember,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Skill) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case skill.FieldIsMember:
			values[i] = new(sql.NullBool)
		case skill.FieldID:
			values[i] = new(sql.NullInt64)
		case skill.FieldName, skill.FieldDescription:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Skill fields.
func (s *Skill) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case skill.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			s.ID = int(value.Int64)
		case skill.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				s.Name = value.String
			}
		case skill.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				s.Description = value.String
			}
		case skill.FieldIsMember:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field isMember", values[i])
			} else if value.Valid {
				s.IsMember = value.Bool
			}
		default:
			s.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Skill.
// This includes values selected through modifiers, order, etc.
func (s *Skill) Value(name string) (ent.Value, error) {
	return s.selectValues.Get(name)
}

// Update returns a builder for updating this Skill.
// Note that you need to call Skill.Unwrap() before calling this method if this Skill
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Skill) Update() *SkillUpdateOne {
	return NewSkillClient(s.config).UpdateOne(s)
}

// Unwrap unwraps the Skill entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Skill) Unwrap() *Skill {
	_tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Skill is not a transactional entity")
	}
	s.config.driver = _tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Skill) String() string {
	var builder strings.Builder
	builder.WriteString("Skill(")
	builder.WriteString(fmt.Sprintf("id=%v, ", s.ID))
	builder.WriteString("name=")
	builder.WriteString(s.Name)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(s.Description)
	builder.WriteString(", ")
	builder.WriteString("isMember=")
	builder.WriteString(fmt.Sprintf("%v", s.IsMember))
	builder.WriteByte(')')
	return builder.String()
}

// Skills is a parsable slice of Skill.
type Skills []*Skill