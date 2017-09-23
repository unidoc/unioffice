// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package document

import "baliance.com/gooxml/schema/soo/wml"

// Cell is a table cell within a document (not a spreadsheet)
type Cell struct {
	d *Document
	x *wml.CT_Tc
}

// X returns the inner wrapped XML type.
func (c Cell) X() *wml.CT_Tc {
	return c.x
}

// AddParagraph adds a paragraph to the table cell.
func (c Cell) AddParagraph() Paragraph {
	ble := wml.NewEG_BlockLevelElts()
	c.x.EG_BlockLevelElts = append(c.x.EG_BlockLevelElts, ble)
	bc := wml.NewEG_ContentBlockContent()
	ble.EG_ContentBlockContent = append(ble.EG_ContentBlockContent, bc)
	p := wml.NewCT_P()
	bc.P = append(bc.P, p)

	return Paragraph{c.d, p}
}

// Properties returns the cell properties.
func (c Cell) Properties() CellProperties {
	if c.x.TcPr == nil {
		c.x.TcPr = wml.NewCT_TcPr()
	}
	return CellProperties{c.x.TcPr}
}
