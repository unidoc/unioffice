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

// Bar3DChart is a 3D bar chart.
type Bar3DChart struct {
	chartBase
	x *crt.CT_Bar3DChart
}

// X returns the inner wrapped XML type.
func (c Bar3DChart) X() *crt.CT_Bar3DChart {
	return c.x
}

// InitializeDefaults the bar chart to its defaults
func (c Bar3DChart) InitializeDefaults() {
	c.SetDirection(crt.ST_BarDirCol)
}

// SetDirection changes the direction of the bar chart (bar or column).
func (c Bar3DChart) SetDirection(d crt.ST_BarDir) {
	c.x.BarDir.ValAttr = d
}

// AddSeries adds a default series to a bar chart.
func (c Bar3DChart) AddSeries() BarChartSeries {
	clr := c.nextColor(len(c.x.Ser))
	ser := crt.NewCT_BarSer()
	c.x.Ser = append(c.x.Ser, ser)
	ser.Idx.ValAttr = uint32(len(c.x.Ser) - 1)
	ser.Order.ValAttr = uint32(len(c.x.Ser) - 1)

	bs := BarChartSeries{ser}
	bs.InitializeDefaults()
	bs.Properties().SetSolidFill(clr)
	return bs
}

func (c Bar3DChart) AddAxis(axis Axis) {
	axisID := crt.NewCT_UnsignedInt()
	axisID.ValAttr = axis.AxisID()
	c.x.AxId = append(c.x.AxId, axisID)
}
