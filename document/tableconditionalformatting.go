// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package document

import "baliance.com/gooxml/schema/soo/wml"

// TableConditionalFormatting controls the conditional formatting within a table
// style.
type TableConditionalFormatting struct {
	x *wml.CT_TblStylePr
}

// X returns the inner wrapped XML type.
func (t TableConditionalFormatting) X() *wml.CT_TblStylePr {
	return t.x
}

// CellProperties returns the cell properties.
func (t TableConditionalFormatting) CellProperties() CellProperties {
	if t.x.TcPr == nil {
		t.x.TcPr = wml.NewCT_TcPr()
	}
	return CellProperties{t.x.TcPr}
}

// RunProperties returns the run properties controlling text formatting within the table.
func (t TableConditionalFormatting) RunProperties() RunProperties {
	if t.x.RPr == nil {
		t.x.RPr = wml.NewCT_RPr()
	}
	return RunProperties{t.x.RPr}
}

// ParagraphProperties returns the paragraph properties controlling text formatting within the table.
func (t TableConditionalFormatting) ParagraphProperties() ParagraphStyleProperties {
	if t.x.PPr == nil {
		t.x.PPr = wml.NewCT_PPrGeneral()
	}
	return ParagraphStyleProperties{t.x.PPr}
}
