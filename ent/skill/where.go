// Code generated by ent, DO NOT EDIT.

package skill

import (
	"runescape_http_server/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Skill {
	return predicate.Skill(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Skill {
	return predicate.Skill(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Skill {
	return predicate.Skill(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Skill {
	return predicate.Skill(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Skill {
	return predicate.Skill(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Skill {
	return predicate.Skill(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Skill {
	return predicate.Skill(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Skill {
	return predicate.Skill(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Skill {
	return predicate.Skill(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Skill {
	return predicate.Skill(sql.FieldEQ(FieldName, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.Skill {
	return predicate.Skill(sql.FieldEQ(FieldDescription, v))
}

// IsMember applies equality check predicate on the "is_member" field. It's identical to IsMemberEQ.
func IsMember(v int) predicate.Skill {
	return predicate.Skill(sql.FieldEQ(FieldIsMember, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Skill {
	return predicate.Skill(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Skill {
	return predicate.Skill(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Skill {
	return predicate.Skill(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Skill {
	return predicate.Skill(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Skill {
	return predicate.Skill(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Skill {
	return predicate.Skill(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Skill {
	return predicate.Skill(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Skill {
	return predicate.Skill(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Skill {
	return predicate.Skill(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Skill {
	return predicate.Skill(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Skill {
	return predicate.Skill(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Skill {
	return predicate.Skill(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Skill {
	return predicate.Skill(sql.FieldContainsFold(FieldName, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.Skill {
	return predicate.Skill(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.Skill {
	return predicate.Skill(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.Skill {
	return predicate.Skill(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.Skill {
	return predicate.Skill(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.Skill {
	return predicate.Skill(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.Skill {
	return predicate.Skill(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.Skill {
	return predicate.Skill(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.Skill {
	return predicate.Skill(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.Skill {
	return predicate.Skill(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.Skill {
	return predicate.Skill(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.Skill {
	return predicate.Skill(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.Skill {
	return predicate.Skill(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.Skill {
	return predicate.Skill(sql.FieldContainsFold(FieldDescription, v))
}

// IsMemberEQ applies the EQ predicate on the "is_member" field.
func IsMemberEQ(v int) predicate.Skill {
	return predicate.Skill(sql.FieldEQ(FieldIsMember, v))
}

// IsMemberNEQ applies the NEQ predicate on the "is_member" field.
func IsMemberNEQ(v int) predicate.Skill {
	return predicate.Skill(sql.FieldNEQ(FieldIsMember, v))
}

// IsMemberIn applies the In predicate on the "is_member" field.
func IsMemberIn(vs ...int) predicate.Skill {
	return predicate.Skill(sql.FieldIn(FieldIsMember, vs...))
}

// IsMemberNotIn applies the NotIn predicate on the "is_member" field.
func IsMemberNotIn(vs ...int) predicate.Skill {
	return predicate.Skill(sql.FieldNotIn(FieldIsMember, vs...))
}

// IsMemberGT applies the GT predicate on the "is_member" field.
func IsMemberGT(v int) predicate.Skill {
	return predicate.Skill(sql.FieldGT(FieldIsMember, v))
}

// IsMemberGTE applies the GTE predicate on the "is_member" field.
func IsMemberGTE(v int) predicate.Skill {
	return predicate.Skill(sql.FieldGTE(FieldIsMember, v))
}

// IsMemberLT applies the LT predicate on the "is_member" field.
func IsMemberLT(v int) predicate.Skill {
	return predicate.Skill(sql.FieldLT(FieldIsMember, v))
}

// IsMemberLTE applies the LTE predicate on the "is_member" field.
func IsMemberLTE(v int) predicate.Skill {
	return predicate.Skill(sql.FieldLTE(FieldIsMember, v))
}

// HasUnlocks applies the HasEdge predicate on the "unlocks" edge.
func HasUnlocks() predicate.Skill {
	return predicate.Skill(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, UnlocksTable, UnlocksColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUnlocksWith applies the HasEdge predicate on the "unlocks" edge with a given conditions (other predicates).
func HasUnlocksWith(preds ...predicate.Unlock) predicate.Skill {
	return predicate.Skill(func(s *sql.Selector) {
		step := newUnlocksStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Skill) predicate.Skill {
	return predicate.Skill(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Skill) predicate.Skill {
	return predicate.Skill(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Skill) predicate.Skill {
	return predicate.Skill(sql.NotPredicates(p))
}
