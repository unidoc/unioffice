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
	"sort"
	"strings"

	"baliance.com/gooxml/spreadsheet/formula"
	"baliance.com/gooxml/spreadsheet/reference"

	"baliance.com/gooxml"
	"baliance.com/gooxml/common"
	"baliance.com/gooxml/schema/soo/sml"
	"baliance.com/gooxml/vmldrawing"
)

// Sheet is a single sheet within a workbook.
type Sheet struct {
	w   *Workbook
	cts *sml.CT_Sheet
	x   *sml.Worksheet
}

func (s Sheet) IsValid() bool {
	return s.x != nil
}

// X returns the inner wrapped XML type.
func (s Sheet) X() *sml.Worksheet {
	return s.x
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
	cref, err := reference.ParseCellReference(cellRef)
	if err != nil {
		gooxml.Log("error parsing cell reference: %s", err)
		return s.AddRow().AddCell()
	}
	return s.Row(cref.RowIdx).Cell(cref.Column)
}

// AddNumberedRow adds a row with a given row number.  If you reuse a row number
// the resulting file will fail validation and fail to open in Office programs. Use
// Row instead which creates a new row or returns an existing row.
func (s Sheet) AddNumberedRow(rowNum uint32) Row {
	r := sml.NewCT_Row()
	r.RAttr = gooxml.Uint32(rowNum)
	s.x.SheetData.Row = append(s.x.SheetData.Row, r)

	// Excel wants the rows to be sorted
	sort.Slice(s.x.SheetData.Row, func(i, j int) bool {
		l := s.x.SheetData.Row[i].RAttr
		r := s.x.SheetData.Row[j].RAttr
		if l == nil {
			return true
		}
		if r == nil {
			return true
		}
		return *l < *r
	})

	return Row{s.w, s.x, r}
}

// addNumberedRowFast is a fast path that can be used when adding consecutive
// rows and not skipping any.
func (s Sheet) addNumberedRowFast(rowNum uint32) Row {
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

	numRows := uint32(len(s.x.SheetData.Row))
	// fast path, adding consecutive rows
	if numRows > 0 && s.x.SheetData.Row[numRows-1].RAttr != nil && *s.x.SheetData.Row[numRows-1].RAttr == numRows {
		return s.addNumberedRowFast(numRows + 1)
	}

	// find the max row number
	for _, r := range s.x.SheetData.Row {
		if r.RAttr != nil && *r.RAttr > maxRowID {
			maxRowID = *r.RAttr
		}
	}

	return s.AddNumberedRow(maxRowID + 1)
}

// InsertRow inserts a new row into a spreadsheet at a particular row number.  This
// row will now be the row number specified, and any rows after it will be renumbed.
func (s Sheet) InsertRow(rowNum int) Row {
	rIdx := uint32(rowNum)

	// Renumber every row after the row we're inserting
	for _, r := range s.Rows() {
		if r.x.RAttr != nil && *r.x.RAttr >= rIdx {
			*r.x.RAttr++
			// and we also need to recompute cell references
			for _, c := range r.Cells() {
				cref, err := reference.ParseCellReference(c.Reference())
				if err != nil {
					continue
				}
				cref.RowIdx++
				c.x.RAttr = gooxml.String(cref.String())
			}
		}
	}

	// check for merged cells, and try to intelligently adjust them
	// issue #212
	for _, mc := range s.MergedCells() {
		from, to, err := reference.ParseRangeReference(mc.Reference())
		if err != nil {
			continue
		}
		if int(from.RowIdx) >= rowNum {
			from.RowIdx++
		}
		if int(to.RowIdx) >= rowNum {
			to.RowIdx++
		}
		ref := fmt.Sprintf("%s:%s", from, to)
		mc.SetReference(ref)
	}

	// finally AddNumberedRow will add and re-sort rows
	return s.AddNumberedRow(rIdx)
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
	validators := []func() error{
		s.validateRowCellNumbers,
		s.validateMergedCells,
		s.validateSheetNames,
	}
	for _, v := range validators {
		if err := v(); err != nil {
			return err
		}
	}
	if err := s.cts.Validate(); err != nil {
		return err
	}
	return s.x.Validate()
}

