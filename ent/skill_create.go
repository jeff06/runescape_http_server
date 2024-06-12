// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"runescape_http_server/ent/skill"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SkillCreate is the builder for creating a Skill entity.
type SkillCreate struct {
	config
	mutation *SkillMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (sc *SkillCreate) SetName(s string) *SkillCreate {
	sc.mutation.SetName(s)
	return sc
}

// SetDescription sets the "description" field.
func (sc *SkillCreate) SetDescription(s string) *SkillCreate {
	sc.mutation.SetDescription(s)
	return sc
}

// SetIsMember sets the "is_member" field.
func (sc *SkillCreate) SetIsMember(i int) *SkillCreate {
	sc.mutation.SetIsMember(i)
	return sc
}

// Mutation returns the SkillMutation object of the builder.
func (sc *SkillCreate) Mutation() *SkillMutation {
	return sc.mutation
}

// Save creates the Skill in the database.
func (sc *SkillCreate) Save(ctx context.Context) (*Skill, error) {
	return withHooks(ctx, sc.sqlSave, sc.mutation, sc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (sc *SkillCreate) SaveX(ctx context.Context) *Skill {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *SkillCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *SkillCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *SkillCreate) check() error {
	if _, ok := sc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Skill.name"`)}
	}
	if _, ok := sc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "Skill.description"`)}
	}
	if _, ok := sc.mutation.IsMember(); !ok {
		return &ValidationError{Name: "is_member", err: errors.New(`ent: missing required field "Skill.is_member"`)}
	}
	return nil
}

func (sc *SkillCreate) sqlSave(ctx context.Context) (*Skill, error) {
	if err := sc.check(); err != nil {
		return nil, err
	}
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	sc.mutation.id = &_node.ID
	sc.mutation.done = true
	return _node, nil
}

func (sc *SkillCreate) createSpec() (*Skill, *sqlgraph.CreateSpec) {
	var (
		_node = &Skill{config: sc.config}
		_spec = sqlgraph.NewCreateSpec(skill.Table, sqlgraph.NewFieldSpec(skill.FieldID, field.TypeInt))
	)
	if value, ok := sc.mutation.Name(); ok {
		_spec.SetField(skill.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := sc.mutation.Description(); ok {
		_spec.SetField(skill.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := sc.mutation.IsMember(); ok {
		_spec.SetField(skill.FieldIsMember, field.TypeInt, value)
		_node.IsMember = value
	}
	return _node, _spec
}

// SkillCreateBulk is the builder for creating many Skill entities in bulk.
type SkillCreateBulk struct {
	config
	err      error
	builders []*SkillCreate
}

// Save creates the Skill entities in the database.
func (scb *SkillCreateBulk) Save(ctx context.Context) ([]*Skill, error) {
	if scb.err != nil {
		return nil, scb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Skill, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SkillMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *SkillCreateBulk) SaveX(ctx context.Context) []*Skill {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *SkillCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *SkillCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}
