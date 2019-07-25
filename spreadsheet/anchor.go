// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheet

import (
	"github.com/unidoc/unioffice/measurement"
)

// Anchor is the interface implemented by anchors. It's modeled after the most
// common anchor (Two cell variant with a from/to position), but will also be
// used for one-cell anchors.  In that case the only non-noop methods are
// TopLeft/MoveTo/SetColOffset/SetRowOffset.
type Anchor interface {
	// BottomRight returns the CellMaker for the bottom right corner of the
	// anchor.
	BottomRight() CellMarker
	// TopLeft returns the CellMaker for the top left corner of the anchor.
	TopLeft() CellMarker
	// MoveTo repositions the anchor without changing the objects size.
	MoveTo(col, row int32)

	// SetWidth sets the width of the anchored object. It is not compatible with
	// SetWidthCells.
	SetWidth(w measurement.Distance)
	// SetWidthCells sets the height the anchored object by moving the right
	// hand side. It is not compatible with SetWidth.
	SetWidthCells(w int32)

	// SetHeight sets the height of the anchored object. It is not compatible
	// with SetHeightCells.
	SetHeight(w measurement.Distance)
	// SetHeightCells sets the height the anchored object by moving the bottom.
	// It is not compatible with SetHeight.
	SetHeightCells(h int32)

	// SetColOffset sets the column offset of the top-left anchor.
	SetColOffset(m measurement.Distance)
	// SetRowOffset sets the row offset of the top-left anchor.
	SetRowOffset(m measurement.Distance)

	// Type returns the type of anchor
	Type() AnchorType
}

// AnchorType is the type of anchor.
type AnchorType byte

// AnchorType constants
const (
	AnchorTypeAbsolute AnchorType = iota
	AnchorTypeOneCell
	AnchorTypeTwoCell
)
