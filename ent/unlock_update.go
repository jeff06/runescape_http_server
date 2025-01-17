// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"runescape_http_server/ent/predicate"
	"runescape_http_server/ent/skill"
	"runescape_http_server/ent/unlock"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UnlockUpdate is the builder for updating Unlock entities.
type UnlockUpdate struct {
	config
	hooks    []Hook
	mutation *UnlockMutation
}

// Where appends a list predicates to the UnlockUpdate builder.
func (uu *UnlockUpdate) Where(ps ...predicate.Unlock) *UnlockUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetIDSkill sets the "id_skill" field.
func (uu *UnlockUpdate) SetIDSkill(i int) *UnlockUpdate {
	uu.mutation.SetIDSkill(i)
	return uu
}

// SetNillableIDSkill sets the "id_skill" field if the given value is not nil.
func (uu *UnlockUpdate) SetNillableIDSkill(i *int) *UnlockUpdate {
	if i != nil {
		uu.SetIDSkill(*i)
	}
	return uu
}

// ClearIDSkill clears the value of the "id_skill" field.
func (uu *UnlockUpdate) ClearIDSkill() *UnlockUpdate {
	uu.mutation.ClearIDSkill()
	return uu
}

// SetName sets the "name" field.
func (uu *UnlockUpdate) SetName(s string) *UnlockUpdate {
	uu.mutation.SetName(s)
	return uu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (uu *UnlockUpdate) SetNillableName(s *string) *UnlockUpdate {
	if s != nil {
		uu.SetName(*s)
	}
	return uu
}

// SetDescription sets the "description" field.
func (uu *UnlockUpdate) SetDescription(s string) *UnlockUpdate {
	uu.mutation.SetDescription(s)
	return uu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (uu *UnlockUpdate) SetNillableDescription(s *string) *UnlockUpdate {
	if s != nil {
		uu.SetDescription(*s)
	}
	return uu
}

// SetOtherRequirement sets the "other_requirement" field.
func (uu *UnlockUpdate) SetOtherRequirement(s string) *UnlockUpdate {
	uu.mutation.SetOtherRequirement(s)
	return uu
}

// SetNillableOtherRequirement sets the "other_requirement" field if the given value is not nil.
func (uu *UnlockUpdate) SetNillableOtherRequirement(s *string) *UnlockUpdate {
	if s != nil {
		uu.SetOtherRequirement(*s)
	}
	return uu
}

// SetLevel sets the "level" field.
func (uu *UnlockUpdate) SetLevel(i int) *UnlockUpdate {
	uu.mutation.ResetLevel()
	uu.mutation.SetLevel(i)
	return uu
}

// SetNillableLevel sets the "level" field if the given value is not nil.
func (uu *UnlockUpdate) SetNillableLevel(i *int) *UnlockUpdate {
	if i != nil {
		uu.SetLevel(*i)
	}
	return uu
}

// AddLevel adds i to the "level" field.
func (uu *UnlockUpdate) AddLevel(i int) *UnlockUpdate {
	uu.mutation.AddLevel(i)
	return uu
}

// SetUnlockIDSkillFkID sets the "unlock_id_skill_fk" edge to the Skill entity by ID.
func (uu *UnlockUpdate) SetUnlockIDSkillFkID(id int) *UnlockUpdate {
	uu.mutation.SetUnlockIDSkillFkID(id)
	return uu
}

// SetNillableUnlockIDSkillFkID sets the "unlock_id_skill_fk" edge to the Skill entity by ID if the given value is not nil.
func (uu *UnlockUpdate) SetNillableUnlockIDSkillFkID(id *int) *UnlockUpdate {
	if id != nil {
		uu = uu.SetUnlockIDSkillFkID(*id)
	}
	return uu
}

// SetUnlockIDSkillFk sets the "unlock_id_skill_fk" edge to the Skill entity.
func (uu *UnlockUpdate) SetUnlockIDSkillFk(s *Skill) *UnlockUpdate {
	return uu.SetUnlockIDSkillFkID(s.ID)
}

// Mutation returns the UnlockMutation object of the builder.
func (uu *UnlockUpdate) Mutation() *UnlockMutation {
	return uu.mutation
}

// ClearUnlockIDSkillFk clears the "unlock_id_skill_fk" edge to the Skill entity.
func (uu *UnlockUpdate) ClearUnlockIDSkillFk() *UnlockUpdate {
	uu.mutation.ClearUnlockIDSkillFk()
	return uu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UnlockUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, uu.sqlSave, uu.mutation, uu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UnlockUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UnlockUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UnlockUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uu *UnlockUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(unlock.Table, unlock.Columns, sqlgraph.NewFieldSpec(unlock.FieldID, field.TypeInt))
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.Name(); ok {
		_spec.SetField(unlock.FieldName, field.TypeString, value)
	}
	if value, ok := uu.mutation.Description(); ok {
		_spec.SetField(unlock.FieldDescription, field.TypeString, value)
	}
	if value, ok := uu.mutation.OtherRequirement(); ok {
		_spec.SetField(unlock.FieldOtherRequirement, field.TypeString, value)
	}
	if value, ok := uu.mutation.Level(); ok {
		_spec.SetField(unlock.FieldLevel, field.TypeInt, value)
	}
	if value, ok := uu.mutation.AddedLevel(); ok {
		_spec.AddField(unlock.FieldLevel, field.TypeInt, value)
	}
	if uu.mutation.UnlockIDSkillFkCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   unlock.UnlockIDSkillFkTable,
			Columns: []string{unlock.UnlockIDSkillFkColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(skill.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.UnlockIDSkillFkIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   unlock.UnlockIDSkillFkTable,
			Columns: []string{unlock.UnlockIDSkillFkColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(skill.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{unlock.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	uu.mutation.done = true
	return n, nil
}

// UnlockUpdateOne is the builder for updating a single Unlock entity.
type UnlockUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UnlockMutation
}

// SetIDSkill sets the "id_skill" field.
func (uuo *UnlockUpdateOne) SetIDSkill(i int) *UnlockUpdateOne {
	uuo.mutation.SetIDSkill(i)
	return uuo
}

// SetNillableIDSkill sets the "id_skill" field if the given value is not nil.
func (uuo *UnlockUpdateOne) SetNillableIDSkill(i *int) *UnlockUpdateOne {
	if i != nil {
		uuo.SetIDSkill(*i)
	}
	return uuo
}

// ClearIDSkill clears the value of the "id_skill" field.
func (uuo *UnlockUpdateOne) ClearIDSkill() *UnlockUpdateOne {
	uuo.mutation.ClearIDSkill()
	return uuo
}

// SetName sets the "name" field.
func (uuo *UnlockUpdateOne) SetName(s string) *UnlockUpdateOne {
	uuo.mutation.SetName(s)
	return uuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (uuo *UnlockUpdateOne) SetNillableName(s *string) *UnlockUpdateOne {
	if s != nil {
		uuo.SetName(*s)
	}
	return uuo
}

// SetDescription sets the "description" field.
func (uuo *UnlockUpdateOne) SetDescription(s string) *UnlockUpdateOne {
	uuo.mutation.SetDescription(s)
	return uuo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (uuo *UnlockUpdateOne) SetNillableDescription(s *string) *UnlockUpdateOne {
	if s != nil {
		uuo.SetDescription(*s)
	}
	return uuo
}

// SetOtherRequirement sets the "other_requirement" field.
func (uuo *UnlockUpdateOne) SetOtherRequirement(s string) *UnlockUpdateOne {
	uuo.mutation.SetOtherRequirement(s)
	return uuo
}

// SetNillableOtherRequirement sets the "other_requirement" field if the given value is not nil.
func (uuo *UnlockUpdateOne) SetNillableOtherRequirement(s *string) *UnlockUpdateOne {
	if s != nil {
		uuo.SetOtherRequirement(*s)
	}
	return uuo
}

// SetLevel sets the "level" field.
func (uuo *UnlockUpdateOne) SetLevel(i int) *UnlockUpdateOne {
	uuo.mutation.ResetLevel()
	uuo.mutation.SetLevel(i)
	return uuo
}

// SetNillableLevel sets the "level" field if the given value is not nil.
func (uuo *UnlockUpdateOne) SetNillableLevel(i *int) *UnlockUpdateOne {
	if i != nil {
		uuo.SetLevel(*i)
	}
	return uuo
}

// AddLevel adds i to the "level" field.
func (uuo *UnlockUpdateOne) AddLevel(i int) *UnlockUpdateOne {
	uuo.mutation.AddLevel(i)
	return uuo
}

// SetUnlockIDSkillFkID sets the "unlock_id_skill_fk" edge to the Skill entity by ID.
func (uuo *UnlockUpdateOne) SetUnlockIDSkillFkID(id int) *UnlockUpdateOne {
	uuo.mutation.SetUnlockIDSkillFkID(id)
	return uuo
}

// SetNillableUnlockIDSkillFkID sets the "unlock_id_skill_fk" edge to the Skill entity by ID if the given value is not nil.
func (uuo *UnlockUpdateOne) SetNillableUnlockIDSkillFkID(id *int) *UnlockUpdateOne {
	if id != nil {
		uuo = uuo.SetUnlockIDSkillFkID(*id)
	}
	return uuo
}

// SetUnlockIDSkillFk sets the "unlock_id_skill_fk" edge to the Skill entity.
func (uuo *UnlockUpdateOne) SetUnlockIDSkillFk(s *Skill) *UnlockUpdateOne {
	return uuo.SetUnlockIDSkillFkID(s.ID)
}

// Mutation returns the UnlockMutation object of the builder.
func (uuo *UnlockUpdateOne) Mutation() *UnlockMutation {
	return uuo.mutation
}

// ClearUnlockIDSkillFk clears the "unlock_id_skill_fk" edge to the Skill entity.
func (uuo *UnlockUpdateOne) ClearUnlockIDSkillFk() *UnlockUpdateOne {
	uuo.mutation.ClearUnlockIDSkillFk()
	return uuo
}

// Where appends a list predicates to the UnlockUpdate builder.
func (uuo *UnlockUpdateOne) Where(ps ...predicate.Unlock) *UnlockUpdateOne {
	uuo.mutation.Where(ps...)
	return uuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UnlockUpdateOne) Select(field string, fields ...string) *UnlockUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated Unlock entity.
func (uuo *UnlockUpdateOne) Save(ctx context.Context) (*Unlock, error) {
	return withHooks(ctx, uuo.sqlSave, uuo.mutation, uuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UnlockUpdateOne) SaveX(ctx context.Context) *Unlock {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UnlockUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UnlockUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uuo *UnlockUpdateOne) sqlSave(ctx context.Context) (_node *Unlock, err error) {
	_spec := sqlgraph.NewUpdateSpec(unlock.Table, unlock.Columns, sqlgraph.NewFieldSpec(unlock.FieldID, field.TypeInt))
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Unlock.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, unlock.FieldID)
		for _, f := range fields {
			if !unlock.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != unlock.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uuo.mutation.Name(); ok {
		_spec.SetField(unlock.FieldName, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Description(); ok {
		_spec.SetField(unlock.FieldDescription, field.TypeString, value)
	}
	if value, ok := uuo.mutation.OtherRequirement(); ok {
		_spec.SetField(unlock.FieldOtherRequirement, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Level(); ok {
		_spec.SetField(unlock.FieldLevel, field.TypeInt, value)
	}
	if value, ok := uuo.mutation.AddedLevel(); ok {
		_spec.AddField(unlock.FieldLevel, field.TypeInt, value)
	}
	if uuo.mutation.UnlockIDSkillFkCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   unlock.UnlockIDSkillFkTable,
			Columns: []string{unlock.UnlockIDSkillFkColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(skill.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.UnlockIDSkillFkIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   unlock.UnlockIDSkillFkTable,
			Columns: []string{unlock.UnlockIDSkillFkColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(skill.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Unlock{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{unlock.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	uuo.mutation.done = true
	return _node, nil
}
