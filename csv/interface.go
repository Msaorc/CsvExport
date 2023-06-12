package csv

type CsvInterface interface {
	NewCsvObject(path string, name string) (*CsvObject, error)
	AddColumns(colunas ...string) *Line
	AddLine(line Line) error
	AddLines(lines []Line) error
	AddHeader(header Line) error
	GenerateCsv() error
}
