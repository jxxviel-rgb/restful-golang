package app

import (
	"database/sql"
	"restful-golang/helper"
	"time"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root:fikri@tcp(localhost:3306)/belajarrestful")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)
	return db
}
