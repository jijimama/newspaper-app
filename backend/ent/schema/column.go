package schema

import (
    "entgo.io/ent"
    "entgo.io/ent/schema/edge"
    "entgo.io/ent/schema/field"
	"time"
)

// Column holds the schema definition for the Column entity.
type Column struct {
	ent.Schema
}

// Fields of the Column.
func (Column) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty(),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

// Edges of the Column.
func (Column) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("newspaper", Newspaper.Type).
			Ref("columns").
			Unique(),
		edge.To("articles", Article.Type),
	}
}
