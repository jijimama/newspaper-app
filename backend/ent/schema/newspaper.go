package schema

import "entgo.io/ent"

// Newspaper holds the schema definition for the Newspaper entity.
type Newspaper struct {
	ent.Schema
}

// Fields of the Newspaper.
func (Newspaper) Fields() []ent.Field {
	return nil
}

// Edges of the Newspaper.
func (Newspaper) Edges() []ent.Edge {
	return nil
}