// validateSheetNames returns an error if any sheet names are too long
func (s Sheet) validateSheetNames() error {
	if len(s.Name()) > 31 {
		return fmt.Errorf("sheet name '%s' has %d characters, max length is 31", s.Name(), len(s.Name()))
	}
	return nil
}

// validateRowCellNumbers returns an error if any row numbers or cell numbers
// within a row are reused
func (s Sheet) validateRowCellNumbers() error {
	// check for re-used row numbers
	usedRows := map[uint32]struct{}{}
	for _, r := range s.x.SheetData.Row {
		if r.RAttr != nil {
			if _, reusedRow := usedRows[*r.RAttr]; reusedRow {
				return fmt.Errorf("'%s' reused row %d", s.Name(), *r.RAttr)
			}
			usedRows[*r.RAttr] = struct{}{}
		}
		// or re-used column labels within a row
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
	return nil
}

// validateMergedCells returns an error if merged cells overlap
func (s Sheet) validateMergedCells() error {
	mergedCells := map[uint64]struct{}{}
	for _, mc := range s.MergedCells() {
		from, to, err := reference.ParseRangeReference(mc.Reference())
		if err != nil {
			return fmt.Errorf("sheet name '%s' has invalid merged cell reference %s", s.Name(), mc.Reference())
		}
		for r := from.RowIdx; r <= to.RowIdx; r++ {
			for c := from.ColumnIdx; c <= to.ColumnIdx; c++ {
				idx := uint64(r)<<32 | uint64(c)
				if _, ok := mergedCells[idx]; ok {
					return fmt.Errorf("sheet name '%s' has overlapping merged cell range", s.Name())
				}
				mergedCells[idx] = struct{}{}
			}
		}

	}
	return nil
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
			rel := rel.AddAutoRelationship(gooxml.DocTypeSpreadsheet, gooxml.WorksheetType, i+1, gooxml.DrawingType)
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

// RangeReference converts a range reference of the form 'A1:A5' to 'Sheet
// 1'!$A$1:$A$5 . Renaming a sheet after calculating a range reference will
// invalidate the reference.
func (s Sheet) RangeReference(n string) string {
	sp := strings.Split(n, ":")
	cref, _ := reference.ParseCellReference(sp[0])
	from := fmt.Sprintf("$%s$%d", cref.Column, cref.RowIdx)
	if len(sp) == 1 {
		return fmt.Sprintf(`'%s'!%s`, s.Name(), from)
	}
	tref, _ := reference.ParseCellReference(sp[1])
	to := fmt.Sprintf("$%s$%d", tref.Column, tref.RowIdx)
	return fmt.Sprintf(`'%s'!%s:%s`, s.Name(), from, to)
}

const autoFilterName = "_xlnm._FilterDatabase"

// ClearAutoFilter removes the autofilters from the sheet.
func (s Sheet) ClearAutoFilter() {
	s.x.AutoFilter = nil
	sn := "'" + s.Name() + "'!"
	// see if we have a defined auto filter name for the sheet
	for _, dn := range s.w.DefinedNames() {
		if dn.Name() == autoFilterName {
			if strings.HasPrefix(dn.Content(), sn) {
				s.w.RemoveDefinedName(dn)
				break
			}
		}
	}
}

// SetAutoFilter creates autofilters on the sheet. These are the automatic
// filters that are common for a header row.  The RangeRef should be of the form
// "A1:C5" and cover the entire range of cells to be filtered, not just the
// header. SetAutoFilter replaces any existing auto filter on the sheet.
func (s Sheet) SetAutoFilter(rangeRef string) {
	// this should have no $ in it
	rangeRef = strings.Replace(rangeRef, "$", "", -1)

	s.x.AutoFilter = sml.NewCT_AutoFilter()
	s.x.AutoFilter.RefAttr = gooxml.String(rangeRef)
	sn := "'" + s.Name() + "'!"
	var sdn DefinedName

	// see if we already have a defined auto filter name for the sheet
	for _, dn := range s.w.DefinedNames() {
		if dn.Name() == autoFilterName {
			if strings.HasPrefix(dn.Content(), sn) {
				sdn = dn
				// name must match, but make sure rangeRef matches as well
				sdn.SetContent(s.RangeReference(rangeRef))
				break
			}
		}
	}
	// no existing name found, so add a new one
	if sdn.X() == nil {
		sdn = s.w.AddDefinedName(autoFilterName, s.RangeReference(rangeRef))
	}

	for i, ws := range s.w.xws {
		if ws == s.x {
			sdn.SetLocalSheetID(uint32(i))
		}
	}
}

// AddMergedCells merges cells within a sheet.
func (s Sheet) AddMergedCells(fromRef, toRef string) MergedCell {
	// TODO: we might need to actually create the merged cells if they don't
	// exist, but it appears to work fine on both Excel and LibreOffice just
	// creating the merged region

	if s.x.MergeCells == nil {
		s.x.MergeCells = sml.NewCT_MergeCells()
	}

	merge := sml.NewCT_MergeCell()
	merge.RefAttr = fmt.Sprintf("%s:%s", fromRef, toRef)

	s.x.MergeCells.MergeCell = append(s.x.MergeCells.MergeCell, merge)
	s.x.MergeCells.CountAttr = gooxml.Uint32(uint32(len(s.x.MergeCells.MergeCell)))
	return MergedCell{s.w, s.x, merge}
}

// MergedCells returns the merged cell regions within the sheet.
func (s Sheet) MergedCells() []MergedCell {
	if s.x.MergeCells == nil {
		return nil
	}
	ret := []MergedCell{}
	for _, c := range s.x.MergeCells.MergeCell {
		ret = append(ret, MergedCell{s.w, s.x, c})
	}
	return ret
}

// RemoveMergedCell removes merging from a cell range within a sheet.  The cells
// that made up the merged cell remain, but are no lon merged.
func (s Sheet) RemoveMergedCell(mc MergedCell) {
	for i, c := range s.x.MergeCells.MergeCell {
		if c == mc.X() {
			copy(s.x.MergeCells.MergeCell[i:], s.x.MergeCells.MergeCell[i+1:])
			s.x.MergeCells.MergeCell[len(s.x.MergeCells.MergeCell)-1] = nil
			s.x.MergeCells.MergeCell = s.x.MergeCells.MergeCell[:len(s.x.MergeCells.MergeCell)-1]
		}
	}

}

func (s Sheet) ExtentsIndex() (string, uint32, string, uint32) {
	var minRow, maxRow, minCol, maxCol uint32 = 1, 1, 0, 0
	for _, r := range s.Rows() {
		if r.RowNumber() < minRow {
			minRow = r.RowNumber()
		} else if r.RowNumber() > maxRow {
			maxRow = r.RowNumber()
		}

		for _, c := range r.Cells() {
			cref, err := reference.ParseCellReference(c.Reference())
			if err == nil {
				// column index is zero based here
				if cref.ColumnIdx < minCol {
					minCol = cref.ColumnIdx
				} else if cref.ColumnIdx > maxCol {
					maxCol = cref.ColumnIdx
				}
			}
		}
	}
	return reference.IndexToColumn(minCol), minRow, reference.IndexToColumn(maxCol), maxRow
}

// Extents returns the sheet extents in the form "A1:B15". This requires
// scanning the entire sheet.
func (s Sheet) Extents() string {
	sc, sr, ec, er := s.ExtentsIndex()
	return fmt.Sprintf("%s%d:%s%d", sc, sr, ec, er)
}

// AddConditionalFormatting adds conditional formatting to the sheet.
func (s Sheet) AddConditionalFormatting(cellRanges []string) ConditionalFormatting {
	cfmt := sml.NewCT_ConditionalFormatting()
	s.x.ConditionalFormatting = append(s.x.ConditionalFormatting, cfmt)

	// TODO: fix generator so this is not a pointer to a slice
	slc := make(sml.ST_Sqref, 0, 0)
	cfmt.SqrefAttr = &slc
	for _, r := range cellRanges {
		*cfmt.SqrefAttr = append(*cfmt.SqrefAttr, r)
	}
	return ConditionalFormatting{cfmt}
}

// Column returns or creates a column that with a given index (1-N).  Columns
// can span multiple column indices, this method will return the column that
// applies to a column index if it exists or create a new column that only
// applies to the index passed in otherwise.
func (s Sheet) Column(idx uint32) Column {
	// scan for any existing column that covers this index
	for _, colSet := range s.x.Cols {
		for _, col := range colSet.Col {
			if idx >= col.MinAttr && idx <= col.MaxAttr {
				return Column{col}
			}
		}
	}
	// does a column set exist?
	var colSet *sml.CT_Cols
	if len(s.x.Cols) == 0 {
		colSet = sml.NewCT_Cols()
		s.x.Cols = append(s.x.Cols, colSet)
	} else {
		colSet = s.x.Cols[0]
	}

	// create our new column
	col := sml.NewCT_Col()
	col.MinAttr = idx
	col.MaxAttr = idx
	colSet.Col = append(colSet.Col, col)
	return Column{col}
}

// Comments returns the comments for a sheet.
func (s Sheet) Comments() Comments {
	for i, wks := range s.w.xws {
		if wks == s.x {
			if s.w.comments[i] == nil {
				s.w.comments[i] = sml.NewComments()
				s.w.xwsRels[i].AddAutoRelationship(gooxml.DocTypeSpreadsheet, gooxml.WorksheetType, i+1, gooxml.CommentsType)
				s.w.ContentTypes.AddOverride(gooxml.AbsoluteFilename(gooxml.DocTypeSpreadsheet, gooxml.CommentsType, i+1), gooxml.CommentsContentType)
			}
			if len(s.w.vmlDrawings) == 0 {
				s.w.vmlDrawings = append(s.w.vmlDrawings, vmldrawing.NewCommentDrawing())
				vmlID := s.w.xwsRels[i].AddAutoRelationship(gooxml.DocTypeSpreadsheet, gooxml.WorksheetType, 1, gooxml.VMLDrawingType)
				if s.x.LegacyDrawing == nil {
					s.x.LegacyDrawing = sml.NewCT_LegacyDrawing()
				}
				s.x.LegacyDrawing.IdAttr = vmlID.ID()
			}
			return Comments{s.w, s.w.comments[i]}
		}
	}

	gooxml.Log("attempted to access comments for non-existent sheet")
	// should never occur
	return Comments{}
}

// SetBorder is a helper function for creating borders across multiple cells. In
// the OOXML spreadsheet format, a border applies to a single cell.  To draw a
// 'boxed' border around multiple cells, you need to apply different styles to
// the cells on the top,left,right,bottom and four corners.  This function
// breaks apart a single border into its components and applies it to cells as
// needed to give the effect of a border applying to multiple cells.
func (s Sheet) SetBorder(cellRange string, border Border) error {
	from, to, err := reference.ParseRangeReference(cellRange)
	if err != nil {
		return err
	}

	topLeftStyle := s.w.StyleSheet.AddCellStyle()
	topLeftBorder := s.w.StyleSheet.AddBorder()
	topLeftStyle.SetBorder(topLeftBorder)
	topLeftBorder.x.Top = border.x.Top
	topLeftBorder.x.Left = border.x.Left

	topRightStyle := s.w.StyleSheet.AddCellStyle()
	topRightBorder := s.w.StyleSheet.AddBorder()
	topRightStyle.SetBorder(topRightBorder)
	topRightBorder.x.Top = border.x.Top
	topRightBorder.x.Right = border.x.Right

	topStyle := s.w.StyleSheet.AddCellStyle()
	topBorder := s.w.StyleSheet.AddBorder()
	topStyle.SetBorder(topBorder)
	topBorder.x.Top = border.x.Top

	leftStyle := s.w.StyleSheet.AddCellStyle()
	leftBorder := s.w.StyleSheet.AddBorder()
	leftStyle.SetBorder(leftBorder)
	leftBorder.x.Left = border.x.Left

	rightStyle := s.w.StyleSheet.AddCellStyle()
	rightBorder := s.w.StyleSheet.AddBorder()
	rightStyle.SetBorder(rightBorder)
	rightBorder.x.Right = border.x.Right

	bottomStyle := s.w.StyleSheet.AddCellStyle()
	bottomBorder := s.w.StyleSheet.AddBorder()
	bottomStyle.SetBorder(bottomBorder)
	bottomBorder.x.Bottom = border.x.Bottom

	bottomLeftStyle := s.w.StyleSheet.AddCellStyle()
	bottomLeftBorder := s.w.StyleSheet.AddBorder()
	bottomLeftStyle.SetBorder(bottomLeftBorder)
	bottomLeftBorder.x.Bottom = border.x.Bottom
	bottomLeftBorder.x.Left = border.x.Left

	bottomRightStyle := s.w.StyleSheet.AddCellStyle()
	bottomRightBorder := s.w.StyleSheet.AddBorder()
	bottomRightStyle.SetBorder(bottomRightBorder)
	bottomRightBorder.x.Bottom = border.x.Bottom
	bottomRightBorder.x.Right = border.x.Right

	tlRowIdx := from.RowIdx
	tlColIdx := from.ColumnIdx
	brRowIdx := to.RowIdx
	brColIdx := to.ColumnIdx

	for row := tlRowIdx; row <= brRowIdx; row++ {
		for col := tlColIdx; col <= brColIdx; col++ {
			ref := fmt.Sprintf("%s%d", reference.IndexToColumn(col), row)
			switch {
			// top corners
			case row == tlRowIdx && col == tlColIdx:
				s.Cell(ref).SetStyle(topLeftStyle)
			case row == tlRowIdx && col == brColIdx:
				s.Cell(ref).SetStyle(topRightStyle)

			// bottom corners
			case row == brRowIdx && col == tlColIdx:
				s.Cell(ref).SetStyle(bottomLeftStyle)
			case row == brRowIdx && col == brColIdx:
				s.Cell(ref).SetStyle(bottomRightStyle)

			// four sides that aren't the corners
			case row == tlRowIdx:
				s.Cell(ref).SetStyle(topStyle)
			case row == brRowIdx:
				s.Cell(ref).SetStyle(bottomStyle)
			case col == tlColIdx:
				s.Cell(ref).SetStyle(leftStyle)
			case col == brColIdx:
				s.Cell(ref).SetStyle(rightStyle)
			}
		}
	}
	return nil
}

// AddDataValidation adds a data validation rule to a sheet.
func (s Sheet) AddDataValidation() DataValidation {
	if s.x.DataValidations == nil {
		s.x.DataValidations = sml.NewCT_DataValidations()
	}
	dv := sml.NewCT_DataValidation()
	dv.ShowErrorMessageAttr = gooxml.Bool(true)
	s.x.DataValidations.DataValidation = append(s.x.DataValidations.DataValidation, dv)
	s.x.DataValidations.CountAttr = gooxml.Uint32(uint32(len(s.x.DataValidations.DataValidation)))
	return DataValidation{dv}
}

// ClearCachedFormulaResults clears any computed formula values that are stored
// in the sheet. This may be required if you modify cells that are used as a
// formula input to force the formulas to be recomputed the next time the sheet
// is opened in Excel.
func (s *Sheet) ClearCachedFormulaResults() {
	for _, r := range s.Rows() {
		for _, c := range r.Cells() {
			if c.X().F != nil {
				c.X().V = nil
			}
		}
	}
}

// RecalculateFormulas re-computes any computed formula values that are stored
// in the sheet. As gooxml formula support is still new and not all functins are
// supported,  if formula execution fails either due to a parse error or missing
// function, or erorr in the result (even if expected) the cached value will be
// left empty allowing Excel to recompute it on load.
func (s *Sheet) RecalculateFormulas() {
	ev := formula.NewEvaluator()
	ctx := s.FormulaContext()
	for _, r := range s.Rows() {
		for _, c := range r.Cells() {
			if c.X().F != nil {
				formStr := c.X().F.Content

				// if the formula is shared, but the content is empty, then it's
				// a cell that a shared formula is applied to so we can ignreo
				// for now.  The value will be calculated later and applied with
				// setShared
				if c.X().F.TAttr == sml.ST_CellFormulaTypeShared && len(formStr) == 0 {
					continue
				}

				res := ev.Eval(ctx, formStr).AsString()
				if res.Type == formula.ResultTypeError {
					gooxml.Log("error evaulating formula %s: %s", formStr, res.ErrorMessage)
					c.X().V = nil
				} else {
					if res.Type == formula.ResultTypeNumber {
						c.X().TAttr = sml.ST_CellTypeN
					} else {
						c.X().TAttr = sml.ST_CellTypeInlineStr
					}
					c.X().V = gooxml.String(res.Value())

					// the formula is of type array, so if the result is also an
					// array we need to expand the array out into cells
					if c.X().F.TAttr == sml.ST_CellFormulaTypeArray && res.Type == formula.ResultTypeArray {
						s.setArray(c.Reference(), res)
					} else if c.X().F.TAttr == sml.ST_CellFormulaTypeShared && c.X().F.RefAttr != nil {
						from, to, err := reference.ParseRangeReference(*c.X().F.RefAttr)
						if err != nil {
							log.Printf("error in shared formula reference: %s", err)
							continue
						}
						s.setShared(c.Reference(), from, to, formStr)
					}
				}
			}
		}
	}
}

func (s *Sheet) setShared(origin string, from, to reference.CellReference, formulaStr string) {
	ctx := s.FormulaContext()
	ev := formula.NewEvaluator()

	for r := from.RowIdx; r <= to.RowIdx; r++ {
		for c := from.ColumnIdx; c <= to.ColumnIdx; c++ {
			rowOffset := r - from.RowIdx
			colOffset := c - from.ColumnIdx

			// add an offset on the contex so that cell references will be
			// offset during formula evaluation
			ctx.SetOffset(colOffset, rowOffset)
			res := ev.Eval(ctx, formulaStr)

			cr := fmt.Sprintf("%s%d", reference.IndexToColumn(c), r)
			c := s.Cell(cr)
			if res.Type == formula.ResultTypeNumber {
				c.X().TAttr = sml.ST_CellTypeN
			} else {
				c.X().TAttr = sml.ST_CellTypeInlineStr
			}
			c.X().V = gooxml.String(res.Value())
		}
	}
	_ = ev
	_ = ctx
}

// setArray expands an array into cached values starting at the origin which
// should be a cell reference of the type "A1". This is used when evaluating
// array type formulas.
func (s *Sheet) setArray(origin string, arr formula.Result) error {
	cref, err := reference.ParseCellReference(origin)
	if err != nil {
		return err
	}
	for ir, row := range arr.ValueArray {
		sr := s.Row(cref.RowIdx + uint32(ir))
		for ic, val := range row {
			cell := sr.Cell(reference.IndexToColumn(cref.ColumnIdx + uint32(ic)))
			cell.SetCachedFormulaResult(val.String())
		}
	}
	return nil
}

// SheetViews returns the sheet views defined.  This is where splits and frozen
// rows/cols are configured.  Multiple sheet views are allowed, but I'm not
// aware of there being a use for more than a single sheet view.
func (s *Sheet) SheetViews() []SheetView {
	if s.x.SheetViews == nil {
		return nil
	}
	r := []SheetView{}
	for _, sv := range s.x.SheetViews.SheetView {
		r = append(r, SheetView{sv})
	}
	return r
}

// AddView adds a sheet view.
func (s *Sheet) AddView() SheetView {
	if s.x.SheetViews == nil {
		s.x.SheetViews = sml.NewCT_SheetViews()
	}
	sv := sml.NewCT_SheetView()
	s.x.SheetViews.SheetView = append(s.x.SheetViews.SheetView, sv)
	return SheetView{sv}
}

// ClearSheetViews clears the list of sheet views.  This will clear the results
// of AddView() or SetFrozen.
func (s *Sheet) ClearSheetViews() {
	s.x.SheetViews = nil
}

// InitialView returns the first defined sheet view. If there are no views, one
// is created and returned.
func (s *Sheet) InitialView() SheetView {
	if s.x.SheetViews == nil || len(s.x.SheetViews.SheetView) == 0 {
		return s.AddView()
	}
	return SheetView{s.x.SheetViews.SheetView[0]}
}

// SetFrozen removes any existing sheet views and creates a new single view with
// either the first row, first column or both frozen.
func (s *Sheet) SetFrozen(firstRow, firstCol bool) {
	s.x.SheetViews = nil
	v := s.AddView()
	v.SetState(sml.ST_PaneStateFrozen)
	switch {
	case firstRow && firstCol:
		v.SetYSplit(1)
		v.SetXSplit(1)
		v.SetTopLeft("B2")
	case firstRow:
		v.SetYSplit(1)
		v.SetTopLeft("A2")
	case firstCol:
		v.SetXSplit(1)
		v.SetTopLeft("B1")
	}

}

// FormulaContext returns a formula evaluation context that can be used to
// evaluate formaulas.
func (s *Sheet) FormulaContext() formula.Context {
	return newEvalContext(s)
}

// ClearProtection removes any protections applied to teh sheet.
func (s *Sheet) ClearProtection() {
	s.x.SheetProtection = nil
}

// Protection controls the protection on an individual sheet.
func (s *Sheet) Protection() SheetProtection {
	if s.x.SheetProtection == nil {
		s.x.SheetProtection = sml.NewCT_SheetProtection()
	}
	return SheetProtection{s.x.SheetProtection}
}

// Sort sorts all of the rows within a sheet by the contents of a column. As the
// file format doesn't suppot indicating that a column should be sorted by the
// viewing/editing program, we actually need to reorder rows and change cell
// references during a sort. If the sheet contains formulas, you should call
// RecalculateFormulas() prior to sorting.  The column is in the form "C" and
// specifies the column to sort by. The firstRow is a 1-based index and
// specifies the firstRow to include in the sort, allowing skipping over a
// header row.
func (s *Sheet) Sort(column string, firstRow uint32, order SortOrder) {
	sheetData := s.x.SheetData.Row
	rows := s.Rows()
	// figure out which row to start the sort at
	for i, r := range rows {
		if r.RowNumber() == firstRow {
			sheetData = s.x.SheetData.Row[i:]
			break
		}
	}

	// perform the sort
	cmp := Comparer{Order: order}
	sort.Slice(sheetData, func(i, j int) bool {
		return cmp.LessRows(column,
			Row{s.w, s.x, sheetData[i]},
			Row{s.w, s.x, sheetData[j]})
	})

	// since we probably moved some rows, we need to go and fix up their row
	// number and cell references
	for i, r := range s.Rows() {
		correctRowNumber := uint32(i + 1)
		if r.RowNumber() != correctRowNumber {
			r.renumberAs(correctRowNumber)
		}
	}
}
