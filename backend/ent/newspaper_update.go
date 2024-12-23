// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/jijimama/newspaper-app/ent/column"
	"github.com/jijimama/newspaper-app/ent/newspaper"
	"github.com/jijimama/newspaper-app/ent/predicate"
)

// NewspaperUpdate is the builder for updating Newspaper entities.
type NewspaperUpdate struct {
	config
	hooks    []Hook
	mutation *NewspaperMutation
}

// Where appends a list predicates to the NewspaperUpdate builder.
func (nu *NewspaperUpdate) Where(ps ...predicate.Newspaper) *NewspaperUpdate {
	nu.mutation.Where(ps...)
	return nu
}

// SetName sets the "name" field.
func (nu *NewspaperUpdate) SetName(s string) *NewspaperUpdate {
	nu.mutation.SetName(s)
	return nu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (nu *NewspaperUpdate) SetNillableName(s *string) *NewspaperUpdate {
	if s != nil {
		nu.SetName(*s)
	}
	return nu
}

// SetCreatedAt sets the "created_at" field.
func (nu *NewspaperUpdate) SetCreatedAt(t time.Time) *NewspaperUpdate {
	nu.mutation.SetCreatedAt(t)
	return nu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (nu *NewspaperUpdate) SetNillableCreatedAt(t *time.Time) *NewspaperUpdate {
	if t != nil {
		nu.SetCreatedAt(*t)
	}
	return nu
}

// SetUpdatedAt sets the "updated_at" field.
func (nu *NewspaperUpdate) SetUpdatedAt(t time.Time) *NewspaperUpdate {
	nu.mutation.SetUpdatedAt(t)
	return nu
}

// AddColumnIDs adds the "columns" edge to the Column entity by IDs.
func (nu *NewspaperUpdate) AddColumnIDs(ids ...int) *NewspaperUpdate {
	nu.mutation.AddColumnIDs(ids...)
	return nu
}

// AddColumns adds the "columns" edges to the Column entity.
func (nu *NewspaperUpdate) AddColumns(c ...*Column) *NewspaperUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return nu.AddColumnIDs(ids...)
}

// Mutation returns the NewspaperMutation object of the builder.
func (nu *NewspaperUpdate) Mutation() *NewspaperMutation {
	return nu.mutation
}

// ClearColumns clears all "columns" edges to the Column entity.
func (nu *NewspaperUpdate) ClearColumns() *NewspaperUpdate {
	nu.mutation.ClearColumns()
	return nu
}

// RemoveColumnIDs removes the "columns" edge to Column entities by IDs.
func (nu *NewspaperUpdate) RemoveColumnIDs(ids ...int) *NewspaperUpdate {
	nu.mutation.RemoveColumnIDs(ids...)
	return nu
}

// RemoveColumns removes "columns" edges to Column entities.
func (nu *NewspaperUpdate) RemoveColumns(c ...*Column) *NewspaperUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return nu.RemoveColumnIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (nu *NewspaperUpdate) Save(ctx context.Context) (int, error) {
	nu.defaults()
	return withHooks(ctx, nu.sqlSave, nu.mutation, nu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (nu *NewspaperUpdate) SaveX(ctx context.Context) int {
	affected, err := nu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (nu *NewspaperUpdate) Exec(ctx context.Context) error {
	_, err := nu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nu *NewspaperUpdate) ExecX(ctx context.Context) {
	if err := nu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (nu *NewspaperUpdate) defaults() {
	if _, ok := nu.mutation.UpdatedAt(); !ok {
		v := newspaper.UpdateDefaultUpdatedAt()
		nu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (nu *NewspaperUpdate) check() error {
	if v, ok := nu.mutation.Name(); ok {
		if err := newspaper.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Newspaper.name": %w`, err)}
		}
	}
	return nil
}

func (nu *NewspaperUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := nu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(newspaper.Table, newspaper.Columns, sqlgraph.NewFieldSpec(newspaper.FieldID, field.TypeInt))
	if ps := nu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := nu.mutation.Name(); ok {
		_spec.SetField(newspaper.FieldName, field.TypeString, value)
	}
	if value, ok := nu.mutation.CreatedAt(); ok {
		_spec.SetField(newspaper.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := nu.mutation.UpdatedAt(); ok {
		_spec.SetField(newspaper.FieldUpdatedAt, field.TypeTime, value)
	}
	if nu.mutation.ColumnsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   newspaper.ColumnsTable,
			Columns: []string{newspaper.ColumnsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(column.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nu.mutation.RemovedColumnsIDs(); len(nodes) > 0 && !nu.mutation.ColumnsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   newspaper.ColumnsTable,
			Columns: []string{newspaper.ColumnsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(column.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nu.mutation.ColumnsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   newspaper.ColumnsTable,
			Columns: []string{newspaper.ColumnsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(column.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, nu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{newspaper.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	nu.mutation.done = true
	return n, nil
}

// NewspaperUpdateOne is the builder for updating a single Newspaper entity.
type NewspaperUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *NewspaperMutation
}

// SetName sets the "name" field.
func (nuo *NewspaperUpdateOne) SetName(s string) *NewspaperUpdateOne {
	nuo.mutation.SetName(s)
	return nuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (nuo *NewspaperUpdateOne) SetNillableName(s *string) *NewspaperUpdateOne {
	if s != nil {
		nuo.SetName(*s)
	}
	return nuo
}

// SetCreatedAt sets the "created_at" field.
func (nuo *NewspaperUpdateOne) SetCreatedAt(t time.Time) *NewspaperUpdateOne {
	nuo.mutation.SetCreatedAt(t)
	return nuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (nuo *NewspaperUpdateOne) SetNillableCreatedAt(t *time.Time) *NewspaperUpdateOne {
	if t != nil {
		nuo.SetCreatedAt(*t)
	}
	return nuo
}

// SetUpdatedAt sets the "updated_at" field.
func (nuo *NewspaperUpdateOne) SetUpdatedAt(t time.Time) *NewspaperUpdateOne {
	nuo.mutation.SetUpdatedAt(t)
	return nuo
}

// AddColumnIDs adds the "columns" edge to the Column entity by IDs.
func (nuo *NewspaperUpdateOne) AddColumnIDs(ids ...int) *NewspaperUpdateOne {
	nuo.mutation.AddColumnIDs(ids...)
	return nuo
}

// AddColumns adds the "columns" edges to the Column entity.
func (nuo *NewspaperUpdateOne) AddColumns(c ...*Column) *NewspaperUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return nuo.AddColumnIDs(ids...)
}

// Mutation returns the NewspaperMutation object of the builder.
func (nuo *NewspaperUpdateOne) Mutation() *NewspaperMutation {
	return nuo.mutation
}

// ClearColumns clears all "columns" edges to the Column entity.
func (nuo *NewspaperUpdateOne) ClearColumns() *NewspaperUpdateOne {
	nuo.mutation.ClearColumns()
	return nuo
}

// RemoveColumnIDs removes the "columns" edge to Column entities by IDs.
func (nuo *NewspaperUpdateOne) RemoveColumnIDs(ids ...int) *NewspaperUpdateOne {
	nuo.mutation.RemoveColumnIDs(ids...)
	return nuo
}

// RemoveColumns removes "columns" edges to Column entities.
func (nuo *NewspaperUpdateOne) RemoveColumns(c ...*Column) *NewspaperUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return nuo.RemoveColumnIDs(ids...)
}

// Where appends a list predicates to the NewspaperUpdate builder.
func (nuo *NewspaperUpdateOne) Where(ps ...predicate.Newspaper) *NewspaperUpdateOne {
	nuo.mutation.Where(ps...)
	return nuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (nuo *NewspaperUpdateOne) Select(field string, fields ...string) *NewspaperUpdateOne {
	nuo.fields = append([]string{field}, fields...)
	return nuo
}

// Save executes the query and returns the updated Newspaper entity.
func (nuo *NewspaperUpdateOne) Save(ctx context.Context) (*Newspaper, error) {
	nuo.defaults()
	return withHooks(ctx, nuo.sqlSave, nuo.mutation, nuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (nuo *NewspaperUpdateOne) SaveX(ctx context.Context) *Newspaper {
	node, err := nuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (nuo *NewspaperUpdateOne) Exec(ctx context.Context) error {
	_, err := nuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nuo *NewspaperUpdateOne) ExecX(ctx context.Context) {
	if err := nuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (nuo *NewspaperUpdateOne) defaults() {
	if _, ok := nuo.mutation.UpdatedAt(); !ok {
		v := newspaper.UpdateDefaultUpdatedAt()
		nuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (nuo *NewspaperUpdateOne) check() error {
	if v, ok := nuo.mutation.Name(); ok {
		if err := newspaper.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Newspaper.name": %w`, err)}
		}
	}
	return nil
}

func (nuo *NewspaperUpdateOne) sqlSave(ctx context.Context) (_node *Newspaper, err error) {
	if err := nuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(newspaper.Table, newspaper.Columns, sqlgraph.NewFieldSpec(newspaper.FieldID, field.TypeInt))
	id, ok := nuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Newspaper.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := nuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, newspaper.FieldID)
		for _, f := range fields {
			if !newspaper.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != newspaper.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := nuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := nuo.mutation.Name(); ok {
		_spec.SetField(newspaper.FieldName, field.TypeString, value)
	}
	if value, ok := nuo.mutation.CreatedAt(); ok {
		_spec.SetField(newspaper.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := nuo.mutation.UpdatedAt(); ok {
		_spec.SetField(newspaper.FieldUpdatedAt, field.TypeTime, value)
	}
	if nuo.mutation.ColumnsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   newspaper.ColumnsTable,
			Columns: []string{newspaper.ColumnsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(column.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nuo.mutation.RemovedColumnsIDs(); len(nodes) > 0 && !nuo.mutation.ColumnsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   newspaper.ColumnsTable,
			Columns: []string{newspaper.ColumnsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(column.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nuo.mutation.ColumnsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   newspaper.ColumnsTable,
			Columns: []string{newspaper.ColumnsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(column.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Newspaper{config: nuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, nuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{newspaper.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	nuo.mutation.done = true
	return _node, nil
}
