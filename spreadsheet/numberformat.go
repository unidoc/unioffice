// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheet

import "github.com/unidoc/unioffice/schema/soo/sml"

// NumberFormat is a number formatting string that can be applied to a cell
// style.
type NumberFormat struct {
	wb *Workbook
	x  *sml.CT_NumFmt
}

// X returns the inner wrapped XML type.
func (n NumberFormat) X() *sml.CT_NumFmt {
	return n.x
}

// SetFormat sets the number format code.
func (n NumberFormat) SetFormat(f string) {
	n.x.FormatCodeAttr = f
}

// GetFormat sets the number format code.
func (n NumberFormat) GetFormat() string {
	return n.x.FormatCodeAttr
}

// ID returns the number format ID.  This is not an index as there are some
// predefined number formats which can be used in cell styles and don't need a
// corresponding NumberFormat.
func (n NumberFormat) ID() uint32 {
	return n.x.NumFmtIdAttr
}
