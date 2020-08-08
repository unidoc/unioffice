// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io.

package spreadsheet

import "github.com/unidoc/unioffice/schema/soo/sml"

type DifferentialStyle struct {
	x   *sml.CT_Dxf
	wb  *Workbook
	dxf *sml.CT_Dxfs
}

// X returns the inner wrapped XML type.
func (d DifferentialStyle) X() *sml.CT_Dxf {
	return d.x
}

// Index returns the index of the differential style.
func (d DifferentialStyle) Index() uint32 {
	for i, dxf := range d.dxf.Dxf {
		if d.x == dxf {
			return uint32(i)
		}
	}
	return 0
}

func (d DifferentialStyle) Fill() Fill {
	if d.x.Fill == nil {
		d.x.Fill = sml.NewCT_Fill()
	}
	return Fill{d.x.Fill, nil}
}
