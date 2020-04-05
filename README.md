# fileoper

多类型文件处理

-- csv excel

-- 生成 下载 修改

生成函数：

func NewWriter(w io.Writer) *Writer

func (w *Writer) Flush()

func (w *Writer) Write(record []string) (err os.Error)

func (w *Writer) WriteAll(records [][]string) (err os.Error)

func (w *Writer)WriteCsvFile(records [][]string) (os.Error)

func (w *Writer)WriteExcelFile(records [][]string) (os.Error)


读取函数：

func NewReader(r io.Reader) *Reader

func (r *Reader) Read() (record []string, err os.Error)

func (r *Reader) ReadAll() (records [][]string, err os.Error)

func (r *Reader) ReadFile(fileName string) ([][]string, os.Error)