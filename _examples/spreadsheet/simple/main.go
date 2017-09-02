// Copyright 2017 Baliance. All rights reserved.
package main

import (
	"fmt"
	"log"

	"baliance.com/gooxml/spreadsheet"
)

func main() {
	ss := spreadsheet.New()
	// add a single sheet
	sheet := ss.AddSheet()

	// rows
	for r := 0; r < 5; r++ {
		row := sheet.AddRow()
		// and cells
		for c := 0; c < 5; c++ {
			cell := row.AddCell()
			cell.SetString(fmt.Sprintf("row %d cell %d", r, c))
		}
	}

	if err := ss.Validate(); err != nil {
		log.Fatalf("error validating sheet: %s", err)
	}

	ss.SaveToFile("simple.xlsx")
}
