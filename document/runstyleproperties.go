// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package document

import (
	"baliance.com/gooxml"
	"baliance.com/gooxml/measurement"
	"baliance.com/gooxml/schema/soo/wml"
)

// RunStyleProperties controls run styling properties
type RunStyleProperties struct {
	x *wml.CT_RPr
}

// X returns the inner wrapped XML type.
func (r RunStyleProperties) X() *wml.CT_RPr {
	return r.x
}

// SetSize sets the font size for a run.
func (r RunStyleProperties) SetSize(size measurement.Distance) {
	r.x.Sz = wml.NewCT_HpsMeasure()
	r.x.Sz.ValAttr.ST_UnsignedDecimalNumber = gooxml.Uint64(uint64(size / measurement.HalfPoint))
	r.x.SzCs = wml.NewCT_HpsMeasure()
	r.x.SzCs.ValAttr.ST_UnsignedDecimalNumber = gooxml.Uint64(uint64(size / measurement.HalfPoint))
}

// SetKerning sets the run's font kerning.
func (r RunStyleProperties) SetKerning(size measurement.Distance) {
	r.x.Kern = wml.NewCT_HpsMeasure()
	r.x.Kern.ValAttr.ST_UnsignedDecimalNumber = gooxml.Uint64(uint64(size / measurement.HalfPoint))
}

// SetCharacterSpacing sets the run's Character Spacing Adjustment.
func (r RunStyleProperties) SetCharacterSpacing(size measurement.Distance) {
	r.x.Spacing = wml.NewCT_SignedTwipsMeasure()
	r.x.Spacing.ValAttr.Int64 = gooxml.Int64(int64(size / measurement.Twips))
}

// Fonts returns the style's Fonts.
func (r RunStyleProperties) Fonts() Fonts {
	if r.x.RFonts == nil {
		r.x.RFonts = wml.NewCT_Fonts()
	}
	return Fonts{r.x.RFonts}
}

// Color returns the style's Color.
func (r RunStyleProperties) Color() Color {
	if r.x.Color == nil {
		r.x.Color = wml.NewCT_Color()
	}
	return Color{r.x.Color}
}
