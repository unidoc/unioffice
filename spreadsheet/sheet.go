// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheet

import (
	"baliance.com/gooxml"
	"baliance.com/gooxml/schema/schemas.openxmlformats.org/spreadsheetml"
)

// Sheet is a single sheet within a workbook.
type Sheet struct {
	w  *Workbook
	x  *spreadsheetml.CT_Sheet
	ws *spreadsheetml.Worksheet
}

// AddRow adds a new row to a sheet.
func (s Sheet) AddRow() Row {
	r := spreadsheetml.NewCT_Row()
	r.RAttr = gooxml.Uint32(uint32(len(s.ws.SheetData.Row) + 1))
	s.ws.SheetData.Row = append(s.ws.SheetData.Row, r)
	return Row{s.w, r}
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
