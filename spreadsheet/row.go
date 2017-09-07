// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheet

import (
	"fmt"

	"baliance.com/gooxml"
	"baliance.com/gooxml/measurement"
	"baliance.com/gooxml/schema/schemas.openxmlformats.org/spreadsheetml"
)

// Row is a row within a spreadsheet.
type Row struct {
	w *Workbook
	s *spreadsheetml.CT_Sheet
	x *spreadsheetml.CT_Row
}

// X returns the inner wrapped XML type.
func (r Row) X() *spreadsheetml.CT_Row {
	return r.x
}

// Number returns the row number, or zero if it is unset.
func (r Row) Number() uint32 {
	if r.x.RAttr != nil {
		return *r.x.RAttr
	}
	return 0
}

// SetHeight sets the row height in points.
func (r Row) SetHeight(d measurement.Distance) {
	r.x.HtAttr = gooxml.Float64(float64(d))
}

// SetHeightAuto sets the row height to be automatically determined.
func (r Row) SetHeightAuto() {
	r.x.HtAttr = nil
}

// IsHidden returns whether the row is hidden or not.
func (r Row) IsHidden() bool {
	return r.x.HiddenAttr != nil && *r.x.HiddenAttr
}

// SetHidden hides or unhides the row
func (r Row) SetHidden(hidden bool) {
	if !hidden {
		r.x.HiddenAttr = nil
	} else {
		r.x.HiddenAttr = gooxml.Bool(true)
	}
}

// AddCell adds a cell to a spreadsheet.
func (r Row) AddCell() Cell {
	c := spreadsheetml.NewCT_Cell()
	r.x.C = append(r.x.C, c)
	return Cell{r.w, r.s, r.x, c}
}

// Cells returns a slice of cells.  The cells can be manipulated, but appending
// to the slice will have no effect.
func (r Row) Cells() []Cell {
	ret := []Cell{}
	for _, c := range r.x.C {
		ret = append(ret, Cell{r.w, r.s, r.x, c})
	}
	return ret
}

// AddNamedCell adds a new named cell to a row and returns it. You should
// normally prefer Cell() as it will return the existing cell if the cell
// already exists, while AddNamedCell will duplicate the cell creating an
// invaild spreadsheet.
func (r Row) AddNamedCell(col string) Cell {
	c := spreadsheetml.NewCT_Cell()
	r.x.C = append(r.x.C, c)
	c.RAttr = gooxml.Stringf("%s%d", col, r.Number())
	return Cell{r.w, r.s, r.x, c}
}

// Cell retrieves or adds a new cell to a row. Col is the column (e.g. 'A', 'B')
func (r Row) Cell(col string) Cell {
	name := fmt.Sprintf("%s%d", col, r.Number())
	for _, c := range r.x.C {
		if c.RAttr != nil && *c.RAttr == name {
			return Cell{r.w, r.s, r.x, c}
		}
	}
	return r.AddNamedCell(col)
}
