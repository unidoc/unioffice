// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheet

import (
	"strings"

	"github.com/unidoc/unioffice/schema/soo/sml"
)

type MergedCell struct {
	wb *Workbook
	ws *sml.Worksheet
	x  *sml.CT_MergeCell
}

// X returns the inner wrapped XML type.
func (s MergedCell) X() *sml.CT_MergeCell {
	return s.x
}

// SetReference sets the regin of cells that the merged cell applies to.
func (s MergedCell) SetReference(ref string) {
	s.x.RefAttr = ref
}

// Reference returns the region of cells that are merged.
func (s MergedCell) Reference() string {
	return s.x.RefAttr
}

// Cell returns the actual cell behind the merged region
func (s MergedCell) Cell() Cell {
	ref := s.Reference()
	if idx := strings.Index(s.Reference(), ":"); idx != -1 {
		ref = ref[0:idx]
		return Sheet{w: s.wb, x: s.ws}.Cell(ref)
	}
	// couldn't find it, log an error?
	return Cell{}
}
