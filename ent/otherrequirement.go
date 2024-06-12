// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"runescape_http_server/ent/otherrequirement"
	"runescape_http_server/ent/unlock"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// OtherRequirement is the model entity for the OtherRequirement schema.
type OtherRequirement struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name"`
	// IsMember holds the value of the "is_member" field.
	IsMember int `json:"is_member"`
	// IsQuest holds the value of the "is_quest" field.
	IsQuest int `json:"is_quest"`
	// IsSkill holds the value of the "is_skill" field.
	IsSkill int `json:"is_skill"`
	// IDUnlock holds the value of the "id_unlock" field.
	IDUnlock int `json:"id_unlock"`
	// IDOfRequirement holds the value of the "id_of_requirement" field.
	IDOfRequirement int `json:"id_of_requirement"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the OtherRequirementQuery when eager-loading is set.
	Edges        OtherRequirementEdges `json:"edges"`
	selectValues sql.SelectValues
}

// OtherRequirementEdges holds the relations/edges for other nodes in the graph.
type OtherRequirementEdges struct {
	// OtherRequirementIDUnlockFk holds the value of the other_requirement_id_unlock_fk edge.
	OtherRequirementIDUnlockFk *Unlock `json:"other_requirement_id_unlock_fk,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// OtherRequirementIDUnlockFkOrErr returns the OtherRequirementIDUnlockFk value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OtherRequirementEdges) OtherRequirementIDUnlockFkOrErr() (*Unlock, error) {
	if e.OtherRequirementIDUnlockFk != nil {
		return e.OtherRequirementIDUnlockFk, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: unlock.Label}
	}
	return nil, &NotLoadedError{edge: "other_requirement_id_unlock_fk"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*OtherRequirement) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case otherrequirement.FieldID, otherrequirement.FieldIsMember, otherrequirement.FieldIsQuest, otherrequirement.FieldIsSkill, otherrequirement.FieldIDUnlock, otherrequirement.FieldIDOfRequirement:
			values[i] = new(sql.NullInt64)
		case otherrequirement.FieldName:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the OtherRequirement fields.
func (or *OtherRequirement) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case otherrequirement.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			or.ID = int(value.Int64)
		case otherrequirement.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				or.Name = value.String
			}
		case otherrequirement.FieldIsMember:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field is_member", values[i])
			} else if value.Valid {
				or.IsMember = int(value.Int64)
			}
		case otherrequirement.FieldIsQuest:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field is_quest", values[i])
			} else if value.Valid {
				or.IsQuest = int(value.Int64)
			}
		case otherrequirement.FieldIsSkill:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field is_skill", values[i])
			} else if value.Valid {
				or.IsSkill = int(value.Int64)
			}
		case otherrequirement.FieldIDUnlock:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field id_unlock", values[i])
			} else if value.Valid {
				or.IDUnlock = int(value.Int64)
			}
		case otherrequirement.FieldIDOfRequirement:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field id_of_requirement", values[i])
			} else if value.Valid {
				or.IDOfRequirement = int(value.Int64)
			}
		default:
			or.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the OtherRequirement.
// This includes values selected through modifiers, order, etc.
func (or *OtherRequirement) Value(name string) (ent.Value, error) {
	return or.selectValues.Get(name)
}

// QueryOtherRequirementIDUnlockFk queries the "other_requirement_id_unlock_fk" edge of the OtherRequirement entity.
func (or *OtherRequirement) QueryOtherRequirementIDUnlockFk() *UnlockQuery {
	return NewOtherRequirementClient(or.config).QueryOtherRequirementIDUnlockFk(or)
}

// Update returns a builder for updating this OtherRequirement.
// Note that you need to call OtherRequirement.Unwrap() before calling this method if this OtherRequirement
// was returned from a transaction, and the transaction was committed or rolled back.
func (or *OtherRequirement) Update() *OtherRequirementUpdateOne {
	return NewOtherRequirementClient(or.config).UpdateOne(or)
}

// Unwrap unwraps the OtherRequirement entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (or *OtherRequirement) Unwrap() *OtherRequirement {
	_tx, ok := or.config.driver.(*txDriver)
	if !ok {
		panic("ent: OtherRequirement is not a transactional entity")
	}
	or.config.driver = _tx.drv
	return or
}

// String implements the fmt.Stringer.
func (or *OtherRequirement) String() string {
	var builder strings.Builder
	builder.WriteString("OtherRequirement(")
	builder.WriteString(fmt.Sprintf("id=%v, ", or.ID))
	builder.WriteString("name=")
	builder.WriteString(or.Name)
	builder.WriteString(", ")
	builder.WriteString("is_member=")
	builder.WriteString(fmt.Sprintf("%v", or.IsMember))
	builder.WriteString(", ")
	builder.WriteString("is_quest=")
	builder.WriteString(fmt.Sprintf("%v", or.IsQuest))
	builder.WriteString(", ")
	builder.WriteString("is_skill=")
	builder.WriteString(fmt.Sprintf("%v", or.IsSkill))
	builder.WriteString(", ")
	builder.WriteString("id_unlock=")
	builder.WriteString(fmt.Sprintf("%v", or.IDUnlock))
	builder.WriteString(", ")
	builder.WriteString("id_of_requirement=")
	builder.WriteString(fmt.Sprintf("%v", or.IDOfRequirement))
	builder.WriteByte(')')
	return builder.String()
}

// OtherRequirements is a parsable slice of OtherRequirement.
type OtherRequirements []*OtherRequirement
