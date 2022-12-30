package bongo

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
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

	app.tree.
		SetRoot(treeNode).
		SetCurrentNode(treeNode).SetGraphics(false).
		SetTopLevel(1).
		SetPrefixes([]string{"> "}).
		SetBorder(true).
		SetTitle("Finder").SetTitleAlign(tview.AlignLeft).
		SetBorderPadding(0, 0, 1, 0).
		SetFocusFunc(func() {
			app.tree.SetBorderColor(blueColor).
				SetTitleColor(blueColor).
				SetTitle("Finder*")
		}).
		SetBlurFunc(func() {
			app.tree.SetBorderColor(whiteColor).
				SetTitleColor(whiteColor).
				SetTitle("Finder")
		})

	app.input.
		SetPlaceholder("Press q to exit, i to insert.").
		SetPlaceholderStyle(tcell.StyleDefault.Attributes(tcell.AttrDim)).
		SetBorderPadding(0, 0, 1, 0).
		SetBorder(true).
		SetTitle("Input").SetTitleAlign(tview.AlignLeft).
		SetInputCapture(app.inputAreaInputHandler).
		SetFocusFunc(func() {
			app.input.SetBorderColor(blueColor).
				SetTitleColor(blueColor).
				SetTitle("Input*")
		}).
		SetBlurFunc(func() {
			app.input.SetBorderColor(whiteColor).
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
		AddItem(app.input, 0, 1, false).SetDirection(tview.FlexRow).
		AddItem(
			tview.NewFlex().SetDirection(tview.FlexColumn).
				AddItem(app.tree, 0, 1, false).
				AddItem(app.preview, 0, 5, false),
			0, 14, false).
		SetBorderPadding(1, 1, 1, 1)

	app.populateFinder(treeNode)
	app.tree.SetSelectedFunc(func(node *tview.TreeNode) {
		node.SetExpanded(!node.IsExpanded())
	}).SetInputCapture(app.treeInputHandler)

	app.app.SetInputCapture(app.appInputHandler)

	app.pages.AddPage("layout", layout, true, true)
	if err := app.app.SetRoot(app.pages, true).
		SetFocus(app.pages).Run(); err != nil {
		panic(err)
	}
}

func (app *App) populateFinder(target *tview.TreeNode) {
	names, _ := app.database.Client.ListDatabaseNames()
	for _, databaseName := range names {
		databaseNode := tview.NewTreeNode(databaseName).
			SetColor(blueColor)
		target.AddChild(databaseNode)

		collections, _ := app.database.Client.ListCollections(databaseName)
		collectionNode := tview.NewTreeNode("Collections").Collapse().
			SetColor(blueColor)
		databaseNode.AddChild(collectionNode)
		for _, collection := range collections {
			collectionNameTreeNode := tview.NewTreeNode(collection).
				SetColor(lightWhite)
			collectionNode.AddChild(collectionNameTreeNode)
		}

		views, _ := app.database.Client.ListViews(databaseName)
		viewsNode := tview.NewTreeNode("Views").Collapse().
			SetColor(blueColor)
		databaseNode.AddChild(viewsNode)
		for _, view := range views {
			viewsTree := tview.NewTreeNode(view).
				SetColor(lightWhite)
			viewsNode.AddChild(viewsTree)
		}

		users, _ := app.database.Client.ListUsers(databaseName)
		usersNode := tview.NewTreeNode("Users").Collapse().
			SetColor(blueColor)
		databaseNode.AddChild(usersNode)
		for _, user := range users {
			userTreeNode := tview.NewTreeNode(user).
				SetColor(lightWhite)
			usersNode.AddChild(userTreeNode)
		}
	}
}

func (app *App) treeInputHandler(event *tcell.EventKey) *tcell.EventKey {
	reference := make(map[bool]string)

	if event.Rune() == 83 || event.Rune() == 115 { // S or s
		if app.tree.GetCurrentNode().GetReference() == nil {
			reference[true] = app.tree.GetCurrentNode().GetText()
			app.tree.GetCurrentNode().SetReference(reference)
		} else {
			reference[false] = app.tree.GetCurrentNode().GetText()
			app.tree.GetCurrentNode().SetReference(reference)
		}
		return nil
	}

	return event
}

func (app *App) appInputHandler(event *tcell.EventKey) *tcell.EventKey {
	if !app.input.HasFocus() {
		switch event.Rune() {
		case 81, 113: // Q or q or ctrl-c
			app.app.Stop()
			return nil
		case 73, 105, 58: //I or i or :
			app.app.SetFocus(app.input)
			return nil
		case 68, 100: // D or d
			app.app.SetFocus(app.tree)
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

		var db string
		referenceInterface := app.tree.GetCurrentNode().GetReference()
		reference := referenceInterface.(map[bool]string)
		if reference == nil {
			db = "admin"
		} else {
			db = reference[true]
		}

		command := app.input.GetText()
		output := app.database.Client.RunCommand(db, command)
		_, _ = fmt.Fprint(app.preview, output)

		return nil
	}

	return event
}
