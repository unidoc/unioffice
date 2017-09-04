// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheet

import (
	sd "baliance.com/gooxml/schema/schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing"
)

// TwoCellAnchor is an anchor that is attached from a top-left to a bottom-right
// cell (from & to cells).
type TwoCellAnchor struct {
	x *sd.CT_TwoCellAnchor
}

// BottomRight returns the CellMaker for the bottom right corner of the anchor.
func (t TwoCellAnchor) BottomRight() CellMarker {
	if t.x.To == nil {
		t.x.To = sd.NewCT_Marker()
	}
	return CellMarker{t.x.To}
}

// TopLeft returns the CellMaker for the top left corner of the anchor.
func (t TwoCellAnchor) TopLeft() CellMarker {
	if t.x.From == nil {
		t.x.From = sd.NewCT_Marker()
	}
	return CellMarker{t.x.From}
}

// MoveTo repositions the anchor without changing the objects size.
func (t TwoCellAnchor) MoveTo(col, row int32) {
	tl := t.TopLeft()
	br := t.BottomRight()
	w := br.Col() - tl.Col()
	h := br.Row() - tl.Row()

	tl.SetCol(col)
	tl.SetRow(row)
	t.SetWidth(w)
	t.SetHeight(h)
}

// SetWidth sets the height the anchored objet by moving the right hand side.
func (t TwoCellAnchor) SetWidth(w int32) {
	tl := t.TopLeft()
	br := t.BottomRight()
	br.SetCol(tl.Col() + w)
}

// SetHeight sets the height the anchored objet by moving the bottom.
func (t TwoCellAnchor) SetHeight(h int32) {
	tl := t.TopLeft()
	br := t.BottomRight()
	br.SetRow(tl.Row() + h)
}
