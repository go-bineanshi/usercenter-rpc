package schema

import (
  "entgo.io/ent"
  "entgo.io/ent/dialect/entsql"
  "entgo.io/ent/schema"
  "entgo.io/ent/schema/edge"
  "entgo.io/ent/schema/field"
  "entgo.io/ent/schema/index"
  "github.com/go-bineanshi/usercenter-rpc/ent/property"
  "github.com/go-bineanshi/usercenter-rpc/ent/schema/mixins"
)

// UserAuth holds the schema definition for the UserAuth entity.
type UserAuth struct {
  ent.Schema
}

func (UserAuth) Mixin() []ent.Mixin {
  return []ent.Mixin{
    mixins.IDMixin{},
    mixins.TimeMixin{},
    mixins.StatusMixin{},
  }
}

// Fields of the UserAuth.
func (UserAuth) Fields() []ent.Field {
  return []ent.Field{
    field.Uint64("user_id").Optional(),
    field.Enum("provider").
      GoType(property.Provider("")).
      Comment("登录提供商").
      Annotations(entsql.WithComments(true)),
    field.String("account").
      Comment("账号").
      NotEmpty().
      Annotations(entsql.WithComments(true)),
  }
}

// Edges of the UserAuth.
func (UserAuth) Edges() []ent.Edge {
  return []ent.Edge{
    edge.From("user", User.Type).
      Ref("auths").
      Field("user_id").
      Unique(),
  }
}

func (UserAuth) Indexes() []ent.Index {
  return []ent.Index{
    // 唯一约束索引
    index.Fields("provider", "account").
      Unique(),
  }
}

func (UserAuth) Annotations() []schema.Annotation {
  return []schema.Annotation{
    entsql.Annotation{Table: "sys_user_auths"},
  }
}
