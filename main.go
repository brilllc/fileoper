package main

import (
	"brill/fileoper/filecsv"
	"brill/fileoper/fileexcel"
	"fmt"
)

func main() {
	var fileName = "./fc.csv"
	strs, err := filecsv.WriteCsvNoFile([][]string{{"12", "13"}, {"aa", "bb"}})
	fmt.Println(strs, err)
	n, err := filecsv.WriteCsvFile([][]string{{"12", "13"}, {"aa", "bb"}}, fileName)
	fmt.Println(n, err)
	strs1, err := filecsv.ReadFileAll(fileName)
	fmt.Println(strs1, err)

	nf, err := filecsv.OpenFile(fileName)
	line, err := nf.Read()
	fmt.Println(line, err)
	line, err = nf.Read()
	fmt.Println(line, err)
	line, err = nf.Read()
	fmt.Println(line, err)

	ex := fileexcel.NewFile()
	sh, err := ex.AddSheet("she1")
	fmt.Println(sh, err)
	row := sh.AddRow()
	ce := row.AddCell()
	ce.Value = "abc"
	ce1 := row.AddCell()
	ce1.Value = "abcd"

	err = ex.Save("./fc.xls")
	fmt.Println("line", err)

	fo, err := fileexcel.OpenFile("./fc.xls")
	fmt.Println(fo, err)
	for i, k := range fo.Sheets {
		fmt.Println(i, k.Relations)
	}

}
