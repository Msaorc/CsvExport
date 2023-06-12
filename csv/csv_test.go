package csv

import (
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
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), csvObject)
	assert.Equal(suite.T(), suite.Path, csvObject.Path)
	assert.Equal(suite.T(), suite.FileName, csvObject.Name)
}
