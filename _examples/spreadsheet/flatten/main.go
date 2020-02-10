// Copyright 2017 FoxyUtils ehf. All rights reserved.
package main
// This example demonstrates flattening all formulas from an input Excel file and outputs the flattened values to a new xlsx.

import (
	"log"
	"fmt"
	"runtime"
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
				// copying cell style
				cellStyle := spreadsheet.CellStyle{}
				x := cell.X()
				if x.SAttr != nil {
					sid := *x.SAttr
					cellStyle = ss.StyleSheet.GetCellStyle(sid)
				}

				// copying value
				c := ctx.Cell(cell.Reference(), formEv)
				value := ""
				if cell.X().V != nil {
					value = *cell.X().V
				}
				cell.Clear()
				setValue(cell, c, value)

				// setting cell style
				if !cellStyle.IsEmpty() {
					cell.SetStyle(cellStyle)
				}
			}
		}
	}
	finish := time.Now().UnixNano()
	fmt.Printf("total time: %d ns\n", finish - start)
	PrintMemUsage()

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

func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Println("Memory usage:")
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
