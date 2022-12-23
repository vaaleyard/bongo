package bongo

import (
	"github.com/rivo/tview"
	"github.com/vaaleyard/bongo/mongo"
)

type App struct {
	app         *tview.Application
	pages       *tview.Pages
	treeView    *tview.TreeView
	preview     *tview.TextView
	inputArea   *tview.TextArea
	mongoClient *mongo.Mongo
}

func Init(uri string) (*App, error) {
	a := App{
		app:       tview.NewApplication(),
		pages:     tview.NewPages(),
		treeView:  tview.NewTreeView(),
		preview:   tview.NewTextView(),
		inputArea: tview.NewTextArea(),
	}

	client, _ := mongo.NewConnection(uri)
	a.mongoClient = mongo.Interface(client)

	return &a, nil
}
