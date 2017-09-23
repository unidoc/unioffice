// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheet

import (
	"baliance.com/gooxml"
	"baliance.com/gooxml/color"
	"baliance.com/gooxml/measurement"
	"baliance.com/gooxml/schema/soo/sml"
)

// RichTextRun is a segment of text within a cell that is directly formatted.
type RichTextRun struct {
	x *sml.CT_RElt
}

// X returns the inner wrapped XML type.
func (r RichTextRun) X() *sml.CT_RElt {
	return r.x
}

// SetText sets the text to be displayed.
func (r RichTextRun) SetText(s string) {
	r.x.T = s
}

func (r RichTextRun) ensureRpr() {
	if r.x.RPr == nil {
		r.x.RPr = sml.NewCT_RPrElt()
	}
}

// SetBold causes the text to be displayed in bold.
func (r RichTextRun) SetBold(b bool) {
	r.ensureRpr()
	r.x.RPr.B = sml.NewCT_BooleanProperty()
	r.x.RPr.B.ValAttr = gooxml.Bool(b)
}

// SetColor sets the text color.
func (r RichTextRun) SetColor(c color.Color) {
	r.ensureRpr()
	r.x.RPr.Color = sml.NewCT_Color()
	r.x.RPr.Color.RgbAttr = c.AsRGBString()
}

// SetItalic causes the text to be displayed in italic.
func (r RichTextRun) SetItalic(b bool) {
	r.ensureRpr()
	r.x.RPr.I = sml.NewCT_BooleanProperty()
	r.x.RPr.I.ValAttr = gooxml.Bool(b)
}

// SetUnderline controls if the run is underlined.
func (r RichTextRun) SetUnderline(u sml.ST_UnderlineValues) {
	r.ensureRpr()
	r.x.RPr.U = sml.NewCT_UnderlineProperty()
	r.x.RPr.U.ValAttr = u
}

// SetSize sets the text size for a rich text run.
func (r RichTextRun) SetSize(m measurement.Distance) {
	r.ensureRpr()
	r.x.RPr.Sz = sml.NewCT_FontSize()
	r.x.RPr.Sz.ValAttr = float64(m / measurement.Point)
}

// SetFont sets the font name for a rich text run.
func (r RichTextRun) SetFont(s string) {
	r.ensureRpr()
	r.x.RPr.RFont = sml.NewCT_FontName()
	r.x.RPr.RFont.ValAttr = s
}
