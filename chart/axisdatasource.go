// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package chart

import (
	crt "baliance.com/gooxml/schema/schemas.openxmlformats.org/drawingml/2006/chart"
)

// AxisDataSource specifies the data for an axis.  It's commonly used with
// SetReference to set the axis data to a range of cells.
type AxisDataSource struct {
	x *crt.CT_AxDataSource
}

// MakeAxisDataSource constructs an AxisDataSource wrapper.
func MakeAxisDataSource(x *crt.CT_AxDataSource) AxisDataSource {
	return AxisDataSource{x}
}

func (a AxisDataSource) SetReference(s string) {
	a.x.Choice = crt.NewCT_AxDataSourceChoice()
	a.x.Choice.StrRef = crt.NewCT_StrRef()
	a.x.Choice.StrRef.F = s
}

func (a AxisDataSource) SetValues(v []string) {
	a.x.Choice = crt.NewCT_AxDataSourceChoice()
	a.x.Choice.StrLit = crt.NewCT_StrData()
	a.x.Choice.StrLit.PtCount = crt.NewCT_UnsignedInt()
	a.x.Choice.StrLit.PtCount.ValAttr = uint32(len(v))

	for i, x := range v {
		a.x.Choice.StrLit.Pt = append(a.x.Choice.StrLit.Pt,
			&crt.CT_StrVal{
				IdxAttr: uint32(i),
				V:       x})
	}

}
