package document

import (
	"github.com/unidoc/unioffice/schema/soo/wml"
)

// Endnote is an individual footnote reference within the document.
type Endnote struct {
	d *Document
	x *wml.CT_FtnEdn
}

// X returns the inner wrapped XML type.
func (e Endnote) X() *wml.CT_FtnEdn {
	return e.x
}

// Paragraphs returns the paragraphs defined in a footnote.
func (e Endnote) Paragraphs() []Paragraph {
	ret := []Paragraph{}
	for _, a := range e.content() {
		for _, p := range a.P {
			ret = append(ret, Paragraph{e.d, p})
		}
	}
	return ret
}

// AddParagraph adds a paragraph to the footnote.
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

// RemoveParagraph removes a paragraph from the footnote.
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

// helper function to get at the content of a footnote
func (e Endnote) content() []*wml.EG_ContentBlockContent {
	var ret []*wml.EG_ContentBlockContent
	for _, ec := range e.x.EG_BlockLevelElts {
		ret = append(ret, ec.EG_ContentBlockContent...)
	}
	return ret
}

// helper function to get the IDAttr of a footnote object
func (e Endnote) id() int64 {
	return e.x.IdAttr
}
