// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased on https://unidoc.io.

package reference_test

import (
	"testing"

	"github.com/unidoc/unioffice/spreadsheet/reference"
)

func TestParseCellReference(t *testing.T) {
	td := []struct {
		Inp      string
		Exp      reference.CellReference
		HasError bool
	}{
		{"A1", reference.CellReference{RowIdx: 1, ColumnIdx: 0, Column: "A"}, false},
		{"Z100", reference.CellReference{RowIdx: 100, ColumnIdx: 25, Column: "Z"}, false},
		{"", reference.CellReference{}, true},
		{"1", reference.CellReference{}, true},
		{"A", reference.CellReference{}, true},
		{"ABC", reference.CellReference{}, true},
		{"A1Z", reference.CellReference{}, true},
		{"$", reference.CellReference{}, true},
		{"$$", reference.CellReference{}, true},
		{"$$A", reference.CellReference{}, true},
		{"$$A1", reference.CellReference{}, true},
		{"$A1", reference.CellReference{RowIdx: 1, ColumnIdx: 0, Column: "A", AbsoluteColumn: true}, false},
		{"A$1", reference.CellReference{RowIdx: 1, ColumnIdx: 0, Column: "A", AbsoluteRow: true}, false},
		{"$A$1", reference.CellReference{RowIdx: 1, ColumnIdx: 0, Column: "A", AbsoluteRow: true, AbsoluteColumn: true}, false},
		{"$D$15", reference.CellReference{RowIdx: 15, ColumnIdx: 3, Column: "D", AbsoluteRow: true, AbsoluteColumn: true}, false},
	}
	for _, tc := range td {
		cref, err := reference.ParseCellReference(tc.Inp)
		if tc.HasError {
			if err == nil {
				t.Errorf("expected error for input %s", tc.Inp)
			}
			// expected an error, so don't check anything else
			continue
		} else if err != nil {
			t.Errorf("expected no error for input %s, got %s", tc.Inp, err)
		}

		if cref.RowIdx != tc.Exp.RowIdx {
			t.Errorf("expected row = %d for %s, got %d", tc.Exp.RowIdx, tc.Inp, cref.RowIdx)
		}
		if cref.ColumnIdx != tc.Exp.ColumnIdx {
			t.Errorf("expected column = %d for %s, got %d", tc.Exp.ColumnIdx, tc.Inp, cref.ColumnIdx)
		}
		if cref.Column != tc.Exp.Column {
			t.Errorf("expected column = %s for %s, got %s", tc.Exp.Column, tc.Inp, cref.Column)
		}
		if cref.AbsoluteRow != tc.Exp.AbsoluteRow {
			t.Errorf("expected absolute-row = %v for %s, got %v", tc.Exp.AbsoluteRow, tc.Inp, cref.AbsoluteRow)
		}
		if cref.AbsoluteColumn != tc.Exp.AbsoluteColumn {
			t.Errorf("expected absolute-column = %v for %s, got %v", tc.Exp.AbsoluteColumn, tc.Inp, cref.AbsoluteColumn)
		}
	}
}

func TestParseCellRangeReference(t *testing.T) {
	td := []struct {
		Inp      string
		ExpFrom  reference.CellReference
		ExpTo    reference.CellReference
		HasError bool
	}{
		{"A1:B2",
			reference.CellReference{RowIdx: 1, ColumnIdx: 0, Column: "A"},
			reference.CellReference{RowIdx: 2, ColumnIdx: 1, Column: "B"},
			false},
		{"$A1:B$2",
			reference.CellReference{RowIdx: 1, ColumnIdx: 0, Column: "A", AbsoluteColumn: true},
			reference.CellReference{RowIdx: 2, ColumnIdx: 1, Column: "B", AbsoluteRow: true},
			false},
		{"A1",
			reference.CellReference{},
			reference.CellReference{},
			true},
		{"A1:",
			reference.CellReference{},
			reference.CellReference{},
			true},
		{"A1:A",
			reference.CellReference{},
			reference.CellReference{},
			true},
		{"ABC:BAC",
			reference.CellReference{},
			reference.CellReference{},
			true},
	}

	for _, tc := range td {
		from, to, err := reference.ParseRangeReference(tc.Inp)
		if tc.HasError {
			if err == nil {
				t.Errorf("expected error for input %s", tc.Inp)
			}
			// expected an error, so don't check anything else
			continue
		} else if err != nil {
			t.Errorf("expected no error for input %s, got %s", tc.Inp, err)
		}
		if from != tc.ExpFrom {
			t.Errorf("expected from=%v, got %v for %s", tc.ExpFrom, from, tc.Inp)
		}
		if to != tc.ExpTo {
			t.Errorf("expected to=%v, got %v for %s", tc.ExpTo, to, tc.Inp)
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

		if got := reference.ColumnToIndex(tc.Inp); got != tc.ExpIndex {
			t.Errorf("expected %s = %d, got %d", tc.Inp, tc.ExpIndex, got)
		}
		if got := reference.IndexToColumn(tc.ExpIndex); got != tc.Inp {
			t.Errorf("expected %d = %s, got %s", tc.ExpIndex, tc.Inp, got)
		}
	}
}

func TestColumnToIndexRoundTrip(t *testing.T) {
	for i := 0; i < 10000; i++ {
		s := reference.IndexToColumn(uint32(i))
		got := reference.ColumnToIndex(s)
		if got != uint32(i) {
			t.Errorf("failed on %d %s", i, s)
		}
	}
}
