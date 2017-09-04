// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheet

// Anchor is the interface implemented by anchors. It's modeled after the most
// common anchor (Two cell variant with a from/to position), but will also be
// used for one-cell anchors.  In that case the only non-noop methods are
// TopLeft() and MoveTo().
type Anchor interface {
	BottomRight() CellMarker
	TopLeft() CellMarker
	MoveTo(col, row int32)
	SetWidth(w int32)
	SetHeight(h int32)
}
