package mixins

import (
  "entgo.io/ent"
  "entgo.io/ent/dialect/entsql"
  "entgo.io/ent/schema/field"
  "entgo.io/ent/schema/mixin"
)

// StatusMixin is the mixin with status field which is uint8 type. It
// is used for multiple status more than 2. You can set more numbers as status.
type StatusMixin struct {
  mixin.Schema
}

func (StatusMixin) Fields() []ent.Field {
  return []ent.Field{
    field.Uint8("status").
      Default(1).
      Optional().
      Comment("Status 1: normal 2: ban | 状态 1 正常 2 禁用").
      Annotations(entsql.WithComments(true)),
  }
}
