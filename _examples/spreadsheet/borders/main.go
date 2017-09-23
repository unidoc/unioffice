// Copyright 2017 Baliance. All rights reserved.
package main

import (
	"log"

	"baliance.com/gooxml/color"
	"baliance.com/gooxml/schema/soo/sml"
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
	bAll.SetLeft(sml.ST_BorderStyleThin, color.Blue)
	bAll.SetRight(sml.ST_BorderStyleThin, color.Blue)
	bAll.SetTop(sml.ST_BorderStyleThin, color.Blue)
	bAll.SetBottom(sml.ST_BorderStyleThin, color.Blue)
	// red dashed line from top left down to bottom right
	bAll.SetDiagonal(sml.ST_BorderStyleDashed, color.Red, false, true)

	// Cell styles and thus border styles only apply to a single cell.  This
	// means to apply a boxed border around multiple cells, you would need to
	// create individual styles for the corners, left, right, top and bottom
	// sides.  There is a helper that can do this for you, ignoring any diagonal
	// borders.
	sheet.SetBorder("B6:D10", bAll)

	if err := ss.Validate(); err != nil {
		log.Fatalf("error validating sheet: %s", err)
	}

	ss.SaveToFile("borders.xlsx")
}
