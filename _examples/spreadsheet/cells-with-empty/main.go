// Copyright 2017 FoxyUtils ehf. All rights reserved.
package main
// This example demonstrates flattening all formulas from an input Excel file and outputs the flattened values to a new xlsx.

import (
	"log"
	"fmt"

	"github.com/unidoc/unioffice/spreadsheet"
)

func main() {
	ss, err := spreadsheet.Open("test.xlsx")
	if err != nil {
		log.Fatalf("error opening document: %s", err)
	}

	s := ss.Sheets()[0]

	maxColumnIdx := s.MaxColumnIdx()
	for _, row := range s.Rows(){
		for _,cell := range row.CellsWithEmpty(maxColumnIdx){
			fmt.Println(cell.Reference(),":", cell.GetFormattedValue())
		}
	}
	fmt.Print("\n\n\n")

	s.Cell("F4").SetString("Hello world")
	maxColumnIdx = s.MaxColumnIdx()
	for _, row := range s.Rows(){
		for _,cell := range row.CellsWithEmpty(maxColumnIdx){
			fmt.Println(cell.Reference(),":", cell.GetFormattedValue())
		}
	}
}
