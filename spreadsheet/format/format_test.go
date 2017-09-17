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

func TestCellFormattingNumber(t *testing.T) {
	td := []struct {
		Inp float64
		Fmt string
		Exp string
	}{
		// General format, gathered by testing with Mac Excel 365
		{1.0, "", "1"},
		{1.23, "", "1.23"},
		{123.456789, "", "123.456789"},
		{12341234.12341234, "", "12341234.12"},
		{12341234.125, "", "12341234.13"},
		{123412341234, "", "1.23412E+11"},
		{12345.12341234, "", "12345.12341"},
		{1e11, "", "1.00E+11"},
		{1e308, "", "1.00E+308"},
		{1234e15, "", "1.23E+18"},
		{1e-10, "", "1.00E-10"},
		{99999999.995, "", "100000000"},
		{.123400, "", "0.1234"},
		{.123412341234, "", "0.123412341"},
		{1.000000000001, "", "1"},
		{.9999999, "", "0.9999999"},
		{.99999999, "", "0.99999999"},
		{0.123999999, "", "0.124"},
		{10.19999999, "", "10.2"},
		{.999999999, "", "1"},
		{.9999999999, "", "1"},
	}
	for _, tc := range td {

		got := format.Number(tc.Inp, tc.Fmt)
		if got != tc.Exp {
			t.Errorf("expected %s, got %s for %f %s", tc.Exp, got, tc.Inp, tc.Fmt)
		}
	}
}
