// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheet

import "baliance.com/gooxml/schema/soo/sml"

type PivotDataField struct {
	x *sml.CT_DataField
}

// X returns the inner wrapped XML type.
func (p PivotDataField) X() *sml.CT_DataField {
	return p.x
}

func (p PivotDataField) SetField(idx uint32) {
	p.x.FldAttr = idx
}

func (p PivotDataField) SetSubtotal(fn sml.ST_DataConsolidateFunction) {
	p.x.SubtotalAttr = fn
}
