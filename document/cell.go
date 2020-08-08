// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io.

package document

import "github.com/unidoc/unioffice/schema/soo/wml"

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

// AddTable adds a table to the table cell.
func (c Cell) AddTable() Table {
	ble := wml.NewEG_BlockLevelElts()
	c.x.EG_BlockLevelElts = append(c.x.EG_BlockLevelElts, ble)
	bc := wml.NewEG_ContentBlockContent()
	ble.EG_ContentBlockContent = append(ble.EG_ContentBlockContent, bc)
	tbl := wml.NewCT_Tbl()
	bc.Tbl = append(bc.Tbl, tbl)

	return Table{c.d, tbl}
}

// Properties returns the cell properties.
func (c Cell) Properties() CellProperties {
	if c.x.TcPr == nil {
		c.x.TcPr = wml.NewCT_TcPr()
	}
	return CellProperties{c.x.TcPr}
}

// Paragraphs returns the paragraphs defined in the cell.
func (c Cell) Paragraphs() []Paragraph {
	ret := []Paragraph{}
	for _, ble := range c.x.EG_BlockLevelElts {
		for _, cbc := range ble.EG_ContentBlockContent {
			for _, p := range cbc.P {
				ret = append(ret, Paragraph{c.d, p})
			}
		}
	}
	return ret
}
