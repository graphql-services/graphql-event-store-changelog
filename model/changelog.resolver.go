package model

import (
	"encoding/json"
	"strings"

	graphql "github.com/graph-gophers/graphql-go"
)

// ChangeLogChange ...
type ChangeLogChange struct {
	IColumn   string
	IOldValue *string
	INewValue *string
}

// Name ...
func (n ChangeLogChange) Column() string {
	return n.IColumn
}

// OldValue ...
func (n ChangeLogChange) OldValue() *string {
	return n.IOldValue
}

// NewValue ...
func (n ChangeLogChange) NewValue() *string {
	return n.INewValue
}

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

// Changes ...
func (n ChangeLog) Changes() []ChangeLogChange {
	var changes []ChangeLogChange
	if err := json.Unmarshal([]byte(n.IChanges), &changes); err != nil {
		panic(err)
	}
	return changes
}

// Columns ...
func (n ChangeLog) Columns() []string {
	return strings.Split(strings.Trim(n.IColumns, "#"), "#,#")
}

// PrincipalID ...
func (n ChangeLog) PrincipalID() *string {
	return n.IPrincipalID
}

// Date ...
func (n ChangeLog) Date() graphql.Time {
	return graphql.Time{Time: n.IDate}
}
