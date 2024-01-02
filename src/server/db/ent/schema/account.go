package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Account holds the schema definition for the Account entity.
type Account struct {
	ent.Schema
}

// Fields of the Account.
func (Account) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}),
		field.String("username").MaxLen(64).Comment("用户名"),
		field.String("password").MaxLen(64).Comment("密码"),
		field.String("salt").MaxLen(64).Comment("加盐"),
	}
}

// Edges of the Account.
func (Account) Edges() []ent.Edge {
	return nil
}
