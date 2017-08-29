// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheet_test

import "testing"
import "baliance.com/gooxml/spreadsheet"
import sml "baliance.com/gooxml/schema/schemas.openxmlformats.org/spreadsheetml"

func TestCell(t *testing.T) {
	wb := spreadsheet.New()
	sheet := wb.AddSheet("test")
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
