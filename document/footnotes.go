// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package document

import (
	"github.com/unidoc/unioffice/schema/soo/wml"
)

// Footnote is an individual footnote reference within the document.
type Footnote struct {
	d *Document
	x *wml.CT_FtnEdn
}

// X returns the inner wrapped XML type.
func (f Footnote) X() *wml.CT_FtnEdn {
	return f.x
}

// Paragraphs returns the paragraphs defined in a footnote.
func (f Footnote) Paragraphs() []Paragraph {
	ret := []Paragraph{}
	for _, a := range f.content() {
		for _, p := range a.P {
			ret = append(ret, Paragraph{f.d, p})
		}
	}
	return ret
}

// AddParagraph adds a paragraph to the footnote.
func (f Footnote) AddParagraph() Paragraph {
	bc := wml.NewEG_ContentBlockContent()
	cbcB4 := len(f.x.EG_BlockLevelElts[0].EG_ContentBlockContent)
	f.x.EG_BlockLevelElts[0].EG_ContentBlockContent = append(
		f.x.EG_BlockLevelElts[0].EG_ContentBlockContent, bc)

	p := wml.NewCT_P()
	var prevStyle *wml.CT_String
	if cbcB4 != 0 {
		parasB4 := len(f.x.EG_BlockLevelElts[0].EG_ContentBlockContent[cbcB4-1].P)
		prevStyle = f.x.EG_BlockLevelElts[0].EG_ContentBlockContent[cbcB4-1].P[parasB4-1].PPr.PStyle
	} else {
		prevStyle = wml.NewCT_String()
		prevStyle.ValAttr = "Footnote"
	}
	bc.P = append(bc.P, p)

	newPara := Paragraph{f.d, p}
	newPara.x.PPr = wml.NewCT_PPr()
	newPara.x.PPr.PStyle = prevStyle
	newPara.x.PPr.RPr = wml.NewCT_ParaRPr()

	return newPara
}

// RemoveParagraph removes a paragraph from the footnote.
func (f Footnote) RemoveParagraph(p Paragraph) {
	for _, ec := range f.content() {
		for i, pa := range ec.P {
			// do we need to remove this paragraph?
			if pa == p.x {
				copy(ec.P[i:], ec.P[i+1:])
				ec.P = ec.P[0 : len(ec.P)-1]
				return
			}
		}
	}
}

// helper function to get at the content of a footnote.
func (f Footnote) content() []*wml.EG_ContentBlockContent {
	var ret []*wml.EG_ContentBlockContent
	for _, fc := range f.x.EG_BlockLevelElts {
		ret = append(ret, fc.EG_ContentBlockContent...)
	}
	return ret
}

// helper function to get the IDAttr of a footnote object.
func (f Footnote) id() int64 {
	return f.x.IdAttr
}
