// Copyright 2017 Baliance. All rights reserved.
package main

import (
	"log"

	"baliance.com/gooxml/spreadsheet"
)

func main() {
	ss := spreadsheet.New()
	sheet := ss.AddSheet()

	// drop-down list that references data from another sheet
	vsheet := ss.AddSheet()
	vsheet.SetName("Validation Data")
	vsheet.Cell("A1").SetString("A")
	vsheet.Cell("A2").SetString("B")
	vsheet.Cell("A3").SetString("C")
	vsheet.Cell("A4").SetString("D")

	sheet.Cell("B1").SetString("references sheet")
	dvCombo := sheet.AddDataValidation()
	dvCombo.SetRange("B2")
	dvList := dvCombo.SetList()
	dvList.SetRange(vsheet.RangeReference("A1:A4"))

	// drop-down list with direct options specified as opposed to referenced
	// from a sheet
	sheet.Cell("C1").SetString("value list")
	dvComboDirect := sheet.AddDataValidation()
	dvComboDirect.SetRange("C2")
	dvListDirect := dvComboDirect.SetList()
	dvListDirect.SetValues([]string{"foo", "bar", "baz"})

	// positive whole numbers
	sheet.Cell("C1").SetString("positive whole numbers")
	dvWhole := sheet.AddDataValidation()
	dvWhole.SetRange("D2")
	dvWholeCmp := dvWhole.SetComparison(spreadsheet.DVCompareTypeWholeNumber, spreadsheet.DVCompareOpGreaterEqual)
	dvWholeCmp.SetValue("0")

	if err := ss.Validate(); err != nil {
		log.Fatalf("error validating sheet: %s", err)
	}

	ss.SaveToFile("validation.xlsx")
}
