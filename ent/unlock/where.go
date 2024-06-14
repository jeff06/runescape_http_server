// Code generated by ent, DO NOT EDIT.

package unlock

import (
	"runescape_http_server/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Unlock {
	return predicate.Unlock(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Unlock {
	return predicate.Unlock(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Unlock {
	return predicate.Unlock(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Unlock {
	return predicate.Unlock(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Unlock {
	return predicate.Unlock(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Unlock {
	return predicate.Unlock(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Unlock {
	return predicate.Unlock(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Unlock {
	return predicate.Unlock(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Unlock {
	return predicate.Unlock(sql.FieldLTE(FieldID, id))
}

// IDSkill applies equality check predicate on the "id_skill" field. It's identical to IDSkillEQ.
func IDSkill(v int) predicate.Unlock {
	return predicate.Unlock(sql.FieldEQ(FieldIDSkill, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Unlock {
	return predicate.Unlock(sql.FieldEQ(FieldName, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.Unlock {
	return predicate.Unlock(sql.FieldEQ(FieldDescription, v))
}

// OtherRequirement applies equality check predicate on the "other_requirement" field. It's identical to OtherRequirementEQ.
func OtherRequirement(v string) predicate.Unlock {
	return predicate.Unlock(sql.FieldEQ(FieldOtherRequirement, v))
}

// Level applies equality check predicate on the "level" field. It's identical to LevelEQ.
func Level(v int) predicate.Unlock {
	return predicate.Unlock(sql.FieldEQ(FieldLevel, v))
}

// IDSkillEQ applies the EQ predicate on the "id_skill" field.
func IDSkillEQ(v int) predicate.Unlock {
	return predicate.Unlock(sql.FieldEQ(FieldIDSkill, v))
}

// IDSkillNEQ applies the NEQ predicate on the "id_skill" field.
func IDSkillNEQ(v int) predicate.Unlock {
	return predicate.Unlock(sql.FieldNEQ(FieldIDSkill, v))
}

// IDSkillIn applies the In predicate on the "id_skill" field.
func IDSkillIn(vs ...int) predicate.Unlock {
	return predicate.Unlock(sql.FieldIn(FieldIDSkill, vs...))
}

// IDSkillNotIn applies the NotIn predicate on the "id_skill" field.
func IDSkillNotIn(vs ...int) predicate.Unlock {
	return predicate.Unlock(sql.FieldNotIn(FieldIDSkill, vs...))
}

// IDSkillIsNil applies the IsNil predicate on the "id_skill" field.
func IDSkillIsNil() predicate.Unlock {
	return predicate.Unlock(sql.FieldIsNull(FieldIDSkill))
}

// IDSkillNotNil applies the NotNil predicate on the "id_skill" field.
func IDSkillNotNil() predicate.Unlock {
	return predicate.Unlock(sql.FieldNotNull(FieldIDSkill))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Unlock {
	return predicate.Unlock(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Unlock {
	return predicate.Unlock(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Unlock {
	return predicate.Unlock(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Unlock {
	return predicate.Unlock(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Unlock {
	return predicate.Unlock(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Unlock {
	return predicate.Unlock(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Unlock {
	return predicate.Unlock(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Unlock {
	return predicate.Unlock(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Unlock {
	return predicate.Unlock(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Unlock {
	return predicate.Unlock(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Unlock {
	return predicate.Unlock(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Unlock {
	return predicate.Unlock(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Unlock {
	return predicate.Unlock(sql.FieldContainsFold(FieldName, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.Unlock {
	return predicate.Unlock(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.Unlock {
	return predicate.Unlock(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.Unlock {
	return predicate.Unlock(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.Unlock {
	return predicate.Unlock(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.Unlock {
	return predicate.Unlock(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.Unlock {
	return predicate.Unlock(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.Unlock {
	return predicate.Unlock(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.Unlock {
	return predicate.Unlock(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.Unlock {
	return predicate.Unlock(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.Unlock {
	return predicate.Unlock(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.Unlock {
	return predicate.Unlock(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.Unlock {
	return predicate.Unlock(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.Unlock {
	return predicate.Unlock(sql.FieldContainsFold(FieldDescription, v))
}

// OtherRequirementEQ applies the EQ predicate on the "other_requirement" field.
func OtherRequirementEQ(v string) predicate.Unlock {
	return predicate.Unlock(sql.FieldEQ(FieldOtherRequirement, v))
}

// OtherRequirementNEQ applies the NEQ predicate on the "other_requirement" field.
func OtherRequirementNEQ(v string) predicate.Unlock {
	return predicate.Unlock(sql.FieldNEQ(FieldOtherRequirement, v))
}

// OtherRequirementIn applies the In predicate on the "other_requirement" field.
func OtherRequirementIn(vs ...string) predicate.Unlock {
	return predicate.Unlock(sql.FieldIn(FieldOtherRequirement, vs...))
}

// OtherRequirementNotIn applies the NotIn predicate on the "other_requirement" field.
func OtherRequirementNotIn(vs ...string) predicate.Unlock {
	return predicate.Unlock(sql.FieldNotIn(FieldOtherRequirement, vs...))
}

// OtherRequirementGT applies the GT predicate on the "other_requirement" field.
func OtherRequirementGT(v string) predicate.Unlock {
	return predicate.Unlock(sql.FieldGT(FieldOtherRequirement, v))
}

// OtherRequirementGTE applies the GTE predicate on the "other_requirement" field.
func OtherRequirementGTE(v string) predicate.Unlock {
	return predicate.Unlock(sql.FieldGTE(FieldOtherRequirement, v))
}

// OtherRequirementLT applies the LT predicate on the "other_requirement" field.
func OtherRequirementLT(v string) predicate.Unlock {
	return predicate.Unlock(sql.FieldLT(FieldOtherRequirement, v))
}

// OtherRequirementLTE applies the LTE predicate on the "other_requirement" field.
func OtherRequirementLTE(v string) predicate.Unlock {
	return predicate.Unlock(sql.FieldLTE(FieldOtherRequirement, v))
}

// OtherRequirementContains applies the Contains predicate on the "other_requirement" field.
func OtherRequirementContains(v string) predicate.Unlock {
	return predicate.Unlock(sql.FieldContains(FieldOtherRequirement, v))
}

// OtherRequirementHasPrefix applies the HasPrefix predicate on the "other_requirement" field.
func OtherRequirementHasPrefix(v string) predicate.Unlock {
	return predicate.Unlock(sql.FieldHasPrefix(FieldOtherRequirement, v))
}

// OtherRequirementHasSuffix applies the HasSuffix predicate on the "other_requirement" field.
func OtherRequirementHasSuffix(v string) predicate.Unlock {
	return predicate.Unlock(sql.FieldHasSuffix(FieldOtherRequirement, v))
}

// OtherRequirementEqualFold applies the EqualFold predicate on the "other_requirement" field.
func OtherRequirementEqualFold(v string) predicate.Unlock {
	return predicate.Unlock(sql.FieldEqualFold(FieldOtherRequirement, v))
}

// OtherRequirementContainsFold applies the ContainsFold predicate on the "other_requirement" field.
func OtherRequirementContainsFold(v string) predicate.Unlock {
	return predicate.Unlock(sql.FieldContainsFold(FieldOtherRequirement, v))
}

// LevelEQ applies the EQ predicate on the "level" field.
func LevelEQ(v int) predicate.Unlock {
	return predicate.Unlock(sql.FieldEQ(FieldLevel, v))
}

// LevelNEQ applies the NEQ predicate on the "level" field.
func LevelNEQ(v int) predicate.Unlock {
	return predicate.Unlock(sql.FieldNEQ(FieldLevel, v))
}

// LevelIn applies the In predicate on the "level" field.
func LevelIn(vs ...int) predicate.Unlock {
	return predicate.Unlock(sql.FieldIn(FieldLevel, vs...))
}

// LevelNotIn applies the NotIn predicate on the "level" field.
func LevelNotIn(vs ...int) predicate.Unlock {
	return predicate.Unlock(sql.FieldNotIn(FieldLevel, vs...))
}

// LevelGT applies the GT predicate on the "level" field.
func LevelGT(v int) predicate.Unlock {
	return predicate.Unlock(sql.FieldGT(FieldLevel, v))
}

// LevelGTE applies the GTE predicate on the "level" field.
func LevelGTE(v int) predicate.Unlock {
	return predicate.Unlock(sql.FieldGTE(FieldLevel, v))
}

// LevelLT applies the LT predicate on the "level" field.
func LevelLT(v int) predicate.Unlock {
	return predicate.Unlock(sql.FieldLT(FieldLevel, v))
}

// LevelLTE applies the LTE predicate on the "level" field.
func LevelLTE(v int) predicate.Unlock {
	return predicate.Unlock(sql.FieldLTE(FieldLevel, v))
}

// HasUnlockIDSkillFk applies the HasEdge predicate on the "unlock_id_skill_fk" edge.
func HasUnlockIDSkillFk() predicate.Unlock {
	return predicate.Unlock(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UnlockIDSkillFkTable, UnlockIDSkillFkColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUnlockIDSkillFkWith applies the HasEdge predicate on the "unlock_id_skill_fk" edge with a given conditions (other predicates).
func HasUnlockIDSkillFkWith(preds ...predicate.Skill) predicate.Unlock {
	return predicate.Unlock(func(s *sql.Selector) {
		step := newUnlockIDSkillFkStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Unlock) predicate.Unlock {
	return predicate.Unlock(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Unlock) predicate.Unlock {
	return predicate.Unlock(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Unlock) predicate.Unlock {
	return predicate.Unlock(sql.NotPredicates(p))
}
