package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello")
	l := widget.NewLabel("Hello Fyne!")

	mm := fyne.NewMainMenu(
		fyne.NewMenu("File",
			fyne.NewMenuItem("New", func() {
				l.SetText("select 'New' menu item.")
			}),
			fyne.NewMenuItem("Quit", func() {
				a.Quit()
			}),
		),
	)
	w.SetMainMenu(mm)
	w.SetContent(
		widget.NewVBox(
			l,
			widget.NewButton("ok", nil),
		),
	)
	w.Resize(fyne.NewSize(300, 200))
	w.ShowAndRun()
}
