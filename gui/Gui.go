package gui

import (
	"excel2csv/config"
	"excel2csv/excel"
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

func TestGui()  {
	app := app.New()
	window := app.NewWindow("ExcelTool")

	left := widget.NewVBox()
	right := widget.NewVBox()

	// 左边
	files := excel.GetAllExcelInDir(config.EXCEL_DIR)
	leftForm := widget.NewForm()
	left.Append(leftForm)
	for _, file := range files{
		leftForm.Append("", widget.NewLabel(file))
	}

	// 右边
	rightForm := widget.NewForm()
	rightForm.Append("", widget.NewLabel("22"))
	right.Append(rightForm)

	// 窗口 菜单
	mainMenu := fyne.NewMainMenu(fyne.NewMenu("file",
		fyne.NewMenuItem("general", func() {
		general()
	})))
	window.SetMainMenu(mainMenu)
	content := fyne.NewContainerWithLayout(layout.NewHBoxLayout(), left, right)
	window.SetContent(content)
	window.Resize(fyne.NewSize(480, 320))
	window.ShowAndRun()
}

func general()  {

}



