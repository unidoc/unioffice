// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package chart

import (
	"github.com/unidoc/unioffice/color"
	"github.com/unidoc/unioffice/drawing"
	"github.com/unidoc/unioffice/measurement"
	"github.com/unidoc/unioffice/schema/soo/dml"
	crt "github.com/unidoc/unioffice/schema/soo/dml/chart"
)

// LineChartSeries is the data series for a line chart.
type LineChartSeries struct {
	x *crt.CT_LineSer
}

// X returns the inner wrapped XML type.
func (c LineChartSeries) X() *crt.CT_LineSer {
	return c.x
}

// Index returns the index of the series
func (c LineChartSeries) Index() uint32 {
	return c.x.Idx.ValAttr
}

// SetIndex sets the index of the series
func (c LineChartSeries) SetIndex(idx uint32) {
	c.x.Idx.ValAttr = idx
}

// Order returns the order of the series
func (c LineChartSeries) Order() uint32 {
	return c.x.Order.ValAttr
}

// SetOrder sets the order of the series
func (c LineChartSeries) SetOrder(idx uint32) {
	c.x.Order.ValAttr = idx
}

// SetText sets the series text
func (c LineChartSeries) SetText(s string) {
	c.x.Tx = crt.NewCT_SerTx()
	c.x.Tx.Choice.V = &s
}

// Properties returns the line chart series shape properties.
func (c LineChartSeries) Properties() drawing.ShapeProperties {
	if c.x.SpPr == nil {
		c.x.SpPr = dml.NewCT_ShapeProperties()
	}
	return drawing.MakeShapeProperties(c.x.SpPr)
}

// Marker returns the marker properties.
func (c LineChartSeries) Marker() Marker {
	if c.x.Marker == nil {
		c.x.Marker = crt.NewCT_Marker()
	}
	return MakeMarker(c.x.Marker)
}

// Labels returns the data label properties.
func (c LineChartSeries) Labels() DataLabels {
	if c.x.DLbls == nil {
		c.x.DLbls = crt.NewCT_DLbls()
	}
	return MakeDataLabels(c.x.DLbls)
}

func (c LineChartSeries) CategoryAxis() CategoryAxisDataSource {
	if c.x.Cat == nil {
		c.x.Cat = crt.NewCT_AxDataSource()
	}
	return MakeAxisDataSource(c.x.Cat)
}

func (c LineChartSeries) Values() NumberDataSource {
	if c.x.Val == nil {
		c.x.Val = crt.NewCT_NumDataSource()
	}
	return MakeNumberDataSource(c.x.Val)
}

func (c LineChartSeries) SetSmooth(b bool) {
	c.x.Smooth = crt.NewCT_Boolean()
	c.x.Smooth.ValAttr = &b
}

func (c LineChartSeries) InitializeDefaults() {
	c.Properties().LineProperties().SetWidth(1 * measurement.Point)
	c.Properties().LineProperties().SetSolidFill(color.Black)
	c.Properties().LineProperties().SetJoin(drawing.LineJoinRound)

	c.Marker().SetSymbol(crt.ST_MarkerStyleNone)
	// don't set position by default as it breaks line 3d chart in Excel
	c.Labels().SetShowLegendKey(false)
	c.Labels().SetShowValue(false)
	c.Labels().SetShowPercent(false)
	c.Labels().SetShowCategoryName(false)
	c.Labels().SetShowSeriesName(false)
	c.Labels().SetShowLeaderLines(false)
}
