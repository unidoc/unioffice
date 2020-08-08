// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io.

package chart

import (
	crt "github.com/unidoc/unioffice/schema/soo/dml/chart"
)

// BubbleChart is a 2D Bubble chart.
type BubbleChart struct {
	chartBase
	x *crt.CT_BubbleChart
}

// X returns the inner wrapped XML type.
func (c BubbleChart) X() *crt.CT_BubbleChart {
	return c.x
}

// InitializeDefaults the Bubble chart to its defaults
func (c BubbleChart) InitializeDefaults() {

}

// AddSeries adds a default series to a Bubble chart.
func (c BubbleChart) AddSeries() BubbleChartSeries {
	clr := c.nextColor(len(c.x.Ser))
	ser := crt.NewCT_BubbleSer()
	c.x.Ser = append(c.x.Ser, ser)
	ser.Idx.ValAttr = uint32(len(c.x.Ser) - 1)
	ser.Order.ValAttr = uint32(len(c.x.Ser) - 1)

	bs := BubbleChartSeries{ser}
	bs.InitializeDefaults()
	bs.Properties().SetSolidFill(clr)
	return bs
}

func (c BubbleChart) AddAxis(axis Axis) {
	axisID := crt.NewCT_UnsignedInt()
	axisID.ValAttr = axis.AxisID()
	c.x.AxId = append(c.x.AxId, axisID)
}
