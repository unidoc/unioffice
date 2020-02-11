// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased on https://unidoc.io.

package formula

import (
	"github.com/unidoc/unioffice/spreadsheet/reference"
	"github.com/unidoc/unioffice/spreadsheet/update"
)

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

// Update makes a reference to point to one of the neighboring cells after removing a row/column with respect to the update type.
func (c CellRef) Update(q *update.UpdateQuery) Expression {
	if q.UpdateCurrentSheet {
		c.s = updateRefStr(c.s, q)
	}
	return c
}

// updateRefStr gets reference string representation like C1, parses it and makes a string representation of a new reference with respect to the update type (e.g. B1 if a column to the left of this reference was removed).
func updateRefStr(refStr string, q *update.UpdateQuery) string {
	ref, err := reference.ParseCellReference(refStr)
	if err != nil {
		return "#REF!"
	}
	if q.UpdateType == update.UpdateActionRemoveColumn {
		columnIdxToRemove := q.ColumnIdx
		columnIdx := ref.ColumnIdx
		if columnIdx < columnIdxToRemove {
			return refStr
		} else if columnIdx == columnIdxToRemove {
			return "#REF!"
		} else {
			return ref.Update(update.UpdateActionRemoveColumn).String()
		}
	}
	return refStr
}
