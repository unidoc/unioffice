// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheet

import (
	"baliance.com/gooxml/schema/soo/sml"
)

type Fill struct {
	x     *sml.CT_Fill
	fills *sml.CT_Fills
}

func (f Fill) Index() uint32 {
	// in differential formats, CT_Fill is not held in a CT_Fills and index
	// doesn't mean anything
	if f.fills == nil {
		return 0
	}

	for i, sf := range f.fills.Fill {
		if f.x == sf {
			return uint32(i)
		}
	}
	return 0
}

func (f Fill) SetPatternFill() PatternFill {
	f.x.GradientFill = nil
	f.x.PatternFill = sml.NewCT_PatternFill()
	f.x.PatternFill.PatternTypeAttr = sml.ST_PatternTypeSolid
	return PatternFill{f.x.PatternFill, f.x}
}
