package filecsv

import (
	"bytes"
	"io"
	"log"
)

func (f *File) GetRows(sheet string) ([][]string, error) {
	rows, err := f.Rows(sheet)
	if err != nil {
		return nil, err
	}
	results := make([][]string, 0, 64)
	for rows.Next() {
		if rows.Error() != nil {
			break
		}
		row, err := rows.Columns()
		if err != nil {
			break
		}
		results = append(results, row)
	}
	return results, nil
}

func (f *File) sharedStringsReader() *xlsxSST {
	var err error

	if f.SharedStrings == nil {
		var sharedStrings xlsxSST
		ss := f.readXML("xl/sharedStrings.xml")
		if len(ss) == 0 {
			ss = f.readXML("xl/SharedStrings.xml")
		}
		if err = f.xmlNewDecoder(bytes.NewReader(namespaceStrictToTransitional(ss))).
			Decode(&sharedStrings); err != nil && err != io.EOF {
			log.Printf("xml decode error: %s", err)
		}
		f.SharedStrings = &sharedStrings
	}

	return f.SharedStrings
}

func (f *File) RemoveRow(sheet string, row int) error {
	if row < 1 {
		return newInvalidRowNumberError(row)
	}

	xlsx, err := f.workSheetReader(sheet)
	if err != nil {
		return err
	}
	if row > len(xlsx.SheetData.Row) {
		return f.adjustHelper(sheet, rows, row, -1)
	}
	keep := 0
	for rowIdx := 0; rowIdx < len(xlsx.SheetData.Row); rowIdx++ {
		v := &xlsx.SheetData.Row[rowIdx]
		if v.R != row {
			xlsx.SheetData.Row[keep] = *v
			keep++
		}
	}
	xlsx.SheetData.Row = xlsx.SheetData.Row[:keep]
	return f.adjustHelper(sheet, rows, row, -1)
}
