// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io.

package chart

import (
	"github.com/unidoc/unioffice/drawing"
	"github.com/unidoc/unioffice/schema/soo/dml"
	crt "github.com/unidoc/unioffice/schema/soo/dml/chart"
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
