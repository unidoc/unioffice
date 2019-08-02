// Copyright 2017 FoxyUtils ehf. All rights reserved.
package main

import (
	"log"

	"github.com/unidoc/unioffice/spreadsheet"
)

func main() {
	ss := spreadsheet.New()
	sheet := ss.AddSheet()

	sheet.Cell("A1").SetNumber(1)
	sheet.Cell("B1").SetNumber(2)
	sheet.Cell("C1").SetNumber(3)
	sheet.Cell("D1").SetNumber(4)
	sheet.Cell("A2").SetNumber(5)
	sheet.Cell("B2").SetNumber(6)
	sheet.Cell("C2").SetNumber(7)
	sheet.Cell("D2").SetNumber(8)
	sheet.Cell("A3").SetNumber(9)
	sheet.Cell("B3").SetNumber(10)
	sheet.Cell("C3").SetNumber(11)
	sheet.Cell("D3").SetNumber(12)

	sheet.Cell("A5").SetFormulaShared("A1+1", 2, 3)
	sheet.Cell("A9").SetFormulaShared("$A1+1", 2, 3)
	sheet.Cell("A13").SetFormulaShared("$A$1+1", 2, 3)

	ss.RecalculateFormulas()
	if err := ss.Validate(); err != nil {
		log.Fatalf("error validating: %s", err)
	}
	ss.SaveToFile("shared-formula.xlsx")
}
