// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package document

import "baliance.com/gooxml/schema/soo/wml"

// CellProperties are a table cells properties within a document.
type CellProperties struct {
	x *wml.CT_TcPr
}

// X returns the inner wrapped XML type.
func (c CellProperties) X() *wml.CT_TcPr {
	return c.x
}

// SetColumnSpan sets the number of Grid Columns Spanned by the Cell
func (c CellProperties) SetColumnSpan(cols int) {
	if cols == 0 {
		c.x.GridSpan = nil
	} else {
		c.x.GridSpan = wml.NewCT_DecimalNumber()
		c.x.GridSpan.ValAttr = int64(cols)
	}
}
