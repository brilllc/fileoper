package fileexcel

import (
	"fmt"
	"time"
)

func (f *File) GetCellValue(sheet, axis string) (string, error) {
	return f.getCellStringFunc(sheet, axis, func(x *xlsxWorksheet, c *xlsxC) (string, bool, error) {
		val, err := c.getValueFrom(f, f.sharedStringsReader())
		if err != nil {
			return val, false, err
		}
		return val, true, err
	})
}

func (f *File) SetCellValue(sheet, axis string, value interface{}) error {
	var err error
	switch v := value.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		err = f.setCellIntFunc(sheet, axis, v)
	case float32:
		err = f.SetCellFloat(sheet, axis, float64(v), -1, 32)
	case float64:
		err = f.SetCellFloat(sheet, axis, v, -1, 64)
	case string:
		err = f.SetCellStr(sheet, axis, v)
	case []byte:
		err = f.SetCellStr(sheet, axis, string(v))
	case time.Duration:
		_, d := setCellDuration(v)
		err = f.SetCellDefault(sheet, axis, d)
		if err != nil {
			return err
		}
		err = f.setDefaultTimeStyle(sheet, axis, 21)
	case time.Time:
		err = f.setCellTimeFunc(sheet, axis, v)
	case bool:
		err = f.SetCellBool(sheet, axis, v)
	case nil:
		err = f.SetCellStr(sheet, axis, "")
	default:
		err = f.SetCellStr(sheet, axis, fmt.Sprint(value))
	}
	return err
}