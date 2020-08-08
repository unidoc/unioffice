// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io.

package spreadsheet

import (
	"strconv"

	"github.com/unidoc/unioffice/spreadsheet/reference"
)

// SortOrder is a column sort order.
//go:generate stringer -type=SortOrder
type SortOrder byte

// SortOrder constants
const (
	SortOrderAscending SortOrder = iota
	SortOrderDescending
)

// Comparer is used to compare rows based off a column and cells based off of
// their value.
type Comparer struct {
	Order SortOrder
}

// LessRows compares two rows based off of a column. If the column doesn't exist
// in one row, that row is 'less'.
func (c Comparer) LessRows(column string, lhs, rhs Row) bool {
	var lhsCell, rhsCell Cell
	for _, c := range lhs.Cells() {
		cref, _ := reference.ParseCellReference(c.Reference())
		if cref.Column == column {
			lhsCell = c
			break
		}
	}

	for _, c := range rhs.Cells() {
		cref, _ := reference.ParseCellReference(c.Reference())
		if cref.Column == column {
			rhsCell = c
			break
		}
	}

	return c.LessCells(lhsCell, rhsCell)
}

// LessCells returns true if the lhs value is less than the rhs value. If the
// cells contain numeric values, their value interpreted as a floating point is
// compared. Otherwise their string contents are compared.
func (c Comparer) LessCells(lhs, rhs Cell) bool {
	if c.Order == SortOrderDescending {
		lhs, rhs = rhs, lhs
	}

	// handle zero-value cells first as we can get those based off of LessRows
	// above
	if lhs.X() == nil {
		if rhs.X() == nil {
			return false
		}
		return true
	}
	if rhs.X() == nil {
		return false
	}

	lhsValue, lhsIsNum := lhs.getRawSortValue()
	rhsValue, rhsIsNum := rhs.getRawSortValue()

	switch {
	// both numbers
	case lhsIsNum && rhsIsNum:
		lf, _ := strconv.ParseFloat(lhsValue, 64)
		rf, _ := strconv.ParseFloat(rhsValue, 64)
		return lf < rf
	// numbers sort before non-numbers
	case lhsIsNum:
		return true
	case rhsIsNum:
		return false
	}

	lhsValue = lhs.GetFormattedValue()
	rhsValue = rhs.GetFormattedValue()
	return lhsValue < rhsValue
}
