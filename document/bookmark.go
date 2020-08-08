// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io.

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
