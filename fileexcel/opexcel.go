package fileexcel

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"unicode/utf8"
)

//File .
type File struct {
	worksheets    map[string]*zip.File
	worksheetRels map[string]*zip.File
	// referenceTable       *RefTable
	Date1904 bool
	// styles               *xlsxStyleSheet
	Sheets []*Sheet
	Sheet  map[string]*Sheet
	// theme                *theme
	// DefinedNames         []*xlsxDefinedName
	cellStoreConstructor CellStoreConstructor
	rowLimit             int
}

//NoRowLimit .
const NoRowLimit int = -1

//FileOption .
type FileOption func(f *File)

// RowLimit .
func RowLimit(n int) FileOption {
	return func(f *File) {
		f.rowLimit = n
	}
}

// NewFile .
func NewFile(options ...FileOption) *File {
	f := &File{
		Sheet:  make(map[string]*Sheet),
		Sheets: make([]*Sheet, 0),
		// DefinedNames:         make([]*xlsxDefinedName, 0),
		rowLimit:             NoRowLimit,
		cellStoreConstructor: NewMemoryCellStore,
	}
	for _, opt := range options {
		opt(f)
	}
	return f
}

// OpenFile .
func OpenFile(fileName string, options ...FileOption) (file *File, err error) {
	var z *zip.ReadCloser
	z, err = zip.OpenReader(fileName)
	if err != nil {
		return nil, err
	}
	defer z.Close()
	return NewFile(options...), nil
}

// Save .
func (f *File) Save(path string) (err error) {
	target, err := os.Create(path)
	if err != nil {
		return err
	}
	err = f.Write(target)
	if err != nil {
		return err
	}
	return target.Close()
}

// Write .
func (f *File) Write(writer io.Writer) (err error) {
	parts := make(map[string]string, 0)
	zipWriter := zip.NewWriter(writer)
	for partName, part := range parts {
		w, err := zipWriter.Create(partName)
		if err != nil {
			return err
		}
		_, err = w.Write([]byte(part))
		if err != nil {
			return err
		}
	}
	return zipWriter.Close()
}

// AddSheet .
func (f *File) AddSheet(sheetName string) (*Sheet, error) {
	return f.AddSheetWithCellStore(sheetName, NewMemoryCellStore)
}

//AddSheetWithCellStore .
func (f *File) AddSheetWithCellStore(sheetName string, constructor CellStoreConstructor) (*Sheet, error) {
	var err error
	if _, exists := f.Sheet[sheetName]; exists {
		return nil, fmt.Errorf("duplicate sheet name '%s'.", sheetName)
	}
	runeLength := utf8.RuneCountInString(sheetName)
	if runeLength > 31 || runeLength == 0 {
		return nil, fmt.Errorf("sheet name must be 31 or fewer characters long.  It is currently '%d' characters long", runeLength)
	}
	// Iterate over the runes
	for _, r := range sheetName {
		// Excel forbids : \ / ? * [ ]
		if r == ':' || r == '\\' || r == '/' || r == '?' || r == '*' || r == '[' || r == ']' {
			return nil, fmt.Errorf("sheet name must not contain any restricted characters : \\ / ? * [ ] but contains '%s'", string(r))
		}
	}
	sheet := &Sheet{
		Name:     sheetName,
		File:     f,
		Selected: len(f.Sheets) == 0,
		// Cols:     &ColStore{},
	}

	// sheet.cellStore, err = constructor()
	if err != nil {
		return nil, err
	}
	f.Sheet[sheetName] = sheet
	f.Sheets = append(f.Sheets, sheet)
	return sheet, nil
}
