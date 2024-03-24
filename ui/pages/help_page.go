package pages

import (
	"github.com/rivo/tview"
)

func CreateHelpForm() *tview.Form {
	helpPage := tview.NewForm()
	helpPage.AddTextView("This App is a Work in progress", "dd", 10, 2, true, true)
	return helpPage
}
