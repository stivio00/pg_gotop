package db

import (
	"log"

	"github.com/stivio00/pg_gotop/model"
)

const (
	GetActivityQuery = "SELECT * FROM pg_catalog.pg_stat_activity;"
)

func (d *DbConnection) GetActivity() []model.PgStatActivity {

	rows, err := d.db.Query(GetActivityQuery)
	if err != nil {
		log.Fatal(err)
	}
	var activity []model.PgStatActivity = make([]model.PgStatActivity, 0)
	for rows.Next() {
		row := model.PgStatActivity{}
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
