// Copyright 2017 Baliance. All rights reserved.
package main

import (
	"log"
	"math/rand"

	"baliance.com/gooxml/spreadsheet"
)

func main() {
	ss := spreadsheet.New()
	sheet := ss.AddSheet()

	row := sheet.AddRow()
	row.AddCell()
	for i := 0; i < 99; i++ {
		row.AddCell().SetString("Header")
	}
	for i := 0; i < 100; i++ {
		row = sheet.AddRow()
		row.AddCell().SetString("Header")
		for j := 0; j < 99; j++ {
			row.AddCell().SetNumber(rand.Float64() * 100)
		}
	}

	// freeze the first row and column
	sheet.SetFrozen(true, true)

	/* this is equivalent to
	v := sheet.InitialView()
	v.SetState(sml.ST_PaneStateFrozen)
	v.SetYSplit(1)
	v.SetXSplit(1)
	v.SetTopLeft("B2")
	*/

	if err := ss.Validate(); err != nil {
		log.Fatalf("error validating sheet: %s", err)
	}

	ss.SaveToFile("freeze-rows-cols.xlsx")
}
