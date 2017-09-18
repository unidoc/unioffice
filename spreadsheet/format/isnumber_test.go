// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package format_test

import (
	"testing"

	"baliance.com/gooxml/spreadsheet/format"
)

func TestIsNumber(t *testing.T) {
	td := []struct {
		Inp string
		Exp bool
	}{
		{"123", true},
		{"1.23", true},
		{"1.23.", false},
		{"1.23E+10", true},
		{"1.23E-10", true},
		{"1.23E10", false},
		{"1213131312312312390", true},
		{"0", true},
		{"", false},
		{"abc", false},
	}
	for _, tc := range td {
		got := format.IsNumber(tc.Inp)
		if got != tc.Exp {
			t.Errorf("expected IsNumber(%s) = %v, got %v", tc.Inp, tc.Exp, got)
		}
	}
}
