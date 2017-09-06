// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheet_test

import (
	"fmt"
	"time"

	"baliance.com/gooxml/spreadsheet"
)

func ExampleSheet_Cell() {
	wb := spreadsheet.New()
	sheet := wb.AddSheet()

	// Cell creates or returns a cell with a given reference
	sheet.Cell("A1").SetNumber(1.23)
}

func ExampleSheet_Row() {
	wb := spreadsheet.New()
	sheet := wb.AddSheet()

	// Row/Cell create or returns a cell with a given reference
	sheet.Row(1).Cell("A").SetNumber(1.23)
}

func ExampleSheet_AddRow() {
	wb := spreadsheet.New()
	sheet := wb.AddSheet()

	// AddRow/AddCell add a new unspecified row/cell.  These will be numbered
	// sequentially in the resulting file.
	sheet.AddRow().AddCell().SetNumber(1.23)
}

func ExampleSheet_AddNumberedRow() {
	wb := spreadsheet.New()
	sheet := wb.AddSheet()

	// AddNumberedRow adds a row with a given number, the difference between Row
	// and AddNumberedRow is that AddNumberedRow doesn't check for an existing
	// row with the same number, while Row will return an existing row if it exists.
	sheet.AddNumberedRow(5).AddCell().SetNumber(1.23)
}

func ExampleSheet_Name() {
	wb := spreadsheet.New()
	sheet := wb.AddSheet()
	sheet.SetName("Sheet 1")
	fmt.Println(sheet.Name())

	// Output: Sheet 1
}

func ExampleCell_SetDate() {
	wb := spreadsheet.New()
	sheet := wb.AddSheet()
	cell := sheet.Cell("A1")
	// set our date value
	cell.SetDate(time.Now())

	// then display it with a date style
	dateStyle := wb.StyleSheet.AddCellStyle()
	dateStyle.SetNumberFormatStandard(spreadsheet.StandardFormatDate)
	cell.SetStyle(dateStyle)
}

func ExampleCell_SetDateWithStyle() {
	wb := spreadsheet.New()
	sheet := wb.AddSheet()
	cell := sheet.Cell("A1")

	cell.SetDateWithStyle(time.Now())
}
