// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package document

import (
	"baliance.com/gooxml"
	"baliance.com/gooxml/measurement"
	"baliance.com/gooxml/schema/soo/ofc/sharedTypes"
	"baliance.com/gooxml/schema/soo/wml"
)

// RowProperties are the properties for a row within a table
type RowProperties struct {
	x *wml.CT_TrPr
}

// SetHeight allows controlling the height of a row within a table.
func (r RowProperties) SetHeight(ht measurement.Distance, rule wml.ST_HeightRule) {
	if rule == wml.ST_HeightRuleUnset {
		r.x.TrHeight = nil
	} else {
		htv := wml.NewCT_Height()
		htv.HRuleAttr = rule
		htv.ValAttr = &sharedTypes.ST_TwipsMeasure{}
		htv.ValAttr.ST_UnsignedDecimalNumber = gooxml.Uint64(uint64(ht / measurement.Twips))
		r.x.TrHeight = []*wml.CT_Height{htv}
	}
}
