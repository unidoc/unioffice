// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package document

import (
	"baliance.com/gooxml"
	"baliance.com/gooxml/color"
	"baliance.com/gooxml/measurement"
	"baliance.com/gooxml/schema/soo/wml"
)

// TableBorders allows manipulation of borders on a table.
type TableBorders struct {
	x *wml.CT_TblBorders
}

// X returns the inner wml.CT_TblBorders
func (b TableBorders) X() *wml.CT_TblBorders {
	return b.x
}

func (b TableBorders) setBorder(brd *wml.CT_Border, t wml.ST_Border, c color.Color, thickness measurement.Distance) {
	brd.ValAttr = t
	brd.ColorAttr = &wml.ST_HexColor{}
	if c.IsAuto() {
		brd.ColorAttr.ST_HexColorAuto = wml.ST_HexColorAutoAuto
	} else {
		brd.ColorAttr.ST_HexColorRGB = c.AsRGBString()
	}
	if thickness != measurement.Zero {
		// sz here is in 1/8'th points, the range is 0.25 to 12 pts
		brd.SzAttr = gooxml.Uint64(uint64(thickness * 8))
	}
}

// SetAll sets all of the borders to a given value.
func (b TableBorders) SetAll(t wml.ST_Border, c color.Color, thickness measurement.Distance) {
	b.SetBottom(t, c, thickness)
	b.SetLeft(t, c, thickness)
	b.SetRight(t, c, thickness)
	b.SetTop(t, c, thickness)
	b.SetInsideHorizontal(t, c, thickness)
	b.SetInsideVertical(t, c, thickness)
}

// SetBottom sets the bottom border to a specified type, color and thickness.
func (b TableBorders) SetBottom(t wml.ST_Border, c color.Color, thickness measurement.Distance) {
	b.x.Bottom = wml.NewCT_Border()
	b.setBorder(b.x.Bottom, t, c, thickness)
}

// SetTop sets the top border to a specified type, color and thickness.
func (b TableBorders) SetTop(t wml.ST_Border, c color.Color, thickness measurement.Distance) {
	b.x.Top = wml.NewCT_Border()
	b.setBorder(b.x.Top, t, c, thickness)
}

// SetLeft sets the left border to a specified type, color and thickness.
func (b TableBorders) SetLeft(t wml.ST_Border, c color.Color, thickness measurement.Distance) {
	b.x.Left = wml.NewCT_Border()
	b.setBorder(b.x.Left, t, c, thickness)
}

// SetRight sets the right border to a specified type, color and thickness.
func (b TableBorders) SetRight(t wml.ST_Border, c color.Color, thickness measurement.Distance) {
	b.x.Right = wml.NewCT_Border()
	b.setBorder(b.x.Right, t, c, thickness)
}

// SetInsideHorizontal sets the interior horizontal borders to a specified type, color and thickness.
func (b TableBorders) SetInsideHorizontal(t wml.ST_Border, c color.Color, thickness measurement.Distance) {
	b.x.InsideH = wml.NewCT_Border()
	b.setBorder(b.x.InsideH, t, c, thickness)
}

// SetInsideVertical sets the interior vertical borders to a specified type, color and thickness.
func (b TableBorders) SetInsideVertical(t wml.ST_Border, c color.Color, thickness measurement.Distance) {
	b.x.InsideV = wml.NewCT_Border()
	b.setBorder(b.x.InsideV, t, c, thickness)
}
