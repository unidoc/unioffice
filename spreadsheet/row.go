// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased on https://unidoc.io.

package spreadsheet

import (
	"fmt"

	"github.com/unidoc/unioffice"
	"github.com/unidoc/unioffice/measurement"
	"github.com/unidoc/unioffice/schema/soo/sml"
	"github.com/unidoc/unioffice/spreadsheet/reference"
)

// Row is a row within a spreadsheet.
type Row struct {
	w     *Workbook
	sheet *Sheet
	x     *sml.CT_Row
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
	r.x.HtAttr = unioffice.Float64(float64(d))
	r.x.CustomHeightAttr = unioffice.Bool(true)
}

// SetHeightAuto sets the row height to be automatically determined.
func (r Row) SetHeightAuto() {
	r.x.HtAttr = nil
	r.x.CustomHeightAttr = nil
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
		r.x.HiddenAttr = unioffice.Bool(true)
	}
}

// AddCell adds a cell to a spreadsheet.
func (r Row) AddCell() Cell {
	numCells := uint32(len(r.x.C))
	var nextCellID *string
	if numCells > 0 {
		prevCellName := unioffice.Stringf("%s%d", reference.IndexToColumn(numCells-1), r.RowNumber())
		// previous cell has an expected name
		if r.x.C[numCells-1].RAttr != nil && *r.x.C[numCells-1].RAttr == *prevCellName {
			nextCellID = unioffice.Stringf("%s%d", reference.IndexToColumn(numCells), r.RowNumber())
		}
	}

	c := sml.NewCT_Cell()
	r.x.C = append(r.x.C, c)

	// fast path failed, so find the last cell and add another
	if nextCellID == nil {
		nextIdx := uint32(0)
		for _, c := range r.x.C {
			if c.RAttr != nil {
				cref, _ := reference.ParseCellReference(*c.RAttr)
				if cref.ColumnIdx >= nextIdx {
					nextIdx = cref.ColumnIdx + 1
				}
			}
		}
		nextCellID = unioffice.Stringf("%s%d", reference.IndexToColumn(nextIdx), r.RowNumber())
	}
	c.RAttr = nextCellID
	return Cell{r.w, r.sheet, r.x, c}
}

// Cells returns a slice of cells.  The cells can be manipulated, but appending
// to the slice will have no effect.
func (r Row) Cells() []Cell {
	ret := []Cell{}
	lastIndex := -1
	for _, c := range r.x.C {
		if c.RAttr == nil {
			unioffice.Log("RAttr is nil for a cell, skipping.")
			continue
		}
		ref, err := reference.ParseCellReference(*c.RAttr)
		if err != nil {
			unioffice.Log("RAttr is incorrect for a cell: " + *c.RAttr + ", skipping.")
			continue
		}
		currentIndex := int(ref.ColumnIdx)
		if currentIndex-lastIndex > 1 {
			for col := lastIndex + 1; col < currentIndex; col++ {
				ret = append(ret, r.Cell(reference.IndexToColumn(uint32(col))))
			}
		}
		lastIndex = currentIndex
		ret = append(ret, Cell{r.w, r.sheet, r.x, c})
	}
	return ret
}

// CellsWithEmpty returns a slice of cells including empty ones from the first column to the last one used in the sheet.
// The cells can be manipulated, but appending to the slice will have no effect.
func (r Row) CellsWithEmpty(lastColIdx uint32) []Cell {
	ret := []Cell{}
	for columnIdx := uint32(0); columnIdx <= lastColIdx; columnIdx++ {
		c := r.Cell(reference.IndexToColumn(columnIdx))
		ret = append(ret, c)
	}
	return ret
}

// AddNamedCell adds a new named cell to a row and returns it. You should
// normally prefer Cell() as it will return the existing cell if the cell
// already exists, while AddNamedCell will duplicate the cell creating an
// invaild spreadsheet.
func (r Row) AddNamedCell(col string) Cell {
	c := sml.NewCT_Cell()
	c.RAttr = unioffice.Stringf("%s%d", col, r.RowNumber())

	indexToInsert := -1
	colIdx := reference.ColumnToIndex(col)
	for i, cell := range r.x.C {
		cr, err := reference.ParseCellReference(*cell.RAttr)
		if err != nil {
			return Cell{}
		}
		if colIdx < cr.ColumnIdx {
			indexToInsert = i
			break
		}
	}
	if indexToInsert == -1 {
		r.x.C = append(r.x.C, c)
	} else {
		r.x.C = append(r.x.C[:indexToInsert], append([]*sml.CT_Cell{c}, r.x.C[indexToInsert:]...)...)
	}

	return Cell{r.w, r.sheet, r.x, c}
}

// Cell retrieves or adds a new cell to a row. Col is the column (e.g. 'A', 'B')
func (r Row) Cell(col string) Cell {
	name := fmt.Sprintf("%s%d", col, r.RowNumber())
	for _, c := range r.x.C {
		if c.RAttr != nil && *c.RAttr == name {
			return Cell{r.w, r.sheet, r.x, c}
		}
	}
	return r.AddNamedCell(col)
}

// renumberAs assigns a new row number and fixes any cell references within the
// row so they refer to the new row number. This is used when sorting to fix up
// moved rows.
func (r Row) renumberAs(rowNumber uint32) {
	r.x.RAttr = unioffice.Uint32(rowNumber)
	for _, c := range r.Cells() {
		cref, err := reference.ParseCellReference(c.Reference())
		if err == nil {
			newRef := fmt.Sprintf("%s%d", cref.Column, rowNumber)
			c.x.RAttr = unioffice.String(newRef)
		}
	}
}
