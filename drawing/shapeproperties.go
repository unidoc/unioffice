// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package drawing

import (
	"baliance.com/gooxml/color"

	"baliance.com/gooxml/schema/soo/dml"
)

type ShapeProperties struct {
	x *dml.CT_ShapeProperties
}

func MakeShapeProperties(x *dml.CT_ShapeProperties) ShapeProperties {
	return ShapeProperties{x}
}

// X returns the inner wrapped XML type.
func (s ShapeProperties) X() *dml.CT_ShapeProperties {
	return s.x
}

func (s ShapeProperties) clearFill() {
	s.x.NoFill = nil
	s.x.BlipFill = nil
	s.x.GradFill = nil
	s.x.GrpFill = nil
	s.x.SolidFill = nil
	s.x.PattFill = nil
}

func (s ShapeProperties) SetNoFill() {
	s.clearFill()
	s.x.NoFill = dml.NewCT_NoFillProperties()
}

func (s ShapeProperties) SetSolidFill(c color.Color) {
	s.clearFill()
	s.x.SolidFill = dml.NewCT_SolidColorFillProperties()
	s.x.SolidFill.SrgbClr = dml.NewCT_SRgbColor()
	s.x.SolidFill.SrgbClr.ValAttr = *c.AsRGBString()
}

func (s ShapeProperties) LineProperties() LineProperties {
	if s.x.Ln == nil {
		s.x.Ln = dml.NewCT_LineProperties()
	}
	return LineProperties{s.x.Ln}
}
