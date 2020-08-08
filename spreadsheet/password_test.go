// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io.

package spreadsheet_test

import (
	"testing"

	"github.com/unidoc/unioffice/spreadsheet"
)

func TestKnownHashes(t *testing.T) {
	td := []struct {
		Inp string
		Exp string
	}{
		{"gooxml", "DD67"},
		{"", "0000"},
	}
	for _, tc := range td {
		if got := spreadsheet.PasswordHash(tc.Inp); got != tc.Exp {
			t.Errorf("expected hash of %s = %s, got %s", tc.Inp, tc.Exp, got)
		}
	}
}
