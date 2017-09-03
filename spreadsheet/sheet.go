// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheet

import (
	"baliance.com/gooxml"
	sml "baliance.com/gooxml/schema/schemas.openxmlformats.org/spreadsheetml"
)

// Sheet is a single sheet within a workbook.
type Sheet struct {
	w  *Workbook
	x  *sml.CT_Sheet
	ws *sml.Worksheet
}

// AddRow adds a new row to a sheet.
func (s Sheet) AddRow() Row {
	r := sml.NewCT_Row()
	r.RAttr = gooxml.Uint32(uint32(len(s.ws.SheetData.Row) + 1))
	s.ws.SheetData.Row = append(s.ws.SheetData.Row, r)
	return Row{s.w, r}
}

// Name returns the sheet name
func (s Sheet) Name() string {
	return s.x.NameAttr
}

// SetName sets the sheet name.
func (s Sheet) SetName(name string) {
	s.x.NameAttr = name
}

// Validate validates the sheet, returning an error if it is found to be invalid.
func (s Sheet) Validate() error {
	return s.x.Validate()
}

// ValidateWithPath validates the sheet passing path informaton for a better
// error message
func (s Sheet) ValidateWithPath(path string) error {
	return s.x.ValidateWithPath(path)
}

// Rows returns all of the rows in a sheet.
func (s Sheet) Rows() []Row {
	ret := []Row{}
	for _, r := range s.ws.SheetData.Row {
		ret = append(ret, Row{s.w, r})
	}
	return ret
}
