package src

import (
	"net/url"
	"strings"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// DB ...
type DB struct {
	db *gorm.DB
}

// NewDB ...
func NewDB(db *gorm.DB) *DB {
	v := DB{db}
	return &v
}

// NewDBWithString ...
func NewDBWithString(urlString string) *DB {
	u, err := url.Parse(urlString)
	if err != nil {
		panic(err)
	}

	urlString = strings.Replace(urlString, u.Scheme+"://", "", 1)

	db, err := gorm.Open(u.Scheme, urlString)
	if err != nil {
		panic(err)
	}
	db.LogMode(true)
	return NewDB(db)
}

// AutoMigrate ...
func (db *DB) AutoMigrate(values ...interface{}) {
	db.db.AutoMigrate(values...)
}

// Close ...
func (db *DB) Close() error {
	return db.db.Close()
}
