package entfw

import (
	"fmt"

	"entgo.io/ent"
)

// EntFields returns a slice of ent.Field which is converted to prevent `json:omitempty` from being set.
//
// Example:
//
//	func (Shop) Fields() []ent.Field {
//		return conv.EntFields(
//			field.Uint32("id"),
//			field.String("name"),
//			// ...,
//		)
//	}
func Fields(fields ...ent.Field) []ent.Field {
	for i := range fields {
		descriptor := fields[i].Descriptor()
		if !descriptor.Sensitive && descriptor.Tag == "" {
			descriptor.Tag = fmt.Sprintf(`json:"%s"`, descriptor.Name)
		}
	}

	return fields
}
