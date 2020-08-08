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
