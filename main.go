package main

import (
	"brill/fileoper/filecsv"
	"brill/fileoper/fileexcel"
)

func main() {
	filecsv.WriteCsvFile([][]string{})
	filecsv.WriteCsvNoFile([][]string{})
	fileexcel.WriteExcelFile([][]string{})
	fileexcel.WriteExcelFile([][]string{})

}
