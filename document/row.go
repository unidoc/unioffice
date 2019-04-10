// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package document

import (
	"baliance.com/gooxml/schema/soo/wml"
)

// Row is a row within a table within a document.
type Row struct {
	d *Document
	x *wml.CT_Row
}

// X returns the inner wrapped XML type.
func (r Row) X() *wml.CT_Row {
	return r.x
}

// AddCell adds a cell to a row and returns it
func (r Row) AddCell() Cell {
	cc := wml.NewEG_ContentCellContent()
	r.x.EG_ContentCellContent = append(r.x.EG_ContentCellContent, cc)
	tc := wml.NewCT_Tc()
	cc.Tc = append(cc.Tc, tc)
	return Cell{r.d, tc}
}

// Properties returns the row properties.
func (r Row) Properties() RowProperties {
	if r.x.TrPr == nil {
		r.x.TrPr = wml.NewCT_TrPr()
	}
	return RowProperties{r.x.TrPr}
}

// Cells returns the cells defined in the table.
func (r Row) Cells() []Cell {
	ret := []Cell{}
	for _, cc := range r.x.EG_ContentCellContent {
		for _, ctCell := range cc.Tc {
			ret = append(ret, Cell{r.d, ctCell})
		}
		if cc.Sdt != nil && cc.Sdt.SdtContent != nil {
			for _, ctCell := range cc.Sdt.SdtContent.Tc {
				ret = append(ret, Cell{r.d, ctCell})
			}
		}
	}
	return ret
}
