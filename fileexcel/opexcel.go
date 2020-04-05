package fileexcel

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"strings"
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
func WriteExcelFile(content [][]string, filename string) error {
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

//ReadFile .
func ReadFile(fileName string) ([][]string, error) {
	cntb, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err.String())
	}
	r2 := csv.NewReader(strings.NewReader(string(cntb)))
	ss, err := r2.ReadAll()
	sz := len(ss)
	for i := 0; i < sz; i++ {
		fmt.Println(ss[i])
	}
	return ss, err
}
