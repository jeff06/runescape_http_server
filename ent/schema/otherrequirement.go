package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// OtherRequirement holds the schema definition for the OtherRequirement entity.
type OtherRequirement struct {
	ent.Schema
}

// Fields of the OtherRequirement.
func (OtherRequirement) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").StructTag(`json:"name"`),
		field.Int("is_member").StructTag(`json:"is_member"`),
		field.Int("is_quest").StructTag(`json:"is_quest"`),
		field.Int("is_skill").StructTag(`json:"is_skill"`),
		field.Int("id_unlock").StructTag(`json:"id_unlock"`).Optional(),
		field.Int("id_of_requirement").StructTag(`json:"id_of_requirement"`),
	}
}

// Edges of the OtherRequirement.
func (OtherRequirement) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("other_requirement_id_unlock_fk", Unlock.Type).
			Ref("other_requirements").
			Unique().
			Field("id_unlock"),
	}
}
