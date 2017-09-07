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
	hdrRow := sheet.AddRow()
	hdrRow.AddCell().SetString("Product Name")
	hdrRow.AddCell().SetString("Quantity")
	hdrRow.AddCell().SetString("Price")
	sheet.SetAutoFilter("A1:C6")

	// rows
	for r := 0; r < 5; r++ {
		row := sheet.AddRow()
		row.AddCell().SetString(fmt.Sprintf("Product %d", r+1))
		row.AddCell().SetNumber(float64(r + 2))
		row.AddCell().SetNumber(float64(3*r + 1))

	}

	if err := ss.Validate(); err != nil {
		log.Fatalf("error validating sheet: %s", err)
	}

	ss.SaveToFile("sort-filter.xlsx")
}
