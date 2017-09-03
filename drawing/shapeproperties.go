// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package drawing

import (
	"baliance.com/gooxml/color"

	dml "baliance.com/gooxml/schema/schemas.openxmlformats.org/drawingml"
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

func (s ShapeProperties) SetSolidFill(c color.Color) {
	s.x.SolidFill = dml.NewCT_SolidColorFillProperties()
	s.x.SolidFill.SrgbClr = dml.NewCT_SRgbColor()
	s.x.SolidFill.SrgbClr.ValAttr = *c.AsRGBAString()
}

func (s ShapeProperties) LineProperties() LineProperties {
	if s.x.Ln == nil {
		s.x.Ln = dml.NewCT_LineProperties()
	}
	return LineProperties{s.x.Ln}
}
