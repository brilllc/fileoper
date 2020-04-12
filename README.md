# fileoper

多类型文件处理

-- csv excel

-- 生成 下载 修改

csv函数：

//不写入文件

strs, err := filecsv.WriteCsvNoFile([][]string{{"12", "13"}, {"aa", "bb"}})

//写入文件

n, err := filecsv.WriteCsvFile([][]string{{"12", "13"}, {"aa", "bb"}}, fileName)

//读取文件全部内容

strs1, err := filecsv.ReadFileAll(fileName)

//打开文件

nf, err := filecsv.OpenFile(fileName)
	
//读取一行内容
line, err := nf.Read()


excel函数：

//创建文件

ex := fileexcel.NewFile()

//设置sheet名

sh, err := ex.AddSheet("she1")

//写入一行内容

row := sh.AddRow()

ce := row.AddCell()

ce.Value = "abc"

ce1 := row.AddCell()

ce1.Value = "abcd"

//把内容写入文件

err = ex.Save("./fc.xls")

//读取文件

fo, err := fileexcel.OpenFile("./fc.xls")
for i, k := range fo.Sheets {
	fmt.Println(i, k.Relations)
}