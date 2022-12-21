package bongo

import "github.com/rivo/tview"

type App struct {
	app   *tview.Application
	pages *tview.Pages
}

func Init() *App {
	return &App{
		app:   tview.NewApplication(),
		pages: tview.NewPages(),
	}
}
