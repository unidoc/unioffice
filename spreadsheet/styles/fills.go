// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package styles

import sml "baliance.com/gooxml/schema/schemas.openxmlformats.org/spreadsheetml"

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
	fill := NewPatternFill(f.x)
	fill.SetPattern(sml.ST_PatternTypeSolid)
	f.x.Fill = append(f.x.Fill, fill.X())
	v := uint32(len(f.x.Fill))
	f.x.CountAttr = &v
	return fill
}
