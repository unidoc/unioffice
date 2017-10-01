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

func TestParseCellReference(t *testing.T) {
	td := []struct {
		Inp      string
		ExpCol   string
		ExpRow   uint32
		HasError bool
	}{
		{"A1", "A", 1, false},
		{"B25", "B", 25, false},
		{"AZ9", "AZ", 9, false},
		{"A", "", 0, true},
		{"1", "", 0, true},
	}
	for _, tc := range td {
		col, row, err := spreadsheet.ParseCellReference(tc.Inp)
		if tc.HasError {
			if err == nil {
				t.Errorf("expected error for input %s", tc.Inp)
			}
		} else if err != nil {
			t.Errorf("expected no error for input %s, got %s", tc.Inp, err)
		}

		if col != tc.ExpCol {
			t.Errorf("expected col = %s for %s, got %s", tc.ExpCol, tc.Inp, col)
		}

		if row != tc.ExpRow {
			t.Errorf("expected row = %d for %s, got %d", tc.ExpRow, tc.Inp, row)
		}
	}
}
