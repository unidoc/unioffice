// Copyright 2017 Baliance. All rights reserved.
package main

import (
	"fmt"
	"log"

	"baliance.com/gooxml/spreadsheet"
)

func main() {
	ss := spreadsheet.New()
	sheet := ss.AddSheet()

	for r := 0; r < 5; r++ {
		row := sheet.AddRow()

		// can't add an un-named cell to row zero here as we also add cell 'A1',
		// meaning the un-naned cell must come before 'A1' which is invalid.
		if r != 0 {
			// an unnamed cell displays in the first available column
			row.AddCell().SetString("unnamed-before")
		}

		// setting these to A, B, C, specifically
		cell := row.AddNamedCell(fmt.Sprintf("%c", 'A'+r))
		cell.SetString(fmt.Sprintf("row %d", r))

		// an un-named cell after a named cell is display immediately after a named cell
		row.AddCell().SetString("unnamed-after")
	}

	sheet.AddNumberedRow(26).AddNamedCell("C").SetString("Cell C26")

	// This line would create an invalid sheet with two identically ID'd rows
	// which would fail validation below
	// sheet.AddNumberedRow(26).AddNamedCell("C27").SetString("Cell C27")

	// so instead use Row which will create or retrieve an existing row
	sheet.Row(26).AddNamedCell("E").SetString("Cell E26")
	sheet.Row(26).Cell("F").SetString("Cell F26")

	// You can also reference cells fully from the sheet.
	sheet.Cell("H1").SetString("Cell H1")
	sheet.Cell("H2").SetString("Cell H2")
	sheet.Cell("H3").SetString("Cell H3")

	if err := ss.Validate(); err != nil {
		log.Fatalf("error validating sheet: %s", err)
	}

	ss.SaveToFile("named-cells.xlsx")
}
