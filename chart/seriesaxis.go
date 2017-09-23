// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package chart

import crt "baliance.com/gooxml/schema/soo/dml/chart"

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
