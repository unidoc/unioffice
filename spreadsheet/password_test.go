// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheet_test

import (
	"testing"

	"baliance.com/gooxml/spreadsheet"
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
