package bongo

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"golang.design/x/clipboard"
	"strings"
)

func (app *App) previewInputHandler(event *tcell.EventKey) *tcell.EventKey {
	if event.Rune() == 89 || event.Rune() == 121 { // Y or y
		content := app.preview.GetText(true)

		err := clipboard.Init()
		if err != nil {
			panic(err)
		}

		clipboard.Write(clipboard.FmtText, []byte(content))
		return nil
	}

	return event
}

func (app *App) highlightSelectedDatabase() {
	databaseName := app.tree.GetCurrentNode().GetText()
	app.tree.GetCurrentNode().SetText(databaseName + " *")
}

// Future improvement: find a way to unhighlight only the past highlighted database
// instead of the whole tree.
func (app *App) unhighlightWholeTree() {
	databaseNodes := app.tree.GetRoot().GetChildren()
	for _, databaseChild := range databaseNodes {
		databaseChild.SetText(strings.Trim(databaseChild.GetText(), "* "))
	}
}

func (app *App) treeInputHandler(event *tcell.EventKey) *tcell.EventKey {
	// This is an interface to store the selected database.
	// It's used by the input block, the command will run in the database
	// stored in this interface.
	reference := make(map[bool]string)

	if event.Rune() == 83 || event.Rune() == 115 { // S or s
		lineSelected := app.tree.GetCurrentNode().GetText()
		databases := app.tree.GetRoot().GetChildren()
		for _, database := range databases {
			// User can select only database nodes and not his children
			if lineSelected == database.GetText() {
				// Don't add * to the reference in case user select it multiple times
				reference[true] = strings.Trim(lineSelected, " *")
				app.tree.GetCurrentNode().SetReference(reference)
				app.unhighlightWholeTree()
				app.highlightSelectedDatabase()

				return nil
			}
		}
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
		if referenceInterface == nil {
			// Default database if user doesn't select one
			db = "admin"
		} else {
			reference := referenceInterface.(map[bool]string)
			db = reference[true]
		}

		command := app.input.GetText()
		output := app.database.Client.RunCommand(db, command)
		_, _ = fmt.Fprint(app.preview, output)

		return nil
	}

	return event
}
