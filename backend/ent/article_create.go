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
	"github.com/jijimama/newspaper-app/ent/newspaper"
)

// ArticleCreate is the builder for creating a Article entity.
type ArticleCreate struct {
	config
	mutation *ArticleMutation
	hooks    []Hook
}

// SetContent sets the "content" field.
func (ac *ArticleCreate) SetContent(s string) *ArticleCreate {
	ac.mutation.SetContent(s)
	return ac
}

// SetYear sets the "year" field.
func (ac *ArticleCreate) SetYear(i int) *ArticleCreate {
	ac.mutation.SetYear(i)
	return ac
}

// SetMonth sets the "month" field.
func (ac *ArticleCreate) SetMonth(i int) *ArticleCreate {
	ac.mutation.SetMonth(i)
	return ac
}

// SetDay sets the "day" field.
func (ac *ArticleCreate) SetDay(i int) *ArticleCreate {
	ac.mutation.SetDay(i)
	return ac
}

// SetCreatedAt sets the "created_at" field.
func (ac *ArticleCreate) SetCreatedAt(t time.Time) *ArticleCreate {
	ac.mutation.SetCreatedAt(t)
	return ac
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ac *ArticleCreate) SetNillableCreatedAt(t *time.Time) *ArticleCreate {
	if t != nil {
		ac.SetCreatedAt(*t)
	}
	return ac
}

// SetUpdatedAt sets the "updated_at" field.
func (ac *ArticleCreate) SetUpdatedAt(t time.Time) *ArticleCreate {
	ac.mutation.SetUpdatedAt(t)
	return ac
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ac *ArticleCreate) SetNillableUpdatedAt(t *time.Time) *ArticleCreate {
	if t != nil {
		ac.SetUpdatedAt(*t)
	}
	return ac
}

// SetNewspaperID sets the "newspaper" edge to the Newspaper entity by ID.
func (ac *ArticleCreate) SetNewspaperID(id int) *ArticleCreate {
	ac.mutation.SetNewspaperID(id)
	return ac
}

// SetNillableNewspaperID sets the "newspaper" edge to the Newspaper entity by ID if the given value is not nil.
func (ac *ArticleCreate) SetNillableNewspaperID(id *int) *ArticleCreate {
	if id != nil {
		ac = ac.SetNewspaperID(*id)
	}
	return ac
}

// SetNewspaper sets the "newspaper" edge to the Newspaper entity.
func (ac *ArticleCreate) SetNewspaper(n *Newspaper) *ArticleCreate {
	return ac.SetNewspaperID(n.ID)
}

// Mutation returns the ArticleMutation object of the builder.
func (ac *ArticleCreate) Mutation() *ArticleMutation {
	return ac.mutation
}

// Save creates the Article in the database.
func (ac *ArticleCreate) Save(ctx context.Context) (*Article, error) {
	ac.defaults()
	return withHooks(ctx, ac.sqlSave, ac.mutation, ac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ac *ArticleCreate) SaveX(ctx context.Context) *Article {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ac *ArticleCreate) Exec(ctx context.Context) error {
	_, err := ac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ac *ArticleCreate) ExecX(ctx context.Context) {
	if err := ac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ac *ArticleCreate) defaults() {
	if _, ok := ac.mutation.CreatedAt(); !ok {
		v := article.DefaultCreatedAt()
		ac.mutation.SetCreatedAt(v)
	}
	if _, ok := ac.mutation.UpdatedAt(); !ok {
		v := article.DefaultUpdatedAt()
		ac.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ac *ArticleCreate) check() error {
	if _, ok := ac.mutation.Content(); !ok {
		return &ValidationError{Name: "content", err: errors.New(`ent: missing required field "Article.content"`)}
	}
	if v, ok := ac.mutation.Content(); ok {
		if err := article.ContentValidator(v); err != nil {
			return &ValidationError{Name: "content", err: fmt.Errorf(`ent: validator failed for field "Article.content": %w`, err)}
		}
	}
	if _, ok := ac.mutation.Year(); !ok {
		return &ValidationError{Name: "year", err: errors.New(`ent: missing required field "Article.year"`)}
	}
	if v, ok := ac.mutation.Year(); ok {
		if err := article.YearValidator(v); err != nil {
			return &ValidationError{Name: "year", err: fmt.Errorf(`ent: validator failed for field "Article.year": %w`, err)}
		}
	}
	if _, ok := ac.mutation.Month(); !ok {
		return &ValidationError{Name: "month", err: errors.New(`ent: missing required field "Article.month"`)}
	}
	if v, ok := ac.mutation.Month(); ok {
		if err := article.MonthValidator(v); err != nil {
			return &ValidationError{Name: "month", err: fmt.Errorf(`ent: validator failed for field "Article.month": %w`, err)}
		}
	}
	if _, ok := ac.mutation.Day(); !ok {
		return &ValidationError{Name: "day", err: errors.New(`ent: missing required field "Article.day"`)}
	}
	if v, ok := ac.mutation.Day(); ok {
		if err := article.DayValidator(v); err != nil {
			return &ValidationError{Name: "day", err: fmt.Errorf(`ent: validator failed for field "Article.day": %w`, err)}
		}
	}
	if _, ok := ac.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Article.created_at"`)}
	}
	if _, ok := ac.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Article.updated_at"`)}
	}
	return nil
}

func (ac *ArticleCreate) sqlSave(ctx context.Context) (*Article, error) {
	if err := ac.check(); err != nil {
		return nil, err
	}
	_node, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	ac.mutation.id = &_node.ID
	ac.mutation.done = true
	return _node, nil
}

func (ac *ArticleCreate) createSpec() (*Article, *sqlgraph.CreateSpec) {
	var (
		_node = &Article{config: ac.config}
		_spec = sqlgraph.NewCreateSpec(article.Table, sqlgraph.NewFieldSpec(article.FieldID, field.TypeInt))
	)
	if value, ok := ac.mutation.Content(); ok {
		_spec.SetField(article.FieldContent, field.TypeString, value)
		_node.Content = value
	}
	if value, ok := ac.mutation.Year(); ok {
		_spec.SetField(article.FieldYear, field.TypeInt, value)
		_node.Year = value
	}
	if value, ok := ac.mutation.Month(); ok {
		_spec.SetField(article.FieldMonth, field.TypeInt, value)
		_node.Month = value
	}
	if value, ok := ac.mutation.Day(); ok {
		_spec.SetField(article.FieldDay, field.TypeInt, value)
		_node.Day = value
	}
	if value, ok := ac.mutation.CreatedAt(); ok {
		_spec.SetField(article.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := ac.mutation.UpdatedAt(); ok {
		_spec.SetField(article.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := ac.mutation.NewspaperIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   article.NewspaperTable,
			Columns: []string{article.NewspaperColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(newspaper.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.newspaper_articles = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ArticleCreateBulk is the builder for creating many Article entities in bulk.
type ArticleCreateBulk struct {
	config
	err      error
	builders []*ArticleCreate
}

// Save creates the Article entities in the database.
func (acb *ArticleCreateBulk) Save(ctx context.Context) ([]*Article, error) {
	if acb.err != nil {
		return nil, acb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*Article, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ArticleMutation)
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
					_, err = mutators[i+1].Mutate(root, acb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, acb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, acb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (acb *ArticleCreateBulk) SaveX(ctx context.Context) []*Article {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acb *ArticleCreateBulk) Exec(ctx context.Context) error {
	_, err := acb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acb *ArticleCreateBulk) ExecX(ctx context.Context) {
	if err := acb.Exec(ctx); err != nil {
		panic(err)
	}
}
