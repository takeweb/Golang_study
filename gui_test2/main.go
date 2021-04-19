package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

type MyEntry struct {
	widget.Entry
	entered func(e *MyEntry)
}

func NewMyEntry(f func(e *MyEntry)) *MyEntry {
	e := &MyEntry{}
	e.ExtendBaseWidget(e)
	e.entered = f
	return e
}

func (e *MyEntry) KeyDown(key *fyne.KeyEvent) {
	switch key.Name {
	case fyne.KeyReturn, fyne.KeyEnter:
		e.entered(e)
	default:
		e.Entry.KeyDown(key)
	}
}

func main() {
	a := app.New()
	w := a.NewWindow("Hello")
	l := widget.NewLabel("Hello Fyne!")
	e := NewMyEntry(func(e *MyEntry) {
		s := e.Text
		e.SetText("")
		l.SetText("you type '" + s + "'.")
	})

	w.SetContent(
		widget.NewVBox(
			l, e,
		),
	)
	a.Settings().SetTheme(theme.LightTheme())
	w.Resize(fyne.NewSize(300, 100))
	w.ShowAndRun()
}

// func total(n int) int {
// 	t := 0
// 	for i := 1; i <= n; i++ {
// 		t += i
// 	}
// 	return t
// }

// func total(n int) int {
// 	t := 0
// 	for i := 1; i <= n; i++ {
// 		t += i
// 	}
// 	return t
// }
