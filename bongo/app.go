package bongo

import "github.com/rivo/tview"

type App struct {
	app       *tview.Application
	pages     *tview.Pages
	treeView  *tview.TreeView
	preview   *tview.TextView
	inputArea *tview.TextArea
}

func Init() *App {
	return &App{
		app:       tview.NewApplication(),
		pages:     tview.NewPages(),
		treeView:  tview.NewTreeView(),
		preview:   tview.NewTextView(),
		inputArea: tview.NewTextArea(),
	}
}
