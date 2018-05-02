package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// AddressPlace holds the schema definition for the AddressPlace entity.
type AddressPlace struct {
	ent.Schema
}

// Fields of the AddressPlace.
func (AddressPlace) Fields() []ent.Field {
	return []ent.Field{
		field.Int32("id").
			Unique(),
		field.Int32("number"),
		field.Int32("orientation_number").Optional(),
		field.String("orientation_number_letter").Optional(),
		field.Int32("zip"),
	}
}

// Edges of the AddressPlace.
func (AddressPlace) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("streets", Street.Type),
	}
}
