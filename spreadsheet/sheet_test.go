// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheet_test

import (
	"math"
	"math/rand"
	"testing"

	"baliance.com/gooxml"
	"baliance.com/gooxml/spreadsheet"
)

func TestRowNumIncreases(t *testing.T) {
	wb := spreadsheet.New()
	sheet := wb.AddSheet()
	if len(sheet.Rows()) != 0 {
		t.Errorf("new sheet must have zero rows, had %d", len(sheet.Rows()))
	}

	// add 5 rows
	for i := 1; i < 5; i++ {
		r := sheet.AddRow()
		if r.RowNumber() != uint32(i) {
			t.Errorf("expected row number %d, got %d", i, r.RowNumber())
		}
	}
}

func TestAddNumberedRow(t *testing.T) {
	wb := spreadsheet.New()
	sheet := wb.AddSheet()
	for i := 1; i < 5; i++ {
		sheet.AddRow()
	}

	r10 := sheet.AddNumberedRow(10)
	if r10.RowNumber() != 10 {
		t.Errorf("expected row number 10, got %d", r10.RowNumber())
	}
	r102 := sheet.Row(10)
	if r102.RowNumber() != 10 {
		t.Errorf("expected row number 10, got %d", r102.RowNumber())
	}
	if r10.X() != r102.X() {
		t.Errorf("rows should wrap the same inner element")
	}

	// next row should be one after the last row
	r11 := sheet.AddRow()
	if r11.RowNumber() != 11 {
		t.Errorf("expected row number 11, got %d", r11.RowNumber())
	}
}

func TestEnsureRow(t *testing.T) {
	wb := spreadsheet.New()
	sheet := wb.AddSheet()

	r101 := sheet.Row(10)
	if r101.RowNumber() != 10 {
		t.Errorf("expected row number 10, got %d", r101.RowNumber())
	}
	r102 := sheet.Row(10)
	if r102.RowNumber() != 10 {
		t.Errorf("expected row number 10, got %d", r102.RowNumber())
	}
	if r101.X() != r102.X() {
		t.Errorf("rows should wrap the same inner element")
	}

}

func TestRowNumberValidation(t *testing.T) {
	wb := spreadsheet.New()
	sheet := wb.AddSheet()
	sheet.AddNumberedRow(2)
	sheet.AddNumberedRow(2)
	if err := sheet.Validate(); err == nil {
		t.Errorf("expected validation error with identically numbered rows")
	}
}

func TestAutoFilter(t *testing.T) {
	wb := spreadsheet.New()
	sheet := wb.AddSheet()
	if len(wb.DefinedNames()) != 0 {
		t.Errorf("expected no defined names for new workbook")
	}
	sheet.SetAutoFilter("A1:C10")
	if len(wb.DefinedNames()) != 1 {
		t.Errorf("expected a new defined names for the autofilter")
	}
	dn := wb.DefinedNames()[0]
	expContent := "'Sheet 1'!$A$1:$C$10"
	if dn.Content() != expContent {
		t.Errorf("expected defined name content = '%s', got %s", expContent, dn.Content())
	}

	sheet.SetAutoFilter("A1:B10")
	expContent = "'Sheet 1'!$A$1:$B$10"
	// setting the filter again should re-write the defined name and not create a new one
	if len(wb.DefinedNames()) != 1 {
		t.Errorf("expected a new defined names for the autofilter")
	}
	dn = wb.DefinedNames()[0]
	// but the content should have changed
	if dn.Content() != expContent {
		t.Errorf("expected defined name content = '%s', got %s", expContent, dn.Content())
	}

	sheet.ClearAutoFilter()
	if len(wb.DefinedNames()) != 0 {
		t.Errorf("clearing the filter should have removed the defined name")
	}

	if sheet.X().AutoFilter != nil {
		t.Errorf("autofilter should have been nil after clear")
	}
}

func TestSheetNameLength(t *testing.T) {
	wb := spreadsheet.New()
	sheet := wb.AddSheet()
	if err := sheet.Validate(); err != nil {
		t.Errorf("expected no validaton error on new sheet, got %s:", err)
	}
	sheet.SetName("01234567890123456789012345678901")
	if err := sheet.Validate(); err == nil {
		t.Errorf("expected validation error with sheet name too long")
	}
}

