package filecsv

import (
	"bytes"
	"context"
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

//WriteCsvNoFile .
func WriteCsvNoFile(content [][]string, filename string) error {
	ctx := context.Background()
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

//WriteCsvFile .
func WriteCsvFile(content [][]string, fileName string) (int, error) {
	ctx := context.Background()
	b := &bytes.Buffer{}
	wr := csv.NewWriter(b)
	for _, cn := range content {
		wr.Write(cn)
	}
	wr.Flush()

	fout, err := os.Create(fileName)
	defer fout.Close()
	if err != nil {
		return
	}
	return fout.WriteString(b.String())
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
