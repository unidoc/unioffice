// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io.

package format_test

import (
	"testing"

	"github.com/unidoc/unioffice/spreadsheet/format"
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
