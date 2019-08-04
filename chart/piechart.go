// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package chart

import crt "github.com/unidoc/unioffice/schema/soo/dml/chart"
import "github.com/unidoc/unioffice"

// PieChart is a Pie chart.
type PieChart struct {
	chartBase
	x *crt.CT_PieChart
}

// X returns the inner wrapped XML type.
func (c PieChart) X() *crt.CT_PieChart {
	return c.x
}

// InitializeDefaults the bar chart to its defaults
func (c PieChart) InitializeDefaults() {
	c.x.VaryColors = crt.NewCT_Boolean()
	c.x.VaryColors.ValAttr = unioffice.Bool(true)

}

// AddSeries adds a default series to an Pie chart.
func (c PieChart) AddSeries() PieChartSeries {
	ser := crt.NewCT_PieSer()
	c.x.Ser = append(c.x.Ser, ser)
	ser.Idx.ValAttr = uint32(len(c.x.Ser) - 1)
	ser.Order.ValAttr = uint32(len(c.x.Ser) - 1)

	bs := PieChartSeries{ser}
	bs.InitializeDefaults()
	return bs
}
