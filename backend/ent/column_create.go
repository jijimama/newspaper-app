// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/jijimama/newspaper-app/ent/article"
	"github.com/jijimama/newspaper-app/ent/column"
	"github.com/jijimama/newspaper-app/ent/newspaper"
)

// ColumnCreate is the builder for creating a Column entity.
type ColumnCreate struct {
	config
	mutation *ColumnMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (cc *ColumnCreate) SetName(s string) *ColumnCreate {
	cc.mutation.SetName(s)
	return cc
}

// SetCreatedAt sets the "created_at" field.
func (cc *ColumnCreate) SetCreatedAt(t time.Time) *ColumnCreate {
	cc.mutation.SetCreatedAt(t)
	return cc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cc *ColumnCreate) SetNillableCreatedAt(t *time.Time) *ColumnCreate {
	if t != nil {
		cc.SetCreatedAt(*t)
	}
	return cc
}

// SetUpdatedAt sets the "updated_at" field.
func (cc *ColumnCreate) SetUpdatedAt(t time.Time) *ColumnCreate {
	cc.mutation.SetUpdatedAt(t)
	return cc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (cc *ColumnCreate) SetNillableUpdatedAt(t *time.Time) *ColumnCreate {
	if t != nil {
		cc.SetUpdatedAt(*t)
	}
	return cc
}

// SetNewspaperID sets the "newspaper" edge to the Newspaper entity by ID.
func (cc *ColumnCreate) SetNewspaperID(id int) *ColumnCreate {
	cc.mutation.SetNewspaperID(id)
	return cc
}

// SetNillableNewspaperID sets the "newspaper" edge to the Newspaper entity by ID if the given value is not nil.
func (cc *ColumnCreate) SetNillableNewspaperID(id *int) *ColumnCreate {
	if id != nil {
		cc = cc.SetNewspaperID(*id)
	}
	return cc
}

// SetNewspaper sets the "newspaper" edge to the Newspaper entity.
func (cc *ColumnCreate) SetNewspaper(n *Newspaper) *ColumnCreate {
	return cc.SetNewspaperID(n.ID)
}

// AddArticleIDs adds the "articles" edge to the Article entity by IDs.
func (cc *ColumnCreate) AddArticleIDs(ids ...int) *ColumnCreate {
	cc.mutation.AddArticleIDs(ids...)
	return cc
}

// AddArticles adds the "articles" edges to the Article entity.
func (cc *ColumnCreate) AddArticles(a ...*Article) *ColumnCreate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return cc.AddArticleIDs(ids...)
}

// Mutation returns the ColumnMutation object of the builder.
func (cc *ColumnCreate) Mutation() *ColumnMutation {
	return cc.mutation
}

// Save creates the Column in the database.
func (cc *ColumnCreate) Save(ctx context.Context) (*Column, error) {
	cc.defaults()
	return withHooks(ctx, cc.sqlSave, cc.mutation, cc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (cc *ColumnCreate) SaveX(ctx context.Context) *Column {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *ColumnCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *ColumnCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cc *ColumnCreate) defaults() {
	if _, ok := cc.mutation.CreatedAt(); !ok {
		v := column.DefaultCreatedAt()
		cc.mutation.SetCreatedAt(v)
	}
	if _, ok := cc.mutation.UpdatedAt(); !ok {
		v := column.DefaultUpdatedAt()
		cc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *ColumnCreate) check() error {
	if _, ok := cc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Column.name"`)}
	}
	if v, ok := cc.mutation.Name(); ok {
		if err := column.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Column.name": %w`, err)}
		}
	}
	if _, ok := cc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Column.created_at"`)}
	}
	if _, ok := cc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Column.updated_at"`)}
	}
	return nil
}

func (cc *ColumnCreate) sqlSave(ctx context.Context) (*Column, error) {
	if err := cc.check(); err != nil {
		return nil, err
	}
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	cc.mutation.id = &_node.ID
	cc.mutation.done = true
	return _node, nil
}

func (cc *ColumnCreate) createSpec() (*Column, *sqlgraph.CreateSpec) {
	var (
		_node = &Column{config: cc.config}
		_spec = sqlgraph.NewCreateSpec(column.Table, sqlgraph.NewFieldSpec(column.FieldID, field.TypeInt))
	)
	if value, ok := cc.mutation.Name(); ok {
		_spec.SetField(column.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := cc.mutation.CreatedAt(); ok {
		_spec.SetField(column.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := cc.mutation.UpdatedAt(); ok {
		_spec.SetField(column.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := cc.mutation.NewspaperIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   column.NewspaperTable,
			Columns: []string{column.NewspaperColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(newspaper.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.newspaper_columns = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.ArticlesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   column.ArticlesTable,
			Columns: []string{column.ArticlesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(article.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ColumnCreateBulk is the builder for creating many Column entities in bulk.
type ColumnCreateBulk struct {
	config
	err      error
	builders []*ColumnCreate
}

// Save creates the Column entities in the database.
func (ccb *ColumnCreateBulk) Save(ctx context.Context) ([]*Column, error) {
	if ccb.err != nil {
		return nil, ccb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Column, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ColumnMutation)
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
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *ColumnCreateBulk) SaveX(ctx context.Context) []*Column {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *ColumnCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *ColumnCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}
