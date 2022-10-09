package app

import (
	"database/sql"
	"time"

	"github.com/jxxviel-rgb/restful-golang/helper"
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
