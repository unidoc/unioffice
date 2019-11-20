// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased on https://unidoc.io.

package document

import (
	"errors"
	"fmt"

	"github.com/unidoc/unioffice"
	"github.com/unidoc/unioffice/common"
	"github.com/unidoc/unioffice/schema/soo/wml"
)

// Footer is a footer for a document section.
type Footer struct {
	d *Document
	x *wml.Ftr
}

// X returns the inner wrapped XML type.
func (f Footer) X() *wml.Ftr {
	return f.x
}

// AddParagraph adds a paragraph to the footer.
func (f Footer) AddParagraph() Paragraph {
	bc := wml.NewEG_ContentBlockContent()
	f.x.EG_ContentBlockContent = append(f.x.EG_ContentBlockContent, bc)
	p := wml.NewCT_P()
	bc.P = append(bc.P, p)
	return Paragraph{f.d, p}
}

// Index returns the index of the footer within the document.  This is used to
// form its zip packaged filename as well as to match it with its relationship
// ID.
func (f Footer) Index() int {
	for i, hdr := range f.d.footers {
		if hdr == f.x {
			return i
		}
	}
	return -1
}

// Tables returns the tables defined in the footer.
func (f Footer) Tables() []Table {
	ret := []Table{}
	if f.x == nil {
		return nil
	}
	for _, c := range f.x.EG_ContentBlockContent {
		for _, t := range f.d.tables(c) {
			ret = append(ret, t)
		}
	}
	return ret
}

// Paragraphs returns the paragraphs defined in a footer.
func (f Footer) Paragraphs() []Paragraph {
	ret := []Paragraph{}
	for _, ec := range f.x.EG_ContentBlockContent {
		for _, p := range ec.P {
			ret = append(ret, Paragraph{f.d, p})
		}
	}

	for _, t := range f.Tables() {
		for _, r := range t.Rows() {
			for _, c := range r.Cells() {
				ret = append(ret, c.Paragraphs()...)
			}
		}
	}
	return ret
}

// RemoveParagraph removes a paragraph from a footer.
func (f Footer) RemoveParagraph(p Paragraph) {
	for _, ec := range f.x.EG_ContentBlockContent {
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

// Clear clears all content within a footer
func (f Footer) Clear() {
	f.x.EG_ContentBlockContent = nil
}

// AddImage adds an image to the document package, returning a reference that
// can be used to add the image to a run and place it in the document contents.
func (f Footer) AddImage(i common.Image) (common.ImageRef, error) {
	var ftrRels common.Relationships
	for i, ftr := range f.d.footers {
		if ftr == f.x {
			ftrRels = f.d.ftrRels[i]
		}
	}

	r := common.MakeImageRef(i, &f.d.DocBase, ftrRels)
	if i.Data == nil && i.Path == "" {
		return r, errors.New("image must have data or a path")
	}

	if i.Format == "" {
		return r, errors.New("image must have a valid format")
	}
	if i.Size.X == 0 || i.Size.Y == 0 {
		return r, errors.New("image must have a valid size")
	}

	f.d.Images = append(f.d.Images, r)
	fn := fmt.Sprintf("media/image%d.%s", len(f.d.Images), i.Format)
	rel := ftrRels.AddRelationship(fn, unioffice.ImageType)
	r.SetRelID(rel.X().IdAttr)
	return r, nil
}
