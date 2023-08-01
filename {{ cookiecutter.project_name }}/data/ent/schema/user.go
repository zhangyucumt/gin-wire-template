package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("username").Unique(),
		field.String("name"),
		field.String("email").Default(""),
		field.String("phone").Default(""),
		field.String("password").Sensitive(),
		field.Time("last_login").Nillable().Optional(),
		field.Bool("is_superuser").Default(false),
		field.Bool("is_active").Default(true),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		// User has many Group.
		edge.From("groups", Group.Type).Ref("users"),
		edge.To("permissions", Permission.Type),
		edge.To("logs", OperateLog.Type),
	}
}

// Mixin of the schema.
func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
