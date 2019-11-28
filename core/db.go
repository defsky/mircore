package core

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"

	"fmt"
	"mircore/utils/log"
	"time"
)

const (
	USERNAME = "root"
	PASSWORD = "mysql"
	NETWORK  = "tcp"
	SERVER   = "localhost"
	PORT     = 3306
	DATABASE = "phpmir2"
)

var slog = log.Core

type DB struct {
	conn *sql.DB
}

func (db *DB) Instance() *DB {
	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)
	DB, err := sql.Open("mysql", dsn)

	if err != nil {
		fmt.Printf("Open mysql failed,err:%v\n", err)
		return nil
	}

	DB.SetConnMaxLifetime(100 * time.Second)
	DB.SetMaxOpenConns(100)
	DB.SetMaxIdleConns(16)

	db.conn = DB

	return db
}
