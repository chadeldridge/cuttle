package main

import (
	"image/color"
	"log"
	"os"
	"runtime"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/chadeldridge/cuttle/connections"
	"github.com/chadeldridge/cuttle/helpers"
)

const (
	defaultWidth   = 600
	defaultHeight  = 400
	maxResultLines = 50
	maxLogLines    = 50
)

var (
	colorGridBackground = color.RGBA{R: 31, G: 31, B: 31, A: 255}
	gridBackground      = canvas.NewRectangle(colorGridBackground)
)

func getPassword() string {
	if p, ok := os.LookupEnv("PASSWORD"); ok {
		return p
	}

	log.Fatal("failed to get env PASSWORD")
	return ""
}

func main() {
	a := app.New()
	w := a.NewWindow("Cuttle")

	if runtime.GOOS != "android" && runtime.GOOS != "ios" {
		w.Resize(fyne.NewSize(defaultWidth, defaultHeight))
	}

	profiles := []string{"test", "Work", "Others"}
	servers := []string{"192.168.50.105", "monserv1", "monserv2"}
	hello := widget.NewLabel("Welcome to Cuttle!")

	conn, err := connections.NewSSH(uint16(22), map[string]interface{}{
		"ip":       servers[0],
		"username": "debian",
		"password": getPassword(),
	})
	if err != nil {
		log.Fatal(err)
	}

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

	cmdPane := container.NewVBox(
		pBox,
		sBox,
		hello,
		widget.NewButton("Test Connection", func() { testConnection(conn) }),
	)

	resultsGrid := widget.NewTextGrid()
	results := helpers.NewQueue(maxResultLines, resultsGrid)
	conn.Results = results
	resultBox := container.NewStack(gridBackground, container.NewVScroll(resultsGrid))

	logsGrid := widget.NewTextGrid()
	logs := helpers.NewQueue(maxLogLines, logsGrid)
	conn.Logs = logs
	logBox := container.NewStack(gridBackground, container.NewVScroll(logsGrid))
	/*
		conn.Results = make([]string, 0)
		resultList := container.NewVBox(widget.NewList(
			func() int { return len(results) },
			func() fyne.CanvasObject { return widget.NewLabel("label") },
			func(i widget.ListItemID, o fyne.CanvasObject) { o.(*widget.Label).SetText(results[i]) },
		))
	*/

	/*
		w.SetContent(container.NewBorder(
			cmdPane,
			// logsBox,
			nil,
			nil,
			resultList,
		))
	*/

	w.SetContent(container.NewVBox(
		container.NewHBox(
			cmdPane,
			resultBox,
		),
		logBox,
	))

	w.ShowAndRun()
}

func testConnection(conn connections.Handler) {
	err := conn.TestConnection()
	if err != nil {
		log.Println(err)
	}
}
