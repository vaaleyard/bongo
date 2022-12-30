package bongo

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/vaaleyard/bongo/mongo"
)

const (
	blueColor  = tcell.ColorCornflowerBlue
	whiteColor = tcell.ColorLightGrey
	lightWhite = tcell.ColorLightGrey
)

func Ui(app *App) {
	treeNode := tview.NewTreeNode(".")
	treeNode.
		SetColor(tcell.ColorGreen)

	app.treeView.
		SetRoot(treeNode).
		SetCurrentNode(treeNode).SetGraphics(false).
		SetTopLevel(1).
		SetPrefixes([]string{"> "}).
		SetBorder(true).
		SetTitle("Finder").SetTitleAlign(tview.AlignLeft).
		SetBorderPadding(0, 0, 1, 0).
		SetFocusFunc(func() {
			app.treeView.SetBorderColor(blueColor).
				SetTitleColor(blueColor).
				SetTitle("Finder*")
		}).
		SetBlurFunc(func() {
			app.treeView.SetBorderColor(whiteColor).
				SetTitleColor(whiteColor).
				SetTitle("Finder")
		})

	app.inputArea.
		SetPlaceholder("Press q to exit, i to insert.").
		SetPlaceholderStyle(tcell.StyleDefault.Attributes(tcell.AttrDim)).
		SetBorderPadding(0, 0, 1, 0).
		SetBorder(true).
		SetTitle("Input").SetTitleAlign(tview.AlignLeft).
		SetInputCapture(app.inputAreaInputHandler).
		SetFocusFunc(func() {
			app.inputArea.SetBorderColor(blueColor).
				SetTitleColor(blueColor).
				SetTitle("Input*")
		}).
		SetBlurFunc(func() {
			app.inputArea.SetBorderColor(whiteColor).
				SetTitleColor(whiteColor).
				SetTitle("Input")
		})

	app.preview.SetBorder(true).
		SetTitle("Preview").SetTitleAlign(tview.AlignCenter).
		SetFocusFunc(func() {
			app.preview.SetBorderColor(blueColor).
				SetTitleColor(blueColor).
				SetTitle("Preview*")
		}).
		SetBlurFunc(func() {
			app.preview.SetBorderColor(whiteColor).
				SetTitleColor(whiteColor).
				SetTitle("Preview")
		})

	layout := tview.NewFlex()
	layout.
		AddItem(app.inputArea, 0, 1, false).SetDirection(tview.FlexRow).
		AddItem(
			tview.NewFlex().SetDirection(tview.FlexColumn).
				AddItem(app.treeView, 0, 1, false).
				AddItem(app.preview, 0, 5, false),
			0, 14, false).
		SetBorderPadding(1, 1, 1, 1)

	app.populateFinder(treeNode, app.mongoClient)
	app.treeView.SetSelectedFunc(selectNode)
	app.treeView.SetInputCapture(app.treeInputHandler)
	app.app.SetInputCapture(app.appInputHandler)

	app.pages.AddPage("layout", layout, true, true)
	if err := app.app.SetRoot(app.pages, true).
		SetFocus(app.pages).Run(); err != nil {
		panic(err)
	}
}

func (app *App) populateFinder(target *tview.TreeNode, mongoClient *mongo.Mongo) {
	dbs, _ := mongoClient.ListDatabaseNames()
	for _, db := range dbs {
		nodeDB := tview.NewTreeNode(db).
			SetColor(blueColor)
		target.AddChild(nodeDB)

		collections, _ := mongoClient.ListCollections(db)
		collectionNode := tview.NewTreeNode("Collections").Collapse().
			SetColor(blueColor)
		nodeDB.AddChild(collectionNode)
		for _, collection := range collections {
			collectionTreeNode := tview.NewTreeNode(collection).
				SetColor(lightWhite)
			collectionNode.AddChild(collectionTreeNode)
		}

		views, _ := mongoClient.ListViews(db)
		viewsNode := tview.NewTreeNode("Views").Collapse().
			SetColor(blueColor)
		nodeDB.AddChild(viewsNode)
		for _, view := range views {
			viewsTree := tview.NewTreeNode(view).
				SetColor(lightWhite)
			viewsNode.AddChild(viewsTree)
		}

		users, _ := mongoClient.ListUsers(db)
		usersNode := tview.NewTreeNode("Users").Collapse().
			SetColor(blueColor)
		nodeDB.AddChild(usersNode)
		for _, user := range users {
			userTreeNode := tview.NewTreeNode(user).
				SetColor(lightWhite)
			usersNode.AddChild(userTreeNode)
		}
	}
}

func selectNode(node *tview.TreeNode) {
	node.SetExpanded(!node.IsExpanded())
}

func (app *App) treeInputHandler(event *tcell.EventKey) *tcell.EventKey {
	reference := make(map[bool]string)

	if event.Rune() == 83 || event.Rune() == 115 { // S or s
		if app.treeView.GetCurrentNode().GetReference() == nil {
			reference[true] = app.treeView.GetCurrentNode().GetText()
			app.treeView.GetCurrentNode().SetReference(reference)
		} else {
			reference[false] = app.treeView.GetCurrentNode().GetText()
			app.treeView.GetCurrentNode().SetReference(reference)
		}
		return nil
	}

	return event
}

func (app *App) appInputHandler(event *tcell.EventKey) *tcell.EventKey {
	if !app.inputArea.HasFocus() {
		switch event.Rune() {
		case 81, 113: // Q or q or ctrl-c
			app.app.Stop()
			return nil
		case 73, 105, 58: //I or i or :
			app.app.SetFocus(app.inputArea)
			return nil
		case 68, 100: // D or d
			app.app.SetFocus(app.treeView)
			return nil
		case 80, 112: // P or p
			app.app.SetFocus(app.preview)
			return nil
		}
	}
	if event.Key() == tcell.KeyESC {
		app.app.SetFocus(app.pages)
		return nil
	}

	return event
}

func (app *App) inputAreaInputHandler(event *tcell.EventKey) *tcell.EventKey {
	if event.Key() == tcell.KeyEnter {
		app.preview.Clear()

		var database string
		refInterface := app.treeView.GetCurrentNode().GetReference()
		reference := refInterface.(map[bool]string)
		if reference == nil {
			database = "admin"
		} else {
			database = reference[true]
		}

		command := app.inputArea.GetText()
		output := app.mongoClient.RunCommand(database, command)
		_, _ = fmt.Fprint(app.preview, output)

		return nil
	}

	return event
}
