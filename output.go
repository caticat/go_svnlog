package main

import (
	"github.com/xuri/excelize/v2"
	"os"
	"fmt"
)

func OutputCSV(path string, sliSvnLog []*SvnLog) error {
	f, error := os.Create(path)
	if error != nil {
		return error
	}
	defer f.Close()

	// 表头
	ptrSvnLog := NewSvnLog()
	ptrSvnLog.comment = "comment"
	ptrSvnLog.version = "version"
	ptrSvnLog.author = "author"
	ptrSvnLog.datetime = "datetime"
	f.WriteString(ptrSvnLog.ToCsv())

	// 数据列表
	for _, ptrSvnLog := range sliSvnLog {
		f.WriteString(ptrSvnLog.ToCsv())
	}

	return nil
}

func OutputExcel(path string, sliSvnLog []*SvnLog) error {
	xlsx := excelize.NewFile()
	xlsx.SetCellValue("Sheet1", "A1", "abc1234")

	// 表头
	ptrSvnLog := NewSvnLog()
	ptrSvnLog.comment = "comment"
	ptrSvnLog.version = "version"
	ptrSvnLog.author = "author"
	ptrSvnLog.datetime = "datetime"
	xlsx.SetCellValue("Sheet1", "A1", ptrSvnLog.comment)
	xlsx.SetCellValue("Sheet1", "B1", ptrSvnLog.author)
	xlsx.SetCellValue("Sheet1", "C1", ptrSvnLog.datetime)
	xlsx.SetCellValue("Sheet1", "D1", ptrSvnLog.version)

	// 数据列表
	i := 1
	for _, ptrSvnLog := range sliSvnLog {
		i++
		xlsx.SetCellValue("Sheet1", fmt.Sprintf("A%v", i), ptrSvnLog.comment)
		xlsx.SetCellValue("Sheet1", fmt.Sprintf("B%v", i), ptrSvnLog.author)
		xlsx.SetCellValue("Sheet1", fmt.Sprintf("C%v", i), ptrSvnLog.datetime)
		xlsx.SetCellValue("Sheet1", fmt.Sprintf("D%v", i), ptrSvnLog.version)
	}

	if err := xlsx.SaveAs(path); err != nil {
		return err
	}

	return nil
}
