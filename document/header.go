// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package document

import (
	"errors"
	"fmt"

	"baliance.com/gooxml"
	"baliance.com/gooxml/common"
	"baliance.com/gooxml/schema/soo/wml"
)

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

// Paragraphs returns the paragraphs defined in a header.
func (h Header) Paragraphs() []Paragraph {
	ret := []Paragraph{}
	for _, ec := range h.x.EG_ContentBlockContent {
		for _, p := range ec.P {
			ret = append(ret, Paragraph{h.d, p})
		}
	}
	return ret
}

// RemoveParagraph removes a paragraph from a footer.
func (h Header) RemoveParagraph(p Paragraph) {
	for _, ec := range h.x.EG_ContentBlockContent {
		for i, pa := range ec.P {
			// do we need to remove this paragraph
			if pa == p.x {
				copy(ec.P[i:], ec.P[i+1:])
				ec.P = ec.P[0 : len(ec.P)-1]
				return
			}
		}
	}
}

// Clear clears all content within a header
func (h Header) Clear() {
	h.x.EG_ContentBlockContent = nil
}

// AddImage adds an image to the document package, returning a reference that
// can be used to add the image to a run and place it in the document contents.
func (h Header) AddImage(i common.Image) (common.ImageRef, error) {
	var hdrRels common.Relationships
	for i, hdr := range h.d.headers {
		if hdr == h.x {
			hdrRels = h.d.hdrRels[i]
		}
	}

	r := common.MakeImageRef(i, &h.d.DocBase, hdrRels)
	if i.Path == "" {
		return r, errors.New("image must have a path")
	}

	if i.Format == "" {
		return r, errors.New("image must have a valid format")
	}
	if i.Size.X == 0 || i.Size.Y == 0 {
		return r, errors.New("image must have a valid size")
	}

	h.d.Images = append(h.d.Images, r)
	fn := fmt.Sprintf("media/image%d.%s", len(h.d.Images), i.Format)
	rel := hdrRels.AddRelationship(fn, gooxml.ImageType)
	r.SetRelID(rel.X().IdAttr)
	return r, nil
}
