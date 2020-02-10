// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased on https://unidoc.io.

package reference

import (
	"errors"
	"regexp"
	"strings"

	"github.com/unidoc/unioffice/spreadsheet/update"
)

// ColumnReference is a parsed reference to a column.  Input is of the form 'A',
// '$C', etc.
type ColumnReference struct {
	ColumnIdx      uint32
	Column         string
	AbsoluteColumn bool
	SheetName      string
}

// String returns a string representation of ColumnReference.
func (c ColumnReference) String() string {
	buf := make([]byte, 0, 4)
	if c.AbsoluteColumn {
		buf = append(buf, '$')
	}
	buf = append(buf, c.Column...)
	return string(buf)
}

var reColumn = regexp.MustCompile(`^[a-zA-Z]([a-zA-Z]?)$`)

// ParseColumnReference parses a column reference of the form 'Sheet1!A' and splits it
// into sheet name and column segments.
func ParseColumnReference(s string) (ColumnReference, error) {
	s = strings.TrimSpace(s)
	if len(s) < 1 {
		return ColumnReference{}, errors.New("column reference must have at least one character")
	}

	r := ColumnReference{}
	sl := strings.Split(s, "!")
	if len(sl) == 2 {
		r.SheetName = sl[0]
		s = sl[1]
	}
	// check for absolute column
	if s[0] == '$' {
		r.AbsoluteColumn = true
		s = s[1:]
	}

	if !reColumn.MatchString(s) {
		return ColumnReference{}, errors.New("column reference must be between A and ZZ")
	}

	r.Column = s

	r.ColumnIdx = ColumnToIndex(r.Column)
	return r, nil
}

// Update updates reference to point one of the neighboring columns with respect to the update type after removing a row/column.
func (ref *ColumnReference) Update(updateType update.UpdateAction) *ColumnReference {
	switch updateType {
	case update.UpdateActionRemoveColumn:
		newRef := ref
		newRef.ColumnIdx = ref.ColumnIdx - 1
		newRef.Column = IndexToColumn(newRef.ColumnIdx)
		return newRef
	default:
		return ref
	}
}
