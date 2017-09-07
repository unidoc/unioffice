// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheet

import (
	"fmt"
	"log"

	"baliance.com/gooxml"
	"baliance.com/gooxml/common"
	sml "baliance.com/gooxml/schema/schemas.openxmlformats.org/spreadsheetml"
)

// Sheet is a single sheet within a workbook.
type Sheet struct {
	w   *Workbook
	cts *sml.CT_Sheet
	x   *sml.Worksheet
}

// Row will return a row with a given row number, creating a new row if
// necessary.
func (s Sheet) Row(rowNum uint32) Row {
	// see if the row exists
	for _, r := range s.x.SheetData.Row {
		if r.RAttr != nil && *r.RAttr == rowNum {
			return Row{s.w, s.x, r}
		}
	}
	// create a new row
	return s.AddNumberedRow(rowNum)
}

// Cell creates or returns a cell given a cell reference of the form 'A10'
func (s Sheet) Cell(cellRef string) Cell {
	col, row, err := ParseCellReference(cellRef)
	if err != nil {
		log.Printf("error parsing cell reference: %s", err)
		return s.AddRow().AddCell()
	}
	return s.Row(row).Cell(col)
}

// AddNumberedRow adds a row with a given row number.  If you reuse a row number
// the resulting file will fail validation and fail to open in Office programs. Use
// Row instead which creates a new row or returns an existing row.
func (s Sheet) AddNumberedRow(rowNum uint32) Row {
	r := sml.NewCT_Row()
	r.RAttr = gooxml.Uint32(rowNum)
	s.x.SheetData.Row = append(s.x.SheetData.Row, r)
	return Row{s.w, s.x, r}
}

// AddRow adds a new row to a sheet.  You can mix this with numbered rows,
// however it will get confusing. You should prefer to use either automatically
// numbered rows with AddRow or manually numbered rows with Row/AddNumberedRow
func (s Sheet) AddRow() Row {
	maxRowID := uint32(0)
	// find the max row number
	for _, r := range s.x.SheetData.Row {
		if r.RAttr != nil && *r.RAttr > maxRowID {
			maxRowID = *r.RAttr
		}
	}

	return s.AddNumberedRow(maxRowID + 1)
}

// Name returns the sheet name
func (s Sheet) Name() string {
	return s.cts.NameAttr
}

// SetName sets the sheet name.
func (s Sheet) SetName(name string) {
	s.cts.NameAttr = name
}

// Validate validates the sheet, returning an error if it is found to be invalid.
func (s Sheet) Validate() error {

	usedRows := map[uint32]struct{}{}
	for _, r := range s.x.SheetData.Row {
		if r.RAttr != nil {
			if _, reusedRow := usedRows[*r.RAttr]; reusedRow {
				return fmt.Errorf("'%s' reused row %d", s.Name(), *r.RAttr)
			}
			usedRows[*r.RAttr] = struct{}{}
		}
		usedCells := map[string]struct{}{}
		for _, c := range r.C {
			if c.RAttr == nil {
				continue
			}

			if _, reusedCell := usedCells[*c.RAttr]; reusedCell {
				return fmt.Errorf("'%s' reused cell %s", s.Name(), *c.RAttr)
			}
			usedCells[*c.RAttr] = struct{}{}
		}
	}
	if err := s.cts.Validate(); err != nil {
		return err
	}
	return s.x.Validate()
}

// ValidateWithPath validates the sheet passing path informaton for a better
// error message
func (s Sheet) ValidateWithPath(path string) error {
	return s.cts.ValidateWithPath(path)
}

// Rows returns all of the rows in a sheet.
func (s Sheet) Rows() []Row {
	ret := []Row{}
	for _, r := range s.x.SheetData.Row {
		ret = append(ret, Row{s.w, s.x, r})
	}
	return ret
}

// SetDrawing sets the worksheet drawing.  A worksheet can have a reference to a
// single drawing, but the drawing can have many charts.
func (s Sheet) SetDrawing(d Drawing) {
	var rel common.Relationships
	for i, wks := range s.w.xws {
		if wks == s.x {
			rel = s.w.xwsRels[i]
			break
		}
	}
	// add relationship from drawing to the sheet
	var drawingID string
	for i, dr := range d.wb.drawings {
		if dr == d.x {
			rel := rel.AddAutoRelationship(gooxml.DocTypeSpreadsheet, i+1, gooxml.DrawingType)
			drawingID = rel.ID()
			break
		}
	}
	s.x.Drawing = sml.NewCT_Drawing()
	s.x.Drawing.IdAttr = drawingID
}

// AddHyperlink adds a hyperlink to a sheet. Adding the hyperlink to the sheet
// and setting it on a cell is more efficient than setting hyperlinks directly
// on a cell.
func (s Sheet) AddHyperlink(url string) common.Hyperlink {
	// store the relationships so we don't need to do a lookup here?
	for i, ws := range s.w.xws {
		if ws == s.x {
			// add a hyperlink relationship in the worksheet relationships file
			return s.w.xwsRels[i].AddHyperlink(url)
		}
	}
	// should never occur
	return common.Hyperlink{}
}
