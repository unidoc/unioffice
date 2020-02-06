// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased on https://unidoc.io.

package formula

import "github.com/unidoc/unioffice/spreadsheet/reference"

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

// Reference returns a string reference value to a cell.
func (c CellRef) Reference(ctx Context, ev Evaluator) Reference {
	return Reference{Type: ReferenceTypeCell, Value: c.s}
}

// String returns a string representation of CellRef.
func (c CellRef) String() string {
	return c.s
}

// MoveLeft makes a reference to point to a cell which one step left from the original one after removing a column.
func (c CellRef) MoveLeft(q *MoveQuery) Expression {
	if q.MoveCurrentSheet {
		c.s = moveCellLeft(c.s, q.ColumnIdx)
	}
	return c
}

func moveCellLeft(refStr string, columnIdxToRemove uint32) string {
	ref, err := reference.ParseCellReference(refStr)
	if err != nil {
		return "#REF!"
	}
	columnIdx := ref.ColumnIdx
	if columnIdx < columnIdxToRemove {
		return refStr
	} else if columnIdx == columnIdxToRemove {
		return "#REF!"
	} else {
		return ref.MoveLeft().String()
	}
}
