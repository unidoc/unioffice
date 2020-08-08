// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io.

package reference

import (
	"strings"
)

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

// IndexToColumn maps a column number to a column name (e.g. 0 = A, 1 = B, 26 = AA)
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