func TestMergedCell(t *testing.T) {
	wb := spreadsheet.New()
	sheet := wb.AddSheet()

	expContent := "testing 123"
	sheet.Cell("A1").SetString(expContent)
	sheet.Cell("B1").SetString("in range, but not visible")
	if len(sheet.MergedCells()) != 0 {
		t.Errorf("new sheet should have no merged cells")
	}
	sheet.AddMergedCells("A1", "C2")
	if len(sheet.MergedCells()) != 1 {
		t.Errorf("sheet should have a single merged cells")
	}

	mc := sheet.MergedCells()[0]
	expRef := "A1:C2"
	if mc.Reference() != expRef {
		t.Errorf("expected merged cell reference %s, got %s", expRef, mc.Reference())
	}

	if mc.Cell().GetString() != expContent {
		t.Errorf("expected merged cell content to be '%s', got '%s'", expContent, mc.Cell().GetString())
	}

	sheet.RemoveMergedCell(mc)
	if len(sheet.MergedCells()) != 0 {
		t.Errorf("after removal, sheet should have no merged cells")
	}
}

func TestSheetExtents(t *testing.T) {
	ss := spreadsheet.New()
	sheet := ss.AddSheet()
	if sheet.Extents() != "A1:A1" {
		t.Errorf("expected 'A1:A1' for empty sheet, got %s", sheet.Extents())
	}

	for r := 0; r < 5; r++ {
		row := sheet.AddRow()
		for c := 0; c < 5; c++ {
			cell := row.AddCell()
			cell.SetNumber(float64(rand.Intn(1000)) / 100.0)
		}
	}

	exp := "A1:E5"
	if sheet.Extents() != exp {
		t.Errorf("expected %s , got %s", exp, sheet.Extents())
	}

}

func TestSheetClearCachedFormula(t *testing.T) {
	ss := spreadsheet.New()
	sheet := ss.AddSheet()
	cell := sheet.Cell("A1")
	cell.SetFormulaRaw("foo")
	cell.X().V = gooxml.String("cached-results")
	sheet.ClearCachedFormulaResults()
	if cell.X().V != nil {
		t.Errorf("cached result not cleared")
	}
}

func TestFormattedCell(t *testing.T) {
	wb, err := spreadsheet.Open("testdata/fmt.xlsx")
	if err != nil {
		t.Fatalf("error reading fmt.xlsx: %s", err)
	}
	// these cells all have the same value with different formatting applied
	td := []struct {
		Cell string
		Exp  string
	}{
		{"A1", "9/18/17"},
		{"A2", "Monday, September 18, 2017"},
		{"A3", "4:47:28 PM"},
		{"A4", "42996.69963"},
		{"A5", "42996.70"},
		{"A6", "4.30E+04"},
		{"A7", "9/18"},
		{"A8", "18-Sep-17"},
	}

	sheet := wb.Sheets()[0]
	for _, tc := range td {
		got := sheet.Cell(tc.Cell).GetFormattedValue()
		if got != tc.Exp {
			t.Errorf("expected %s in cell %s, got %s", tc.Exp, tc.Cell, got)
		}
	}
}

func TestInfNan(t *testing.T) {
	wb := spreadsheet.New()
	sheet := wb.AddSheet()
	sheet.Cell("A1").SetNumber(math.NaN())

	rv, err := sheet.Cell("A1").GetRawValue()
	if err != nil {
		t.Errorf("got error: %s", err)
	}
	if rv != "#NUM!" {
		t.Error("expected error for NaN")
	}

	sheet.Cell("A1").SetNumber(math.Inf(1))
	rv, err = sheet.Cell("A1").GetRawValue()
	if err != nil {
		t.Errorf("got error: %s", err)
	}
	if rv != "#NUM!" {
		t.Error("expected error for NaN")
	}
}

func TestMergedCellValidation(t *testing.T) {
	wb := spreadsheet.New()
	sheet := wb.AddSheet()
	sheet.AddMergedCells("A1", "B5")
	sheet.AddMergedCells("A3", "B9")
	if err := sheet.Validate(); err == nil {
		t.Errorf("expected validation error due to overlapping merged cells")
	}

}
