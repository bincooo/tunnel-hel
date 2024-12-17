package database

import (
	"github.com/jinzhu/gorm"
	"tunnel-hel/config"
)
import _ "github.com/jinzhu/gorm/dialects/mysql"

var (
	db *gorm.DB
)

func DB() *gorm.DB {
	if db == nil {
		db = NewDB()
	}
	return db
}

func NewDB() (db *gorm.DB) {
	c := config.Get().Database
	db, err := gorm.Open(c.Dialect, c.Args)
	if err != nil {
		panic(err)
	}
	db.SingularTable(true)
	db.DB().SetMaxOpenConns(c.MaxOpen)
	db.DB().SetMaxIdleConns(c.MaxIdle)
	return db
}
