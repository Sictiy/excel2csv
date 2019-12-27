package excel

import (
	"excel2csv/config"
	"excel2csv/entry"
	"excel2csv/log"
	"fmt"
	"github.com/xuri/excelize"
	"io/ioutil"
	"path"
	"strings"
)

func GetAllExcelInDir(filePath string) []string {
	var files []string
	allFile, err := ioutil.ReadDir(filePath)
	if log.CheckError(err) {
		return files
	}
	for _, file := range allFile{
		if file.IsDir() {
			files = append(files, GetAllExcelInDir(filePath + file.Name() + "\\")...)
		} else {
			if path.Ext(file.Name()) == ".xlsx" {
				files = append(files, strings.TrimSuffix(file.Name(), ".xlsx"))
			}
		}
	}
	return files
}

func ExcelToTable(fileName string) *entry.Table {
	xlsx, err := excelize.OpenFile(fileName + ".xlsx")
	if log.CheckError(err) {
		return nil
	}
	sheets := xlsx.GetSheetMap()
	if len(sheets) <= 0 {
		return nil
	}
	rows, err := xlsx.GetRows(sheets[1])
	if log.CheckError(err) {
		return nil
	}
	return RowsToTable(rows, fileName)
}

func RowsToTable(rows [][]string, fileName string) *entry.Table {
	if len(rows) < 3 {
		return nil
	}
	table := &entry.Table{}
	table.Name = getFileNameByFilePath(fileName)
	table.Comment = fileName
	rowNum := len(rows[0])
	columns := make([]*entry.Column, rowNum)
	for index, row := range rows {
		if index < 3 {
			if columns[0] != nil {
				continue
			}
			for j, cell := range row {
				column := &entry.Column{}
				column.Name = cell
				column.Comment = rows[index+1][j]
				column.JavaType = rows[index+2][j]
				columns[j] = column
			}
			table.ColumnData = rows[0:3]
			continue
		}
		table.AddRow(row)
	}
	table.Columns = columns
	table.Package = config.JAVA_BEAN_PACKAGE
	table.JavaName = Capitalize(table.Name)
	table.Imports = GetImportFromColumn(table)
	return table
}

func GetImportFromColumn(table *entry.Table) map[string]string {
	imports := make(map[string]string)
	for _, column := range table.Columns {
		if column.JavaType == "Date"{
			imports[column.JavaType] = "java.util.Date"
		}
		// 后续有其他的需要再加上
	}
	return imports
}

func getFileNameByFilePath(path string) string {
	paths := strings.Split(path, "/")
	splitString := strings.Split(paths[len(paths) - 1], ".")
	return splitString[0]
}

// Capitalize 字符首字母大写
func Capitalize(str string) string {
	var upperStr string
	vv := []rune(str)   // 后文有介绍
	for i := 0; i < len(vv); i++ {
		if i == 0 {
			if vv[i] >= 97 && vv[i] <= 122 {  // 后文有介绍
				vv[i] -= 32 // string的码表相差32位
				upperStr += string(vv[i])
			} else {
				fmt.Println("Not begins with lowercase letter,")
				return str
			}
		} else {
			upperStr += string(vv[i])
		}
	}
	return upperStr
}