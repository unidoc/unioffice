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
	"baliance.com/gooxml/schema/soo/sml"
)

// Row is a row within a spreadsheet.
type Row struct {
	w *Workbook
	s *sml.Worksheet
	x *sml.CT_Row
}

// X returns the inner wrapped XML type.
func (r Row) X() *sml.CT_Row {
	return r.x
}

// RowNumber returns the row number (1-N), or zero if it is unset.
func (r Row) RowNumber() uint32 {
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
	numCells := uint32(len(r.x.C))
	var nextCellID *string
	if numCells > 0 {
		prevCellName := gooxml.Stringf("%s%d", IndexToColumn(numCells-1), r.RowNumber())
		// previous cell has an expected name
		if r.x.C[numCells-1].RAttr != nil && *r.x.C[numCells-1].RAttr == *prevCellName {
			nextCellID = gooxml.Stringf("%s%d", IndexToColumn(numCells), r.RowNumber())
		}
	}

	c := sml.NewCT_Cell()
	r.x.C = append(r.x.C, c)

	// fast path failed, so find the last cell and add another
	if nextCellID == nil {
		nextIdx := uint32(0)
		for _, c := range r.x.C {
			if c.RAttr != nil {
				col, _, _ := ParseCellReference(*c.RAttr)
				if col := ColumnToIndex(col); col >= nextIdx {
					nextIdx = col + 1
				}
			}
		}
		nextCellID = gooxml.Stringf("%s%d", IndexToColumn(nextIdx), r.RowNumber())
	}
	c.RAttr = nextCellID
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
	c := sml.NewCT_Cell()

	r.x.C = append(r.x.C, c)
	c.RAttr = gooxml.Stringf("%s%d", col, r.RowNumber())
	return Cell{r.w, r.s, r.x, c}
}

// Cell retrieves or adds a new cell to a row. Col is the column (e.g. 'A', 'B')
func (r Row) Cell(col string) Cell {
	name := fmt.Sprintf("%s%d", col, r.RowNumber())
	for _, c := range r.x.C {
		if c.RAttr != nil && *c.RAttr == name {
			return Cell{r.w, r.s, r.x, c}
		}
	}
	return r.AddNamedCell(col)
}
