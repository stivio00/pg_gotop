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

func (d *DbConnection) GetActivity() []PgStatActivity {

	rows, err := d.db.Query(GetActivityQuery)
	if err != nil {
		log.Fatal(err)
	}
	var activity []PgStatActivity = make([]PgStatActivity, 0)
	for rows.Next() {
		row := PgStatActivity{}
		rows.Scan(&row.Datid,
			&row.Datname,
			&row.Pid,
			&row.LeaderPid,
			&row.Usesysid,
			&row.Usename,
			&row.ApplicationName,
			&row.ClientAddr,
			&row.ClientHostname,
			&row.ClientPort,
			&row.BackendStart,
			&row.XactStart,
			&row.QueryStart,
			&row.StateChange,
			&row.WaitEventType,
			&row.WaitEvent,
			&row.State,
			&row.BackendXid,
			&row.BackendXmin,
			&row.QueryId,
			&row.Query,
			&row.BackendType)
		activity = append(activity, row)
	}
	return activity
}
