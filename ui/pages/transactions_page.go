package pages

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/stivio00/pg_gotop/db"
)

func CreateTransactionsPage(db *db.DbConnection) *tview.List {
	list := tview.NewList()
	list.SetBackgroundColor(tcell.ColorDarkGreen)
	list.SetBorder(true)
	transactions := db.GetTransactions()

	if len(transactions) == 0 {
		list.AddItem("No current transactions running", "", ' ', func() {})
		return list
	}

	for _, t := range transactions {
		line := fmt.Sprintf("Id: %s, Started:%s (%s),  Pid: %d ", t.Xactid, t.Started, t.Duration, t.Pid)
		list.AddItem(line, t.Sql, ' ', func() {})
	}

	return list
}
