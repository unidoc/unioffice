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

// Endnote is an individual endnote reference within the document.
type Endnote struct {
	d *Document
	x *wml.CT_FtnEdn
}

// X returns the inner wrapped XML type.
func (e Endnote) X() *wml.CT_FtnEdn {
	return e.x
}

// Paragraphs returns the paragraphs defined in an endnote.
func (e Endnote) Paragraphs() []Paragraph {
	ret := []Paragraph{}
	for _, a := range e.content() {
		for _, p := range a.P {
			ret = append(ret, Paragraph{e.d, p})
		}
	}
	return ret
}

// AddParagraph adds a paragraph to the endnote.
func (e Endnote) AddParagraph() Paragraph {
	bc := wml.NewEG_ContentBlockContent()
	cbcB4 := len(e.x.EG_BlockLevelElts[0].EG_ContentBlockContent)
	e.x.EG_BlockLevelElts[0].EG_ContentBlockContent = append(
		e.x.EG_BlockLevelElts[0].EG_ContentBlockContent, bc)

	p := wml.NewCT_P()
	var prevStyle *wml.CT_String
	if cbcB4 != 0 {
		parasB4 := len(e.x.EG_BlockLevelElts[0].EG_ContentBlockContent[cbcB4-1].P)
		prevStyle = e.x.EG_BlockLevelElts[0].EG_ContentBlockContent[cbcB4-1].P[parasB4-1].PPr.PStyle
	} else {
		prevStyle = wml.NewCT_String()
		prevStyle.ValAttr = "Endnote"
	}
	bc.P = append(bc.P, p)

	newPara := Paragraph{e.d, p}
	newPara.x.PPr = wml.NewCT_PPr()
	newPara.x.PPr.PStyle = prevStyle
	newPara.x.PPr.RPr = wml.NewCT_ParaRPr()

	return newPara
}

// RemoveParagraph removes a paragraph from the endnote.
func (e Endnote) RemoveParagraph(p Paragraph) {
	for _, ec := range e.content() {
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

// helper function to get at the content of an endnote.
func (e Endnote) content() []*wml.EG_ContentBlockContent {
	var ret []*wml.EG_ContentBlockContent
	for _, ec := range e.x.EG_BlockLevelElts {
		ret = append(ret, ec.EG_ContentBlockContent...)
	}
	return ret
}

// helper function to get the IDAttr of an endnote object.
func (e Endnote) id() int64 {
	return e.x.IdAttr
}
