// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package document

import (
	"baliance.com/gooxml/color"
	"baliance.com/gooxml/schema/soo/wml"
)

// CellProperties are a table cells properties within a document.
type CellProperties struct {
	x *wml.CT_TcPr
}

// X returns the inner wrapped XML type.
func (c CellProperties) X() *wml.CT_TcPr {
	return c.x
}

// SetColumnSpan sets the number of Grid Columns Spanned by the Cell
func (c CellProperties) SetColumnSpan(cols int) {
	if cols == 0 {
		c.x.GridSpan = nil
	} else {
		c.x.GridSpan = wml.NewCT_DecimalNumber()
		c.x.GridSpan.ValAttr = int64(cols)
	}
}

// SetShading controls the cell shading.
func (c CellProperties) SetShading(shd wml.ST_Shd, foreground, fill color.Color) {
	if shd == wml.ST_ShdUnset {
		c.x.Shd = nil
	} else {
		c.x.Shd = wml.NewCT_Shd()
		c.x.Shd.ValAttr = shd
		c.x.Shd.ColorAttr = &wml.ST_HexColor{}
		if foreground.IsAuto() {
			c.x.Shd.ColorAttr.ST_HexColorAuto = wml.ST_HexColorAutoAuto
		} else {
			c.x.Shd.ColorAttr.ST_HexColorRGB = foreground.AsRGBString()
		}
		c.x.Shd.FillAttr = &wml.ST_HexColor{}
		if fill.IsAuto() {
			c.x.Shd.FillAttr.ST_HexColorAuto = wml.ST_HexColorAutoAuto
		} else {
			c.x.Shd.FillAttr.ST_HexColorRGB = fill.AsRGBString()
		}
	}
}
