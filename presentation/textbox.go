// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased on https://unidoc.io.

package presentation

import (
	"github.com/unidoc/unioffice/drawing"
	"github.com/unidoc/unioffice/schema/soo/dml"
	"github.com/unidoc/unioffice/schema/soo/pml"
)

// TextBox is a text box within a slide.
type TextBox struct {
	x *pml.CT_Shape
}

// AddParagraph adds a paragraph to the text box
func (t TextBox) AddParagraph() drawing.Paragraph {
	p := dml.NewCT_TextParagraph()
	t.x.TxBody.P = append(t.x.TxBody.P, p)
	return drawing.MakeParagraph(p)
}

// Properties returns the properties of the TextBox.
func (t TextBox) Properties() drawing.ShapeProperties {
	if t.x.SpPr == nil {
		t.x.SpPr = dml.NewCT_ShapeProperties()
	}
	return drawing.MakeShapeProperties(t.x.SpPr)
}

// SetTextAnchor controls the text anchoring
func (t TextBox) SetTextAnchor(a dml.ST_TextAnchoringType) {
	t.x.TxBody.BodyPr = dml.NewCT_TextBodyProperties()
	t.x.TxBody.BodyPr.AnchorAttr = a
}
