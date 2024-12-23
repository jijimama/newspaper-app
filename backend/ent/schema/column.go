package schema

import "entgo.io/ent"

// Column holds the schema definition for the Column entity.
type Column struct {
	ent.Schema
}

// Fields of the Column.
func (Column) Fields() []ent.Field {
	return nil
}

// Edges of the Column.
func (Column) Edges() []ent.Edge {
	return nil
}
