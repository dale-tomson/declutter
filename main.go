package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"

	"github.com/dale-tomson/declutter/internal/icon"
	"github.com/dale-tomson/declutter/internal/theme"
	"github.com/dale-tomson/declutter/internal/ui"
)

func main() {
	a := app.NewWithID("com.github.dale-tomson.declutter")
	a.Settings().SetTheme(theme.New())
	a.SetIcon(icon.Resource())

	w := a.NewWindow("Declutter")
	w.Resize(fyne.NewSize(700, 500))
	w.CenterOnScreen()

	appUI := ui.New(w)
	w.SetContent(appUI.GetContent())

	w.SetCloseIntercept(func() {
		w.Close()
		a.Quit()
	})

	w.ShowAndRun()
}
