// Copyright 2017 Baliance. All rights reserved.
package main

import (
	"log"

	"baliance.com/gooxml/color"
	sml "baliance.com/gooxml/schema/schemas.openxmlformats.org/spreadsheetml"
	"baliance.com/gooxml/spreadsheet"
)

func main() {
	ss := spreadsheet.New()
	sheet := ss.AddSheet()

	// set some cell text
	sheet.Cell("C4").SetString("all sides")
	// create and set a style
	cs := ss.StyleSheet.AddCellStyle()
	sheet.Cell("C4").SetStyle(cs)

	// add some borders to the style (ordering isn't important, we could just as
	// easily construct the cell style and then apply it to the cell)
	bAll := ss.StyleSheet.AddBorder()
	cs.SetBorder(bAll)
	bAll.SetLeft(sml.ST_BorderStyleMedium, color.Blue)
	bAll.SetRight(sml.ST_BorderStyleMedium, color.Blue)
	bAll.SetTop(sml.ST_BorderStyleMedium, color.Blue)
	bAll.SetBottom(sml.ST_BorderStyleMedium, color.Blue)

	// red dashed line from top left down to bottom right
	bAll.SetDiagonal(sml.ST_BorderStyleDashed, color.Red, false, true)

	if err := ss.Validate(); err != nil {
		log.Fatalf("error validating sheet: %s", err)
	}

	ss.SaveToFile("borders.xlsx")
}
