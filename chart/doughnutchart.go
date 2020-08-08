// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io.

package chart

import crt "github.com/unidoc/unioffice/schema/soo/dml/chart"
import "github.com/unidoc/unioffice"

// DoughnutChart is a Doughnut chart.
type DoughnutChart struct {
	chartBase
	x *crt.CT_DoughnutChart
}

// X returns the inner wrapped XML type.
func (c DoughnutChart) X() *crt.CT_DoughnutChart {
	return c.x
}

// InitializeDefaults the bar chart to its defaults
func (c DoughnutChart) InitializeDefaults() {
	c.x.VaryColors = crt.NewCT_Boolean()
	c.x.VaryColors.ValAttr = unioffice.Bool(true)
	c.x.HoleSize = crt.NewCT_HoleSize()
	c.x.HoleSize.ValAttr = &crt.ST_HoleSize{}
	c.x.HoleSize.ValAttr.ST_HoleSizeUByte = unioffice.Uint8(50)
}

// SetHoleSize controls the hole size in the pie chart and is measured in percent.
func (c DoughnutChart) SetHoleSize(pct uint8) {
	if c.x.HoleSize == nil {
		c.x.HoleSize = crt.NewCT_HoleSize()
	}
	if c.x.HoleSize.ValAttr == nil {
		c.x.HoleSize.ValAttr = &crt.ST_HoleSize{}
	}
	c.x.HoleSize.ValAttr.ST_HoleSizeUByte = &pct
}

// AddSeries adds a default series to an Doughnut chart.
func (c DoughnutChart) AddSeries() PieChartSeries {
	ser := crt.NewCT_PieSer()
	c.x.Ser = append(c.x.Ser, ser)
	ser.Idx.ValAttr = uint32(len(c.x.Ser) - 1)
	ser.Order.ValAttr = uint32(len(c.x.Ser) - 1)

	bs := PieChartSeries{ser}
	bs.InitializeDefaults()
	return bs
}
