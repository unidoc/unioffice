// Copyright 2017 FoxyUtils ehf. All rights reserved.
package main

import (
	"log"

	"github.com/unidoc/unioffice/spreadsheet"
)

var lorem = `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Proin lobortis, lectus dictum feugiat tempus, sem neque finibus enim, sed eleifend sem nunc ac diam. Vestibulum tempus sagittis elementum`

func main() {
	ss := spreadsheet.New()
	// add a single sheet
	sheet := ss.AddSheet()

	row := sheet.AddRow()
	cell := row.AddCell()

	wrapped := ss.StyleSheet.AddCellStyle()
	wrapped.SetWrapped(true)
	cell.SetString(lorem)
	cell.SetStyle(wrapped)

	if err := ss.Validate(); err != nil {
		log.Fatalf("error validating sheet: %s", err)
	}

	ss.SaveToFile("wrapped.xlsx")
}
