// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package drawing

import (
	"baliance.com/gooxml"
	"baliance.com/gooxml/schema/soo/dml"
)

// ParagraphProperties allows controlling paragraph properties.
type ParagraphProperties struct {
	x *dml.CT_TextParagraphProperties
}

// MakeParagraphProperties constructs a new ParagraphProperties wrapper.
func MakeParagraphProperties(x *dml.CT_TextParagraphProperties) ParagraphProperties {
	return ParagraphProperties{x}
}

// SetLevel sets the level of indentation of a paragraph.
func (p ParagraphProperties) SetLevel(idx int32) {
	p.x.LvlAttr = gooxml.Int32(idx)
}

// SetNumbered controls if bullets are numbered or not.
func (p ParagraphProperties) SetNumbered(scheme dml.ST_TextAutonumberScheme) {
	if scheme == dml.ST_TextAutonumberSchemeUnset {
		p.x.BuAutoNum = nil
	} else {
		p.x.BuAutoNum = dml.NewCT_TextAutonumberBullet()
		p.x.BuAutoNum.TypeAttr = scheme
	}
}
