// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased on https://unidoc.io.

package reference

import (
	"errors"
	"strings"
)

// ParseRangeReference splits a range reference of the form "A1:B5" into its
// components.
func ParseRangeReference(s string) (from, to CellReference, err error) {
	sheetName := ""
	sp0 := strings.Split(s, "!")
	if len(sp0) == 2 {
		sheetName = sp0[0]
		s = sp0[1]
	}
	sp := strings.Split(s, ":")
	if len(sp) != 2 {
		return CellReference{}, CellReference{}, errors.New("invalid range format")
	}

	if sheetName != "" {
		sp[0] = sheetName + "!" + sp[0]
		sp[1] = sheetName + "!" + sp[1]
	}
	fromRef, err := ParseCellReference(sp[0])
	if err != nil {
		return CellReference{}, CellReference{}, err
	}

	toRef, err := ParseCellReference(sp[1])
	if err != nil {
		return CellReference{}, CellReference{}, err
	}
	return fromRef, toRef, nil
}

// ParseColumnRangeReference splits a range reference of the form "A:B" into its
// components.
func ParseColumnRangeReference(s string) (from, to ColumnReference, err error) {
	sheetName := ""
	sp0 := strings.Split(s, "!")
	if len(sp0) == 2 {
		sheetName = sp0[0]
		s = sp0[1]
	}
	sp := strings.Split(s, ":")
	if len(sp) != 2 {
		return ColumnReference{}, ColumnReference{}, errors.New("invalid range format")
	}

	if sheetName != "" {
		sp[0] = sheetName + "!" + sp[0]
		sp[1] = sheetName + "!" + sp[1]
	}
	fromRef, err := ParseColumnReference(sp[0])
	if err != nil {
		return ColumnReference{}, ColumnReference{}, err
	}

	toRef, err := ParseColumnReference(sp[1])
	if err != nil {
		return ColumnReference{}, ColumnReference{}, err
	}
	return fromRef, toRef, nil
}
