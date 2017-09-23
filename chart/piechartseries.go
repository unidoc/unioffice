// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package chart

import (
	"baliance.com/gooxml/drawing"
	"baliance.com/gooxml/schema/soo/dml"
	crt "baliance.com/gooxml/schema/soo/dml/chart"
)

// PieChartSeries is a series to be used on an Pie chart.
type PieChartSeries struct {
	x *crt.CT_PieSer
}

// X returns the inner wrapped XML type.
func (c PieChartSeries) X() *crt.CT_PieSer {
	return c.x
}

// InitializeDefaults initializes an Pie series to the default values.
func (c PieChartSeries) InitializeDefaults() {
}

// SetExplosion sets the value that the segements of the pie are 'exploded' by
func (c PieChartSeries) SetExplosion(v uint32) {
	c.x.Explosion = crt.NewCT_UnsignedInt()
	c.x.Explosion.ValAttr = v
}

// SetText sets the series text.
func (c PieChartSeries) SetText(s string) {
	c.x.Tx = crt.NewCT_SerTx()
	c.x.Tx.Choice.V = &s
}

// CategoryAxis returns the category data source.
func (c PieChartSeries) CategoryAxis() CategoryAxisDataSource {
	if c.x.Cat == nil {
		c.x.Cat = crt.NewCT_AxDataSource()
	}
	return MakeAxisDataSource(c.x.Cat)
}

// Values returns the value data source.
func (c PieChartSeries) Values() NumberDataSource {
	if c.x.Val == nil {
		c.x.Val = crt.NewCT_NumDataSource()
	}
	return MakeNumberDataSource(c.x.Val)
}

// Properties returns the bar chart series shape properties.
func (c PieChartSeries) Properties() drawing.ShapeProperties {
	if c.x.SpPr == nil {
		c.x.SpPr = dml.NewCT_ShapeProperties()
	}
	return drawing.MakeShapeProperties(c.x.SpPr)
}
