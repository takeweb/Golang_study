package main

import (
	"strconv"

	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func main() {
	c := 0
	a := app.New()
	w := a.NewWindow("Hello")
	l := widget.NewLabel("Hello Fyne!")
	w.SetContent(
		widget.NewVBox(
			l,
			widget.NewButton("Click me!", func() {
				c++
				l.SetText("count:" + strconv.Itoa(c))
			}),
		),
	)
	w.ShowAndRun()
}
