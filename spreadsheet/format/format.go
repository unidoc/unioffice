// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package format

import (
	"math"
	"strconv"
)

const maxGeneric = 1e11
const minGeneric = 1e-10

// NNumber is used to format a number with a format string.  If the format
// string is empty, then General number formatting is used which attempts to mimic
// Excel's general formatting.
func Number(v float64, fmt string) string {
	if fmt == "" {
		return NumberGeneric(v)
	}
	buf := make([]byte, 0, 15)
	buf = strconv.AppendFloat(buf, v, 'f', -1, 64)
	buf = trimTrailingZeros(buf)
	return string(buf)
}

// NumberGeneric formats the number with the generic format which attemps to
// mimic Excel's general formatting.
func NumberGeneric(v float64) string {
	if math.Abs(v) >= maxGeneric || math.Abs(v) <= minGeneric {
		return formatExpNumberGeneric(v)
	}

	b := make([]byte, 0, 15)
	b = strconv.AppendFloat(b, v, 'f', -1, 64)
	if len(b) > 11 {
		ntlDigit := b[11] - '0'
		// if after truncating, do we need to round?
		if ntlDigit >= 5 && ntlDigit <= 9 {
			b[10]++
			b = b[0:11]
			b = performCarries(b)
		}
		b = b[0:11]
	} else if len(b) == 11 {
		// Excel rounds up a series of consecutive 9's if the total length is 11 digits
		if b[len(b)-1] == '9' {
			b[len(b)-1]++
			b = performCarries(b)
		}
	}

	b = trimTrailingZeros(b)
	return string(b)
}

func formatExpNumberGeneric(v float64) string {
	s1 := strconv.FormatFloat(v, 'E', -1, 64)
	s2 := strconv.FormatFloat(v, 'E', 5, 64)
	if len(s1) < len(s2) {
		return strconv.FormatFloat(v, 'E', 2, 64)
	}
	return s2
}

// trim any trailing zeros for numbers like '1.23000000' => '1.23' while not
// triming zeros on things like '10000'
func trimTrailingZeros(b []byte) []byte {
	end := len(b)
	sawDecimal := false
	sawDigit := false
	for i := len(b) - 1; i >= 0; i-- {
		if b[i] == '0' && !sawDigit && !sawDecimal {
			end = i
		} else if b[i] == '.' {
			sawDecimal = true
		} else {
			sawDigit = true
		}
	}
	if sawDecimal && sawDigit {
		if b[end-1] == '.' {
			end--
		}
		return b[0:end]
	}
	return b
}

func performCarries(b []byte) []byte {
	// perform any required carries caused by rounding
	for i := len(b) - 1; i > 0; i-- {
		if b[i] == '9'+1 {
			b[i] = '0'
			if b[i-1] == '.' {
				i--
			}
			b[i-1]++
		}
	}
	// perform carry on initial digit by prepending a 1
	if b[0] == '9'+1 {
		b[0] = '0'
		copy(b[1:], b[0:])
		b[0] = '1'
	}
	return b
}
