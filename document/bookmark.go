// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased on https://unidoc.io.

package document

import "github.com/unidoc/unioffice/schema/soo/wml"

// Bookmark is a bookmarked location within a document that can be referenced
// with a hyperlink.
type Bookmark struct {
	x *wml.CT_Bookmark
}

// X returns the inner wrapped XML type.
func (b Bookmark) X() *wml.CT_Bookmark {
	return b.x
}

// SetName sets the name of the bookmark. This is the name that is used to
// reference the bookmark from hyperlinks.
func (b Bookmark) SetName(name string) {
	b.x.NameAttr = name
}

// Name returns the name of the bookmark whcih is the document unique ID that
// identifies the bookmark.
func (b Bookmark) Name() string {
	return b.x.NameAttr
}
