package bongo

import "github.com/gdamore/tcell/v2"

const (
	blueColor  = tcell.ColorCornflowerBlue
	whiteColor = tcell.ColorLightGrey
	focusColor = tcell.ColorDodgerBlue
    blackColor = tcell.ColorBlack
)

func (app *App) colorize() {
    app.tree.
		SetBackgroundColor(blackColor.TrueColor()).
		SetTitleColor(whiteColor.TrueColor()).
		SetBorderColor(whiteColor.TrueColor())
	
    app.input.
		SetTextStyle(tcell.StyleDefault.Background(blackColor.TrueColor()).Foreground(whiteColor.TrueColor())).
		SetPlaceholderStyle(tcell.StyleDefault.Dim(true).Background(tcell.ColorBlack.TrueColor())).
        SetBackgroundColor(tcell.ColorBlack.TrueColor()).
		SetTitleColor(whiteColor.TrueColor()).
		SetBorderColor(whiteColor.TrueColor())

    app.preview.
		SetTextColor(whiteColor.TrueColor()).
        SetBackgroundColor(tcell.ColorBlack.TrueColor())
}
