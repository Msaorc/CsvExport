package csv

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type CsvTestSuite struct {
	suite.Suite
	Line1    *Line
	Line2    *Line
	Header   *Line
	Path     string
	FileName string
}

func (suite *CsvTestSuite) SetupTest() {
	suite.Line1 = GenerateLine("Marcos", "marcos@marcos.com")
	suite.Line2 = GenerateLine("Duda", "duda@marcos.com")
	suite.Header = GenerateLine("Nome", "Email")
	suite.Path = "/home/marcao/projetos/ExportacaoExcel/"
	suite.FileName = "CsvTest"
}

func (suite *CsvTestSuite) TestNewCSVObjetct() {
	csvObject, err := NewCsvObject(suite.Path, suite.FileName)
	assert.NotNil(suite.T(), err)
	assert.NotNil(suite.T(), csvObject)
	assert.Equal(suite.T(), suite.Path, csvObject.Path)
	assert.Equal(suite.T(), suite.FileName, csvObject.Name)
}

func (suite *CsvTestSuite) TestAddLine() {
	csvObject, err := NewCsvObject(suite.Path, suite.FileName)
	assert.Nil(suite.T(), err)
	err = csvObject.AddLine(*suite.Line1)
	assert.Nil(suite.T(), err)
	assert.Len(suite.T(), csvObject.File.Lines, 1)
	err = csvObject.AddLine(*suite.Line2)
	assert.Nil(suite.T(), err)
	assert.Len(suite.T(), csvObject.File.Lines, 2)
}

func (suite *CsvTestSuite) TestGenereteCsv() {
	csvObject, err := NewCsvObject(suite.Path, suite.FileName)
	assert.Nil(suite.T(), err)
	err = csvObject.AddLine(*suite.Line1)
	assert.Nil(suite.T(), err)
	err = csvObject.AddLine(*suite.Line2)
	assert.Nil(suite.T(), err)
	err = csvObject.GenerateCsv()
	suite.Nil(err)
}

func (suite *CsvTestSuite) TestValueInCsv() {
	csvObject, err := NewCsvObject(suite.Path, suite.FileName)
	assert.Nil(suite.T(), err)
	err = csvObject.AddLine(*suite.Line1)
	assert.Nil(suite.T(), err)
	err = csvObject.AddLine(*suite.Line2)
	assert.Nil(suite.T(), err)
	err = csvObject.AddHeader(*suite.Header)
	assert.Nil(suite.T(), err)
	err = csvObject.GenerateCsv()
	assert.Nil(suite.T(), err)

	fullPath := suite.Path + suite.FileName + ".csv"
	file, err := ReadCsv(fullPath)
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), file)
	assert.Equal(suite.T(), "Nome", file[0].Name)
	assert.Equal(suite.T(), "Email", file[0].Email)
	assert.Equal(suite.T(), "Marcos", file[1].Name)
	assert.Equal(suite.T(), "marcos@marcos.com", file[1].Email)
	assert.Equal(suite.T(), "Duda", file[2].Name)
	assert.Equal(suite.T(), "duda@marcos.com", file[2].Email)
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(CsvTestSuite))
}
