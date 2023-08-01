package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// OperateLog holds the schema definition for the OperateLog entity.
type OperateLog struct {
	ent.Schema
}

// Fields of the OperateLog.
func (OperateLog) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("path"),
		field.String("method"),
		field.String("ip"),
		field.Int("status_code"),
	}
}

// Edges of the OperateLog.
func (OperateLog) Edges() []ent.Edge {
	return []ent.Edge{edge.From("user", User.Type).Ref("logs").Unique()}
}

// Mixin of the schema.
func (OperateLog) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
