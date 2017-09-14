// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package formula

// CellRef is a reference to a single cell
type CellRef struct {
	s string
}

// NewCellRef constructs a new cell reference.
func NewCellRef(v string) Expression {
	return CellRef{v}
}

// Eval evaluates and returns the result of the cell reference.
func (c CellRef) Eval(ctx Context, ev Evaluator) Result {
	return ctx.Cell(c.s, ev)
}

func (c CellRef) Reference(ctx Context, ev Evaluator) Reference {
	return Reference{Type: ReferenceTypeCell, Value: c.s}
}
