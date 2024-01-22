package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/a10adotapp/entfw/v1"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return entfw.Fields(
		field.String("name"),

		field.String("password").
			Sensitive(),

		field.Uint32("email_address").
			StructTag(`json:"email"`),
	)
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
