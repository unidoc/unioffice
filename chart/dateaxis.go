// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package chart

import (
	"github.com/unidoc/unioffice/drawing"
	"github.com/unidoc/unioffice/schema/soo/dml"
	crt "github.com/unidoc/unioffice/schema/soo/dml/chart"
)

type DateAxis struct {
	x *crt.CT_DateAx
}

// X returns the inner wrapped XML type.
func (v DateAxis) X() *crt.CT_DateAx {
	return v.x
}

func (v DateAxis) AxisID() uint32 {
	return v.x.AxId.ValAttr
}

func (v DateAxis) SetPosition(p crt.ST_AxPos) {
	v.x.AxPos = crt.NewCT_AxPos()
	v.x.AxPos.ValAttr = p
}

func (v DateAxis) MajorGridLines() GridLines {
	if v.x.MajorGridlines == nil {
		v.x.MajorGridlines = crt.NewCT_ChartLines()
	}
	return GridLines{v.x.MajorGridlines}
}

func (v DateAxis) SetMajorTickMark(m crt.ST_TickMark) {
	if m == crt.ST_TickMarkUnset {
		v.x.MajorTickMark = nil
	} else {
		v.x.MajorTickMark = crt.NewCT_TickMark()
		v.x.MajorTickMark.ValAttr = m
	}
}

func (v DateAxis) SetMinorTickMark(m crt.ST_TickMark) {
	if m == crt.ST_TickMarkUnset {
		v.x.MinorTickMark = nil
	} else {
		v.x.MinorTickMark = crt.NewCT_TickMark()
		v.x.MinorTickMark.ValAttr = m
	}

}

func (v DateAxis) SetTickLabelPosition(p crt.ST_TickLblPos) {
	if p == crt.ST_TickLblPosUnset {
		v.x.TickLblPos = nil
	} else {
		v.x.TickLblPos = crt.NewCT_TickLblPos()
		v.x.TickLblPos.ValAttr = p
	}
}

func (v DateAxis) Properties() drawing.ShapeProperties {
	if v.x.SpPr == nil {
		v.x.SpPr = dml.NewCT_ShapeProperties()
	}
	return drawing.MakeShapeProperties(v.x.SpPr)
}

func (v DateAxis) SetCrosses(axis Axis) {
	v.x.CrossAx.ValAttr = axis.AxisID()
}
