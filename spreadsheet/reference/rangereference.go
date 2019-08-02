// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package reference

import (
	"errors"
	"strings"
)

// ParseRangeReference splits a range reference of the form "A1:B5" into its
// components.
func ParseRangeReference(s string) (from, to CellReference, err error) {
	sp := strings.Split(s, ":")
	if len(sp) != 2 {
		return CellReference{}, CellReference{}, errors.New("invalid range format")
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
