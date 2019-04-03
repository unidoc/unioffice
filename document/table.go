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

// Table is a table within a document.
type Table struct {
	d *Document
	x *wml.CT_Tbl
}

// X returns the inner wrapped XML type.
func (t Table) X() *wml.CT_Tbl {
	return t.x
}

// Properties returns the table properties.
func (t Table) Properties() TableProperties {
	if t.x.TblPr == nil {
		t.x.TblPr = wml.NewCT_TblPr()
	}
	return TableProperties{t.x.TblPr}
}

// AddRow adds a row to a table.
func (t Table) AddRow() Row {
	c := wml.NewEG_ContentRowContent()
	t.x.EG_ContentRowContent = append(t.x.EG_ContentRowContent, c)
	tr := wml.NewCT_Row()
	c.Tr = append(c.Tr, tr)
	return Row{t.d, tr}
}

// InsertRowAfter inserts a row after another row
func (t Table) InsertRowAfter(r Row) Row {
	for i, rc := range t.x.EG_ContentRowContent {
		if len(rc.Tr) > 0 && r.X() == rc.Tr[0] {
			c := wml.NewEG_ContentRowContent()
			if len(t.x.EG_ContentRowContent) <= i+2 {
				return t.AddRow()
			}
			t.x.EG_ContentRowContent = append(t.x.EG_ContentRowContent, nil)
			copy(t.x.EG_ContentRowContent[i+2:], t.x.EG_ContentRowContent[i+1:])
			t.x.EG_ContentRowContent[i+1] = c
			tr := wml.NewCT_Row()
			c.Tr = append(c.Tr, tr)
			return Row{t.d, tr}
		}
	}
	return t.AddRow()
}

// InsertRowBefore inserts a row before another row
func (t Table) InsertRowBefore(r Row) Row {
	for i, rc := range t.x.EG_ContentRowContent {
		if len(rc.Tr) > 0 && r.X() == rc.Tr[0] {
			c := wml.NewEG_ContentRowContent()
			t.x.EG_ContentRowContent = append(t.x.EG_ContentRowContent, nil)
			copy(t.x.EG_ContentRowContent[i+1:], t.x.EG_ContentRowContent[i:])
			t.x.EG_ContentRowContent[i] = c
			tr := wml.NewCT_Row()
			c.Tr = append(c.Tr, tr)
			return Row{t.d, tr}
		}
	}
	return t.AddRow()
}

// Rows returns the rows defined in the table.
func (t Table) Rows() []Row {
	ret := []Row{}
	for _, rc := range t.x.EG_ContentRowContent {
		for _, ctRow := range rc.Tr {
			ret = append(ret, Row{t.d, ctRow})
		}
		if rc.Sdt != nil && rc.Sdt.SdtContent != nil {
			for _, ctRow := range rc.Sdt.SdtContent.Tr {
				ret = append(ret, Row{t.d, ctRow})
			}
		}
	}
	return ret
}
