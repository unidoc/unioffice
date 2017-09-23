// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package document

import "baliance.com/gooxml/schema/soo/wml"

// Header is a header for a document section.
type Header struct {
	d *Document
	x *wml.Hdr
}

// X returns the inner wrapped XML type.
func (h Header) X() *wml.Hdr {
	return h.x
}

// AddParagraph adds a paragraph to the header.
func (h Header) AddParagraph() Paragraph {
	bc := wml.NewEG_ContentBlockContent()
	h.x.EG_ContentBlockContent = append(h.x.EG_ContentBlockContent, bc)
	p := wml.NewCT_P()
	bc.P = append(bc.P, p)
	return Paragraph{h.d, p}
}

// Index returns the index of the header within the document.  This is used to
// form its zip packaged filename as well as to match it with its relationship
// ID.
func (h Header) Index() int {
	for i, hdr := range h.d.headers {
		if hdr == h.x {
			return i
		}
	}
	return -1
}
