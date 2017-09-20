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

func TestDefaultFormats(t *testing.T) {
	td := []struct {
		Inp float64
		Fmt spreadsheet.StandardFormat
		Exp string
	}{
		{0, spreadsheet.StandardFormatGeneral, "0"},
		{1.23, spreadsheet.StandardFormatGeneral, "1.23"},
		{5, spreadsheet.StandardFormatWholeNumber, "5"},
		{5, spreadsheet.StandardFormat2, "5.00"},
		{5, spreadsheet.StandardFormat3, "5"},
		{5123, spreadsheet.StandardFormat3, "5,123"},
		{5123, spreadsheet.StandardFormat4, "5,123.00"},
		{0, spreadsheet.StandardFormatPercent, "0%"},
		{.25, spreadsheet.StandardFormatPercent, "25%"},
		{1.2, spreadsheet.StandardFormatPercent, "120%"},
		{0, spreadsheet.StandardFormat10, "0.00%"},
		{.2502, spreadsheet.StandardFormat10, "25.02%"},
		{1.2, spreadsheet.StandardFormat10, "120.00%"},
	}
	wb := spreadsheet.New()
	sheet := wb.AddSheet()
	row := sheet.AddRow()
	for _, tc := range td {
		cell := row.AddCell()
		cell.SetNumber(tc.Inp)

		cs := wb.StyleSheet.AddCellStyle()
		cs.SetNumberFormatStandard(tc.Fmt)

		cell.SetStyle(cs)

		got := cell.GetFormattedValue()
		if got != tc.Exp {
			t.Errorf("expected %s for %f/%s, got %s", tc.Exp, tc.Inp, tc.Fmt, got)
		}
	}
}
