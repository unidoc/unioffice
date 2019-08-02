// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheet

import "github.com/unidoc/unioffice/schema/soo/sml"

// RichText is a container for the rich text within a cell. It's similar to a
// paragaraph for a document, except a cell can only contain one rich text item.
type RichText struct {
	x *sml.CT_Rst
}

// X returns the inner wrapped XML type.
func (r RichText) X() *sml.CT_Rst {
	return r.x
}

// AddRun adds a new run of text to the cell.
func (r RichText) AddRun() RichTextRun {
	elt := sml.NewCT_RElt()
	r.x.R = append(r.x.R, elt)
	return RichTextRun{elt}
}
