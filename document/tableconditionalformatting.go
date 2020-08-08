// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io.

package document

import "github.com/unidoc/unioffice/schema/soo/wml"

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
