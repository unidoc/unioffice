// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io.

package algo

import (
	"strconv"
)

func isdigit(c byte) bool {
	return c >= '0' && c <= '9'
}

// NaturalLess compares two strings in a human manner so rId2 sorts less than rId10
func NaturalLess(lhs, rhs string) bool {
	lidx, ridx := 0, 0
	for lidx < len(lhs) && ridx < len(rhs) {
		lc := lhs[lidx]
		rc := rhs[ridx]
		ldigit := isdigit(lc)
		rdigit := isdigit(rc)
		switch {
		// digits sort before characters
		case ldigit && !rdigit:
			return true
		// characters after digits
		case !ldigit && rdigit:
			return false
		// no digits, so compare the characters
		case !ldigit && !rdigit:
			if lc != rc {
				return lc < rc
			}
			lidx++
			ridx++
		// both digits, so parse and compare
		default:
			lend := lidx + 1
			rend := ridx + 1

			for lend < len(lhs) && isdigit(lhs[lend]) {
				lend++
			}
			for rend < len(rhs) && isdigit(rhs[rend]) {
				rend++
			}
			lv, _ := strconv.ParseUint(lhs[lidx:lend], 10, 64)
			rv, _ := strconv.ParseUint(rhs[lidx:rend], 10, 64)
			if lv != rv {
				return lv < rv
			}
			// digits are equal, so keep looking
			lidx = lend
			ridx = rend
		}
	}
	// fall back to comparing length
	return len(lhs) < len(rhs)
}
