// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheet

import (
	"baliance.com/gooxml"
	sml "baliance.com/gooxml/schema/schemas.openxmlformats.org/spreadsheetml"
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

func (f Fills) AddPatternFill() PatternFill {
	pfill := NewPatternFill(f.x)
	pfill.SetPattern(sml.ST_PatternTypeSolid)
	f.x.Fill = append(f.x.Fill, pfill.f)
	f.x.CountAttr = gooxml.Uint32(uint32(len(f.x.Fill)))
	return pfill
}
