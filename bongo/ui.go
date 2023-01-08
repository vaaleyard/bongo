package bongo

import (
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
		SetSelectedFunc(func(node *tview.TreeNode) {
			node.SetExpanded(!node.IsExpanded())
		}).
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
				SetTitle("Databases*")
		}).
		SetBlurFunc(func() {
			app.tree.SetBorderColor(whiteColor).
				SetTitleColor(whiteColor).
				SetTitle("Databases")
		}).SetInputCapture(app.treeInputHandler)

	app.input.
		SetPlaceholder("Press q to exit, i to insert.").
		SetPlaceholderStyle(tcell.StyleDefault.Attributes(tcell.AttrDim)).
		SetBorderPadding(0, 0, 1, 0).
		SetBorder(true).
		SetTitle("Input").SetTitleAlign(tview.AlignLeft).
		SetFocusFunc(func() {
			app.input.SetBorderColor(blueColor).
				SetTitleColor(blueColor).
				SetTitle("Input*")
		}).
		SetBlurFunc(func() {
			app.input.SetBorderColor(whiteColor).
				SetTitleColor(whiteColor).
				SetTitle("Input")
		}).SetInputCapture(app.inputAreaInputHandler)

	app.preview.
		SetBorder(true).
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
		}).SetInputCapture(app.previewInputHandler)

	layout := tview.NewFlex()
	layout.
		AddItem(app.input, 0, 1, false).SetDirection(tview.FlexRow).
		AddItem(
			tview.NewFlex().SetDirection(tview.FlexColumn).
				AddItem(app.tree, 0, 1, false).
				AddItem(app.preview, 0, 5, false),
			0, 14, false).
		SetBorderPadding(1, 1, 1, 1)

	app.populateTree(treeNode)

	app.app.SetInputCapture(app.appInputHandler)

	app.pages.AddPage("layout", layout, true, true)
	if err := app.app.SetRoot(app.pages, true).
		SetFocus(app.pages).Run(); err != nil {
		panic(err)
	}
}

func (app *App) populateTree(target *tview.TreeNode) {
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
