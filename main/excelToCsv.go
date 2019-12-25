package main

import (
	"fmt"
	"github.com/xuri/excelize"
	"fyne.io/fyne/widget"
	"fyne.io/fyne/app"
)

func main()  {
	testGui()
	processCsv()
}

func testGui()  {
	app := app.New()

	w := app.NewWindow("Hello")
	w.SetContent(widget.NewVBox(
		widget.NewLabel("Hello Fyne!"),
		widget.NewButton("Quit", func() {
			app.Quit()
		}),
	))

	w.ShowAndRun()
}

func processCsv()  {
	var fileName string
	_, err := fmt.Scanln(&fileName)
	if CheckError(err) {
		return
	}
	Log(fileName)
	excel2csv(fileName)
}

func excel2csv(fileName string) {
	xlsx, err := excelize.OpenFile(fileName)
	if CheckError(err) {
		return
	}
	sheets := xlsx.GetSheetMap()
	if len(sheets) <= 0 {
		return
	}
	rows, err := xlsx.GetRows(sheets[1])
	if CheckError(err) {
		return	
	}
	for _, row := range rows {
		var rowSting string
		for _, cell := range row {
			rowSting = rowSting + cell + ","
		}
		Log(rowSting)
	}
}
