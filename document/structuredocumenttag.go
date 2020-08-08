// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io.

package document

import "github.com/unidoc/unioffice/schema/soo/wml"

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
