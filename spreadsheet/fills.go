// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io.

package spreadsheet

import (
	"github.com/unidoc/unioffice"
	"github.com/unidoc/unioffice/schema/soo/sml"
)

type Fills struct {
	x *sml.CT_Fills
}

func NewFills() Fills {
	return Fills{sml.NewCT_Fills()}
}

func (f Fills) X() *sml.CT_Fills {
	return f.x
}

func (f Fills) AddFill() Fill {
	fill := sml.NewCT_Fill()
	f.x.Fill = append(f.x.Fill, fill)
	f.x.CountAttr = unioffice.Uint32(uint32(len(f.x.Fill)))
	return Fill{fill, f.x}
}
