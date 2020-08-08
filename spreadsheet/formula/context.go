// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io.

package formula

import "time"

// Context is a formula execution context.  Formula evaluation uses the context
// to retreive information from sheets.
type Context interface {
	// Cell returns the result of evaluating a cell.
	Cell(ref string, ev Evaluator) Result

	// Sheet returns an evaluation context for a given sheet name.  This is used
	// when evaluating cells that pull data from other sheets (e.g. ='Sheet 2'!A1).
	Sheet(name string) Context

	// GetEpoch returns the time epoch of the context's Workbook.
	GetEpoch() time.Time

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

	// HasFormula returns if cell contains formula.
	HasFormula(cellRef string) bool

	// IsBool returns if cell contains boolean value.
	IsBool(cellRef string) bool

	// IsDBCS returns if workbook default language is among DBCS.
	IsDBCS() bool

	// LastColumn returns the name of last column which contains data in range of context sheet's given rows.
	LastColumn(rowFrom, rowTo int) string

	// LastRow returns the name of last row which contains data in range of context sheet's given columns.
	LastRow(colFrom string) int

	// SetLocked returns sets cell's protected attribute.
	SetLocked(cellRef string, locked bool)

	// NamedRange returns a named range.
	NamedRange(name string) Reference

	// SetOffset is used so that the Context can evaluate cell references
	// differently when they are not absolute (e.g. not like '$A$5').  See the
	// shared formula support in Cell for usage.
	SetOffset(col, row uint32)
}
