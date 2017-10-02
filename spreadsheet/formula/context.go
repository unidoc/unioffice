// Copyright 2017 Baliance. All rights reserved.
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
	// when evaluating cells that pull data from other sheets (e.g. ='Sheet 2'!A1)
	Sheet(name string) Context

	// NamedRange returns a named range.
	NamedRange(name string) Reference

	// SetOffset is used so that the Context can evaluate cell references
	// differently when they are not absolute (e.g. not like '$A$5').  See the
	// shared formula support in Cell for usage.
	SetOffset(col, row uint32)
}
