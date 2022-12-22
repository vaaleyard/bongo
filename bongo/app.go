package bongo

import "github.com/rivo/tview"

type App struct {
	app      *tview.Application
	pages    *tview.Pages
	treeNode *tview.TreeNode
	treeView *tview.TreeView
}

func Init() *App {
	return &App{
		app:      tview.NewApplication(),
		pages:    tview.NewPages(),
		treeNode: tview.NewTreeNode("."),
		treeView: tview.NewTreeView(),
	}
}
