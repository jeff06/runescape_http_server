// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"runescape_http_server/ent/skill"
	"runescape_http_server/ent/unlock"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Unlock is the model entity for the Unlock schema.
type Unlock struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// IDSkill holds the value of the "id_skill" field.
	IDSkill int `json:"idSkill"`
	// Name holds the value of the "name" field.
	Name string `json:"name"`
	// Description holds the value of the "description" field.
	Description string `json:"description"`
	// OtherRequirement holds the value of the "other_requirement" field.
	OtherRequirement string `json:"other_requirement"`
	// Level holds the value of the "level" field.
	Level int `json:"level"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UnlockQuery when eager-loading is set.
	Edges        UnlockEdges `json:"edges"`
	selectValues sql.SelectValues
}

// UnlockEdges holds the relations/edges for other nodes in the graph.
type UnlockEdges struct {
	// UnlockIDSkillFk holds the value of the unlock_id_skill_fk edge.
	UnlockIDSkillFk *Skill `json:"unlock_id_skill_fk,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// UnlockIDSkillFkOrErr returns the UnlockIDSkillFk value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UnlockEdges) UnlockIDSkillFkOrErr() (*Skill, error) {
	if e.UnlockIDSkillFk != nil {
		return e.UnlockIDSkillFk, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: skill.Label}
	}
	return nil, &NotLoadedError{edge: "unlock_id_skill_fk"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Unlock) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case unlock.FieldID, unlock.FieldIDSkill, unlock.FieldLevel:
			values[i] = new(sql.NullInt64)
		case unlock.FieldName, unlock.FieldDescription, unlock.FieldOtherRequirement:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Unlock fields.
func (u *Unlock) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case unlock.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			u.ID = int(value.Int64)
		case unlock.FieldIDSkill:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field id_skill", values[i])
			} else if value.Valid {
				u.IDSkill = int(value.Int64)
			}
		case unlock.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				u.Name = value.String
			}
		case unlock.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				u.Description = value.String
			}
		case unlock.FieldOtherRequirement:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field other_requirement", values[i])
			} else if value.Valid {
				u.OtherRequirement = value.String
			}
		case unlock.FieldLevel:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field level", values[i])
			} else if value.Valid {
				u.Level = int(value.Int64)
			}
		default:
			u.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Unlock.
// This includes values selected through modifiers, order, etc.
func (u *Unlock) Value(name string) (ent.Value, error) {
	return u.selectValues.Get(name)
}

// QueryUnlockIDSkillFk queries the "unlock_id_skill_fk" edge of the Unlock entity.
func (u *Unlock) QueryUnlockIDSkillFk() *SkillQuery {
	return NewUnlockClient(u.config).QueryUnlockIDSkillFk(u)
}

// Update returns a builder for updating this Unlock.
// Note that you need to call Unlock.Unwrap() before calling this method if this Unlock
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *Unlock) Update() *UnlockUpdateOne {
	return NewUnlockClient(u.config).UpdateOne(u)
}

// Unwrap unwraps the Unlock entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *Unlock) Unwrap() *Unlock {
	_tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: Unlock is not a transactional entity")
	}
	u.config.driver = _tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *Unlock) String() string {
	var builder strings.Builder
	builder.WriteString("Unlock(")
	builder.WriteString(fmt.Sprintf("id=%v, ", u.ID))
	builder.WriteString("id_skill=")
	builder.WriteString(fmt.Sprintf("%v", u.IDSkill))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(u.Name)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(u.Description)
	builder.WriteString(", ")
	builder.WriteString("other_requirement=")
	builder.WriteString(u.OtherRequirement)
	builder.WriteString(", ")
	builder.WriteString("level=")
	builder.WriteString(fmt.Sprintf("%v", u.Level))
	builder.WriteByte(')')
	return builder.String()
}

// Unlocks is a parsable slice of Unlock.
type Unlocks []*Unlock
