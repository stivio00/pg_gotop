package pages

import (
	"strconv"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/stivio00/pg_gotop/db"
)

func CreateActivityPage(db *db.DbConnection) *tview.Table {
	table := tview.NewTable()
	table.SetBorders(true)

	table.SetCell(0, 0, tview.NewTableCell("database").SetTextColor(tcell.ColorYellow).SetAlign(tview.AlignCenter))
	table.SetCell(0, 1, tview.NewTableCell("username").SetTextColor(tcell.ColorYellow).SetAlign(tview.AlignCenter))
	table.SetCell(0, 2, tview.NewTableCell("app").SetTextColor(tcell.ColorYellow).SetAlign(tview.AlignCenter))
	table.SetCell(0, 3, tview.NewTableCell("pid").SetTextColor(tcell.ColorYellow).SetAlign(tview.AlignCenter))
	table.SetCell(0, 4, tview.NewTableCell("state").SetTextColor(tcell.ColorYellow).SetAlign(tview.AlignCenter))
	table.SetCell(0, 5, tview.NewTableCell("type").SetTextColor(tcell.ColorYellow).SetAlign(tview.AlignCenter))

	pgActivity := db.GetActivity()
	rows := len(pgActivity)
	color := tcell.ColorWhite
	for r := 0; r < rows; r++ {
		table.SetCell(r+1, 0, tview.NewTableCell(pgActivity[r].Datname).SetTextColor(color).SetAlign(tview.AlignCenter))
		table.SetCell(r+1, 1, tview.NewTableCell(pgActivity[r].Usename).SetTextColor(color).SetAlign(tview.AlignCenter))
		table.SetCell(r+1, 2, tview.NewTableCell(pgActivity[r].ApplicationName).SetTextColor(color).SetAlign(tview.AlignCenter))
		table.SetCell(r+1, 3, tview.NewTableCell(strconv.Itoa(pgActivity[r].Pid)).SetTextColor(color).SetAlign(tview.AlignCenter))
		table.SetCell(r+1, 4, tview.NewTableCell(pgActivity[r].State).SetTextColor(color).SetAlign(tview.AlignCenter))
		table.SetCell(r+1, 5, tview.NewTableCell(pgActivity[r].BackendType).SetTextColor(color).SetAlign(tview.AlignCenter))
	}
	table.Select(0, 0).SetFixed(1, 1).SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEnter {
			table.SetSelectable(true, true)
		}
	}).SetSelectedFunc(func(row int, column int) {
		table.GetCell(row, column).SetTextColor(tcell.ColorRed)
		table.SetSelectable(false, false)
	})

	return table
}
