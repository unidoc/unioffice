// Copyright 2017 FoxyUtils ehf. All rights reserved.
package main

import (
	"fmt"
	"log"

	"github.com/unidoc/unioffice/spreadsheet"
	"github.com/unidoc/unioffice/spreadsheet/formula"
)

func main() {
	ss, err := spreadsheet.Open("test_for_unidoc.xlsx")
	if err != nil {
		log.Fatalf("error opening document: %s", err)
	}

	sheet, err := ss.GetSheet("Sheet2")
	if err != nil {
		log.Fatalf("error opening sheet: %s", err)
	}
	//sheet.RecalculateFormulas()
	formEv := formula.NewEvaluator()

	// the formula context allows the formula evaluator to pull data from a
	// sheet
	cell := sheet.FormulaContext().Cell("A5", formEv)
	if cell.Type == formula.ResultTypeError {
		fmt.Println("ERROR:", cell.ErrorMessage)
	}
	fmt.Println("A3 is", cell.Value())
}
