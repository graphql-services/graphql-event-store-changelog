package model

import (
	"strings"
	"time"

	graphql "github.com/graph-gophers/graphql-go"
	uuid "github.com/satori/go.uuid"
)

// ChangeLog ...
type ChangeLog struct {
	// gorm.Model
	IID          uuid.UUID `gorm:"primary_key"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time `sql:"index"`
	IEntity      string     `gorm:"column:entity"`
	IEntityID    string     `gorm:"column:entityId"`
	IType        string     `gorm:"column:type"`
	IColumns     string     `gorm:"column:columns"`
	IPrincipalID *string    `gorm:"column:principalId"`
	IDate        time.Time  `gorm:"column:date"`
}

// ChangeLogInput ...
type ChangeLogInput struct {
	Entity      string
	EntityID    string
	Type        string
	Columns     []string
	PrincipalID *string
	Date        graphql.Time
}

// NewChangeLog ...
func NewChangeLog(i ChangeLogInput) ChangeLog {
	id := uuid.Must(uuid.NewV4())
	return ChangeLog{
		IID:          id,
		IEntity:      i.Entity,
		IEntityID:    i.EntityID,
		IType:        i.Type,
		IColumns:     "#" + strings.Join(i.Columns, "#,#") + "#",
		IPrincipalID: i.PrincipalID,
		IDate:        i.Date.Time,
	}
}
