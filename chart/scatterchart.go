// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package chart

import (
	crt "baliance.com/gooxml/schema/soo/dml/chart"
)

type ScatterChart struct {
	chartBase
	x *crt.CT_ScatterChart
}

// X returns the inner wrapped XML type.
func (c ScatterChart) X() *crt.CT_ScatterChart {
	return c.x
}

func (c ScatterChart) InitializeDefaults() {
	c.x.ScatterStyle.ValAttr = crt.ST_ScatterStyleMarker
}

// AddSeries adds a default series to a Scatter chart.
func (c ScatterChart) AddSeries() ScatterChartSeries {
	color := c.nextColor(len(c.x.Ser))
	ser := crt.NewCT_ScatterSer()
	c.x.Ser = append(c.x.Ser, ser)
	ser.Idx.ValAttr = uint32(len(c.x.Ser) - 1)
	ser.Order.ValAttr = uint32(len(c.x.Ser) - 1)

	ls := ScatterChartSeries{ser}
	ls.InitializeDefaults()
	ls.Marker().Properties().LineProperties().SetSolidFill(color)
	ls.Marker().Properties().SetSolidFill(color)
	return ls
}

// AddAxis adds an axis to a Scatter chart.
func (c ScatterChart) AddAxis(axis Axis) {
	axisID := crt.NewCT_UnsignedInt()
	axisID.ValAttr = axis.AxisID()
	c.x.AxId = append(c.x.AxId, axisID)
}
