// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package drawing

import (
	"baliance.com/gooxml/schema/soo/dml"
)

// MakeParagraph constructs a new paragraph wrapper.
func MakeParagraph(x *dml.CT_TextParagraph) Paragraph {
	return Paragraph{x}
}

// Paragraph is a paragraph within a document.
type Paragraph struct {
	x *dml.CT_TextParagraph
}

// X returns the inner wrapped XML type.
func (p Paragraph) X() *dml.CT_TextParagraph {
	return p.x
}

// AddRun adds a new run to a paragraph.
func (p Paragraph) AddRun() Run {
	r := MakeRun(dml.NewEG_TextRun())
	p.x.EG_TextRun = append(p.x.EG_TextRun, r.X())
	return r
}

// AddBreak adds a new line break to a paragraph.
func (p Paragraph) AddBreak() {
	r := dml.NewEG_TextRun()
	r.Br = dml.NewCT_TextLineBreak()
	p.x.EG_TextRun = append(p.x.EG_TextRun, r)
}

// Properties returns the paragraph properties.
func (p Paragraph) Properties() ParagraphProperties {
	if p.x.PPr == nil {
		p.x.PPr = dml.NewCT_TextParagraphProperties()
	}
	return MakeParagraphProperties(p.x.PPr)
}
