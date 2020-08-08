// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io.

package chart

import crt "github.com/unidoc/unioffice/schema/soo/dml/chart"

// RadarChart is an Radar chart that has a shaded Radar underneath a curve.
type RadarChart struct {
	chartBase
	x *crt.CT_RadarChart
}

// X returns the inner wrapped XML type.
func (c RadarChart) X() *crt.CT_RadarChart {
	return c.x
}

// InitializeDefaults the bar chart to its defaults
func (c RadarChart) InitializeDefaults() {
	c.x.RadarStyle.ValAttr = crt.ST_RadarStyleMarker
}

// AddSeries adds a default series to an Radar chart.
func (c RadarChart) AddSeries() RadarChartSeries {
	clr := c.nextColor(len(c.x.Ser))
	ser := crt.NewCT_RadarSer()
	c.x.Ser = append(c.x.Ser, ser)
	ser.Idx.ValAttr = uint32(len(c.x.Ser) - 1)
	ser.Order.ValAttr = uint32(len(c.x.Ser) - 1)

	bs := RadarChartSeries{ser}
	bs.InitializeDefaults()
	bs.Properties().SetSolidFill(clr)
	return bs
}

func (c RadarChart) AddAxis(axis Axis) {
	axisID := crt.NewCT_UnsignedInt()
	axisID.ValAttr = axis.AxisID()
	c.x.AxId = append(c.x.AxId, axisID)
}
