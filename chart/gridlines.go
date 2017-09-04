// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package chart

import (
	"baliance.com/gooxml/drawing"
	dml "baliance.com/gooxml/schema/schemas.openxmlformats.org/drawingml"
	crt "baliance.com/gooxml/schema/schemas.openxmlformats.org/drawingml/2006/chart"
)

type GridLines struct {
	x *crt.CT_ChartLines
}

// X returns the inner wrapped XML type.
func (g GridLines) X() *crt.CT_ChartLines {
	return g.x
}

func (g GridLines) Properties() drawing.ShapeProperties {
	if g.x.SpPr == nil {
		g.x.SpPr = dml.NewCT_ShapeProperties()
	}
	return drawing.MakeShapeProperties(g.x.SpPr)
}
