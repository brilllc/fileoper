package fileexcel

// Sheet is a high level structure intended to provide user access to
// the contents of a particular sheet within an XLSX file.
type Sheet struct {
	Name string
	File *File
	// Cols        *ColStore
	MaxRow      int
	MaxCol      int
	Hidden      bool
	Selected    bool
	SheetViews  []SheetView
	SheetFormat SheetFormat
	AutoFilter  *AutoFilter
	Relations   []Relation
	// DataValidations  []*xlsxDataValidation
	cellStore        CellStore
	streamedRowCount int
	currentRow       *Row
}

//SheetView .
type SheetView struct {
	Pane *Pane
}

//Pane .
type Pane struct {
	XSplit      float64
	YSplit      float64
	TopLeftCell string
	ActivePane  string
	State       string
}

//SheetFormat .
type SheetFormat struct {
	DefaultColWidth  float64
	DefaultRowHeight float64
	OutlineLevelCol  uint8
	OutlineLevelRow  uint8
}

//AutoFilter .
type AutoFilter struct {
	TopLeftCell     string
	BottomRightCell string
}

//Relation .
type Relation struct {
	// Type       RelationshipType
	Target string
	// TargetMode RelationshipTargetMode
}

// AddRow .
func (s *Sheet) AddRow() *Row {
	if s.currentRow != nil {
		s.cellStore.WriteRow(s.currentRow)
	}
	row := &Row{Sheet: s, num: s.MaxRow}
	s.setCurrentRow(row)
	s.MaxRow++
	return row
}

func (s *Sheet) setCurrentRow(r *Row) {
	if r == nil {
		return
	}
	if r.num > s.MaxRow {
		s.MaxRow = r.num + 1
	}
	if s.currentRow != nil {
		err := s.cellStore.WriteRow(s.currentRow)
		if err != nil {
			panic(err)
		}
	}
	s.currentRow = r
}
