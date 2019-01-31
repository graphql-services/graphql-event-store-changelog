package src

import (
	"encoding/json"

	"github.com/golang/glog"

	"github.com/graphql-services/go-saga/eventstore"
	"github.com/graphql-services/graphql-event-store-changelog/model"
)

const (
	latestEventMetaName = "latestEvent"
)

// GetLatestEvent returns [Event,found,error]
func GetLatestEvent(db *DB) (*eventstore.Event, bool, error) {
	m, err := db.GetMeta(latestEventMetaName)
	if err != nil {
		return nil, false, err
	}

	// event not found
	if m == nil {
		return nil, false, nil
	}

	// meta found, but event is empty
	if m.Value == "null" {
		return nil, true, nil
	}

	var e eventstore.Event
	err = json.Unmarshal([]byte(m.Value), &e)
	if err != nil {
		return nil, false, err
	}
	return &e, true, nil
}

func storeLatestEvent(e *eventstore.Event, db *DB) error {
	value, err := json.Marshal(e)
	if err != nil {
		return err
	}
	return db.SaveMeta(latestEventMetaName, string(value))
}

// ImportEvents ...
func ImportEvents(events []eventstore.Event, db *DB) error {
	glog.Info("Importing events", len(events))
	if len(events) == 0 {
		return storeLatestEvent(nil, db)
	}

	tx := db.db.Begin()
	for _, event := range events {
		cl := model.NewChangeLog(event)
		err := db.db.Create(&cl).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	if err := tx.Commit().Error; err != nil {
		return err
	}

	return storeLatestEvent(&events[len(events)-1], db)
}
