package fileexcel

import (
	"errors"
	"fmt"
	"strings"
)

//CellStore .
type CellStore interface {
	ReadRow(key string) (*Row, error)
	WriteRow(r *Row) error
	MoveRow(r *Row, newIndex int) error
	RemoveRow(key string) error
	Close() error
}

//CellStoreConstructor .
type CellStoreConstructor func() (CellStore, error)

//CellVisitorFunc .
type CellVisitorFunc func(c *Cell) error

// MemoryCellStore .
type MemoryCellStore struct {
	rows map[string]*Row
}

// UseMemoryCellStore .
func UseMemoryCellStore(f *File) {
	f.cellStoreConstructor = NewMemoryCellStore
}

// NewMemoryCellStore .
func NewMemoryCellStore() (CellStore, error) {
	cs := &MemoryCellStore{
		rows: make(map[string]*Row),
	}
	return cs, nil
}

// Close .
func (mcs *MemoryCellStore) Close() error {
	return nil
}

// ReadRow .
func (mcs *MemoryCellStore) ReadRow(key string) (*Row, error) {
	r, ok := mcs.rows[key]
	if !ok {
		return nil, errors.New("No such row")
	}
	return r, nil
}

// WriteRow .
func (mcs *MemoryCellStore) WriteRow(r *Row) error {
	if r != nil {
		mcs.rows[r.key()] = r
	}
	return nil
}

// MoveRow .
func (mcs *MemoryCellStore) MoveRow(r *Row, index int) error {
	oldKey := r.key()
	r.num = index
	newKey := r.key()
	if _, exists := mcs.rows[newKey]; exists {
		return fmt.Errorf("Target index for row (%d) would overwrite a row already exists", index)
	}
	delete(mcs.rows, oldKey)
	mcs.rows[newKey] = r
	return nil
}

// RemoveRow .
func (mcs *MemoryCellStore) RemoveRow(key string) error {
	delete(mcs.rows, key)
	return nil
}

// Extract
func keyToRowKey(key string) string {
	parts := strings.Split(key, ":")
	return parts[0] + ":" + parts[1]
}
