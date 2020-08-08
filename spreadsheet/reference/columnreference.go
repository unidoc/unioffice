// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io.

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
