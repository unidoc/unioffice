// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheet

import (
	"fmt"
	"time"

	"github.com/unidoc/unioffice/spreadsheet/formula"
	"github.com/unidoc/unioffice/spreadsheet/reference"
)

func newEvalContext(s *Sheet) *evalContext {
	return &evalContext{s: s, evaluating: make(map[string]struct{})}
}

type evalContext struct {
	s              *Sheet
	colOff, rowOff uint32
	evaluating     map[string]struct{}
}

func (e *evalContext) Cell(ref string, ev formula.Evaluator) formula.Result {
	cr, err := reference.ParseCellReference(ref)
	if err != nil {
		return formula.MakeErrorResult(fmt.Sprintf("error parsing %s: %s", ref, err))
	}

	// offsets are used in shared formulas so that references like 'A1', '$A1',
	// 'A$1', '$A$1' will behave differently according to the offset
	if e.colOff != 0 && !cr.AbsoluteColumn {
		cr.ColumnIdx += e.colOff
		cr.Column = reference.IndexToColumn(cr.ColumnIdx)
	}
	if e.rowOff != 0 && !cr.AbsoluteRow {
		cr.RowIdx += e.rowOff
	}

	c := e.s.Cell(cr.String())

	// if we have a formula, evaluate it
	if c.HasFormula() {
		if _, ok := e.evaluating[ref]; ok {
			// recursively evaluating, so bail out
			return formula.MakeErrorResult("recursion detected during evaluation of " + ref)
		}
		e.evaluating[ref] = struct{}{}
		res := ev.Eval(e, c.GetFormula())
		delete(e.evaluating, ref)
		return res
	}

	if c.IsEmpty() {
		return formula.MakeEmptyResult()
	} else if c.IsNumber() {
		v, _ := c.GetValueAsNumber()
		return formula.MakeNumberResult(v)
	} else if c.IsBool() {
		v, _ := c.GetValueAsBool()
		return formula.MakeBoolResult(v)
	}

	v, _ := c.GetRawValue()
	return formula.MakeStringResult(v)
}

func (e *evalContext) Sheet(name string) formula.Context {
	for _, sheet := range e.s.w.Sheets() {
		if sheet.Name() == name {
			return sheet.FormulaContext()
		}
	}
	return formula.InvalidReferenceContext
}

func (e *evalContext) NamedRange(ref string) formula.Reference {
	for _, dn := range e.s.w.DefinedNames() {
		if dn.Name() == ref {
			return formula.MakeRangeReference(dn.Content())
		}
	}
	for _, tbl := range e.s.w.Tables() {
		if tbl.Name() == ref {
			return formula.MakeRangeReference(tbl.Reference())
		}
	}
	return formula.ReferenceInvalid
}

func (e *evalContext) SetOffset(col, row uint32) {
	e.colOff = col
	e.rowOff = row
}

// GetFilename returns the filename of the context's workbook.
func (e *evalContext) GetFilename() string {
	return e.s.w.GetFilename()
}

// GetFormat returns a cell data format.
func (e *evalContext) GetFormat(cellRef string) string {
	return e.s.Cell(cellRef).getFormat()
}

// GetLabelPrefix returns label prefix which depends on the cell's horizontal alignment.
func (e *evalContext) GetLabelPrefix(cellRef string) string {
	return e.s.Cell(cellRef).getLabelPrefix()
}

// GetLocked returns true if the cell is locked.
func (e *evalContext) GetLocked(cellRef string) bool {
	return e.s.Cell(cellRef).getLocked()
}

// HasFormula returns true if the cell contains formula.
func (e *evalContext) HasFormula(cellRef string) bool {
	return e.s.Cell(cellRef).HasFormula()
}

// IsBool returns true if the cell boolean value.
func (e *evalContext) IsBool(cellRef string) bool {
	return e.s.Cell(cellRef).IsBool()
}

// SetLocked sets cell locked or not.
func (e *evalContext) SetLocked(cellRef string, locked bool) {
	e.s.Cell(cellRef).setLocked(locked)
}

// GetWidth returns a worksheet's column width.
func (e *evalContext) GetWidth(colIdx int) float64 {
	colIdx++
	for _, c := range e.s.X().Cols[0].Col {
		if int(c.MinAttr) <= colIdx && colIdx <= int(c.MaxAttr) {
			return float64(int(*c.WidthAttr))
		}
	}
	return 0
}

// GetEpoch returns a workbook's time epoch.
func (e *evalContext) GetEpoch() time.Time {
	return e.s.w.Epoch()
}

//dbcs is a list of languages which use DBCS or double-byte character set (Chinese dialects, Japanese and Korean)
var dbcs []string = []string{
	"zh-HK",
	"zh-MO",
	"zh-CN",
	"zh-SG",
	"zh-TW",
	"ja-JP",
	"ko-KR",
}

// IsDBCS returns if a workbook's default language is among DBCS.
func (e *evalContext) IsDBCS() bool {
	language := e.s.w.CoreProperties.X().Language
	if language == nil {
		return false
	}
	defaultLanguage := string(language.Data)
	for _, lang := range dbcs {
		if defaultLanguage == lang {
			return true
		}
	}
	return false
}
