// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"runescape_http_server/ent/otherrequirement"
	"runescape_http_server/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// OtherRequirementDelete is the builder for deleting a OtherRequirement entity.
type OtherRequirementDelete struct {
	config
	hooks    []Hook
	mutation *OtherRequirementMutation
}

// Where appends a list predicates to the OtherRequirementDelete builder.
func (ord *OtherRequirementDelete) Where(ps ...predicate.OtherRequirement) *OtherRequirementDelete {
	ord.mutation.Where(ps...)
	return ord
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ord *OtherRequirementDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, ord.sqlExec, ord.mutation, ord.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ord *OtherRequirementDelete) ExecX(ctx context.Context) int {
	n, err := ord.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ord *OtherRequirementDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(otherrequirement.Table, sqlgraph.NewFieldSpec(otherrequirement.FieldID, field.TypeInt))
	if ps := ord.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ord.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	ord.mutation.done = true
	return affected, err
}

// OtherRequirementDeleteOne is the builder for deleting a single OtherRequirement entity.
type OtherRequirementDeleteOne struct {
	ord *OtherRequirementDelete
}

// Where appends a list predicates to the OtherRequirementDelete builder.
func (ordo *OtherRequirementDeleteOne) Where(ps ...predicate.OtherRequirement) *OtherRequirementDeleteOne {
	ordo.ord.mutation.Where(ps...)
	return ordo
}

// Exec executes the deletion query.
func (ordo *OtherRequirementDeleteOne) Exec(ctx context.Context) error {
	n, err := ordo.ord.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{otherrequirement.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ordo *OtherRequirementDeleteOne) ExecX(ctx context.Context) {
	if err := ordo.Exec(ctx); err != nil {
		panic(err)
	}
}