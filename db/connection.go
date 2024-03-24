package db

import (
	"database/sql"
	"log"
)

type DbConnection struct {
	db *sql.DB
}

func New(connectionString string) *DbConnection {
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	return &DbConnection{
		db: db,
	}
}

func (d *DbConnection) GetDbVersion() string {
	row, err := d.db.Query("SELECT version();")
	if err != nil {
		log.Fatal(err)
	}
	var dbVersion string
	row.Next()
	row.Scan(&dbVersion)
	return dbVersion
}
