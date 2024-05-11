package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Cuttle")
	profiles := []string{"My Servers", "Work", "Others"}
	hello := widget.NewLabel("Hello Fyne!")

	pTitle := widget.NewLabel("Profiles: ")
	pSelect := widget.NewSelect(profiles,
		func(value string) {
			hello.SetText(value)
		})
	pBox := container.NewHBox(pTitle, pSelect)

	w.SetContent(container.NewVBox(
		pBox,
		hello,
		widget.NewButton("Hi!", func() {
			hello.SetText("Welcome :D")
		}),
	))

	w.ShowAndRun()
}
