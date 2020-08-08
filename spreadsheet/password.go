// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io.

package spreadsheet

import (
	"fmt"
)

// PasswordHash returns the password hash for a workbook using the modified
// spreadsheetML password hash that is compatible with Excel.
func PasswordHash(s string) string {
	hash := uint16(0)
	if len(s) > 0 {
		for i := len(s) - 1; i >= 0; i-- {
			c := s[i]
			hash = ((hash >> 14) & 0x01) | ((hash << 1) & 0x7fff)
			hash ^= uint16(c)
		}
		hash = ((hash >> 14) & 0x01) | ((hash << 1) & 0x7fff)
		hash ^= uint16(len(s))
		hash ^= (0x8000 | ('N' << 8) | 'K')
	}
	return fmt.Sprintf("%04X", uint64(hash))
}
