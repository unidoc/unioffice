// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package chart

import (
	crt "github.com/unidoc/unioffice/schema/soo/dml/chart"
)

// CategoryAxisDataSource specifies the data for an axis.  It's commonly used with
// SetReference to set the axis data to a range of cells.
type CategoryAxisDataSource struct {
	x *crt.CT_AxDataSource
}

// MakeAxisDataSource constructs an AxisDataSource wrapper.
func MakeAxisDataSource(x *crt.CT_AxDataSource) CategoryAxisDataSource {
	return CategoryAxisDataSource{x}
}

// SetLabelReference is used to set the source data to a range of cells
// containing strings.
func (a CategoryAxisDataSource) SetLabelReference(s string) {
	a.x.Choice = crt.NewCT_AxDataSourceChoice()
	a.x.Choice.StrRef = crt.NewCT_StrRef()
	a.x.Choice.StrRef.F = s
}

// SetNumberReference is used to set the source data to a range of cells containing
// numbers.
func (a CategoryAxisDataSource) SetNumberReference(s string) {
	a.x.Choice = crt.NewCT_AxDataSourceChoice()
	a.x.Choice.NumRef = crt.NewCT_NumRef()
	a.x.Choice.NumRef.F = s
}

// SetValues is used to set the source data to a set of values.
func (a CategoryAxisDataSource) SetValues(v []string) {
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
