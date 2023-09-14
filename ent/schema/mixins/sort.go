package mixins

import (
  "entgo.io/ent"
  "entgo.io/ent/dialect/entsql"
  "entgo.io/ent/schema/field"
  "entgo.io/ent/schema/mixin"
)

// SortMixin is the mixin with sort field which
// is used to sort the data.
type SortMixin struct {
  mixin.Schema
}

func (SortMixin) Fields() []ent.Field {
  return []ent.Field{
    field.Uint32("sort").
      Default(1).
      Comment("Sort Number | 排序编号").
      Annotations(entsql.WithComments(true)),
  }
}
