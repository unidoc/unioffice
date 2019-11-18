// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package formula

// Context is a formula execution context.  Formula evaluation uses the context
// to retreive information from sheets.
type Context interface {
	// Cell returns the result of evaluating a cell.
	Cell(ref string, ev Evaluator) Result

	// Sheet returns an evaluation context for a given sheet name.  This is used
	// when evaluating cells that pull data from other sheets (e.g. ='Sheet 2'!A1).
	Sheet(name string) Context

	// GetFilename returns the full filename of the context's Workbook.
	GetFilename() string

	// GetWidth returns a worksheet's column width.
	GetWidth(colIdx int) float64

	// GetFormat returns a cell's format.
	GetFormat(cellRef string) string

	// GetLabelPrefix returns cell's label prefix dependent on cell horizontal alignment.
	GetLabelPrefix(cellRef string) string

	// GetFormat returns if cell is protected.
	GetLocked(cellRef string) bool

	// GetFormat returns sets cell's protected attribute.
	SetLocked(cellRef string, locked bool)

	// NamedRange returns a named range.
	NamedRange(name string) Reference

	// SetOffset is used so that the Context can evaluate cell references
	// differently when they are not absolute (e.g. not like '$A$5').  See the
	// shared formula support in Cell for usage.
	SetOffset(col, row uint32)
}
