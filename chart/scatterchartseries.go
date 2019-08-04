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

// ScatterChartSeries is the data series for a scatter chart.
type ScatterChartSeries struct {
	x *crt.CT_ScatterSer
}

// X returns the inner wrapped XML type.
func (c ScatterChartSeries) X() *crt.CT_ScatterSer {
	return c.x
}

// Index returns the index of the series
func (c ScatterChartSeries) Index() uint32 {
	return c.x.Idx.ValAttr
}

// SetIndex sets the index of the series
func (c ScatterChartSeries) SetIndex(idx uint32) {
	c.x.Idx.ValAttr = idx
}

// Order returns the order of the series
func (c ScatterChartSeries) Order() uint32 {
	return c.x.Order.ValAttr
}

// SetOrder sets the order of the series
func (c ScatterChartSeries) SetOrder(idx uint32) {
	c.x.Order.ValAttr = idx
}

// SetText sets the series text
func (c ScatterChartSeries) SetText(s string) {
	c.x.Tx = crt.NewCT_SerTx()
	c.x.Tx.Choice.V = &s
}

// Properties returns the line chart series shape properties.
func (c ScatterChartSeries) Properties() drawing.ShapeProperties {
	if c.x.SpPr == nil {
		c.x.SpPr = dml.NewCT_ShapeProperties()
	}
	return drawing.MakeShapeProperties(c.x.SpPr)
}

// Marker returns the marker properties.
func (c ScatterChartSeries) Marker() Marker {
	if c.x.Marker == nil {
		c.x.Marker = crt.NewCT_Marker()
	}
	return MakeMarker(c.x.Marker)
}

// Labels returns the data label properties.
func (c ScatterChartSeries) Labels() DataLabels {
	if c.x.DLbls == nil {
		c.x.DLbls = crt.NewCT_DLbls()
	}
	return MakeDataLabels(c.x.DLbls)
}

func (c ScatterChartSeries) CategoryAxis() CategoryAxisDataSource {
	if c.x.XVal == nil {
		c.x.XVal = crt.NewCT_AxDataSource()
	}
	return MakeAxisDataSource(c.x.XVal)
}

func (c ScatterChartSeries) Values() NumberDataSource {
	if c.x.YVal == nil {
		c.x.YVal = crt.NewCT_NumDataSource()
	}
	return MakeNumberDataSource(c.x.YVal)
}

func (c ScatterChartSeries) SetSmooth(b bool) {
	c.x.Smooth = crt.NewCT_Boolean()
	c.x.Smooth.ValAttr = &b
}

func (c ScatterChartSeries) InitializeDefaults() {
	// turn off the line
	c.Properties().LineProperties().SetNoFill()
	c.Marker().SetSymbol(crt.ST_MarkerStyleAuto)
	c.Labels().SetShowLegendKey(false)
	c.Labels().SetShowValue(true)
	c.Labels().SetShowPercent(false)
	c.Labels().SetShowCategoryName(false)
	c.Labels().SetShowSeriesName(false)
	c.Labels().SetShowLeaderLines(false)
}
