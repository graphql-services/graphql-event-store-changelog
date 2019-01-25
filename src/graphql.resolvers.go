package src

import (
	"github.com/graphql-services/graphql-event-store-changelog/model"
)

type Query struct {
	db *DB
}

func NewQuery(db *DB) Query {
	return Query{db}
}

type changelogParams struct {
	PrincipalID, Entity, EntityID, Type *string
	Columns                             *[]string
	Limit                               int32
}

// ChangeLog ...
func (q *Query) ChangeLog(params changelogParams) ([]model.ChangeLog, error) {
	var items []model.ChangeLog
	query := q.db.db
	if params.Entity != nil {
		query = query.Where(&model.ChangeLog{IEntity: *params.Entity})
	}
	if params.EntityID != nil {
		query = query.Where(&model.ChangeLog{IEntityID: *params.EntityID})
	}
	if params.Type != nil {
		query = query.Where(&model.ChangeLog{IType: *params.Type})
	}
	query = query.Where(&model.ChangeLog{IPrincipalID: params.PrincipalID})
	// if params.Columns != nil {
	// 	query = query.Where(&model.ChangeLog{IColumns: params.Columns})
	// }
	query = query.Limit(params.Limit)
	query = query.Order("IDate")
	query.Find(&items)
	return items, query.Error
}
