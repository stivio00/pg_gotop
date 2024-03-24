package pages

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/stivio00/pg_gotop/ui"
)

func CreateHelpForm() *tview.Form {
	helpPage := tview.NewForm()
	helpPage.AddTextView("help", "dd", 10, 2, true, true)
	return helpPage
}

func CreateHelpPage() *ui.Page {
	return &ui.Page{
		Title:          "Help",
		HotKeySelector: tcell.KeyF1,
		Content:        CreateHelpForm().Box,
	}
}
