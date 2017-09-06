// Copyright 2017 Baliance. All rights reserved.
package main

import (
	"fmt"
	"log"

	"baliance.com/gooxml/color"
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
			//cell.SetString(fmt.Sprintf("row %d cell %d", r, c))
			rt := cell.SetRichTextString()
			run := rt.AddRun()
			run.SetText(fmt.Sprintf("row %d ", r))
			run.SetBold(true)
			run.SetColor(color.Red)

			run = rt.AddRun()
			run.SetSize(16)
			run.SetItalic(true)
			run.SetFont("Courier")
			run.SetText(fmt.Sprintf("cell %d", c))

		}
	}

	if err := ss.Validate(); err != nil {
		log.Fatalf("error validating sheet: %s", err)
	}

	ss.SaveToFile("rich-text.xlsx")
}
