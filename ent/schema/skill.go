package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Skill holds the schema definition for the Skill entity.
type Skill struct {
	ent.Schema
}

// Fields of the Skill.
func (Skill) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Default("unknown"),
		field.String("description").Default("unknown"),
		field.Bool("isMember").Default(false),
	}
}

// Edges of the Skill.
func (Skill) Edges() []ent.Edge {
	return nil
}
