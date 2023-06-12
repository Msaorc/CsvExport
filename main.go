package main

import "github.com/Msaorc/Csv/csv"

func main() {
	csvObject, err := csv.NewCsvObject("/home/marcao/projetos/ExportacaoExcel/", "Marcao")
	if err != nil {
		panic(err)
	}

	header := csv.GenerateLine("Nome", "Email");
	line1 := csv.GenerateLine("marcos", "msaorc@hotmail.com")
	line2 := csv.GenerateLine("duda", "duda@gaspareti.com")
	csvObject.AddLine(*line1)
	csvObject.AddLine(*line2)
	csvObject.AddHeader(*header)
	err = csvObject.GenerateCsv()
	if err != nil {
		panic(err)
	}
}
