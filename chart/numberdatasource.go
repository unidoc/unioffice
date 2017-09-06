// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package chart

import (
	"fmt"

	crt "baliance.com/gooxml/schema/schemas.openxmlformats.org/drawingml/2006/chart"
)

type NumberDataSource struct {
	x *crt.CT_NumDataSource
}

func MakeNumberDataSource(x *crt.CT_NumDataSource) NumberDataSource {
	return NumberDataSource{x}
}

func (n NumberDataSource) SetReference(s string) {
	n.x.Choice = crt.NewCT_NumDataSourceChoice()
	n.x.Choice.NumRef = crt.NewCT_NumRef()
	n.x.Choice.NumRef.F = s
}

func (n NumberDataSource) SetValues(v []float64) {
	n.x.Choice = crt.NewCT_NumDataSourceChoice()
	n.x.Choice.NumLit = crt.NewCT_NumData()
	n.x.Choice.NumLit.PtCount = crt.NewCT_UnsignedInt()
	n.x.Choice.NumLit.PtCount.ValAttr = uint32(len(v))

	for i, x := range v {
		n.x.Choice.NumLit.Pt = append(n.x.Choice.NumLit.Pt,
			&crt.CT_NumVal{
				IdxAttr: uint32(i),
				V:       fmt.Sprintf("%g", x)})
	}

}
