// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheet_test

import (
	"testing"

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
		if r.Number() != uint32(i) {
			t.Errorf("expected row number %d, got %d", i, r.Number())
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
	if r10.Number() != 10 {
		t.Errorf("expected row number 10, got %d", r10.Number())
	}
	r102 := sheet.EnsureRow(10)
	if r102.Number() != 10 {
		t.Errorf("expected row number 10, got %d", r102.Number())
	}
	if r10.X() != r102.X() {
		t.Errorf("rows should wrap the same inner element")
	}

	// next row should be one after the last row
	r11 := sheet.AddRow()
	if r11.Number() != 11 {
		t.Errorf("expected row number 11, got %d", r11.Number())
	}
}

func TestEnsureRow(t *testing.T) {
	wb := spreadsheet.New()
	sheet := wb.AddSheet()

	r101 := sheet.EnsureRow(10)
	if r101.Number() != 10 {
		t.Errorf("expected row number 10, got %d", r101.Number())
	}
	r102 := sheet.EnsureRow(10)
	if r102.Number() != 10 {
		t.Errorf("expected row number 10, got %d", r102.Number())
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
