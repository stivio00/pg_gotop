package ui

import (
	"fmt"
	"log"
	"strconv"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/stivio00/pg_gotop/db"
	"github.com/stivio00/pg_gotop/ui/pages"
)

type App struct {
	status Status
	app    *tview.Application
	pages  *tview.Pages
	db     *db.DbConnection
	//internal widgets
	statusBar *tview.TextView
}

type Status int

const (
	ConnectingStatus  Status = 0
	ConnectedStatus   Status = iota
	DisconectedStatus Status = iota
)

const (
	HelpPageId        string = "help"
	ActivityPageId    string = "activity"
	TransactionPageId string = "xact"
	IOPageId          string = "io"
	MemPageId         string = "mem"
	DiskPageId        string = "disk"
	InfoPageId        string = "info"
)

func New(db *db.DbConnection) *App {
	app := tview.NewApplication()
	// Pages Container
	pages := tview.NewPages()
	pages.SetTitle("pg_gotop v1.0")
	pages.SetBorder(true)

	// Labels
	helpLabel := tview.NewTextView().
		SetText("F1-Help | F2-Activity | F3-Transactions | F4-IO | F5-Refresh | F6-Mem | F7-Disk | F8-Info")
	statusLabel := tview.NewTextView().
		SetText("Loading...")

	// Main Layout - Flex
	mainLayout := tview.NewFlex()
	mainLayout.SetTitle("pg_gotop v1.0")
	mainLayout.SetDirection(tview.FlexRow)

	mainLayout.AddItem(helpLabel, 1, 0, false)
	mainLayout.AddItem(pages, 0, 100, true)
	mainLayout.AddItem(statusLabel, 1, 0, false)

	app.SetRoot(mainLayout, true).SetFocus(pages)

	return &App{
		app:       app,
		status:    DisconectedStatus,
		db:        db,
		pages:     pages,
		statusBar: statusLabel,
	}
}

func (a *App) BuildLayout() {

	// Help page
	help := pages.CreateHelpForm()
	a.pages.AddPage(HelpPageId, help, true, true)

	// Activity Page
	table := tview.NewTable()
	table.SetBorders(true)

	table.SetCell(0, 0, tview.NewTableCell("database").SetTextColor(tcell.ColorYellow).SetAlign(tview.AlignCenter))
	table.SetCell(0, 1, tview.NewTableCell("username").SetTextColor(tcell.ColorYellow).SetAlign(tview.AlignCenter))
	table.SetCell(0, 2, tview.NewTableCell("app").SetTextColor(tcell.ColorYellow).SetAlign(tview.AlignCenter))
	table.SetCell(0, 3, tview.NewTableCell("pid").SetTextColor(tcell.ColorYellow).SetAlign(tview.AlignCenter))
	table.SetCell(0, 4, tview.NewTableCell("state").SetTextColor(tcell.ColorYellow).SetAlign(tview.AlignCenter))
	table.SetCell(0, 5, tview.NewTableCell("type").SetTextColor(tcell.ColorYellow).SetAlign(tview.AlignCenter))

	pgActivity := a.db.GetActivity()
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
		if key == tcell.KeyEscape {
			a.app.Stop()
		}
		if key == tcell.KeyEnter {
			table.SetSelectable(true, true)
		}
	}).SetSelectedFunc(func(row int, column int) {
		table.GetCell(row, column).SetTextColor(tcell.ColorRed)
		table.SetSelectable(false, false)
	})

	a.pages.AddPage(ActivityPageId, table, true, true)

	// Transaction Page

	transactionslist := pages.CreateTransactionsPage(a.db)
	a.pages.AddPage(TransactionPageId, transactionslist, true, false)

	// Info Page
	tree := pages.CreateInfoTree(a.db)
	a.pages.AddPage(InfoPageId, tree, true, false)

	a.bindKeys()
}

func (a App) bindKeys() {
	a.app.SetInputCapture(func(e *tcell.EventKey) *tcell.EventKey {
		switch e.Key() {
		case tcell.KeyF1:
			a.pages.SwitchToPage(HelpPageId)
		case tcell.KeyF2:
			a.pages.SwitchToPage(ActivityPageId)
		case tcell.KeyF3:
			a.pages.SwitchToPage(TransactionPageId)
		case tcell.KeyF4:
			a.pages.SwitchToPage(IOPageId)
		case tcell.KeyF5:
			//Todo: Refresh
			a.pages.SwitchToPage(HelpPageId)
		case tcell.KeyF6:
			a.pages.SwitchToPage(MemPageId)
		case tcell.KeyF7:
			a.pages.SwitchToPage(DiskPageId)
		case tcell.KeyF8:
			a.pages.SwitchToPage(InfoPageId)
		}
		return e
	})
}

func (a App) AddPage(page Page) {

}

func (a *App) setStatus(statusText string) {
	a.statusBar.SetText(statusText)
}

func (a App) Run() {
	current := a.db.GetDbCurrentConnection()
	status := fmt.Sprintf("Connected to %s as %s - %s (%s)", current.Database, current.User, current.Version, current.Size)
	a.setStatus(status)
	if err := a.app.Run(); err != nil {
		log.Fatal(err)
	}
}
