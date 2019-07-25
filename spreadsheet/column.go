// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheet

import (
	"github.com/unidoc/unioffice"
	"github.com/unidoc/unioffice/measurement"
	"github.com/unidoc/unioffice/schema/soo/sml"
)

// Column represents a column within a sheet. It's only used for formatting
// purposes, so it's possible to construct a sheet without configuring columns.
type Column struct {
	x *sml.CT_Col
}

// X returns the inner wrapped XML type.
func (c Column) X() *sml.CT_Col {
	return c.x
}

// SetWidth controls the width of a column.
func (c Column) SetWidth(w measurement.Distance) {
	c.x.WidthAttr = unioffice.Float64(float64(w / measurement.Character))
}

// SetStyle sets the cell style for an entire column.
func (c Column) SetStyle(cs CellStyle) {
	c.x.StyleAttr = unioffice.Uint32(cs.Index())
}

// SetHidden controls the visibility of a column.
func (c Column) SetHidden(b bool) {
	if !b {
		c.x.HiddenAttr = nil
	} else {
		c.x.HiddenAttr = unioffice.Bool(true)
	}
}
