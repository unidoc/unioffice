// Copyright 2017 Baliance. All rights reserved.
package main

import (
	"fmt"
	"log"

	"baliance.com/gooxml/spreadsheet"

	"baliance.com/gooxml/schema/soo/sml"
)

func main() {
	ss := spreadsheet.New()
	sheet := ss.AddSheet()

	sheet.Cell("A1").SetString("Hello World!")
	sheet.Cell("B1").SetString("will not be visible") // as it's not the first cell within a merged range Excel warns you when you do this through the UI
	sheet.AddMergedCells("A1", "C2")

	centered := ss.StyleSheet.AddCellStyle()
	centered.SetHorizontalAlignment(sml.ST_HorizontalAlignmentCenter)
	centered.SetVerticalAlignment(sml.ST_VerticalAlignmentCenter)
	sheet.Cell("A1").SetStyle(centered)

	for _, m := range sheet.MergedCells() {
		fmt.Println("merged region", m.Reference(), "has contents", m.Cell().GetString())
	}

	if err := ss.Validate(); err != nil {
		log.Fatalf("error validating sheet: %s", err)
	}

	ss.SaveToFile("merged.xlsx")
}
