// Copyright 2017 FoxyUtils ehf. All rights reserved.
package main
// This example demonstrates flattening all formulas from an input Excel file and outputs the flattened values to a new xlsx.

import (
	"log"
	"fmt"
	"time"

	"github.com/unidoc/unioffice/spreadsheet"
	"github.com/unidoc/unioffice/spreadsheet/formula"
)

func main() {
	ss, err := spreadsheet.Open("formulas.xlsx")
	if err != nil {
		log.Fatalf("error opening document: %s", err)
	}

	sheets := ss.Sheets()

	start := time.Now().UnixNano()
	formEv := formula.NewEvaluator()
	for _, sheet := range sheets {
		fmt.Println("Sheet name:", sheet.Name())
		ctx := sheet.FormulaContext()
		for _, row := range sheet.Rows() {
			for _, cell := range row.Cells() {
				c := ctx.Cell(cell.Reference(), formEv)
				value := ""
				if cell.X().V != nil {
					value = *cell.X().V
				}
				cell.Clear()
				setValue(cell, c, value)
			}
		}
	}
	finish := time.Now().UnixNano()
	fmt.Printf("total: %d ns\n", finish - start)

	ss.SaveToFile("values.xlsx")
}

func setValue(cell spreadsheet.Cell, c formula.Result, value string) {
	switch c.Type {
	case formula.ResultTypeNumber:
		if c.IsBoolean {
			cell.SetBool(value != "0")
		} else {
			cell.SetNumber(c.ValueNumber)
		}
	case formula.ResultTypeString:
		cell.SetString(c.ValueString)
	case formula.ResultTypeList:
		setValue(cell, c.ValueList[0], value)
	case formula.ResultTypeArray:
		setValue(cell, c.ValueArray[0][0], value)
	case formula.ResultTypeError:
		cell.SetError(c.ValueString)
	}
}
