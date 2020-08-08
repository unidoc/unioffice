// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io.

package spreadsheet

import (
	"github.com/unidoc/unioffice/color"

	"github.com/unidoc/unioffice/schema/soo/sml"
)

type PatternFill struct {
	x *sml.CT_PatternFill
	f *sml.CT_Fill
}

func NewPatternFill(fills *sml.CT_Fills) PatternFill {
	x := sml.NewCT_Fill()
	x.PatternFill = sml.NewCT_PatternFill()
	return PatternFill{x.PatternFill, x}
}

func (f PatternFill) X() *sml.CT_PatternFill {
	return f.x
}

// SetPattern sets the pattern of the fill.
func (f PatternFill) SetPattern(p sml.ST_PatternType) {
	f.x.PatternTypeAttr = p
}

func (f PatternFill) ClearBgColor() {
	f.x.BgColor = nil
}
func (f PatternFill) SetBgColor(c color.Color) {
	f.x.BgColor = sml.NewCT_Color()
	f.x.BgColor.RgbAttr = c.AsRGBAString()
}
func (f PatternFill) ClearFgColor() {
	f.x.FgColor = nil
}

// SetFgColor sets the *fill* foreground color.  As an example, the solid pattern foreground color becomes the
// background color of the cell when applied.
func (f PatternFill) SetFgColor(c color.Color) {
	f.x.FgColor = sml.NewCT_Color()
	f.x.FgColor.RgbAttr = c.AsRGBAString()
}
