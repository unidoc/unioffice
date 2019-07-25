// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package chart

import (
	"github.com/unidoc/unioffice"
	"github.com/unidoc/unioffice/drawing"
	"github.com/unidoc/unioffice/schema/soo/dml"
	crt "github.com/unidoc/unioffice/schema/soo/dml/chart"
)

type Legend struct {
	x *crt.CT_Legend
}

func MakeLegend(l *crt.CT_Legend) Legend {
	return Legend{l}
}

// X returns the inner wrapped XML type.
func (l Legend) X() *crt.CT_Legend {
	return l.x
}
func (l Legend) InitializeDefaults() {
	l.SetPosition(crt.ST_LegendPosR)
	l.SetOverlay(false)
	l.Properties().SetNoFill()
	l.Properties().LineProperties().SetNoFill()
}

func (l Legend) SetPosition(p crt.ST_LegendPos) {
	if p == crt.ST_LegendPosUnset {
		l.x.LegendPos = nil
	} else {
		l.x.LegendPos = crt.NewCT_LegendPos()
		l.x.LegendPos.ValAttr = p
	}
}

func (l Legend) SetOverlay(b bool) {
	l.x.Overlay = crt.NewCT_Boolean()
	l.x.Overlay.ValAttr = unioffice.Bool(b)
}

func (l Legend) Properties() drawing.ShapeProperties {
	if l.x.SpPr == nil {
		l.x.SpPr = dml.NewCT_ShapeProperties()
	}
	return drawing.MakeShapeProperties(l.x.SpPr)
}
