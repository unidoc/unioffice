// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheet

import (
	"baliance.com/gooxml/color"
	"baliance.com/gooxml/drawing"
	"baliance.com/gooxml/measurement"
	crt "baliance.com/gooxml/schema/schemas.openxmlformats.org/drawingml/2006/chart"
)

type LineChart struct {
	x *crt.CT_LineChart
}

// X returns the inner wrapped XML type.
func (c LineChart) X() *crt.CT_LineChart {
	return c.x
}

func (c LineChart) AddSeries() LineChartSeries {
	ser := crt.NewCT_LineSer()
	c.x.Ser = append(c.x.Ser, ser)
	ser.Idx.ValAttr = uint32(len(c.x.Ser))
	ser.Order.ValAttr = uint32(len(c.x.Ser))
	ls := LineChartSeries{ser}
	ls.Properties().LineProperties().SetWidth(2 * measurement.Point)
	ls.Properties().LineProperties().SetSolidFill(color.Green)
	ls.Properties().LineProperties().SetJoin(drawing.LineJoinRound)
	return ls
}
