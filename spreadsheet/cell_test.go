// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheet_test

import (
	"fmt"
	"strconv"
	"testing"
	"time"

	"baliance.com/gooxml/schema/soo/sml"
	"baliance.com/gooxml/spreadsheet"
)

func TestCell(t *testing.T) {
	wb := spreadsheet.New()
	sheet := wb.AddSheet()
	row := sheet.AddRow()
	cell := row.AddCell()

	cell.SetBool(true)
	if cell.X().TAttr != sml.ST_CellTypeB {
		t.Errorf("expected boolean cell type, got %s", cell.X().TAttr)
	}
	if *cell.X().V != "1" {
		t.Errorf("expected 1, got %s", *cell.X().V)
	}
	cell.SetBool(false)
	if *cell.X().V != "0" {
		t.Errorf("expected 0, got %s", *cell.X().V)
	}

	cell.SetInlineString("test123")
	if cell.X().TAttr != sml.ST_CellTypeInlineStr {
		t.Errorf("expected boolean cell type, got %s", cell.X().TAttr)
	}
	if *cell.X().Is.T != "test123" {
		t.Errorf("expected test123, got %s", *cell.X().V)
	}

	cell.SetNumber(1.23)
	if cell.X().TAttr != sml.ST_CellTypeN {
		t.Errorf("expected number cell type, got %s", cell.X().TAttr)
	}
	if *cell.X().V != "1.23" {
		t.Errorf("expected 1.23, got %s", *cell.X().V)
	}

}

func TestCellGetNumber(t *testing.T) {
	wb := spreadsheet.New()
	sheet := wb.AddSheet()
	row := sheet.AddRow()
	cell := row.AddCell()

	cell.SetNumber(1.234)
	f, err := cell.GetValueAsNumber()
	if err != nil {
		t.Errorf("expected no error")
	}
	if f != 1.234 {
		t.Errorf("expected f = 1.234, got %f", f)
	}
}

func TestCellGetNumberFromText(t *testing.T) {
	wb := spreadsheet.New()
	sheet := wb.AddSheet()
	row := sheet.AddRow()
	cell := row.AddCell()

	cell.SetString("foo")
	_, err := cell.GetValueAsNumber()
	if err == nil {
		t.Errorf("expected an error")
	}
}

func TestCellGetBool(t *testing.T) {
	wb := spreadsheet.New()
	sheet := wb.AddSheet()
	row := sheet.AddRow()
	cell := row.AddCell()

	cell.SetBool(true)
	b, err := cell.GetValueAsBool()
	if err != nil {
		t.Errorf("expected no error")
	}
	if !b {
		t.Errorf("expected b = true, got false")
	}
}

func TestCellGetDate(t *testing.T) {
	wb := spreadsheet.New()
	sheet := wb.AddSheet()
	row := sheet.AddRow()
	cell := row.AddCell()

	tm := time.Date(1991, time.April, 8, 1, 2, 3, 0, time.Local)

	cell.SetDate(tm)
	f, err := cell.GetValueAsTime()
	if err != nil {
		t.Errorf("expected no error")
	}

	// SetDate truncates time
	exp := time.Date(1991, time.April, 8, 0, 0, 0, 0, time.Local)
	if !f.Equal(exp) {
		t.Errorf("expected f = %s, got %s", exp, f)
	}
}

func TestCellGetTime(t *testing.T) {
	wb := spreadsheet.New()
	sheet := wb.AddSheet()
	row := sheet.AddRow()
	cell := row.AddCell()

	tm := time.Date(1991, time.April, 8, 1, 2, 3, 0, time.Local)
	cell.SetTime(tm)
	f, err := cell.GetValueAsTime()
	if err != nil {
		t.Errorf("expected no error")
	}
	if !f.Equal(tm) {
		t.Errorf("expected f = %s, got %s", tm, f)
	}
}

func TestCellClear(t *testing.T) {
	wb := spreadsheet.New()
	sheet := wb.AddSheet()
	row := sheet.AddRow()
	cell := row.AddCell()

	cell.SetInlineString("a")
	if cell.X().Is == nil {
		t.Errorf("expected is non nil")
	}

	cell.SetFormulaRaw("=1+2")
	if cell.X().F == nil {
		t.Errorf("expected f != nilnil")
	}

	cell.SetDate(time.Now())
	if cell.X().V == nil {
		t.Errorf("expected v != nil")
	}

	cell.Clear()

	if cell.X().F != nil {
		t.Errorf("expected f = nil")
	}
	if cell.X().Is != nil {
		t.Errorf("expected is = nil")
	}
	if cell.X().V != nil {
		t.Errorf("expected v = nil")
	}
}

func TestCellRichTextString(t *testing.T) {
	wb := spreadsheet.New()
	sheet := wb.AddSheet()
	row := sheet.AddRow()
	cell := row.AddCell()
	rt := cell.SetRichTextString()
	if rt.X() != cell.X().Is {
		t.Errorf("rich text should wrap cell Is")
	}
}

func TestCellStringByID(t *testing.T) {
	wb := spreadsheet.New()
	sheet := wb.AddSheet()
	row := sheet.AddRow()
	cell := row.AddCell()
	// this isn't proper usage of SetStringByID, but it verifies
	// the functionality
	cell.SetStringByID(1)

	v, err := strconv.ParseUint(*cell.X().V, 10, 32)
	if v != 1 || err != nil {
		t.Errorf("expected 1 and no error, got %d %s", v, err)
	}
}

func TestCellReference(t *testing.T) {
	wb := spreadsheet.New()
	sheet := wb.AddSheet()
	for r := 1; r <= 5; r++ {
		row := sheet.AddRow()
		for i := 0; i < 5; i++ {
			expRef := fmt.Sprintf("%c%d", 'A'+i, r)
			cell := row.AddCell()
			if cell.Reference() != expRef {
				t.Errorf("expected ref=%s, got %s", expRef, cell.Reference())
			}
		}
	}
}
