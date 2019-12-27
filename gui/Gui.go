package gui

import (
	"excel2csv/config"
	"excel2csv/entry"
	"excel2csv/excel"
	"excel2csv/general"
	"excel2csv/log"
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

type ToolGui struct {
	// gui组件
	app fyne.App
	window fyne.Window
	allCheck []*widget.Check
	form *widget.Form
	labels []*widget.Label
	// 缓存表数据
	tables map[string]*entry.Table
}

func (gui *ToolGui)Run() bool {
	gui.app = app.New()
	gui.window = gui.app.NewWindow("ExcelTool")
	// 左边
	left := widget.NewVBox()
	left.Append(widget.NewButton("general", func() {
		gui.general()
	}))
	files := excel.GetAllExcelInDir(config.EXCEL_DIR)
	for _, file := range files{
		check := widget.NewCheck(file, func(b bool) {
			gui.freshForm()
		})
		gui.allCheck = append(gui.allCheck, check)
		left.Append(check)
	}

	// 右边
	right := widget.NewVBox()
	gui.form = widget.NewForm()
	right.Append(gui.form)

	content := fyne.NewContainerWithLayout(layout.NewHBoxLayout(), left, right)
	gui.window.SetContent(content)
	gui.window.Resize(fyne.NewSize(480, 320))
	gui.window.ShowAndRun()
	return true
}

func (gui *ToolGui)general()  {
	var selectFiles []string
	for _, check := range gui.allCheck{
		if check.Checked {
			selectFiles = append(selectFiles, check.Text)
		}
	}
	if len(selectFiles) <= 0 {
		return
	}
	gui.GeneralFiles(selectFiles)
}

func (gui *ToolGui)freshForm() {
	for _, check := range gui.allCheck{
		if check.Focused(){
			table := gui.getTable(check.Text)
			if table == nil{
				log.Log("to table failed! table is nil")
			}
			gui.clearForm()
			for i, column := range table.Columns{
				if len(gui.form.Items) <= i {
					label := widget.NewLabel(column.JavaType)
					gui.form.Append(column.Name, label)
					gui.labels = append(gui.labels, label)
				} else {
					gui.form.Items[i].Text = column.Name
					gui.labels[i].Text = column.JavaType
					gui.form.Refresh()
				}
			}
			return
		}
	}
}

func (gui *ToolGui)clearForm() {
	for _, item := range gui.form.Items{
		item.Text = ""
	}
	for _, label := range gui.labels{
		label.Text = ""
		label.Refresh()
	}
	gui.form.Refresh()
}

func (gui *ToolGui)getTable(fileName string) *entry.Table {
	if gui.tables == nil {
		gui.tables = make(map[string]*entry.Table)
	}
	if gui.tables[fileName] == nil {
		gui.tables[fileName] = excel.ExcelToTable(fileName)
	}
	return gui.tables[fileName]
}

/**处理已选择文件***********************/
func (gui *ToolGui)GeneralFiles(files []string)  {
	for _, file := range files{
		gui.processFile(file)
	}
}

func (gui *ToolGui)processFile(fileName string)  {
	table := gui.getTable(fileName)
	if table == nil {
		log.Log("to table failed! table is nil")
		return
	}
	log.Log(table.ToString())
	general.GeneralToJavaBean(table)
	general.GeneralToCsv(table)
}



