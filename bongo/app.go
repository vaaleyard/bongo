package bongo

import "github.com/rivo/tview"

type App struct {
	app      *tview.Application
	pages    *tview.Pages
	treeView *tview.TreeView
	preview  *tview.Box
	inputBox *tview.Box
}

func Init() *App {
	return &App{
		app:      tview.NewApplication(),
		pages:    tview.NewPages(),
		treeView: tview.NewTreeView(),
		preview:  tview.NewBox(),
		inputBox: tview.NewBox(),
	}
}
