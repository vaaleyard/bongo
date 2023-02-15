package bongo

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func Ui(app *App) {
	treeNode := tview.NewTreeNode(".")

	app.tree.
		SetSelectedFunc(func(node *tview.TreeNode) {
			node.SetExpanded(!node.IsExpanded())
		}).
		SetRoot(treeNode).
		SetCurrentNode(treeNode).SetGraphics(false).
		SetTopLevel(1).
		SetPrefixes([]string{"> ", "> ", "- "}).
		SetBorder(true).
		SetTitle("Databases").SetTitleAlign(tview.AlignLeft).
		SetBorderPadding(0, 0, 1, 0).
		SetFocusFunc(func() {
			app.tree.SetBorderColor(focusColor.TrueColor()).
				SetTitleColor(focusColor).
				SetTitle("Databases*")
		}).
		SetBlurFunc(func() {
			app.tree.SetBorderColor(whiteColor).
				SetTitleColor(whiteColor).
				SetTitle("Databases")
		}).SetInputCapture(app.treeInputHandler)

	app.input.
		SetPlaceholder("Press q to exit, i to insert.").
		SetBorderPadding(0, 0, 1, 0).
		SetBorder(true).
		SetTitle("Input").SetTitleAlign(tview.AlignLeft).
		SetFocusFunc(func() {
			app.input.SetBorderColor(focusColor).
				SetBackgroundColor(tcell.ColorBlack.TrueColor()).
				SetTitleColor(focusColor).
				SetTitle("Input*")
		}).
		SetBlurFunc(func() {
			app.input.SetBorderColor(whiteColor).
				SetBackgroundColor(tcell.ColorBlack.TrueColor()).
				SetTitleColor(whiteColor).
				SetTitle("Input")
		}).SetInputCapture(app.inputAreaInputHandler)

	app.preview.
		SetBorder(true).
		SetTitle("Preview").SetTitleAlign(tview.AlignCenter).
		SetFocusFunc(func() {
			app.preview.SetBorderColor(focusColor).
				SetTitleColor(focusColor).
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
		SetBorderPadding(1, 1, 1, 1).SetBackgroundColor(tcell.ColorBlack.TrueColor()).
		SetBorderColor(tcell.ColorBlack.TrueColor())


	app.populateTree(treeNode)
    app.colorize()

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
				SetColor(tcell.ColorSteelBlue.TrueColor())
			collectionNode.AddChild(collectionNameTreeNode)
		}

		views, _ := app.database.Client.ListViews(databaseName)
		viewsNode := tview.NewTreeNode("Views").Collapse().
			SetColor(blueColor)
		databaseNode.AddChild(viewsNode)
		for _, view := range views {
			viewsTree := tview.NewTreeNode(view).
				SetColor(tcell.ColorSteelBlue.TrueColor())
			viewsNode.AddChild(viewsTree)
		}

		users, _ := app.database.Client.ListUsers(databaseName)
		usersNode := tview.NewTreeNode("Users").Collapse().
			SetColor(blueColor)
		databaseNode.AddChild(usersNode)
		for _, user := range users {
			userTreeNode := tview.NewTreeNode(user).
				SetColor(tcell.ColorSteelBlue.TrueColor())
			usersNode.AddChild(userTreeNode)
		}
	}
}
