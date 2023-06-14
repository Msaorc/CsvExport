package main

import (
	"github.com/Msaorc/Csv/csv"
)

func main() {
	csvObject, err := csv.NewCsvObject("/home/marcao/projetos/ExportacaoExcel/", "Marcao")
	if err != nil {
		panic(err)
	}

	header := csv.GenerateLine("Nome", "Email")
	line1 := csv.GenerateLine("marcos", "msaorc@hotmail.com")
	line2 := csv.GenerateLine("duda", "duda@gaspareti.com")
	csvObject.AddLine(*line1)
	csvObject.AddLine(*line2)
	csvObject.AddHeader(*header)
	err = csvObject.GenerateCsv()
	if err != nil {
		panic(err)
	}

	// _, err := csv.ReadCsv("/home/marcao/projetos/ExportacaoExcel/CsvTest.csv")
	// if err != nil {
	// 	panic(err)
	// }

	// file, err := os.Open("/home/marcao/projetos/ExportacaoExcel/CsvTest.csv")
	// if err != nil {
	// 	panic(err)
	// }
	// defer file.Close()
	// reader := csv.NewReader(file)
	// reader.Comma = ';'
	// fileOpen, err := reader.ReadAll()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Print(fileOpen)
}
