// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package document

import "baliance.com/gooxml/schema/soo/wml"

// Footer is a footer for a document section.
type Footer struct {
	d *Document
	x *wml.Ftr
}

// X returns the inner wrapped XML type.
func (h Footer) X() *wml.Ftr {
	return h.x
}

// AddParagraph adds a paragraph to the footer.
func (h Footer) AddParagraph() Paragraph {
	bc := wml.NewEG_ContentBlockContent()
	h.x.EG_ContentBlockContent = append(h.x.EG_ContentBlockContent, bc)
	p := wml.NewCT_P()
	bc.P = append(bc.P, p)
	return Paragraph{h.d, p}
}

// Index returns the index of the footer within the document.  This is used to
// form its zip packaged filename as well as to match it with its relationship
// ID.
func (h Footer) Index() int {
	for i, hdr := range h.d.footers {
		if hdr == h.x {
			return i
		}
	}
	return -1
}
