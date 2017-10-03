// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package drawing

import (
	"baliance.com/gooxml"
	"baliance.com/gooxml/color"
	"baliance.com/gooxml/measurement"
	"baliance.com/gooxml/schema/soo/dml"
)

// RunProperties controls the run properties.
type RunProperties struct {
	x *dml.CT_TextCharacterProperties
}

// MakeRunProperties constructs a new RunProperties wrapper.
func MakeRunProperties(x *dml.CT_TextCharacterProperties) RunProperties {
	return RunProperties{x}
}

// SetSize sets the font size of the run text
func (r RunProperties) SetSize(sz measurement.Distance) {
	r.x.SzAttr = gooxml.Int32(int32(sz / measurement.HundredthPoint))
}

// SetBold controls the bolding of a run.
func (r RunProperties) SetBold(b bool) {
	r.x.BAttr = gooxml.Bool(b)
}

// SetSolidFill controls the text color of a run.
func (r RunProperties) SetSolidFill(c color.Color) {
	r.x.NoFill = nil
	r.x.BlipFill = nil
	r.x.GradFill = nil
	r.x.GrpFill = nil
	r.x.PattFill = nil
	r.x.SolidFill = dml.NewCT_SolidColorFillProperties()
	r.x.SolidFill.SrgbClr = dml.NewCT_SRgbColor()
	r.x.SolidFill.SrgbClr.ValAttr = *c.AsRGBString()
}

// SetFont controls the font of a run.
func (r RunProperties) SetFont(s string) {
	r.x.Latin = dml.NewCT_TextFont()
	r.x.Latin.TypefaceAttr = s
}
