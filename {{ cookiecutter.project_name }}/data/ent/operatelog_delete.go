// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"{{ cookiecutter.project_name }}/data/ent/operatelog"
	"{{ cookiecutter.project_name }}/data/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// OperateLogDelete is the builder for deleting a OperateLog entity.
type OperateLogDelete struct {
	config
	hooks    []Hook
	mutation *OperateLogMutation
}

// Where appends a list predicates to the OperateLogDelete builder.
func (old *OperateLogDelete) Where(ps ...predicate.OperateLog) *OperateLogDelete {
	old.mutation.Where(ps...)
	return old
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (old *OperateLogDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, OperateLogMutation](ctx, old.sqlExec, old.mutation, old.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (old *OperateLogDelete) ExecX(ctx context.Context) int {
	n, err := old.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (old *OperateLogDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(operatelog.Table, sqlgraph.NewFieldSpec(operatelog.FieldID, field.TypeInt))
	if ps := old.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, old.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	old.mutation.done = true
	return affected, err
}

// OperateLogDeleteOne is the builder for deleting a single OperateLog entity.
type OperateLogDeleteOne struct {
	old *OperateLogDelete
}

// Where appends a list predicates to the OperateLogDelete builder.
func (oldo *OperateLogDeleteOne) Where(ps ...predicate.OperateLog) *OperateLogDeleteOne {
	oldo.old.mutation.Where(ps...)
	return oldo
}

// Exec executes the deletion query.
func (oldo *OperateLogDeleteOne) Exec(ctx context.Context) error {
	n, err := oldo.old.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{operatelog.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (oldo *OperateLogDeleteOne) ExecX(ctx context.Context) {
	if err := oldo.Exec(ctx); err != nil {
		panic(err)
	}
}
