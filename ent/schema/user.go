package schema

import (
  "entgo.io/ent"
  "entgo.io/ent/dialect/entsql"
  "entgo.io/ent/schema"
  "entgo.io/ent/schema/edge"
  "entgo.io/ent/schema/field"
  "github.com/go-bineanshi/usercenter-rpc/ent/schema/mixins"
)

// User holds the schema definition for the User entity.
type User struct {
  ent.Schema
}

func (User) Mixin() []ent.Mixin {
  return []ent.Mixin{
    mixins.IDMixin{},
    mixins.TimeMixin{},
    mixins.StatusMixin{},
  }
}

// Fields of the User.
func (User) Fields() []ent.Field {
  return []ent.Field{
    field.String("nickname").Unique().
      Comment("用户名称").
      Annotations(entsql.WithComments(true)),
    field.String("password").
      Comment("密码").
      Sensitive().
      Annotations(entsql.WithComments(true)),
    field.String("avater").
      Comment("用户头像").
      Optional().
      Annotations(entsql.WithComments(true)),
  }
}

// Edges of the User.
func (User) Edges() []ent.Edge {
  return []ent.Edge{
    edge.To("auths", UserAuth.Type),
  }
}

func (User) Annotations() []schema.Annotation {
  return []schema.Annotation{
    entsql.Annotation{Table: "sys_users"},
  }
}
