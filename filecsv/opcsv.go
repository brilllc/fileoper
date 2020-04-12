package filecsv

import (
	"bytes"
	"encoding/csv"
	"io/ioutil"
	"os"
	"strings"
)

//WriteCsvNoFile .
func WriteCsvNoFile(content [][]string) (string, error) {
	// ctx := context.Background()
	b := &bytes.Buffer{}
	wr := csv.NewWriter(b)
	for _, cn := range content {
		wr.Write(cn)
	}
	wr.Flush()
	// ctx.Header("Content-Type", "text/csv")
	// ctx.Header("Content-Disposition", "attachment; filename="+filename)
	tet := b.String()
	// ctx.String(200, tet)
	return tet, nil
}

//WriteCsvFile .
func WriteCsvFile(content [][]string, fileName string) (int, error) {
	b := &bytes.Buffer{}
	wr := csv.NewWriter(b)
	for _, cn := range content {
		wr.Write(cn)
	}
	wr.Flush()

	fout, err := os.Create(fileName)
	defer fout.Close()
	if err != nil {
		return 0, err
	}
	return fout.WriteString(b.String())
}

//ReadFileAll .
func ReadFileAll(fileName string) ([][]string, error) {
	cntb, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	r2 := csv.NewReader(strings.NewReader(string(cntb)))
	ss, err := r2.ReadAll()
	// sz := len(ss)
	// for i := 0; i < sz; i++ {
	// 	fmt.Println(ss[i])
	// }
	return ss, err
}

//ReadFileRow .
func ReadFileRow(fr *csv.Reader, fileName string) ([]string, error) {
	ss, err := fr.Read()
	return ss, err
}

//OpenFile .
func OpenFile(fileName string) (*csv.Reader, error) {
	cntb, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	r2 := csv.NewReader(strings.NewReader(string(cntb)))
	return r2, nil
}
