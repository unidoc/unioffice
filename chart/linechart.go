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

type LineChart struct {
	chartBase
	x *crt.CT_LineChart
}

// X returns the inner wrapped XML type.
func (c LineChart) X() *crt.CT_LineChart {
	return c.x
}

// AddSeries adds a default series to a line chart.
func (c LineChart) AddSeries() LineChartSeries {
	color := c.nextColor(len(c.x.Ser))
	ser := crt.NewCT_LineSer()
	c.x.Ser = append(c.x.Ser, ser)
	ser.Idx.ValAttr = uint32(len(c.x.Ser) - 1)
	ser.Order.ValAttr = uint32(len(c.x.Ser) - 1)

	ls := LineChartSeries{ser}
	ls.InitializeDefaults()
	ls.Properties().LineProperties().SetSolidFill(color)
	return ls
}

// AddAxis adds an axis to a line chart.
func (c LineChart) AddAxis(axis Axis) {
	axisID := crt.NewCT_UnsignedInt()
	axisID.ValAttr = axis.AxisID()
	c.x.AxId = append(c.x.AxId, axisID)
}
