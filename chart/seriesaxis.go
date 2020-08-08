// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io.

package chart

import crt "github.com/unidoc/unioffice/schema/soo/dml/chart"

type SeriesAxis struct {
	x *crt.CT_SerAx
}

func MakeSeriesAxis(x *crt.CT_SerAx) SeriesAxis {
	return SeriesAxis{x}
}

// X returns the inner wrapped XML type.
func (s SeriesAxis) X() *crt.CT_SerAx {
	return s.x
}

func (s SeriesAxis) InitializeDefaults() {

}

func (s SeriesAxis) AxisID() uint32 {
	return s.x.AxId.ValAttr
}

func (s SeriesAxis) SetCrosses(axis Axis) {
	s.x.CrossAx.ValAttr = axis.AxisID()
}
