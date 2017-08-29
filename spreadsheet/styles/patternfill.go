// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package styles

import (
	"baliance.com/gooxml/color"

	sml "baliance.com/gooxml/schema/schemas.openxmlformats.org/spreadsheetml"
)

type Fill struct {
	x     *sml.CT_Fill
	fills *sml.CT_Fills
}

func (f Fill) Index() uint32 {
	for i, sf := range f.fills.Fill {
		if f.x == sf {
			return uint32(i)
		}
	}
	return 0
}

type PatternFill struct {
	Fill
}

func NewPatternFill(fills *sml.CT_Fills) PatternFill {
	x := sml.NewCT_Fill()
	x.PatternFill = sml.NewCT_PatternFill()
	return PatternFill{Fill{x, fills}}
}

// SetPattern sets the pattern of the fill.
func (f PatternFill) SetPattern(p sml.ST_PatternType) {
	f.x.PatternFill.PatternTypeAttr = p
}

func (f PatternFill) ClearBgColor() {
	f.x.PatternFill.BgColor = nil
}
func (f PatternFill) SetBgColor(c color.Color) {
	f.x.PatternFill.BgColor = sml.NewCT_Color()
	f.x.PatternFill.BgColor.RgbAttr = c.AsRGBAString()
}
func (f PatternFill) ClearFgColor() {
	f.x.PatternFill.FgColor = nil
}

// SetFgColor sets the *fill* foreground color.  As an example, the solid pattern foreground color becomes the
// background color of the cell when applied.
func (f PatternFill) SetFgColor(c color.Color) {
	f.x.PatternFill.FgColor = sml.NewCT_Color()
	f.x.PatternFill.FgColor.RgbAttr = c.AsRGBAString()
}

func (f PatternFill) X() *sml.CT_Fill {
	return f.x
}
