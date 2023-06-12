package csv

import (
	"encoding/csv"
	"errors"
	"os"
	"regexp"
)

var errPathInvalid = errors.New("path is invalid")
var errPahtIsRequired = errors.New("paht is required")
var errNameIsRequired = errors.New("name is required")
var errLineIsEmpty = errors.New("line is empty")

type Line struct {
	Columns []string
}

type File struct {
	Lines []Line
}

type CsvObject struct {
	Path   string
	Name   string
	Header Line
	File   File
}

func NewCsvObject(path string, name string) (*CsvObject, error) {
	csv := &CsvObject{
		Path: path,
		Name: removeCaracterInvalid(name),
	}
	if err := csv.validate(); err != nil {
		return nil, err
	}
	return csv, nil
}

func GenerateLine(columns ...string) *Line {
	var line Line
	line.Columns = append(line.Columns, columns...)
	return &line
}

func (o *CsvObject) AddLine(line Line) error {
	if line.Columns == nil {
		return errLineIsEmpty
	}
	o.File.Lines = append(o.File.Lines, line)
	return nil
}

func (o *CsvObject) AddLines(line []Line) error {
	if line == nil {
		return errLineIsEmpty
	}
	o.File.Lines = append(o.File.Lines, line...)
	return nil
}

func (o *CsvObject) AddHeader(header Line) error {
	if header.Columns == nil {
		return errLineIsEmpty
	}
	o.Header = header
	return nil
}

func (o *CsvObject) GenerateCsv() error {
	file, err := createFile(o.Name, o.Path)
	if err != nil {
		panic(err)
	}
	writer := csv.NewWriter(file)
	writer.Comma = ';'
	defer file.Close()
	defer writer.Flush()
	if err = addHeaderInFile(writer, o.Header); err != nil {
		return err
	}
	if err = addBodyInFile(writer, o.File.Lines); err != nil {
		return err
	}
	return nil
}

func (o *CsvObject) validate() error {
	if o.Path[len(o.Path)-1:] != "/" {
		return errPathInvalid
	}
	if o.Path == "" {
		return errPahtIsRequired
	}
	if o.Name == "" {
		return errNameIsRequired
	}
	return nil
}

func removeCaracterInvalid(name string) string {
	var nameValid string
	r, _ := regexp.Compile("[a-zA-Z]")
	nameSlice := r.FindAllString(name, -1)
	for _, char := range nameSlice {
		nameValid += char
	}
	return nameValid
}

func createFile(name string, path string) (*os.File, error) {
	return os.Create(path + "/" + name + ".csv")
}

func addHeaderInFile(file *csv.Writer, header Line) error {
	return file.Write(header.Columns)
}

func addBodyInFile(file *csv.Writer, lines []Line) error {
	for _, line := range lines {
		if err := file.Write(line.Columns); err != nil {
			return err
		}
	}
	return nil
}
