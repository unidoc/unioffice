// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package styles

import sml "baliance.com/gooxml/schema/schemas.openxmlformats.org/spreadsheetml"

type Borders struct {
	x *sml.CT_Borders
}

func NewBorders() Borders {
	b := Borders{sml.NewCT_Borders()}
	cnt := uint32(1)
	b.x.CountAttr = &cnt
	r := sml.NewCT_Border()
	r.Left = sml.NewCT_BorderPr()
	r.Bottom = sml.NewCT_BorderPr()
	r.Right = sml.NewCT_BorderPr()
	r.Top = sml.NewCT_BorderPr()
	r.Diagonal = sml.NewCT_BorderPr()
	b.x.Border = append(b.x.Border, r)
	return b
}

func (b Borders) X() *sml.CT_Borders {
	return b.x
}
