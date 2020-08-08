// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io.

package chart

import (
	"github.com/unidoc/unioffice"
	"github.com/unidoc/unioffice/drawing"
	"github.com/unidoc/unioffice/schema/soo/dml"
	crt "github.com/unidoc/unioffice/schema/soo/dml/chart"
)

// Surface3DChart is a 3D view of a surface chart.
type Surface3DChart struct {
	chartBase
	x *crt.CT_Surface3DChart
}

// X returns the inner wrapped XML type.
func (c Surface3DChart) X() *crt.CT_Surface3DChart {
	return c.x
}

func (c Surface3DChart) InitializeDefaults() {
	c.x.Wireframe = crt.NewCT_Boolean()
	c.x.Wireframe.ValAttr = unioffice.Bool(false)

	c.x.BandFmts = crt.NewCT_BandFmts()
	for i := 0; i < 15; i++ {
		bfmt := crt.NewCT_BandFmt()
		bfmt.Idx.ValAttr = uint32(i)
		bfmt.SpPr = dml.NewCT_ShapeProperties()

		sp := drawing.MakeShapeProperties(bfmt.SpPr)
		sp.SetSolidFill(c.nextColor(i))
		c.x.BandFmts.BandFmt = append(c.x.BandFmts.BandFmt, bfmt)
	}
}

// AddSeries adds a default series to a Surface chart.
func (c Surface3DChart) AddSeries() SurfaceChartSeries {
	color := c.nextColor(len(c.x.Ser))
	ser := crt.NewCT_SurfaceSer()
	c.x.Ser = append(c.x.Ser, ser)
	ser.Idx.ValAttr = uint32(len(c.x.Ser) - 1)
	ser.Order.ValAttr = uint32(len(c.x.Ser) - 1)

	ls := SurfaceChartSeries{ser}
	ls.InitializeDefaults()
	ls.Properties().LineProperties().SetSolidFill(color)
	return ls
}

// AddAxis adds an axis to a Surface chart.
func (c Surface3DChart) AddAxis(axis Axis) {
	axisID := crt.NewCT_UnsignedInt()
	axisID.ValAttr = axis.AxisID()
	c.x.AxId = append(c.x.AxId, axisID)
}
