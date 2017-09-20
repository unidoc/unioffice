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

var standardFmtTestData = []struct {
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
	{0.5, spreadsheet.StandardFormat11, "5.00E-01"},
	{0.5, spreadsheet.StandardFormat12, "1/2"},
	{1.5, spreadsheet.StandardFormat12, "1 1/2"},
	{3.25, spreadsheet.StandardFormat12, "3 1/4"},
	{1.5, spreadsheet.StandardFormat13, "1 1/2"},
	// Excel gives 2 9/31 here, but 2 20/69 is closer to 2.29 and is what we
	// compute. I'm not sure what logic Excel is using, so I'm not going to
	// bother investigating for now.
	{2.29, spreadsheet.StandardFormat13, "2 20/69"},

	{42996.6996269676, spreadsheet.StandardFormat14, "9/18/17"},
	{42996.6996269676, spreadsheet.StandardFormat15, "18-Sep-17"},
	{42996.6996269676, spreadsheet.StandardFormat16, "18-Sep"},
	{42996.6996269676, spreadsheet.StandardFormat17, "Sep-17"},
	{42996.6996269676, spreadsheet.StandardFormat18, "4:47 PM"},
	{42996.6996269676, spreadsheet.StandardFormat19, "4:47:28 PM"},
	{42996.6996269676, spreadsheet.StandardFormat20, "4:47"},
	{42996.6996269676, spreadsheet.StandardFormat21, "4:47:28"},
	{42996.6996269676, spreadsheet.StandardFormat22, "9/18/17 4:47"},

	{1234, spreadsheet.StandardFormat37, "1,234 "},
	{-1234, spreadsheet.StandardFormat37, "(1,234)"},
	{1234, spreadsheet.StandardFormat38, "1,234 "},
	{-1234, spreadsheet.StandardFormat38, "(1,234)"},
	{1234, spreadsheet.StandardFormat39, "1,234.00"},
	{-1234, spreadsheet.StandardFormat39, "(1,234.00)"},
	{1234, spreadsheet.StandardFormat40, "1,234.00"},
	{-1234, spreadsheet.StandardFormat40, "(1,234.00)"},

	{1.23, spreadsheet.StandardFormat45, "31:12"},
	{1.2, spreadsheet.StandardFormat46, "28:48:00"},
	{1.234, spreadsheet.StandardFormat47, "36:57.6"},
	{1234, spreadsheet.StandardFormat48, "1.2E+3"},
}

func TestDefaultFormats(t *testing.T) {
	wb := spreadsheet.New()
	sheet := wb.AddSheet()
	row := sheet.AddRow()
	for _, tc := range standardFmtTestData {
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

func BenchmarkDefaultFormat(b *testing.B) {
	wb := spreadsheet.New()
	sheet := wb.AddSheet()
	row := sheet.AddRow()
	cell := row.AddCell()
	cell.SetNumber(1.234)

	for _, tc := range standardFmtTestData {
		cs := wb.StyleSheet.AddCellStyle()
		cs.SetNumberFormatStandard(tc.Fmt)
		cell.SetStyle(cs)
		for i := 0; i < b.N; i++ {
			cell.GetFormattedValue()
		}
	}
}
