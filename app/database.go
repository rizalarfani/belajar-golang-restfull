package app

import (
	"database/sql"
	"rizalarfani/belajar-restful-api/helper"
	"time"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "bela_golang:CcNehbWTzzy7%RVe@tcp(172.17.0.1:3306)/bela_golang")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxLifetime(10 * time.Minute)

	return db
}
