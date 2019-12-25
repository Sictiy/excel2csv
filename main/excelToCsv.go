package main

import (
	"excel2csv/config"
	"excel2csv/excel"
	"excel2csv/general"
	"excel2csv/gui"
	"excel2csv/log"
	"fmt"
)

func main()  {
	gui.TestGui()
	//testProcess()
}

func testProcess()  {
	var fileName string
	files := excel.GetAllExcelInDir(config.EXCEL_DIR)
	log.Log(files)
	_, err := fmt.Scanln(&fileName)
	if log.CheckError(err) {
		return
	}
	log.Log(fileName)
	processFile(fileName)
	log.Log("success")
}

func processFile(fileName string)  {
	table := excel.ExcelToTable(fileName)
	if table == nil {
		log.Log("to table failed!")
		return
	}
	log.Log(table.ToString())
	general.GeneralToJavaBean(table)
	general.GeneralToCsv(table)
}

