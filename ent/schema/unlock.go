package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Unlock holds the schema definition for the Unlock entity.
type Unlock struct {
	ent.Schema
}

// Fields of the Unlock.
func (Unlock) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id_skill").StructTag(`json:"idSkill"`).Optional(),
		field.String("name").StructTag(`json:"name"`),
		field.String("description").StructTag(`json:"description"`),
		field.Int("is_member").StructTag(`json:"is_member"`),
		field.Int("level").StructTag(`json:"level"`),
	}
}

// Edges of the Unlock.
func (Unlock) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("unlock_id_skill_fk", Skill.Type).
			Ref("unlocks").
			Unique().
			Field("id_skill"),
		edge.To("other_requirements", OtherRequirement.Type),
	}
}
