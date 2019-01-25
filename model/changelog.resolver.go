package model

import (
	"strings"

	graphql "github.com/graph-gophers/graphql-go"
)

// ID ...
func (n ChangeLog) ID() graphql.ID {
	return graphql.ID(n.IID.String())
}

// Entity ...
func (n ChangeLog) Entity() string {
	return n.IEntity
}

// EntityID ...
func (n ChangeLog) EntityID() string {
	return n.IEntityID
}

// Type ...
func (n ChangeLog) Type() string {
	return n.IType
}

// Columns ...
func (n ChangeLog) Columns() []string {
	return strings.Split(n.IColumns, ",")
}

// PrincipalID ...
func (n ChangeLog) PrincipalID() *string {
	return n.IPrincipalID
}

// Date ...
func (n ChangeLog) Date() graphql.Time {
	return graphql.Time{Time: n.IDate}
}
