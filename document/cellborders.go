// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io.

package document

import (
	"github.com/unidoc/unioffice/color"
	"github.com/unidoc/unioffice/measurement"
	"github.com/unidoc/unioffice/schema/soo/wml"
)

// CellBorders are the borders for an individual
type CellBorders struct {
	x *wml.CT_TcBorders
}

// X returns the inner wrapped type
func (b CellBorders) X() *wml.CT_TcBorders {
	return b.x
}

// SetAll sets all of the borders to a given value.
func (b CellBorders) SetAll(t wml.ST_Border, c color.Color, thickness measurement.Distance) {
	b.SetBottom(t, c, thickness)
	b.SetLeft(t, c, thickness)
	b.SetRight(t, c, thickness)
	b.SetTop(t, c, thickness)
	b.SetInsideHorizontal(t, c, thickness)
	b.SetInsideVertical(t, c, thickness)
}

// SetBottom sets the bottom border to a specified type, color and thickness.
func (b CellBorders) SetBottom(t wml.ST_Border, c color.Color, thickness measurement.Distance) {
	b.x.Bottom = wml.NewCT_Border()
	setBorder(b.x.Bottom, t, c, thickness)
}

// SetTop sets the top border to a specified type, color and thickness.
func (b CellBorders) SetTop(t wml.ST_Border, c color.Color, thickness measurement.Distance) {
	b.x.Top = wml.NewCT_Border()
	setBorder(b.x.Top, t, c, thickness)
}

// SetLeft sets the left border to a specified type, color and thickness.
func (b CellBorders) SetLeft(t wml.ST_Border, c color.Color, thickness measurement.Distance) {
	b.x.Left = wml.NewCT_Border()
	setBorder(b.x.Left, t, c, thickness)
}

// SetRight sets the right border to a specified type, color and thickness.
func (b CellBorders) SetRight(t wml.ST_Border, c color.Color, thickness measurement.Distance) {
	b.x.Right = wml.NewCT_Border()
	setBorder(b.x.Right, t, c, thickness)
}

// SetInsideHorizontal sets the interior horizontal borders to a specified type, color and thickness.
func (b CellBorders) SetInsideHorizontal(t wml.ST_Border, c color.Color, thickness measurement.Distance) {
	b.x.InsideH = wml.NewCT_Border()
	setBorder(b.x.InsideH, t, c, thickness)
}

// SetInsideVertical sets the interior vertical borders to a specified type, color and thickness.
func (b CellBorders) SetInsideVertical(t wml.ST_Border, c color.Color, thickness measurement.Distance) {
	b.x.InsideV = wml.NewCT_Border()
	setBorder(b.x.InsideV, t, c, thickness)
}
