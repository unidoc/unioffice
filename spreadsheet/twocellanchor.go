// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheet

import (
	"baliance.com/gooxml/measurement"
	sd "baliance.com/gooxml/schema/soo/dml/spreadsheetDrawing"
)

// TwoCellAnchor is an anchor that is attached to a top-left cell with a fixed
// width/height in cells.
type TwoCellAnchor struct {
	x *sd.CT_TwoCellAnchor
}

// BottomRight returns the CellMaker for the bottom right corner of the anchor.
func (t TwoCellAnchor) BottomRight() CellMarker {
	return CellMarker{t.x.To}
}

// TopLeft returns the CellMaker for the top left corner of the anchor.
func (t TwoCellAnchor) TopLeft() CellMarker {
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
	br.SetCol(col + w)
	br.SetRow(row + h)
}

// SetWidthCells sets the height the anchored object by moving the right hand
// side. It is not compatible with SetWidth.
func (t TwoCellAnchor) SetWidthCells(w int32) {
	tl := t.TopLeft()
	br := t.BottomRight()
	br.SetCol(tl.Col() + w)
}

// SetHeightCells sets the height the anchored object by moving the bottom.  It
// is not compatible with SetHeight.
func (t TwoCellAnchor) SetHeightCells(h int32) {
	t.SetHeight(0)
	tl := t.TopLeft()
	br := t.BottomRight()
	br.SetRow(tl.Row() + h)
}

// SetWidth is a no-op.
func (t TwoCellAnchor) SetWidth(w measurement.Distance) {
}

// SetHeight is a nop-op.
func (t TwoCellAnchor) SetHeight(h measurement.Distance) {
}

// SetColOffset sets the column offset of the two cell anchor.
func (t TwoCellAnchor) SetColOffset(m measurement.Distance) {
	delta := m - t.TopLeft().ColOffset()
	t.TopLeft().SetColOffset(m)
	t.BottomRight().SetColOffset(t.BottomRight().ColOffset() + delta)
}

// SetRowOffset sets the row offset of the two cell anchor
func (t TwoCellAnchor) SetRowOffset(m measurement.Distance) {
	delta := m - t.TopLeft().RowOffset()
	t.TopLeft().SetRowOffset(m)
	t.BottomRight().SetRowOffset(t.BottomRight().RowOffset() + delta)
}

// Type returns the type of anchor
func (t TwoCellAnchor) Type() AnchorType {
	return AnchorTypeTwoCell
}
