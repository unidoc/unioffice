// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheet

import (
	"baliance.com/gooxml"
	"baliance.com/gooxml/measurement"
	"baliance.com/gooxml/schema/schemas.openxmlformats.org/spreadsheetml"
)

// Row is a row within a spreadsheet.
type Row struct {
	w *Workbook
	x *spreadsheetml.CT_Row
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
	return Cell{r.w, c}
}

func (r Row) Cells() []Cell {
	ret := []Cell{}
	for _, c := range r.x.C {
		ret = append(ret, Cell{r.w, c})
	}
	return ret
}
