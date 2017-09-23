// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package chart

import (
	"baliance.com/gooxml/color"
	"baliance.com/gooxml/drawing"
	"baliance.com/gooxml/schema/soo/dml"
	crt "baliance.com/gooxml/schema/soo/dml/chart"
)

type CategoryAxis struct {
	x *crt.CT_CatAx
}

func MakeCategoryAxis(x *crt.CT_CatAx) CategoryAxis {
	return CategoryAxis{x}
}

func (c CategoryAxis) MajorGridLines() GridLines {
	if c.x.MajorGridlines == nil {
		c.x.MajorGridlines = crt.NewCT_ChartLines()
	}
	return GridLines{c.x.MajorGridlines}
}

func (c CategoryAxis) AxisID() uint32 {
	return c.x.AxId.ValAttr
}

func (c CategoryAxis) Properties() drawing.ShapeProperties {
	if c.x.SpPr == nil {
		c.x.SpPr = dml.NewCT_ShapeProperties()
	}
	return drawing.MakeShapeProperties(c.x.SpPr)
}

func (c CategoryAxis) SetCrosses(axis Axis) {
	c.x.Choice = crt.NewEG_AxSharedChoice()

	// TODO: remove this?
	c.x.Choice.Crosses = crt.NewCT_Crosses()
	c.x.Choice.Crosses.ValAttr = crt.ST_CrossesAutoZero

	c.x.CrossAx.ValAttr = axis.AxisID()
}

func (c CategoryAxis) InitializeDefaults() {
	c.SetPosition(crt.ST_AxPosB)
	c.SetMajorTickMark(crt.ST_TickMarkOut)
	c.SetMinorTickMark(crt.ST_TickMarkIn)
	c.SetTickLabelPosition(crt.ST_TickLblPosNextTo)
	c.MajorGridLines().Properties().LineProperties().SetSolidFill(color.LightGray)
	c.Properties().LineProperties().SetSolidFill(color.Black)
}

func (c CategoryAxis) SetPosition(p crt.ST_AxPos) {
	c.x.AxPos = crt.NewCT_AxPos()
	c.x.AxPos.ValAttr = p
}

func (c CategoryAxis) SetMajorTickMark(m crt.ST_TickMark) {
	if m == crt.ST_TickMarkUnset {
		c.x.MajorTickMark = nil
	} else {
		c.x.MajorTickMark = crt.NewCT_TickMark()
		c.x.MajorTickMark.ValAttr = m
	}
}

func (c CategoryAxis) SetMinorTickMark(m crt.ST_TickMark) {
	if m == crt.ST_TickMarkUnset {
		c.x.MinorTickMark = nil
	} else {
		c.x.MinorTickMark = crt.NewCT_TickMark()
		c.x.MinorTickMark.ValAttr = m
	}

}

func (c CategoryAxis) SetTickLabelPosition(p crt.ST_TickLblPos) {
	if p == crt.ST_TickLblPosUnset {
		c.x.TickLblPos = nil
	} else {
		c.x.TickLblPos = crt.NewCT_TickLblPos()
		c.x.TickLblPos.ValAttr = p
	}
}
