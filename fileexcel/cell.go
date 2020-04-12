package fileexcel

//Cell .
type Cell struct {
	Row   *Row
	Value string
	// RichText       []RichTextRun
	formula string
	// style          *Style
	NumFmt string
	// parsedNumFmt   *parsedNumberFormat
	date1904 bool
	Hidden   bool
	HMerge   int
	VMerge   int
	// cellType       CellType
	// DataValidation *xlsxDataValidation
	// Hyperlink      Hyperlink
	num int
}

// NewCell .
func newCell(r *Row, num int) *Cell {
	cell := &Cell{Row: r, num: num}
	return cell
}
