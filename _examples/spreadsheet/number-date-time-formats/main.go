// Copyright 2017 Baliance. All rights reserved.
package main

import (
	"log"
	"time"

	"baliance.com/gooxml/spreadsheet"
)

func main() {
	ss := spreadsheet.New()
	sheet := ss.AddSheet()

	row := sheet.AddRow()
	cell := row.AddCell()

	// If no formatting/styles are applied, then the 'general' format is used.
	cell.SetNumber(1.234)

	// You can also apply a format at the same time you are setting a number.
	cell = row.AddCell()
	cell.SetNumberWithStyle(0.95, spreadsheet.StandardFormatPercent)

	// But that involves a few lookups, so if you're creating many, many cells
	// it wil be faster to
	cell = row.AddCell()
	// create the style
	dateStyle := ss.StyleSheet.AddCellStyle()
	// set its format
	dateStyle.SetNumberFormatStandard(spreadsheet.StandardFormatDate)
	// and apply it to a cell
	cell.SetDate(time.Now())
	cell.SetStyle(dateStyle)

	// It's even faster if repeatedly applying a style to apply the style index
	// directly. This is probably not worth the hassle most of the time, and
	// will generate the same content as calling setXWithStyle
	cs := ss.StyleSheet.AddCellStyle()
	cs.SetNumberFormatStandard(spreadsheet.StandardFormatTime)
	idx := cs.Index()
	for i := 0; i < 5; i++ {
		cell = row.AddCell()
		cell.SetDate(time.Now())
		cell.SetStyleIndex(idx)
	}

	// completely custom number formats can also be used
	customStyle := ss.StyleSheet.AddCellStyle()
	customStyle.SetNumberFormat("$#,##0.00")
	cell = row.AddCell()
	cell.SetNumber(1.234)
	cell.SetStyle(customStyle)

	if err := ss.Validate(); err != nil {
		log.Fatalf("error validating sheet: %s", err)
	}

	ss.SaveToFile("number-date-time-formats.xlsx")
}
