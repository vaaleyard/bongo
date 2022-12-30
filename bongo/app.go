package bongo

import (
	"github.com/rivo/tview"
	"github.com/vaaleyard/bongo/database"
	m "github.com/vaaleyard/bongo/database/mongo"
)

type App struct {
	app       *tview.Application
	pages     *tview.Pages
	treeView  *tview.TreeView
	preview   *tview.TextView
	inputArea *tview.TextArea
	database  *database.Service
}

func Init(uri string) (*App, error) {
	a := App{
		app:       tview.NewApplication(),
		pages:     tview.NewPages(),
		treeView:  tview.NewTreeView(),
		preview:   tview.NewTextView(),
		inputArea: tview.NewTextArea(),
	}

	mongoConnection := m.NewConnection(uri)
	dbService := database.New(mongoConnection)
	a.database = dbService

	return &a, nil
}
