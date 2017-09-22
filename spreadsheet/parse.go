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

// ColumnToIndex maps a column to a zero based index (e.g. A = 0, B = 1, AA = 26)
func ColumnToIndex(col string) uint32 {
	col = strings.ToUpper(col)
	res := uint32(0)
	for _, c := range col {
		res *= 26
		res += uint32(c - 'A' + 1)
	}
	return res - 1
}

// IndexToColumn maps a column number to a coumn name (e.g. 0 = A, 1 = B, 26 = AA)
func IndexToColumn(col uint32) string {
	var a [64 + 1]byte
	i := len(a)
	u := col
	const b = 26
	for u >= b {
		i--
		q := u / b
		a[i] = byte('A' + uint(u-q*b))
		u = q - 1
	}
	i--
	a[i] = byte('A' + uint(u))

	return string(a[i:])
}
