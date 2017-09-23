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

// OneCellAnchor is anchored to a top-left cell with a fixed with/height
// in distance.
type OneCellAnchor struct {
	x *sd.CT_OneCellAnchor
}

// SetColOffset sets the column offset of the top-left anchor.
func (o OneCellAnchor) SetColOffset(m measurement.Distance) {
	o.TopLeft().SetColOffset(m)
}

// SetRowOffset sets the row offset of the top-left anchor.
func (o OneCellAnchor) SetRowOffset(m measurement.Distance) {
	o.TopLeft().SetRowOffset(m)
}

// SetHeight sets the height of the anchored object.
func (o OneCellAnchor) SetHeight(h measurement.Distance) {
	o.x.Ext.CyAttr = int64(h / measurement.EMU)
}

// SetWidth sets the width of the anchored object.
func (o OneCellAnchor) SetWidth(w measurement.Distance) {
	o.x.Ext.CxAttr = int64(w / measurement.EMU)
}

// MoveTo moves the top-left of the anchored object.
func (o OneCellAnchor) MoveTo(col, row int32) {
	o.TopLeft().SetCol(col)
	o.TopLeft().SetRow(row)
}

// TopLeft returns the top-left corner of the anchored object.
func (o OneCellAnchor) TopLeft() CellMarker {
	return CellMarker{o.x.From}
}

// BottomRight is a no-op.
func (o OneCellAnchor) BottomRight() CellMarker {
	return CellMarker{}
}

// SetHeightCells is a no-op.
func (o OneCellAnchor) SetHeightCells(int32) {
}

// SetWidthCells is a no-op.
func (o OneCellAnchor) SetWidthCells(int32) {
}

// Type returns the type of anchor
func (o OneCellAnchor) Type() AnchorType {
	return AnchorTypeOneCell
}
