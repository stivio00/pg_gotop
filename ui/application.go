package ui

import (
	"fmt"
	"log"

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
	table := pages.CreateActivityPage(a.db)
	a.pages.AddPage(ActivityPageId, table, true, true)

	// Transaction Page
	transactionslist := pages.CreateTransactionsPage(a.db)
	a.pages.AddPage(TransactionPageId, transactionslist, true, false)

	// IO Page
	io := pages.CreateIoPage(a.db)
	a.pages.AddPage(IOPageId, io, true, false)

	// Mem Page
	mem := pages.CreateMemPage(a.db)
	a.pages.AddPage(MemPageId, mem, true, false)

	// Mem Page
	disk := pages.CreateDiskPage(a.db)
	a.pages.AddPage(DiskPageId, disk, true, false)

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
