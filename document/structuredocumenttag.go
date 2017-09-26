// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package document

import "baliance.com/gooxml/schema/soo/wml"

// StructuredDocumentTag are a tagged bit of content in a document.
type StructuredDocumentTag struct {
	d *Document
	x *wml.CT_SdtBlock
}

// Paragraphs returns the paragraphs within a structured document tag.
func (s StructuredDocumentTag) Paragraphs() []Paragraph {
	if s.x.SdtContent == nil {
		return nil
	}
	ret := []Paragraph{}
	for _, p := range s.x.SdtContent.P {
		ret = append(ret, Paragraph{s.d, p})
	}
	return ret
}
