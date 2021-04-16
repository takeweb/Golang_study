package main

import (
	"strconv"

	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello")
	l := widget.NewLabel("Hello Fyne!")
	e := widget.NewEntry()
	e.SetText("0")
	w.SetContent(
		widget.NewVBox(
			l, e,
			widget.NewButton("Click me!", func() {
				n, _ := strconv.Atoi(e.Text)
				l.SetText("count:" + strconv.Itoa(total(n)))
			}),
		),
	)
	w.ShowAndRun()
}

func total(n int) int {
	t := 0
	for i := 1; i <= n; i++ {
		t += i
	}
	return t
}
