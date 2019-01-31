package model

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/graphql-services/go-saga/eventstore"

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
	IChanges     string     `gorm:"column:changes;type:text"`
	IPrincipalID *string    `gorm:"column:principalId"`
	IDate        time.Time  `gorm:"column:date"`
}

// ChangeLogInput ...
type ChangeLogInput struct {
	Entity      string
	EntityID    string
	Type        string
	Columns     []string
	Changes     string
	PrincipalID *string
	Date        graphql.Time
}

// NewChangeLog ...
func NewChangeLog(e eventstore.Event) ChangeLog {
	changes := []ChangeLogChange{}

	for _, column := range e.Columns {
		var oldValue *string
		var newValue *string

		for _, val := range e.OldValues {
			if val.Name == column {
				oldValue = val.Value
				break
			}
		}
		for _, val := range e.NewValues {
			if val.Name == column {
				newValue = val.Value
				break
			}
		}

		ch := ChangeLogChange{
			IColumn:   column,
			IOldValue: oldValue,
			INewValue: newValue,
		}
		changes = append(changes, ch)
	}

	changesJSON, err := json.Marshal(changes)
	if err != nil {
		panic(err)
	}
	id := uuid.Must(uuid.NewV4())
	return ChangeLog{
		IID:          id,
		IEntity:      e.Entity,
		IEntityID:    e.EntityID,
		IType:        e.Type,
		IColumns:     "#" + strings.Join(e.Columns, "#,#") + "#",
		IChanges:     string(changesJSON),
		IPrincipalID: e.PrincipalID,
		IDate:        e.Date,
	}
}
