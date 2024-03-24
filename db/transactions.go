package db

import (
	"log"

	"github.com/stivio00/pg_gotop/model"
)

func (db *DbConnection) GetTransactions() []model.Transaction {

	transactions := make([]model.Transaction, 0)
	rows, err := db.db.Query("SELECT  backend_xid::text, pid, xact_start::text, query FROM pg_stat_activity WHERE backend_xid IS NOT NULL;")
	if err != nil {
		log.Println(err)
		return transactions
	}

	for rows.Next() {
		t := model.Transaction{}
		rows.Scan(&t.Xactid, &t.Pid, &t.Started, &t.Sql)
		transactions = append(transactions, t)
	}

	return transactions
}
