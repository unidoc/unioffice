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
		// can't add an un-named cell to row zero here as we also add cell 'A1',
		// meaning the un-naned cell must come before 'A1' which is invalid.
		if r != 0 {
			// an unnamed cell displays in the first available column
			row.AddCell().SetString("unnamed-before")
		}

		// setting these to A1, B2, C3, specifically
		cell := row.AddNamedCell(fmt.Sprintf("%c%d", 'A'+r, r+1))
		cell.SetString(fmt.Sprintf("row %d", r))

		// an un-named cell after a named cell is display immediately after a named cell
		row.AddCell().SetString("unnamed-after")
	}

	sheet.AddNumberedRow(26).AddNamedCell("C26").SetString("Cell C26")

	// This line would create an invalid sheet with two identically ID'd rows
	// which would fail validation below
	// sheet.AddNumberedRow(26).AddNamedCell("C27").SetString("Cell C27")

	// so instead use EnsureRow which will create or retrieve an existing row
	sheet.EnsureRow(26).AddNamedCell("E26").SetString("Cell E26")

	if err := ss.Validate(); err != nil {
		log.Fatalf("error validating sheet: %s", err)
	}

	ss.SaveToFile("named-cells.xlsx")
}
