// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheet

import sd "baliance.com/gooxml/schema/schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing"

// CellMarker represents a cell position
type CellMarker struct {
	x *sd.CT_Marker
}

// X returns the inner wrapped XML type.
func (c CellMarker) X() *sd.CT_Marker {
	return c.x
}

func (c CellMarker) Col() int32 {
	return c.x.Col
}
func (c CellMarker) SetCol(col int32) {
	c.x.Col = col
}

func (c CellMarker) Row() int32 {
	return c.x.Row
}
func (c CellMarker) SetRow(row int32) {
	c.x.Row = row
}
