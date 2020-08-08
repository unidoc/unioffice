// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io.

package chart

import crt "github.com/unidoc/unioffice/schema/soo/dml/chart"
import "github.com/unidoc/unioffice"

// Pie3DChart is a Pie3D chart.
type Pie3DChart struct {
	chartBase
	x *crt.CT_Pie3DChart
}

// X returns the inner wrapped XML type.
func (c Pie3DChart) X() *crt.CT_Pie3DChart {
	return c.x
}

// InitializeDefaults the bar chart to its defaults
func (c Pie3DChart) InitializeDefaults() {
	c.x.VaryColors = crt.NewCT_Boolean()
	c.x.VaryColors.ValAttr = unioffice.Bool(true)
}

// AddSeries adds a default series to an Pie3D chart.
func (c Pie3DChart) AddSeries() PieChartSeries {
	ser := crt.NewCT_PieSer()
	c.x.Ser = append(c.x.Ser, ser)
	ser.Idx.ValAttr = uint32(len(c.x.Ser) - 1)
	ser.Order.ValAttr = uint32(len(c.x.Ser) - 1)

	bs := PieChartSeries{ser}
	bs.InitializeDefaults()
	return bs
}
