// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package formula_test

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/unidoc/unioffice/measurement"
	"github.com/unidoc/unioffice/schema/soo/sml"
	"github.com/unidoc/unioffice/spreadsheet"
	"github.com/unidoc/unioffice/spreadsheet/formula"
)

func TestCell(t *testing.T) {
	td := []struct {
		Inp string
		Exp string
	}{
		{`=CELL("address",A1)`, `$A$1 ResultTypeString`},
		{`=CELL("col",B1)`, `2 ResultTypeNumber`},
		{`=CELL("row",A1)`, `1 ResultTypeNumber`},
		{`=CELL("color",A1)`, `1 ResultTypeNumber`},
		{`=CELL("color",A2)`, `0 ResultTypeNumber`},
		{`=CELL("contents",A1)`, `-12345.6789 ResultTypeNumber`},
		{`=CELL("contents",B1)`, `Hello World ResultTypeString`},
		{`=CELL("filename",B1)`, ` ResultTypeString`},
		{`=CELL("format",A1)`, `.5- ResultTypeString`},
		{`=CELL("format",A2)`, `F0 ResultTypeString`},
		{`=CELL("format",A3)`, `.0 ResultTypeString`},
		{`=CELL("format",A4)`, `.0() ResultTypeString`},
		{`=CELL("format",A5)`, `F2 ResultTypeString`},
		{`=CELL("format",A6)`, `.2 ResultTypeString`},
		{`=CELL("format",A7)`, `.2 ResultTypeString`},
		{`=CELL("format",A8)`, `.2() ResultTypeString`},
		{`=CELL("format",A9)`, `P2 ResultTypeString`},
		{`=CELL("format",A10)`, `C0 ResultTypeString`},
		{`=CELL("format",A11)`, `C2 ResultTypeString`},
		{`=CELL("format",A12)`, `C2 ResultTypeString`},
		{`=CELL("format",A13)`, `D1 ResultTypeString`},
		{`=CELL("format",A14)`, `D1 ResultTypeString`},
		{`=CELL("format",A15)`, `D2 ResultTypeString`},
		{`=CELL("format",A16)`, `D3 ResultTypeString`},
		{`=CELL("format",A17)`, `D4 ResultTypeString`},
		{`=CELL("format",A18)`, `D4 ResultTypeString`},
		{`=CELL("format",A19)`, `D5 ResultTypeString`},
		{`=CELL("format",A20)`, `D6 ResultTypeString`},
		{`=CELL("format",A21)`, `D7 ResultTypeString`},
		{`=CELL("format",A22)`, `D8 ResultTypeString`},
		{`=CELL("format",A23)`, `D9 ResultTypeString`},
		{`=CELL("format",A24)`, `S2 ResultTypeString`},
		{`=CELL("format",A25)`, `G ResultTypeString`},
		{`=CELL("format",C1)`, `.2() ResultTypeString`},
		{`=CELL("parentheses",A1)`, `0 ResultTypeNumber`},
		{`=CELL("parentheses",C1)`, `1 ResultTypeNumber`},
		{`=CELL("prefix",B1)`, ` ResultTypeString`},
		{`=CELL("prefix",B2)`, `' ResultTypeString`},
		{`=CELL("prefix",B3)`, `" ResultTypeString`},
		{`=CELL("prefix",B4)`, `^ ResultTypeString`},
		{`=CELL("prefix",B5)`, `\ ResultTypeString`},
		{`=CELL("protect",A1)`, `1 ResultTypeNumber`},
		{`=CELL("protect",B1)`, `0 ResultTypeNumber`},
		{`=CELL("type",A1)`, `v ResultTypeString`},
		{`=CELL("type",B1)`, `l ResultTypeString`},
		{`=CELL("type",D1)`, `b ResultTypeString`},
		{`=CELL("width",A1)`, `15 ResultTypeNumber`},
		{`=CELL("width",B1)`, `25 ResultTypeNumber`},
	}

	ss := spreadsheet.New()
	sheet := ss.AddSheet()

// cells with number, needed for testing different formats
	for i := 1; i <= 25; i++ {
		sheet.Cell("A"+strconv.Itoa(i)).SetNumber(-12345.6789)
	}

// cells with string values, needed for testing different alignments
	sheet.Cell("B1").SetString("Hello World")
	sheet.Cell("B2").SetString("Hello World Left")
	sheet.Cell("B3").SetString("Hello World Right")
	sheet.Cell("B4").SetString("Hello World Centered")
	sheet.Cell("B5").SetString("Hello World Fill")

// for testing "color" function
	redStyle := ss.StyleSheet.AddCellStyle()
	redStyle.SetNumberFormat("#,##0.00000;[RED]-#,##0.00000")
	sheet.Cell("A1").SetStyle(redStyle)

// for testing "parentheses" function
	parStyle := ss.StyleSheet.AddCellStyle()
	parStyle.SetNumberFormat("#,##0.00_);(#,##0.00)")
	sheet.Cell("C1").SetStyle(parStyle)

// for testing "format" function
	integerStyle := ss.StyleSheet.AddCellStyle()
	integerStyle.SetNumberFormat("00000")
	sheet.Cell("A2").SetStyle(integerStyle)

	intSepStyle := ss.StyleSheet.AddCellStyle()
	intSepStyle.SetNumberFormat("000,000,000")
	sheet.Cell("A3").SetStyle(intSepStyle)

	intParStyle := ss.StyleSheet.AddCellStyle()
	intParStyle.SetNumberFormat("#,##0_);(#,##0)")
	sheet.Cell("A4").SetStyle(intParStyle)

	financeStyle := ss.StyleSheet.AddCellStyle()
	financeStyle.SetNumberFormat("0.00")
	sheet.Cell("A5").SetStyle(financeStyle)

	decimalStyle := ss.StyleSheet.AddCellStyle()
	decimalStyle.SetNumberFormat("#,##0.00")
	sheet.Cell("A6").SetStyle(decimalStyle)

	decJustStyle := ss.StyleSheet.AddCellStyle()
	decJustStyle.SetNumberFormat("_-* #,##0.00_-;-* #,##0.00_-;_-* \"-\"??_-;_-@_-")
	sheet.Cell("A7").SetStyle(decJustStyle)

	decParStyle := ss.StyleSheet.AddCellStyle()
	decParStyle.SetNumberFormat("#,##0.00_);(#,##0.00)")
	sheet.Cell("A8").SetStyle(decParStyle)

	percentStyle := ss.StyleSheet.AddCellStyle()
	percentStyle.SetNumberFormat("0.00%")
	sheet.Cell("A9").SetStyle(percentStyle)

	intCurStyle := ss.StyleSheet.AddCellStyle()
	intCurStyle.SetNumberFormat("[$$-409]#,##0;-[$$-409]#,##0")
	sheet.Cell("A10").SetStyle(intCurStyle)

	curStyle := ss.StyleSheet.AddCellStyle()
	curStyle.SetNumberFormat("[$$-409]#,##0.00;-[$$-409]#,##0.00")
	sheet.Cell("A11").SetStyle(curStyle)

	curLabelStyle := ss.StyleSheet.AddCellStyle()
	curLabelStyle.SetNumberFormat("#,##0.00 [$USD];-#,##0.00 [$USD]")
	sheet.Cell("A12").SetStyle(curLabelStyle)

	mdyStyle := ss.StyleSheet.AddCellStyle()
	mdyStyle.SetNumberFormat("MM/DD/YY")
	sheet.Cell("A13").SetStyle(mdyStyle)

	dmyStyle := ss.StyleSheet.AddCellStyle()
	dmyStyle.SetNumberFormat("D. MMMM YYYY")
	sheet.Cell("A14").SetStyle(dmyStyle)

	d2Style := ss.StyleSheet.AddCellStyle()
	d2Style.SetNumberFormat("MMM DD")
	sheet.Cell("A15").SetStyle(d2Style)

	d3Style := ss.StyleSheet.AddCellStyle()
	d3Style.SetNumberFormat("MM/YY")
	sheet.Cell("A16").SetStyle(d3Style)

	d4Style := ss.StyleSheet.AddCellStyle()
	d4Style.SetNumberFormat("MM/DD/YY\\ HH:MM\\ AM/PM")
	sheet.Cell("A17").SetStyle(d4Style)

	d4Style = ss.StyleSheet.AddCellStyle()
	d4Style.SetNumberFormat("MM/DD/YYYY\\ HH:MM:SS")
	sheet.Cell("A18").SetStyle(d4Style)

	d5Style := ss.StyleSheet.AddCellStyle()
	d5Style.SetNumberFormat("MM\\-DD")
	sheet.Cell("A19").SetStyle(d5Style)

	d6Style := ss.StyleSheet.AddCellStyle()
	d6Style.SetNumberFormat("HH:MM:SS\\ AM/PM")
	sheet.Cell("A20").SetStyle(d6Style)

	d7Style := ss.StyleSheet.AddCellStyle()
	d7Style.SetNumberFormat("HH:MM\\ AM/PM")
	sheet.Cell("A21").SetStyle(d7Style)

	d8Style := ss.StyleSheet.AddCellStyle()
	d8Style.SetNumberFormat("HH:MM:SS")
	sheet.Cell("A22").SetStyle(d8Style)

	d9Style := ss.StyleSheet.AddCellStyle()
	d9Style.SetNumberFormat("HH:MM")
	sheet.Cell("A23").SetStyle(d9Style)

	sciStyle := ss.StyleSheet.AddCellStyle()
	sciStyle.SetNumberFormat("##0.00E+00")
	sheet.Cell("A24").SetStyle(sciStyle)

	incorrectStyle := ss.StyleSheet.AddCellStyle()
	incorrectStyle.SetNumberFormat("incorrect style")
	sheet.Cell("A25").SetStyle(incorrectStyle)

// for testing alignments ("prefix" function)
	leftStyle := ss.StyleSheet.AddCellStyle()
	leftStyle.SetHorizontalAlignment(sml.ST_HorizontalAlignmentLeft)
	sheet.Cell("B2").SetStyle(leftStyle)

	rightStyle := ss.StyleSheet.AddCellStyle()
	rightStyle.SetHorizontalAlignment(sml.ST_HorizontalAlignmentRight)
	sheet.Cell("B3").SetStyle(rightStyle)

	centerStyle := ss.StyleSheet.AddCellStyle()
	centerStyle.SetHorizontalAlignment(sml.ST_HorizontalAlignmentCenter)
	sheet.Cell("B4").SetStyle(centerStyle)

	fillStyle := ss.StyleSheet.AddCellStyle()
	fillStyle.SetHorizontalAlignment(sml.ST_HorizontalAlignmentFill)
	sheet.Cell("B5").SetStyle(fillStyle)

	ctx := sheet.FormulaContext()

// for testing protected cells
	ctx.SetLocked("A1", true)
	ctx.SetLocked("B1", false)

// for testing widths
	sheet.Column(1).SetWidth(1.5 * measurement.Inch)
	sheet.Column(2).SetWidth(2.5 * measurement.Inch)

	ev := formula.NewEvaluator()
	for _, tc := range td {
		t.Run(tc.Inp, func(t *testing.T) {
			p := formula.Parse(strings.NewReader(tc.Inp))
			if p == nil {
				t.Errorf("error parsing %s", tc.Inp)
				return
			}
			result := p.Eval(ctx, ev)
			got := fmt.Sprintf("%s %s", result.Value(), result.Type)
			if got != tc.Exp {
				t.Errorf("expected %s = %s, got %s", tc.Inp, tc.Exp, got)
			}
		})
	}
}
