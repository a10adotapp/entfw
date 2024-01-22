# entfw

`entfw` is a library to remove `json:omitempty` from entity structs generated with `go generate ./ent`.

### Quick Start

Download the module:

```console
go get -u github.com/a10adotapp/entfw
```

Wrap fields in the schema with `entfw.Fields`

``` go
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
```

Run the code generation:

``` console
go generate ./ent
```

### What happens

See generated ent/user.go 

- `ID` has `json:"omitempty"` because the `id` is not defined in the schema file.
- `Name` does not have `json:"omitempty"` because the `name` is defined in the schema file wrapped by `entfw.Fields`.
- `Password` has `json:"-"` because it is defined as `Sensitive` in the schema.
- `EmailAddress` has `json:"email"`. You can overwrite the struct tag using `StructTag` in the schema.

``` go
// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name"`
	// Password holds the value of the "password" field.
	Password string `json:"-"`
	// EmailAddress holds the value of the "email_address" field.
	EmailAddress uint32 `json:"email"`
	selectValues sql.SelectValues
}
```

What if **without** entfw

- `ID` same.
- `Name` has `json:"omitempty"`.
- `Password` same.
- `EmailAddress` same.

``` go
// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Password holds the value of the "password" field.
	Password string `json:"-"`
	// EmailAddress holds the value of the "email_address" field.
	EmailAddress uint32 `json:"email"`
	selectValues sql.SelectValues
}
```