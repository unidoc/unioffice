// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package chart

import crt "baliance.com/gooxml/schema/soo/dml/chart"

// Area3DChart is an area chart that has a shaded area underneath a curve.
type Area3DChart struct {
	chartBase
	x *crt.CT_Area3DChart
}

// X returns the inner wrapped XML type.
func (c Area3DChart) X() *crt.CT_Area3DChart {
	return c.x
}

// InitializeDefaults the bar chart to its defaults
func (c Area3DChart) InitializeDefaults() {
}

// AddSeries adds a default series to an area chart.
func (c Area3DChart) AddSeries() AreaChartSeries {
	clr := c.nextColor(len(c.x.Ser))
	ser := crt.NewCT_AreaSer()
	c.x.Ser = append(c.x.Ser, ser)
	ser.Idx.ValAttr = uint32(len(c.x.Ser) - 1)
	ser.Order.ValAttr = uint32(len(c.x.Ser) - 1)

	bs := AreaChartSeries{ser}
	bs.InitializeDefaults()
	bs.Properties().SetSolidFill(clr)
	return bs
}

func (c Area3DChart) AddAxis(axis Axis) {
	axisID := crt.NewCT_UnsignedInt()
	axisID.ValAttr = axis.AxisID()
	c.x.AxId = append(c.x.AxId, axisID)
}
