// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/jijimama/newspaper-app/ent/newspaper"
	"github.com/jijimama/newspaper-app/ent/predicate"
)

// NewspaperDelete is the builder for deleting a Newspaper entity.
type NewspaperDelete struct {
	config
	hooks    []Hook
	mutation *NewspaperMutation
}

// Where appends a list predicates to the NewspaperDelete builder.
func (nd *NewspaperDelete) Where(ps ...predicate.Newspaper) *NewspaperDelete {
	nd.mutation.Where(ps...)
	return nd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (nd *NewspaperDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, nd.sqlExec, nd.mutation, nd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (nd *NewspaperDelete) ExecX(ctx context.Context) int {
	n, err := nd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (nd *NewspaperDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(newspaper.Table, sqlgraph.NewFieldSpec(newspaper.FieldID, field.TypeInt))
	if ps := nd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, nd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	nd.mutation.done = true
	return affected, err
}

// NewspaperDeleteOne is the builder for deleting a single Newspaper entity.
type NewspaperDeleteOne struct {
	nd *NewspaperDelete
}

// Where appends a list predicates to the NewspaperDelete builder.
func (ndo *NewspaperDeleteOne) Where(ps ...predicate.Newspaper) *NewspaperDeleteOne {
	ndo.nd.mutation.Where(ps...)
	return ndo
}

// Exec executes the deletion query.
func (ndo *NewspaperDeleteOne) Exec(ctx context.Context) error {
	n, err := ndo.nd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{newspaper.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ndo *NewspaperDeleteOne) ExecX(ctx context.Context) {
	if err := ndo.Exec(ctx); err != nil {
		panic(err)
	}
}
