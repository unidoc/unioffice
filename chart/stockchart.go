// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package chart

import (
	"github.com/unidoc/unioffice"
	crt "github.com/unidoc/unioffice/schema/soo/dml/chart"
)

// StockChart is a 2D Stock chart.
type StockChart struct {
	chartBase
	x *crt.CT_StockChart
}

// X returns the inner wrapped XML type.
func (c StockChart) X() *crt.CT_StockChart {
	return c.x
}

// InitializeDefaults the Stock chart to its defaults
func (c StockChart) InitializeDefaults() {
	c.x.HiLowLines = crt.NewCT_ChartLines()
	c.x.UpDownBars = crt.NewCT_UpDownBars()
	c.x.UpDownBars.GapWidth = crt.NewCT_GapAmount()
	c.x.UpDownBars.GapWidth.ValAttr = &crt.ST_GapAmount{}
	c.x.UpDownBars.GapWidth.ValAttr.ST_GapAmountUShort = unioffice.Uint16(150)
	c.x.UpDownBars.UpBars = crt.NewCT_UpDownBar()
	c.x.UpDownBars.DownBars = crt.NewCT_UpDownBar()
}

// AddSeries adds a default series to a Stock chart.
func (c StockChart) AddSeries() LineChartSeries {
	ser := crt.NewCT_LineSer()
	c.x.Ser = append(c.x.Ser, ser)
	ser.Idx.ValAttr = uint32(len(c.x.Ser) - 1)
	ser.Order.ValAttr = uint32(len(c.x.Ser) - 1)

	bs := LineChartSeries{ser}
	bs.Values().CreateEmptyNumberCache()
	// don't use defaults as the stock chart needs special
	// formatting
	bs.Properties().LineProperties().SetNoFill()
	return bs
}

func (c StockChart) AddAxis(axis Axis) {
	axisID := crt.NewCT_UnsignedInt()
	axisID.ValAttr = axis.AxisID()
	c.x.AxId = append(c.x.AxId, axisID)
}
