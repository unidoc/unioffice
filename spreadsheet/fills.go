// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

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
