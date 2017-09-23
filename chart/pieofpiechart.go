// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package chart

import (
	"baliance.com/gooxml"
	"baliance.com/gooxml/color"
	"baliance.com/gooxml/drawing"
	"baliance.com/gooxml/schema/soo/dml"
	crt "baliance.com/gooxml/schema/soo/dml/chart"
)

// PieOfPieChart is a Pie chart with an extra Pie chart.
type PieOfPieChart struct {
	chartBase
	x *crt.CT_OfPieChart
}

// X returns the inner wrapped XML type.
func (c PieOfPieChart) X() *crt.CT_OfPieChart {
	return c.x
}

// InitializeDefaults the bar chart to its defaults
func (c PieOfPieChart) InitializeDefaults() {
	c.x.VaryColors = crt.NewCT_Boolean()
	c.x.VaryColors.ValAttr = gooxml.Bool(true)
	c.SetType(crt.ST_OfPieTypePie)
	c.x.SecondPieSize = crt.NewCT_SecondPieSize()
	c.x.SecondPieSize.ValAttr = &crt.ST_SecondPieSize{}
	c.x.SecondPieSize.ValAttr.ST_SecondPieSizeUShort = gooxml.Uint16(75)
	cl := crt.NewCT_ChartLines()
	cl.SpPr = dml.NewCT_ShapeProperties()
	sp := drawing.MakeShapeProperties(cl.SpPr)
	sp.LineProperties().SetSolidFill(color.Auto)
	c.x.SerLines = append(c.x.SerLines, cl)
}

// SetType sets the type the secone pie to either pie or bar
func (c PieOfPieChart) SetType(t crt.ST_OfPieType) {
	c.x.OfPieType.ValAttr = t
}

// AddSeries adds a default series to an Pie chart.
func (c PieOfPieChart) AddSeries() PieChartSeries {
	ser := crt.NewCT_PieSer()
	c.x.Ser = append(c.x.Ser, ser)
	ser.Idx.ValAttr = uint32(len(c.x.Ser) - 1)
	ser.Order.ValAttr = uint32(len(c.x.Ser) - 1)

	bs := PieChartSeries{ser}
	bs.InitializeDefaults()
	return bs
}
