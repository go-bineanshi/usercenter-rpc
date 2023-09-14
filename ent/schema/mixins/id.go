package mixins

import (
  "entgo.io/ent"
  "entgo.io/ent/schema/field"
  "entgo.io/ent/schema/mixin"
)

// IDMixin is the mixin with uint64 type ID field
// and the created_at, updated_at fields.
type IDMixin struct {
  mixin.Schema
}

func (IDMixin) Fields() []ent.Field {
  return []ent.Field{
    field.Uint64("id"),
  }
}
