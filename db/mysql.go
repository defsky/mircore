package db

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var RealmDB *gorm.DB
var WorldDB *gorm.DB

func init() {
	dbc, err := Mysql(RealmDBConf)
	if err != nil {
		panic(err)
	}
	RealmDB = dbc

	dbc, err = Mysql(WorldDBConf)
	if err != nil {
		panic(err)
	}
	WorldDB = dbc
}

func Mysql(c *DbConfig) (*gorm.DB, error) {
	db, err := gorm.Open("mysql", c.String())
	if err != nil {
		return nil, err
	}

	pool := db.DB()
	pool.SetMaxIdleConns(5)
	pool.SetConnMaxLifetime(2 * time.Minute)
	pool.SetMaxOpenConns(20)

	return db, nil
}
