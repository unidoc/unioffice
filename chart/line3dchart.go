// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package chart

import (
	crt "github.com/unidoc/unioffice/schema/soo/dml/chart"
)

type Line3DChart struct {
	chartBase
	x *crt.CT_Line3DChart
}

// X returns the inner wrapped XML type.
func (c Line3DChart) X() *crt.CT_Line3DChart {
	return c.x
}

// AddSeries adds a default series to a line chart.
func (c Line3DChart) AddSeries() LineChartSeries {
	color := c.nextColor(len(c.x.Ser))
	ser := crt.NewCT_LineSer()
	c.x.Ser = append(c.x.Ser, ser)
	ser.Idx.ValAttr = uint32(len(c.x.Ser) - 1)
	ser.Order.ValAttr = uint32(len(c.x.Ser) - 1)

	ls := LineChartSeries{ser}
	ls.InitializeDefaults()
	ls.Properties().LineProperties().SetSolidFill(color)
	ls.Properties().SetSolidFill(color)
	return ls
}

// AddAxis adds an axis to a line chart.
func (c Line3DChart) AddAxis(axis Axis) {
	axisID := crt.NewCT_UnsignedInt()
	axisID.ValAttr = axis.AxisID()
	c.x.AxId = append(c.x.AxId, axisID)
}
