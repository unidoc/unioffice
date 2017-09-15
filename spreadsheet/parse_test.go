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

func TestColumnToIndex(t *testing.T) {
	td := []struct {
		Inp      string
		ExpIndex uint32
	}{
		{"A", 0},
		{"B", 1},
		{"Z", 25},
		{"AA", 26},
		{"AB", 27},
		{"AC", 28},
		{"BZ", 77},
		{"CA", 78},
		{"GOOXML", 90304485},
	}
	for _, tc := range td {

		if got := spreadsheet.ColumnToIndex(tc.Inp); got != tc.ExpIndex {
			t.Errorf("expected %s = %d, got %d", tc.Inp, tc.ExpIndex, got)
		}
		if got := spreadsheet.IndexToColumn(tc.ExpIndex); got != tc.Inp {
			t.Errorf("expected %d = %s, got %s", tc.ExpIndex, tc.Inp, got)
		}
	}
}

func TestColumnToIndexRoundTrip(t *testing.T) {
	for i := 0; i < 10000; i++ {
		s := spreadsheet.IndexToColumn(uint32(i))
		got := spreadsheet.ColumnToIndex(s)
		if got != uint32(i) {
			t.Errorf("failed on %d %s", i, s)
		}
	}
}
