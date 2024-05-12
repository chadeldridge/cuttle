package main

import (
	"runtime"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

const (
	defaultWidth  = 400
	defaultHeight = 600
)

func main() {
	a := app.New()
	w := a.NewWindow("Cuttle")

	if runtime.GOOS != "android" && runtime.GOOS != "ios" {
		w.Resize(fyne.NewSize(defaultWidth, defaultHeight))
	}

	profiles := []string{"My Servers", "Work", "Others"}
	servers := []string{"myserver01", "monserv1", "monserv2"}
	hello := widget.NewLabel("Hello Fyne!")

	pTitle := widget.NewLabel("Profiles: ")
	pSelect := widget.NewSelect(profiles,
		func(value string) {
			hello.SetText(value)
		})
	pBox := container.NewHBox(pTitle, pSelect)

	sTitle := widget.NewLabel("Server/Group: ")
	sSelect := widget.NewSelect(servers,
		func(value string) {
			hello.SetText(value)
		})
	sBox := container.NewHBox(sTitle, sSelect)

	w.SetContent(container.NewVBox(
		pBox,
		sBox,
		hello,
		widget.NewButton("Hi!", func() {
			hello.SetText("Welcome :D")
		}),
	))

	w.ShowAndRun()
}
