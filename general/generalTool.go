package general

import (
	"encoding/csv"
	"excel2csv/config"
	"excel2csv/entry"
	"excel2csv/log"
	"os"
	"strings"
	"text/template"
	"time"
)

func GeneralToJavaBean(table *entry.Table)  {
	//tmpl, err := template.New("test").Parse("{{.Name}} of {{.Comment}}")
	funcs := template.FuncMap{"getDate": GetDate}
	tmpl, err := template.New("bean.tmpl").Funcs(funcs).ParseFiles("bean.tmpl")
	if log.CheckError(err) {
		return
	}
	tmpl.Funcs(funcs)
	newFileName := strings.Join([]string {table.Name, "java"}, ".")
	filePath := config.JAVA_BEAN_DIR + newFileName
	iWriter, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if log.CheckError(err) {
		return
	}
	defer iWriter.Close()
	err = tmpl.Execute(iWriter, table)
	if log.CheckError(err) {
		return
	}
}

func GeneralToCsv(table *entry.Table)  {
	newFileName := strings.Join([]string {table.Name, "csv"}, ".")
	filePath := config.CSV_DIR + newFileName
	iWriter, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if log.CheckError(err) {
		return
	}
	defer iWriter.Close()
	//iWriter.Seek(0, io.SeekEnd)
	_, _ = iWriter.WriteString("\xEF\xBB\xBF") // 写入UTF-8 BOM，防止中文乱码
	csvWriter := csv.NewWriter(iWriter)
	csvWriter.Comma = ','
	csvWriter.UseCRLF = true
	allRows := append(table.ColumnData, table.Data...)
	err = csvWriter.WriteAll(allRows)
	if log.CheckError(err) {
		return
	}
}

func GetDate() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
