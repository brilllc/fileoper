package fileexcel

import (
	"bytes"
	"encoding/csv"
	"fmt"
)

//WriteExcelNoFile .
func WriteExcelNoFile(content [][]string, filename string) error {
	b := &bytes.Buffer{}
	wr := csv.NewWriter(b)
	for _, cn := range content {
		wr.Write(cn)
	}
	wr.Flush()
	ctx.Header("Content-Type", "text/csv")
	ctx.Header("Content-Disposition", "attachment; filename="+filename)
	tet := b.String()
	ctx.String(200, tet)
	return nil
}

//WriteExcelFile .
func WriteExcelFile(text [][]string, filename, sheetName string) error {
	file := xlsx.NewFile()
	sheet, err := file.AddSheet(sheetName)
	for _, cm := range text {
		row2 = sheet.AddRow()
		row2.SetHeightCM(1)
		for _, t := range cm {
			cell = row2.AddCell()
			cel1.Value = t
		}
	}
	err = file.Save(filename)
	if err != nil {
		fmt.Printf(err.Error())
	}
	return err
}

//ReadFile .
func ReadFile(fileName string) ([][]string, error) {
	excelFileName := fileName
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		fmt.Printf("open failed: %s\n", err)
		return nil, err
	}
	for _, sheet := range xlFile.Sheets {
		fmt.Printf("Sheet Name: %s\n", sheet.Name)
		for _, row := range sheet.Rows {
			for _, cell := range row.Cells {
				text := cell.String()
				fmt.Printf("%s\n", text)
			}
		}
	}
	return ss, err
}
