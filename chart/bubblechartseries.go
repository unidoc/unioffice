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

// BubbleChartSeries is a series to be used on a Bubble chart.
type BubbleChartSeries struct {
	x *crt.CT_BubbleSer
}

// X returns the inner wrapped XML type.
func (c BubbleChartSeries) X() *crt.CT_BubbleSer {
	return c.x
}

// InitializeDefaults initializes a Bubble chart series to the default values.
func (c BubbleChartSeries) InitializeDefaults() {
}

// SetText sets the series text.
func (c BubbleChartSeries) SetText(s string) {
	c.x.Tx = crt.NewCT_SerTx()
	c.x.Tx.Choice.V = &s
}

// CategoryAxis returns the category data source.
func (c BubbleChartSeries) CategoryAxis() CategoryAxisDataSource {
	if c.x.XVal == nil {
		c.x.XVal = crt.NewCT_AxDataSource()
	}
	return MakeAxisDataSource(c.x.XVal)
}

// Values returns the value data source.
func (c BubbleChartSeries) Values() NumberDataSource {
	if c.x.YVal == nil {
		c.x.YVal = crt.NewCT_NumDataSource()
	}
	return MakeNumberDataSource(c.x.YVal)
}

// Values returns the bubble size data source.
func (c BubbleChartSeries) BubbleSizes() NumberDataSource {
	if c.x.BubbleSize == nil {
		c.x.BubbleSize = crt.NewCT_NumDataSource()
	}
	return MakeNumberDataSource(c.x.BubbleSize)
}

// Properties returns the Bubble chart series shape properties.
func (c BubbleChartSeries) Properties() drawing.ShapeProperties {
	if c.x.SpPr == nil {
		c.x.SpPr = dml.NewCT_ShapeProperties()
	}
	return drawing.MakeShapeProperties(c.x.SpPr)
}
