package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func makeUI() (*widget.Label, *widget.Entry) {
	out := widget.NewLabel("Hello, World!")
	in := widget.NewEntry()

	in.OnChanged = func(content string) {
		out.SetText("Hello " + content + "!")
	}

	return out, in
}

func main() {
	a := app.New()
	w := a.NewWindow("Hello Person")

	w.SetContent(container.NewVBox(makeUI()))
	w.ShowAndRun()
}
