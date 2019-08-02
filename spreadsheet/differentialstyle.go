// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

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
