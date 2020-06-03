// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased on https://unidoc.io.

package presentation

import (
	"errors"

	"github.com/unidoc/unioffice"
	"github.com/unidoc/unioffice/drawing"
	"github.com/unidoc/unioffice/schema/soo/dml"
	"github.com/unidoc/unioffice/schema/soo/pml"
)

// PlaceHolder is a place holder from a slide.
type PlaceHolder struct {
	x   *pml.CT_Shape
	sld *pml.Sld
}

// X returns the inner wrapped XML type.
func (s PlaceHolder) X() *pml.CT_Shape {
	return s.x
}

// Type returns the placeholder type
func (s PlaceHolder) Type() pml.ST_PlaceholderType {
	return s.x.NvSpPr.NvPr.Ph.TypeAttr
}

// Index returns the placeholder index
func (s PlaceHolder) Index() uint32 {
	if s.x.NvSpPr.NvPr.Ph.IdxAttr == nil {
		return 0
	}
	return *s.x.NvSpPr.NvPr.Ph.IdxAttr
}

// ClearAll completely clears a placeholder. To be useable, at least one
// paragraph must be added after ClearAll via AddParagraph.
func (s PlaceHolder) ClearAll() {
	s.x.SpPr = dml.NewCT_ShapeProperties()
	s.x.TxBody = dml.NewCT_TextBody()
	s.x.TxBody.LstStyle = dml.NewCT_TextListStyle()
}

// Clear clears the placeholder contents and adds a single empty paragraph.  The
// empty paragrah is required by PowerPoint or it will report the file as being
// invalid.
func (s PlaceHolder) Clear() {
	s.ClearAll()
	para := dml.NewCT_TextParagraph()
	s.x.TxBody.P = []*dml.CT_TextParagraph{para}
	para.EndParaRPr = dml.NewCT_TextCharacterProperties()
	para.EndParaRPr.LangAttr = unioffice.String("en-US")
}

// Remove removes a placeholder from a presentation.
func (s PlaceHolder) Remove() error {
	for i, spChc := range s.sld.CSld.SpTree.Choice {
		for _, sp := range spChc.Sp {
			if sp == s.x {
				copy(s.sld.CSld.SpTree.Choice[i:], s.sld.CSld.SpTree.Choice[i+1:])
				s.sld.CSld.SpTree.Choice = s.sld.CSld.SpTree.Choice[0 : len(s.sld.CSld.SpTree.Choice)-1]
				return nil
			}
		}
	}
	return errors.New("placeholder not found in slide")
}

// SetText sets the text of a placeholder for the initial paragraph. This is a
// shortcut method that is useful for things like titles which only contain a
// single paragraph.
func (s PlaceHolder) SetText(text string) {
	s.Clear()
	tr := dml.NewEG_TextRun()
	tr.R = dml.NewCT_RegularTextRun()
	tr.R.T = text
	if len(s.x.TxBody.P) == 0 {
		s.x.TxBody.P = append(s.x.TxBody.P, dml.NewCT_TextParagraph())
	}
	s.x.TxBody.P[0].EG_TextRun = nil
	s.x.TxBody.P[0].EG_TextRun = append(s.x.TxBody.P[0].EG_TextRun, tr)
}

// Paragraphs returns the paragraphs defined in the placeholder.
func (s PlaceHolder) Paragraphs() []drawing.Paragraph {
	ret := []drawing.Paragraph{}
	for _, p := range s.x.TxBody.P {
		ret = append(ret, drawing.MakeParagraph(p))
	}
	return ret
}

// AddParagraph adds a new paragraph to a placeholder.
func (s PlaceHolder) AddParagraph() drawing.Paragraph {
	p := drawing.MakeParagraph(dml.NewCT_TextParagraph())
	s.x.TxBody.P = append(s.x.TxBody.P, p.X())
	return p
}
