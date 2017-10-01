// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheet

import (
	"fmt"
	"strconv"
	"strings"
)

// ParseRangeReference splits a range reference of the form "A1:B5" into its
// components.
func ParseRangeReference(s string) (from, to string, err error) {
	sp := strings.Split(s, ":")
	if len(sp) == 2 {
		return sp[0], sp[1], nil
	}
	return "", "", fmt.Errorf("invaid range reference: %s", s)
}

// ParseCellReference parses a cell reference of the form 'A10' and splits it
// into column/row segments.
func ParseCellReference(s string) (col string, row uint32, err error) {
	s = strings.Replace(s, "$", "", -1)
	split := -1
lfor:
	for i := 0; i < len(s); i++ {
		switch {
		case s[i] >= '0' && s[i] <= '9':
			split = i
			break lfor
		}
	}
	switch split {
	case 0:
		return col, row, fmt.Errorf("no letter prefix in %s", s)
	case -1:
		return col, row, fmt.Errorf("no digits in %s", s)
	}

	col = s[0:split]
	r64, err := strconv.ParseUint(s[split:], 10, 32)
	row = uint32(r64)
	return col, row, err
}
