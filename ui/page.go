package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Page struct {
	Title          string
	HotKeySelector tcell.Key
	Content        *tview.Box
}
