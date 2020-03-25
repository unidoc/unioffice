// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased on https://unidoc.io.

package formula_test

import (
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/unidoc/unioffice/measurement"
	"github.com/unidoc/unioffice/schema/soo/sml"
	"github.com/unidoc/unioffice/spreadsheet"
	"github.com/unidoc/unioffice/spreadsheet/formula"
)

// Input is an input formula string.
// Expected is the expected output of the formula as a string of format: "value type". It depends on Input and workbook that is being worked with.
type testStruct struct {
	Input    string
	Expected string
}

func runTests(t *testing.T, ctx formula.Context, td []testStruct) {
	for _, tc := range td {
		t.Run(tc.Input, func(t *testing.T) {
			ev := formula.NewEvaluator()
			result := ev.Eval(ctx, tc.Input)
			got := fmt.Sprintf("%s %s", result.Value(), result.Type)
			if got != tc.Expected {
				t.Errorf("expected %s = %s, got %s", tc.Input, tc.Expected, got)
			}
		})
	}
}

func TestCell(t *testing.T) {
	td := []testStruct{
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
		sheet.Cell("A" + strconv.Itoa(i)).SetNumber(-12345.6789)
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

	runTests(t, ctx, td)
}

func TestChoose(t *testing.T) {
	td := []testStruct{
		{`=CHOOSE(A1,B1,B2,B3)`, `value1 ResultTypeString`},
		{`=CHOOSE(A2,B1,B2,B3)`, `value2 ResultTypeString`},
		{`=CHOOSE(A3,B1,B2,B3)`, `value3 ResultTypeString`},
		{`=CHOOSE(A3,B1,B2)`, `#VALUE! ResultTypeError`},
	}

	ss := spreadsheet.New()
	sheet := ss.AddSheet()

	sheet.Cell("A1").SetNumber(1)
	sheet.Cell("A2").SetNumber(2)
	sheet.Cell("A3").SetNumber(3)

	sheet.Cell("B1").SetString("value1")
	sheet.Cell("B2").SetString("value2")
	sheet.Cell("B3").SetString("value3")

	ctx := sheet.FormulaContext()

	runTests(t, ctx, td)
}

func TestColumn(t *testing.T) {
	td := []testStruct{
		{`=COLUMN(A1)`, `1 ResultTypeNumber`},
		{`=COLUMN(A2)`, `1 ResultTypeNumber`},
		{`=COLUMN(B1)`, `2 ResultTypeNumber`},
	}

	ss := spreadsheet.New()
	sheet := ss.AddSheet()

	ctx := sheet.FormulaContext()

	runTests(t, ctx, td)
}

func TestColumns(t *testing.T) {
	td := []testStruct{
		{`=COLUMNS(A1:E8)`, `5 ResultTypeNumber`},
		{`=COLUMNS(E8:A1)`, `#VALUE! ResultTypeError`},
	}

	ss := spreadsheet.New()
	sheet := ss.AddSheet()

	ctx := sheet.FormulaContext()

	runTests(t, ctx, td)
}

func TestCountIf(t *testing.T) {
	td := []testStruct{
		{`=COUNTIF(B1:B10,A1)`, `2 ResultTypeNumber`},
		{`=COUNTIF(B1:B10,A2)`, `3 ResultTypeNumber`},
	}

	ss := spreadsheet.New()
	sheet := ss.AddSheet()

	sheet.Cell("A1").SetNumber(1234.4321)
	sheet.Cell("A2").SetString("value1")

	sheet.Cell("B1").SetString("value1")
	sheet.Cell("B2").SetString("value2")
	sheet.Cell("B3").SetString("value3")
	sheet.Cell("B4").SetNumber(1234.4321)
	sheet.Cell("B5").SetString("value1")
	sheet.Cell("B6").SetString("value2")
	sheet.Cell("B7").SetString("value3")
	sheet.Cell("B8").SetString("value1")
	sheet.Cell("B9").SetNumber(1234.4322)
	sheet.Cell("B10").SetNumber(1234.4321)

	ctx := sheet.FormulaContext()

	runTests(t, ctx, td)
}

func TestCountIfs(t *testing.T) {
	td := []testStruct{
		{`=COUNTIFS(A1:E1,">1")`, `1 ResultTypeNumber`},
		{`=COUNTIFS(A2:E2,">1")`, `0 ResultTypeNumber`},
		{`=COUNTIFS(A3:C4,">1")`, `2 ResultTypeNumber`},
		{`=COUNTIFS(A5:C6,"a")`, `2 ResultTypeNumber`},
		{`=COUNTIFS(A7:B7,"1",A8:B8,"2")`, `0 ResultTypeNumber`},
		{`=COUNTIFS(A9:A10,"1",B9:B10,"2")`, `1 ResultTypeNumber`},
	}

	ss := spreadsheet.New()
	sheet := ss.AddSheet()

	sheet.Cell("A1").SetNumber(1)
	sheet.Cell("C1").SetNumber(3)
	sheet.Cell("D1").SetString("a")
	sheet.Cell("E1").SetString("")

	sheet.Cell("A2").SetNumber(1)
	sheet.Cell("C2").SetNumber(0)
	sheet.Cell("D2").SetString("a")
	sheet.Cell("E2").SetString("")

	sheet.Cell("A3").SetNumber(1)
	sheet.Cell("C3").SetNumber(3)
	sheet.Cell("A4").SetString("a")
	sheet.Cell("B4").SetNumber(4)
	sheet.Cell("C4").SetString("c")

	sheet.Cell("A5").SetNumber(1)
	sheet.Cell("C5").SetString("a")
	sheet.Cell("A6").SetString("a")
	sheet.Cell("B6").SetNumber(4)
	sheet.Cell("C6").SetString("c")

	sheet.Cell("A7").SetNumber(1)
	sheet.Cell("B8").SetNumber(2)

	sheet.Cell("A9").SetNumber(1)
	sheet.Cell("B9").SetNumber(2)
	sheet.Cell("B10").SetNumber(1)

	ctx := sheet.FormulaContext()

	runTests(t, ctx, td)
}

func TestSumIf(t *testing.T) {
	td := []testStruct{
		{`=SUMIF(A1:E1,">2",A3:E3)`, `11100 ResultTypeNumber`},
		{`=SUMIF(A2:E2,"*ound",A3:E3)`, `10100 ResultTypeNumber`},
	}

	ss := spreadsheet.New()
	sheet := ss.AddSheet()

	sheet.Cell("A1").SetNumber(1)
	sheet.Cell("B1").SetNumber(2)
	sheet.Cell("C1").SetNumber(3)
	sheet.Cell("D1").SetNumber(4)
	sheet.Cell("E1").SetNumber(5)

	sheet.Cell("A2").SetString("What")
	sheet.Cell("B2").SetString("goes")
	sheet.Cell("C2").SetString("around")
	sheet.Cell("D2").SetString("comes")
	sheet.Cell("E2").SetString("around")

	sheet.Cell("A3").SetNumber(1)
	sheet.Cell("B3").SetNumber(10)
	sheet.Cell("C3").SetNumber(100)
	sheet.Cell("D3").SetNumber(1000)
	sheet.Cell("E3").SetNumber(10000)

	ctx := sheet.FormulaContext()

	runTests(t, ctx, td)
}

func TestSumIfs(t *testing.T) {
	td := []testStruct{
		{`=SUMIFS(A3:E3,A1:E1,">2",A2:E2,"*ound")`, `100 ResultTypeNumber`},
		{`=SUMIFS(A3:E3,A1:E1,">3",A2:E2,"*ound")`, `0 ResultTypeNumber`},
	}

	ss := spreadsheet.New()
	sheet := ss.AddSheet()

	sheet.Cell("A1").SetNumber(5)
	sheet.Cell("B1").SetNumber(4)
	sheet.Cell("C1").SetNumber(3)
	sheet.Cell("D1").SetNumber(2)
	sheet.Cell("E1").SetNumber(1)

	sheet.Cell("A2").SetString("What")
	sheet.Cell("B2").SetString("goes")
	sheet.Cell("C2").SetString("around")
	sheet.Cell("D2").SetString("comes")
	sheet.Cell("E2").SetString("around")

	sheet.Cell("A3").SetNumber(1)
	sheet.Cell("B3").SetNumber(10)
	sheet.Cell("C3").SetNumber(100)
	sheet.Cell("D3").SetNumber(1000)
	sheet.Cell("E3").SetNumber(10000)

	ctx := sheet.FormulaContext()

	runTests(t, ctx, td)
}

func TestMinIfs(t *testing.T) {
	td := []testStruct{
		{`=MINIFS(C1:C5,A1:A5,">2")`, `-1000 ResultTypeNumber`},
		{`=MINIFS(C1:C5,B1:B5,"????")`, `-1000 ResultTypeNumber`},
		{`=MINIFS(C1:C5,B1:B5,"*ound")`, `10 ResultTypeNumber`},
		{`=MINIFS(C1:C5,A1:A5,">3",B1:B5,"????")`, `-1000 ResultTypeNumber`},
		{`=MINIFS(C1:C5,A1:A5,">3",B1:B5,"*ound")`, `0 ResultTypeNumber`},
	}

	ss := spreadsheet.New()
	sheet := ss.AddSheet()

	sheet.Cell("A1").SetNumber(5)
	sheet.Cell("A2").SetNumber(4)
	sheet.Cell("A3").SetNumber(3)
	sheet.Cell("A4").SetNumber(2)
	sheet.Cell("A5").SetNumber(1)

	sheet.Cell("B1").SetString("What")
	sheet.Cell("B2").SetString("goes")
	sheet.Cell("B3").SetString("around")
	sheet.Cell("B4").SetString("comes")
	sheet.Cell("B5").SetString("around")

	sheet.Cell("C1").SetNumber(-1000)
	sheet.Cell("C2").SetNumber(-100)
	sheet.Cell("C3").SetNumber(10)
	sheet.Cell("C4").SetNumber(100)
	sheet.Cell("C5").SetNumber(1000)

	ctx := sheet.FormulaContext()

	runTests(t, ctx, td)
}

func TestMaxIfs(t *testing.T) {
	td := []testStruct{
		{`=MAXIFS(C1:C5,A1:A5,">2")`, `10 ResultTypeNumber`},
		{`=MAXIFS(C1:C5,B1:B5,"????")`, `-100 ResultTypeNumber`},
		{`=MAXIFS(C1:C5,B1:B5,"*ound")`, `1000 ResultTypeNumber`},
		{`=MAXIFS(C1:C5,A1:A5,">2",B1:B5,"*es")`, `-100 ResultTypeNumber`},
		{`=MAXIFS(C1:C5,A1:A5,">3",B1:B5,"*ound")`, `0 ResultTypeNumber`},
	}

	ss := spreadsheet.New()
	sheet := ss.AddSheet()

	sheet.Cell("A1").SetNumber(5)
	sheet.Cell("A2").SetNumber(4)
	sheet.Cell("A3").SetNumber(3)
	sheet.Cell("A4").SetNumber(2)
	sheet.Cell("A5").SetNumber(1)

	sheet.Cell("B1").SetString("What")
	sheet.Cell("B2").SetString("goes")
	sheet.Cell("B3").SetString("around")
	sheet.Cell("B4").SetString("comes")
	sheet.Cell("B5").SetString("around")

	sheet.Cell("C1").SetNumber(-1000)
	sheet.Cell("C2").SetNumber(-100)
	sheet.Cell("C3").SetNumber(10)
	sheet.Cell("C4").SetNumber(100)
	sheet.Cell("C5").SetNumber(1000)

	ctx := sheet.FormulaContext()

	runTests(t, ctx, td)
}

func TestValue(t *testing.T) {
	td := []testStruct{
		{`=VALUE(A1)`, `5000 ResultTypeNumber`},
		{`=VALUE(A2)`, `4000 ResultTypeNumber`},
		{`=VALUE(A1)-VALUE(A2)`, `1000 ResultTypeNumber`},
		{`=VALUE(A3)-VALUE(A4)`, `0.25 ResultTypeNumber`},
		{`=VALUE(A5)`, `#VALUE! ResultTypeError`},
	}

	ss := spreadsheet.New()
	sheet := ss.AddSheet()

	sheet.Cell("A1").SetNumber(5000)
	sheet.Cell("A2").SetString("4e+03")
	sheet.Cell("A3").SetTime(time.Date(2019, time.November, 4, 16, 0, 0, 0, time.UTC))
	sheet.Cell("A4").SetTime(time.Date(2019, time.November, 4, 10, 0, 0, 0, time.UTC))
	sheet.Cell("A5").SetString("abcde")

	ctx := sheet.FormulaContext()

	runTests(t, ctx, td)
}

func TestMatch(t *testing.T) {
	td := []testStruct{
		{`MATCH("??ny",A1:A5)`, `2 ResultTypeNumber`},
		{`MATCH("*nny",A1:A5)`, `4 ResultTypeNumber`},
		{`MATCH("*nny",A1:A5,)`, `4 ResultTypeNumber`},
		{`=MATCH(5,B1:B5,1)`, `2 ResultTypeNumber`},
		{`=MATCH(5,C1:C5,-1)`, `3 ResultTypeNumber`},
	}

	ss := spreadsheet.New()
	sheet := ss.AddSheet()

	sheet.Cell("A1").SetString("John")
	sheet.Cell("A2").SetString("Tony")
	sheet.Cell("A3").SetString("Tony")
	sheet.Cell("A4").SetString("Benny")
	sheet.Cell("A5").SetString("Willy")

	sheet.Cell("B1").SetNumber(2)
	sheet.Cell("B2").SetNumber(4)
	sheet.Cell("B3").SetNumber(6)
	sheet.Cell("B4").SetNumber(8)
	sheet.Cell("B5").SetNumber(10)

	sheet.Cell("C1").SetNumber(10)
	sheet.Cell("C2").SetNumber(8)
	sheet.Cell("C3").SetNumber(6)
	sheet.Cell("C4").SetNumber(4)
	sheet.Cell("C5").SetNumber(2)

	ctx := sheet.FormulaContext()

	runTests(t, ctx, td)
}

func TestMax(t *testing.T) {
	ss := spreadsheet.New()
	sheet := ss.AddSheet()

	sheet.Cell("A1").SetNumber(0.1)
	sheet.Cell("B1").SetNumber(0.2)

	sheet.Cell("A2").SetNumber(0.4)
	sheet.Cell("B2").SetNumber(0.8)

	sheet.Cell("A3").SetBool(true)
	sheet.Cell("B3").SetBool(false)

	ctx := sheet.FormulaContext()

	td := []testStruct{
		{`MAX(A1:B3)`, `0.8 ResultTypeNumber`},
	}

	runTests(t, ctx, td)
}

func TestMaxA(t *testing.T) {
	ss := spreadsheet.New()
	sheet := ss.AddSheet()

	sheet.Cell("A1").SetNumber(0.1)
	sheet.Cell("B1").SetNumber(0.2)

	sheet.Cell("A2").SetNumber(0.4)
	sheet.Cell("B2").SetNumber(0.8)

	sheet.Cell("A3").SetBool(true)
	sheet.Cell("B3").SetBool(false)

	ctx := sheet.FormulaContext()

	td := []testStruct{
		{`MAXA(A1:B3)`, `1 ResultTypeNumber`},
	}

	runTests(t, ctx, td)
}

func TestMin(t *testing.T) {
	ss := spreadsheet.New()
	sheet := ss.AddSheet()

	sheet.Cell("A1").SetNumber(0.1)
	sheet.Cell("B1").SetNumber(0.2)

	sheet.Cell("A2").SetNumber(0.4)
	sheet.Cell("B2").SetNumber(0.8)

	sheet.Cell("A3").SetBool(true)
	sheet.Cell("B3").SetBool(false)

	ctx := sheet.FormulaContext()

	td := []testStruct{
		{`MIN(A1:B3)`, `0.1 ResultTypeNumber`},
	}

	runTests(t, ctx, td)
}

func TestMinA(t *testing.T) {
	ss := spreadsheet.New()
	sheet := ss.AddSheet()

	sheet.Cell("A1").SetNumber(0.1)
	sheet.Cell("B1").SetNumber(0.2)

	sheet.Cell("A2").SetNumber(0.4)
	sheet.Cell("B2").SetNumber(0.8)

	sheet.Cell("A3").SetBool(true)
	sheet.Cell("B3").SetBool(false)

	ctx := sheet.FormulaContext()

	td := []testStruct{
		{`MINA(A1:B3)`, `0 ResultTypeNumber`},
	}

	runTests(t, ctx, td)
}

func TestIfs(t *testing.T) {
	ss := spreadsheet.New()
	sheet := ss.AddSheet()

	sheet.Cell("B1").SetNumber(1)
	sheet.Cell("B2").SetString("a")
	sheet.Cell("B3").SetNumber(2)
	sheet.Cell("B4").SetString("b")
	sheet.Cell("B5").SetNumber(3)
	sheet.Cell("B6").SetString("c")
	sheet.Cell("B7").SetNumber(4)
	sheet.Cell("B8").SetString("d")
	sheet.Cell("B9").SetNumber(5)
	sheet.Cell("B10").SetString("e")

	ctx := sheet.FormulaContext()

	td := []testStruct{
		{`=IFS(B1>3,"B1",B2="a","B2",B3>2,"B3",B4="c","B4",B5>4,"B5",B6="d","B6",B7=4,"B7",B8="d","B8",B9<=4,"B9",B10="e","B10")`, `B2 ResultTypeString`},
		{`=IFS(B1>3,"B1",B2="b","B2",B3>2,"B3",B4="c","B4",B5>4,"B5",B6="d","B6",B7=4,"B7",B8="d","B8",B9<=4,"B9",B10="e","B10")`, `B7 ResultTypeString`},
		{`=IFS(B1>3,"B1",B2="b","B2",B3>2,"B3",B4="c","B4",B5>4,"B5",B6="d","B6",B7=5,"B7",B8="e","B8",B9<=4,"B9",B10="f","B10")`, `#N/A ResultTypeError`},
	}

	runTests(t, ctx, td)
}

func TestOffset(t *testing.T) {
	ss := spreadsheet.New()
	sheet := ss.AddSheet()

	sheet.Cell("B1").SetNumber(1)
	sheet.Cell("C1").SetNumber(2)
	sheet.Cell("D1").SetNumber(3)
	sheet.Cell("E1").SetNumber(4)
	sheet.Cell("B2").SetNumber(5)
	sheet.Cell("C2").SetNumber(6)
	sheet.Cell("D2").SetNumber(7)
	sheet.Cell("E2").SetNumber(8)
	sheet.Cell("B3").SetNumber(9)
	sheet.Cell("C3").SetNumber(10)
	sheet.Cell("D3").SetNumber(11)
	sheet.Cell("E3").SetNumber(12)

	ctx := sheet.FormulaContext()

	td := []testStruct{
		{`=OFFSET(B1,1,2)`, `7 ResultTypeNumber`},
		{`=SUM(OFFSET(E3,-1,-1,2,2))`, `38 ResultTypeNumber`},
		{`=AVERAGE(OFFSET(B1,1,2,-2,-2))`, `4.5 ResultTypeNumber`},
		{`=SUM(OFFSET(B1,1,2,0,0))`, `#REF! ResultTypeError`},
	}

	runTests(t, ctx, td)
}

func TestIsBlank(t *testing.T) {
	ss := spreadsheet.New()
	sheet := ss.AddSheet()

	sheet.Cell("B1").SetString(" ")

	ctx := sheet.FormulaContext()

	td := []testStruct{
		{`=ISBLANK(B1)`, `0 ResultTypeNumber`},
		{`=ISBLANK(C1)`, `1 ResultTypeNumber`},
	}

	runTests(t, ctx, td)
}

func TestIsErr(t *testing.T) {
	ss := spreadsheet.New()
	sheet := ss.AddSheet()

	sheet.Cell("A1").SetNumber(-1)
	sheet.Cell("A2").SetNumber(2)
	sheet.Cell("A3").SetNumber(0)
	sheet.Cell("A4").SetString("abcde")

	sheet.Cell("B1").SetFormulaRaw("1/A1")
	sheet.Cell("B2").SetFormulaRaw("1/A2")
	sheet.Cell("B3").SetFormulaRaw("1/A3")
	sheet.Cell("B4").SetFormulaRaw("1/A4")
	sheet.Cell("B5").SetFormulaRaw("MATCH(1,C1:C5)")

	td := []testStruct{
		{`=ISERR(B1)`, `0 ResultTypeNumber`},
		{`=ISERR(B2)`, `0 ResultTypeNumber`},
		{`=ISERR(B3)`, `1 ResultTypeNumber`},
		{`=ISERR(B4)`, `1 ResultTypeNumber`},
		{`=ISERR(B5)`, `0 ResultTypeNumber`},
	}

	ctx := sheet.FormulaContext()
	runTests(t, ctx, td)
}

func TestIsError(t *testing.T) {
	ss := spreadsheet.New()
	sheet := ss.AddSheet()

	sheet.Cell("A1").SetNumber(-1)
	sheet.Cell("A2").SetNumber(2)
	sheet.Cell("A3").SetNumber(0)
	sheet.Cell("A4").SetString("abcde")

	sheet.Cell("B1").SetFormulaRaw("1/A1")
	sheet.Cell("B2").SetFormulaRaw("1/A2")
	sheet.Cell("B3").SetFormulaRaw("1/A3")
	sheet.Cell("B4").SetFormulaRaw("1/A4")
	sheet.Cell("B5").SetFormulaRaw("MATCH(1,C1:C5)")

	td := []testStruct{
		{`=ISERROR(B1)`, `0 ResultTypeNumber`},
		{`=ISERROR(B2)`, `0 ResultTypeNumber`},
		{`=ISERROR(B3)`, `1 ResultTypeNumber`},
		{`=ISERROR(B4)`, `1 ResultTypeNumber`},
		{`=ISERROR(B5)`, `1 ResultTypeNumber`},
	}

	ctx := sheet.FormulaContext()
	runTests(t, ctx, td)
}

func TestIsEven(t *testing.T) {
	ss := spreadsheet.New()
	sheet := ss.AddSheet()

	sheet.Cell("B1").SetNumber(0)
	sheet.Cell("B2").SetNumber(2)
	sheet.Cell("B3").SetNumber(123.456)
	sheet.Cell("B4").SetNumber(2.789)

	td := []testStruct{
		{`=ISEVEN(B1)`, `1 ResultTypeNumber`},
		{`=ISEVEN(B2)`, `1 ResultTypeNumber`},
		{`=ISEVEN(B3)`, `0 ResultTypeNumber`},
		{`=ISEVEN(B4)`, `1 ResultTypeNumber`},
	}

	ctx := sheet.FormulaContext()
	runTests(t, ctx, td)
}

func TestIsFormula(t *testing.T) {
	ss := spreadsheet.New()
	sheet := ss.AddSheet()

	sheet.Cell("A1").SetNumber(0)
	sheet.Cell("A2").SetString("2")
	sheet.Cell("A3").SetFormulaRaw("=A2=A1")

	td := []testStruct{
		{`=_xlfn.ISFORMULA(A1)`, `0 ResultTypeNumber`},
		{`=_xlfn.ISFORMULA(A2)`, `0 ResultTypeNumber`},
		{`=_xlfn.ISFORMULA(A3)`, `1 ResultTypeNumber`},
	}

	ctx := sheet.FormulaContext()
	runTests(t, ctx, td)
}

func TestIsLeapYear(t *testing.T) {
	ss := spreadsheet.New()
	sheet := ss.AddSheet()

	sheet.Cell("A1").SetDate(time.Date(2016, 1, 1, 0, 0, 0, 0, time.UTC))
	sheet.Cell("A2").SetDate(time.Date(2017, 1, 1, 0, 0, 0, 0, time.UTC))
	sheet.Cell("A3").SetDate(time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC))
	sheet.Cell("A4").SetDate(time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC))

	td := []testStruct{
		{`=ORG.OPENOFFICE.ISLEAPYEAR(A1)`, `1 ResultTypeNumber`},
		{`=ORG.OPENOFFICE.ISLEAPYEAR(A2)`, `0 ResultTypeNumber`},
		{`=ORG.OPENOFFICE.ISLEAPYEAR(A3)`, `0 ResultTypeNumber`},
		{`=ORG.OPENOFFICE.ISLEAPYEAR(A4)`, `0 ResultTypeNumber`},
	}

	ctx := sheet.FormulaContext()
	runTests(t, ctx, td)
}

func TestIsNonText(t *testing.T) {
	ss := spreadsheet.New()
	sheet := ss.AddSheet()

	sheet.Cell("B2").SetString(" ")
	sheet.Cell("B3").SetNumber(123.456)
	sheet.Cell("B4").SetString("123.456")

	td := []testStruct{
		{`=ISNONTEXT(B1)`, `1 ResultTypeNumber`},
		{`=ISNONTEXT(B2)`, `0 ResultTypeNumber`},
		{`=ISNONTEXT(B3)`, `1 ResultTypeNumber`},
		{`=ISNONTEXT(B4)`, `0 ResultTypeNumber`},
	}

	ctx := sheet.FormulaContext()
	runTests(t, ctx, td)
}

func TestIsNumber(t *testing.T) {
	ss := spreadsheet.New()
	sheet := ss.AddSheet()

	sheet.Cell("B2").SetString(" ")
	sheet.Cell("B3").SetNumber(123.456)
	sheet.Cell("B4").SetString("123.456")

	td := []testStruct{
		{`=ISNUMBER(B1)`, `0 ResultTypeNumber`},
		{`=ISNUMBER(B2)`, `0 ResultTypeNumber`},
		{`=ISNUMBER(B3)`, `1 ResultTypeNumber`},
		{`=ISNUMBER(B4)`, `0 ResultTypeNumber`},
	}

	ctx := sheet.FormulaContext()
	runTests(t, ctx, td)
}

func TestIsOdd(t *testing.T) {
	ss := spreadsheet.New()
	sheet := ss.AddSheet()

	sheet.Cell("B1").SetNumber(0)
	sheet.Cell("B2").SetNumber(2)
	sheet.Cell("B3").SetNumber(123.456)
	sheet.Cell("B4").SetNumber(2.789)

	td := []testStruct{
		{`=ISODD(B1)`, `0 ResultTypeNumber`},
		{`=ISODD(B2)`, `0 ResultTypeNumber`},
		{`=ISODD(B3)`, `1 ResultTypeNumber`},
		{`=ISODD(B4)`, `0 ResultTypeNumber`},
	}

	ctx := sheet.FormulaContext()
	runTests(t, ctx, td)
}

func TestIsText(t *testing.T) {
	ss := spreadsheet.New()
	sheet := ss.AddSheet()

	sheet.Cell("B2").SetString(" ")
	sheet.Cell("B3").SetNumber(123.456)
	sheet.Cell("B4").SetString("123.456")

	td := []testStruct{
		{`=ISTEXT(B1)`, `0 ResultTypeNumber`},
		{`=ISTEXT(B2)`, `1 ResultTypeNumber`},
		{`=ISTEXT(B3)`, `0 ResultTypeNumber`},
		{`=ISTEXT(B4)`, `1 ResultTypeNumber`},
	}

	ctx := sheet.FormulaContext()
	runTests(t, ctx, td)
}

func TestIsLogical(t *testing.T) {
	ss := spreadsheet.New()
	sheet := ss.AddSheet()

	sheet.Cell("A1").SetBool(true)
	sheet.Cell("A2").SetBool(false)
	sheet.Cell("A3").SetNumber(0)
	sheet.Cell("A4").SetString("FALSE")
	sheet.Cell("A5").SetFormulaRaw("=1=2")

	td := []testStruct{
		{`=ISLOGICAL(A1)`, `1 ResultTypeNumber`},
		{`=ISLOGICAL(A2)`, `1 ResultTypeNumber`},
		{`=ISLOGICAL(A3)`, `0 ResultTypeNumber`},
		{`=ISLOGICAL(A4)`, `0 ResultTypeNumber`},
		{`=ISLOGICAL(A5)`, `1 ResultTypeNumber`},
	}

	ctx := sheet.FormulaContext()
	runTests(t, ctx, td)
}

func TestIsNA(t *testing.T) {
	ss := spreadsheet.New()
	sheet := ss.AddSheet()

	sheet.Cell("A1").SetFormulaRaw("MATCH(1,B1:B5)")
	sheet.Cell("A2").SetString("#N/A")
	sheet.Cell("A3").SetNumber(0)

	td := []testStruct{
		{`=ISNA(A1)`, `1 ResultTypeNumber`},
		{`=ISNA(A2)`, `0 ResultTypeNumber`},
		{`=ISNA(A3)`, `0 ResultTypeNumber`},
		{`=ISNA(A4)`, `0 ResultTypeNumber`},
	}

	ctx := sheet.FormulaContext()
	runTests(t, ctx, td)
}

func TestIsRef(t *testing.T) {
	ss := spreadsheet.New()
	sheet := ss.AddSheet()

	sheet.Cell("A1").SetFormulaRaw("=ISREF(B1)")
	sheet.Cell("A2").SetFormulaRaw("=ISREF(B1:Z999)")
	sheet.Cell("A3").SetFormulaRaw("=ISREF(A1048577)")
	sheet.Cell("A4").SetFormulaRaw("=ISREF(ZZA0)")
	sheet.Cell("A5").SetFormulaRaw("=ISREF(ZZ0)")
	sheet.Cell("A6").SetFormulaRaw("=ISREF(AA:ZZ)")
	sheet.Cell("A7").SetFormulaRaw("=ISREF(1:4)")
	sheet.Cell("A8").SetFormulaRaw("=4/0")
	sheet.Cell("A9").SetFormulaRaw("=ISREF(A7)")
	sheet.Cell("A10").SetString("A1")

	td := []testStruct{
		{`A1`, `1 ResultTypeNumber`},
		{`A2`, `1 ResultTypeNumber`},
		{`A3`, `0 ResultTypeNumber`},
		{`A4`, `0 ResultTypeNumber`},
		{`A5`, `1 ResultTypeNumber`},
		{`A6`, `1 ResultTypeNumber`},
		{`A7`, `1 ResultTypeNumber`},
		{`A8`, `#DIV/0! ResultTypeError`},
		{`A9`, `1 ResultTypeNumber`},
		{`A10`, `A1 ResultTypeString`},
	}

	ctx := sheet.FormulaContext()
	runTests(t, ctx, td)
}

func TestFind(t *testing.T) {
	ss := spreadsheet.New()
	sheet := ss.AddSheet()

	sheet.Cell("A1").SetString("abcde")
	sheet.Cell("B1").SetString("b")
	sheet.Cell("B2").SetString("c")

	sheet.Cell("C1").SetString("\u4f60\u597d\uff0c\u4e16\u754c")
	sheet.Cell("D1").SetString("\u4f60")
	sheet.Cell("D2").SetString("\u4e16\u754c")

	td := []testStruct{
		{`FIND("",A1)`, `1 ResultTypeNumber`},
		{`FIND("",A1,)`, `1 ResultTypeNumber`},
		{`FIND(B1,A1)`, `2 ResultTypeNumber`},
		{`FIND(B2,A1,3)`, `3 ResultTypeNumber`},
		{`FIND(B2,A1,4)`, `#VALUE! ResultTypeError`},
		{`FIND(D1,C1)`, `1 ResultTypeNumber`},
		{`FIND(D2,C1,3)`, `4 ResultTypeNumber`},
		{`FIND(D2,C1,5)`, `#VALUE! ResultTypeError`},
	}

	ctx := sheet.FormulaContext()
	runTests(t, ctx, td)
}

func TestFindb(t *testing.T) {
	ss := spreadsheet.New()
	ss.CoreProperties.SetLanguage("zh-TW") // set DBCS language
	sheet := ss.AddSheet()

	sheet.Cell("A1").SetString("abcde")
	sheet.Cell("B1").SetString("b")
	sheet.Cell("B2").SetString("c")

	sheet.Cell("C1").SetString("\u4f60\u597d\uff0c\u4e16\u754c")
	sheet.Cell("D1").SetString("\u4f60")
	sheet.Cell("D2").SetString("\u4e16\u754c")

	td := []testStruct{
		{`FINDB("",A1)`, `1 ResultTypeNumber`},
		{`FINDB(B1,A1)`, `2 ResultTypeNumber`},
		{`FINDB(B1,A1,)`, `2 ResultTypeNumber`},
		{`FINDB(B2,A1,3)`, `3 ResultTypeNumber`},
		{`FINDB(B2,A1,4)`, `#VALUE! ResultTypeError`},
		{`FINDB(D1,C1)`, `1 ResultTypeNumber`},
		{`FINDB(D2,C1,3)`, `7 ResultTypeNumber`},
		{`FINDB(D2,C1,8)`, `#VALUE! ResultTypeError`},
	}

	ctx := sheet.FormulaContext()
	runTests(t, ctx, td)
}

func TestSearch(t *testing.T) {
	ss := spreadsheet.New()
	sheet := ss.AddSheet()

	sheet.Cell("A1").SetString("ABCDE")
	sheet.Cell("B1").SetString("b")
	sheet.Cell("B2").SetString("*c")
	sheet.Cell("B3").SetString("??d")

	sheet.Cell("C1").SetString("\u4f60\u597d\uff0c\u4e16\u754c")
	sheet.Cell("D1").SetString("*\u4f60")
	sheet.Cell("D2").SetString("??\u4e16\u754c")

	td := []testStruct{
		{`SEARCH("",A1)`, `1 ResultTypeNumber`},
		{`SEARCH(B1,A1)`, `2 ResultTypeNumber`},
		{`SEARCH(B1,A1,)`, `2 ResultTypeNumber`},
		{`SEARCH(B2,A1,3)`, `3 ResultTypeNumber`},
		{`SEARCH(B2,A1,4)`, `#VALUE! ResultTypeError`},
		{`SEARCH(B3,A1,2)`, `2 ResultTypeNumber`},
		{`SEARCH(D1,C1)`, `1 ResultTypeNumber`},
		{`SEARCH(D2,C1,2)`, `2 ResultTypeNumber`},
		{`SEARCH(D2,C1,5)`, `#VALUE! ResultTypeError`},
	}

	ctx := sheet.FormulaContext()
	runTests(t, ctx, td)
}

func TestSearchb(t *testing.T) {
	ss := spreadsheet.New()
	ss.CoreProperties.SetLanguage("zh-TW") // set DBCS language
	sheet := ss.AddSheet()

	sheet.Cell("A1").SetString("ABCDE")
	sheet.Cell("B1").SetString("b")
	sheet.Cell("B2").SetString("*c")
	sheet.Cell("B3").SetString("??d")

	sheet.Cell("C1").SetString("\u4f60\u597d\uff0c\u4e16\u754c")
	sheet.Cell("D1").SetString("*\u4f60")
	sheet.Cell("D2").SetString("??\u4e16\u754c")

	td := []testStruct{
		{`SEARCHB("",A1)`, `1 ResultTypeNumber`},
		{`SEARCHB(B1,A1)`, `2 ResultTypeNumber`},
		{`SEARCHB(B1,A1,)`, `2 ResultTypeNumber`},
		{`SEARCHB(B2,A1,3)`, `3 ResultTypeNumber`},
		{`SEARCHB(B2,A1,4)`, `#VALUE! ResultTypeError`},
		{`SEARCHB(D1,C1)`, `1 ResultTypeNumber`},
		{`SEARCHB(D2,C1,3)`, `3 ResultTypeNumber`},
		{`SEARCHB(D2,C1,8)`, `#VALUE! ResultTypeError`},
	}

	ctx := sheet.FormulaContext()
	runTests(t, ctx, td)
}

func TestConcat(t *testing.T) {
	ss := spreadsheet.New()
	sheet := ss.AddSheet()

	sheet.Cell("A1").SetBool(true)
	sheet.Cell("A2").SetBool(false)

	td := []testStruct{
		{`CONCAT("Hello"," ","world")`, `Hello world ResultTypeString`},
		{`CONCAT("Hello"," my ","world")`, `Hello my world ResultTypeString`},
		{`CONCAT("1","one")`, `1one ResultTypeString`},
		{`CONCAT(A1,"yes")`, `TRUEyes ResultTypeString`},
		{`CONCAT(A2,"no")`, `FALSEno ResultTypeString`},
	}

	ctx := sheet.FormulaContext()
	runTests(t, ctx, td)
}

func TestYear(t *testing.T) {
	td := []testStruct{
		{`=YEAR(A1)`, `2019 ResultTypeNumber`},
		{`=YEAR(A2)`, `#VALUE! ResultTypeError`},
	}

	ss := spreadsheet.New()
	sheet := ss.AddSheet()

	sheet.Cell("A1").SetTime(time.Date(2019, time.November, 4, 16, 0, 0, 0, time.UTC))

	ctx := sheet.FormulaContext()

	runTests(t, ctx, td)
}

func TestYearFrac(t *testing.T) {
	ss := spreadsheet.New()
	sheet := ss.AddSheet()

	sheet.Cell("A1").SetTime(time.Date(1900, time.January, 1, 0, 0, 0, 0, time.UTC))
	sheet.Cell("A2").SetTime(time.Date(1900, time.January, 2, 0, 0, 0, 0, time.UTC))
	sheet.Cell("A3").SetTime(time.Date(1900, time.January, 31, 0, 0, 0, 0, time.UTC))
	sheet.Cell("A4").SetTime(time.Date(1900, time.March, 31, 0, 0, 0, 0, time.UTC))
	sheet.Cell("A5").SetTime(time.Date(1900, time.February, 1, 0, 0, 0, 0, time.UTC))
	sheet.Cell("A6").SetTime(time.Date(1904, time.January, 1, 0, 0, 0, 0, time.UTC))
	sheet.Cell("A7").SetTime(time.Date(1904, time.January, 2, 0, 0, 0, 0, time.UTC))
	sheet.Cell("A8").SetTime(time.Date(1905, time.January, 1, 0, 0, 0, 0, time.UTC))
	sheet.Cell("A9").SetString("Hello")
	sheet.Cell("A10").SetString("World")

	td := []testStruct{
		{`=YEARFRAC(A1,A2)`, `0.00277777777 ResultTypeNumber`},
		{`=YEARFRAC(A3,A4)`, `0.16666666666 ResultTypeNumber`},
		{`=YEARFRAC(A3,A5)`, `0.00277777777 ResultTypeNumber`},
		{`=YEARFRAC(A3,A5,)`, `0.00277777777 ResultTypeNumber`},
		{`=YEARFRAC(A1,A2,1)`, `0.00273972602 ResultTypeNumber`},
		{`=YEARFRAC(A6,A7,1)`, `0.00273224043 ResultTypeNumber`},
		{`=YEARFRAC(A6,A8,1)`, `1 ResultTypeNumber`},
		{`=YEARFRAC(A1,A2,2)`, `0.00277777777 ResultTypeNumber`},
		{`=YEARFRAC(A1,A2,3)`, `0.00273972602 ResultTypeNumber`},
		{`=YEARFRAC(A1,A2,4)`, `0.00277777777 ResultTypeNumber`},
		{`=YEARFRAC(A9,A2)`, `#VALUE! ResultTypeError`},
		{`=YEARFRAC(A1,A10)`, `#VALUE! ResultTypeError`},
	}

	ctx := sheet.FormulaContext()
	runTests(t, ctx, td)
}

func TestTime(t *testing.T) {
	ss := spreadsheet.New()
	sheet := ss.AddSheet()

	td := []testStruct{
		{`=TIME(6,0,0)`, `0.25 ResultTypeNumber`},
		{`=TIME(12,0,0)`, `0.5 ResultTypeNumber`},
		{`=TIME(2,24,0)`, `0.1 ResultTypeNumber`},
		{`=TIME(7,-60,0)`, `0.25 ResultTypeNumber`},
		{`=TIME(1,-120,0)`, `#NUM! ResultTypeError`},
	}

	ctx := sheet.FormulaContext()
	runTests(t, ctx, td)
}

func TestTimeValue(t *testing.T) {
	ss := spreadsheet.New()
	sheet := ss.AddSheet()

	td := []testStruct{
		{`=TIMEVALUE("1/1/1900 00:00:00")`, `0 ResultTypeNumber`},
		{`=TIMEVALUE("1/1/1900 12:00:00")`, `0.5 ResultTypeNumber`},
		{`=TIMEVALUE("1/1/1900")`, `0 ResultTypeNumber`},
		{`=TIMEVALUE("02:12:18 PM")`, `0.591875 ResultTypeNumber`},
		{`=TIMEVALUE("a")`, `#VALUE! ResultTypeError`},
	}

	ctx := sheet.FormulaContext()
	runTests(t, ctx, td)
}

func TestDay(t *testing.T) {
	ss := spreadsheet.New()
	sheet := ss.AddSheet()

	td := []testStruct{
		{`=DAY("02-29-2019")`, `#VALUE! ResultTypeError`},
		{`=DAY("02-29-2020")`, `29 ResultTypeNumber`},
		{`=DAY("01/03/2019 12:14:16")`, `3 ResultTypeNumber`},
		{`=DAY("January 25, 2020 01:03 AM")`, `25 ResultTypeNumber`},
	}

	ctx := sheet.FormulaContext()
	runTests(t, ctx, td)
}

func TestDays(t *testing.T) {
	ss := spreadsheet.New()
	sheet := ss.AddSheet()

	sheet.Cell("A1").SetTime(time.Date(2021, time.February, 28, 0, 0, 0, 0, time.UTC))
	sheet.Cell("A2").SetTime(time.Date(1900, time.January, 25, 0, 0, 0, 0, time.UTC))
	sheet.Cell("A3").SetString("02/28/2021")
	sheet.Cell("A4").SetString("01/25/1900")

	td := []testStruct{
		{`=DAYS(A1,A2)`, `44230 ResultTypeNumber`},
		{`=DAYS(A3,A4)`, `44230 ResultTypeNumber`},
		{`=DAYS(A3,"02/29/1900")`, `#VALUE! ResultTypeError`},
	}

	ctx := sheet.FormulaContext()
	runTests(t, ctx, td)
}

func TestDate(t *testing.T) {
	ss := spreadsheet.New()
	sheet := ss.AddSheet()

	td := []testStruct{
		{`=DATE(1899,1,1)`, `693598 ResultTypeNumber`},
		{`=DATE(10000,1,1)`, `#NUM! ResultTypeError`},
		{`=DATE(2020,3,31)`, `43921 ResultTypeNumber`},
	}

	ctx := sheet.FormulaContext()
	runTests(t, ctx, td)
}

func TestDateValue(t *testing.T) {
	ss := spreadsheet.New()
	sheet := ss.AddSheet()

	td := []testStruct{
		{`=DATEVALUE("1/1/1900 00:00:00")`, `1 ResultTypeNumber`},
		{`=DATEVALUE("1/1/1900 7:00 AM")`, `1 ResultTypeNumber`},
		{`=DATEVALUE("1/1/1900")`, `1 ResultTypeNumber`},
		{`=DATEVALUE("11/27/2019")`, `43796 ResultTypeNumber`},
		{`=DATEVALUE("25-Jan-2019")`, `43490 ResultTypeNumber`},
		{`=DATEVALUE("02:12:18 PM")`, `0 ResultTypeNumber`},
		{`=DATEVALUE("a")`, `#VALUE! ResultTypeError`},
	}

	ctx := sheet.FormulaContext()
	runTests(t, ctx, td)
}

func TestDateDif(t *testing.T) {
	ss := spreadsheet.New()
	sheet := ss.AddSheet()

	sheet.Cell("A1").SetTime(time.Date(2019, time.February, 7, 0, 0, 0, 0, time.UTC))
	sheet.Cell("A2").SetTime(time.Date(2021, time.April, 6, 0, 0, 0, 0, time.UTC))

	td := []testStruct{
		{`=DATEDIF(A1,A2,"y")`, `2 ResultTypeNumber`},
		{`=DATEDIF(A1,A2,"m")`, `25 ResultTypeNumber`},
		{`=DATEDIF(A1,A2,"d")`, `789 ResultTypeNumber`},
		{`=DATEDIF(A1,A2,"ym")`, `1 ResultTypeNumber`},
		{`=DATEDIF(A1,A2,"yd")`, `58 ResultTypeNumber`},
		{`=DATEDIF(A1,A2,"md")`, `30 ResultTypeNumber`},
		{`=DATEDIF(A2,A1,"y")`, `#NUM! ResultTypeError`},
		{`=DATEDIF(A1,A2,"yy")`, `#NUM! ResultTypeError`},
	}

	ctx := sheet.FormulaContext()
	runTests(t, ctx, td)
}

func TestMinute(t *testing.T) {
	ss := spreadsheet.New()
	sheet := ss.AddSheet()

	td := []testStruct{
		{`=MINUTE("02-29-2019 14:18:16")`, `#VALUE! ResultTypeError`},
		{`=MINUTE("02-29-2020 14:18:16")`, `18 ResultTypeNumber`},
		{`=MINUTE("01/03/2019 12:14")`, `14 ResultTypeNumber`},
		{`=MINUTE("January 25, 2020 01:03 AM")`, `3 ResultTypeNumber`},
	}

	ctx := sheet.FormulaContext()
	runTests(t, ctx, td)
}

func TestMonth(t *testing.T) {
	ss := spreadsheet.New()
	sheet := ss.AddSheet()

	td := []testStruct{
		{`=MONTH("02-29-2019")`, `#VALUE! ResultTypeError`},
		{`=MONTH("02-29-2020")`, `2 ResultTypeNumber`},
		{`=MONTH("01/03/2019 12:14:16")`, `1 ResultTypeNumber`},
		{`=MONTH("February 25, 2020 01:03 AM")`, `2 ResultTypeNumber`},
		{`=MONTH("12:14:16")`, `1 ResultTypeNumber`},
	}

	ctx := sheet.FormulaContext()
	runTests(t, ctx, td)
}

func TestEdate(t *testing.T) {
	ss := spreadsheet.New()
	sheet := ss.AddSheet()

	td := []testStruct{
		{`=EDATE("02-29-2019",-6)`, `#VALUE! ResultTypeError`},
		{`=EDATE("02-29-2020",-6)`, `43706 ResultTypeNumber`},
		{`=EDATE("06/30/1900 12:14:16",-6)`, `#NUM! ResultTypeError`},
		{`=EDATE("07/01/1900 12:14:16",-6)`, `1 ResultTypeNumber`},
		{`=EDATE("01:03 AM",-6)`, `#NUM! ResultTypeError`},
	}

	ctx := sheet.FormulaContext()
	runTests(t, ctx, td)
}

func TestEomonth(t *testing.T) {
	ss := spreadsheet.New()
	sheet := ss.AddSheet()

	td := []testStruct{
		{`=EOMONTH("02-29-2019",-6)`, `#VALUE! ResultTypeError`},
		{`=EOMONTH("02-29-2020",-6)`, `43708 ResultTypeNumber`},
		{`=EOMONTH("06/30/1900 12:14:16",-6)`, `#NUM! ResultTypeError`},
		{`=EOMONTH("07/01/1900 12:14:16",-6)`, `31 ResultTypeNumber`},
		{`=EOMONTH("01:03 AM",-6)`, `#NUM! ResultTypeError`},
	}

	ctx := sheet.FormulaContext()
	runTests(t, ctx, td)
}

func TestDuration(t *testing.T) {
	ss := spreadsheet.New()
	sheet := ss.AddSheet()

	sheet.Cell("A1").SetTime(time.Date(2018, time.July, 1, 0, 0, 0, 0, time.UTC))    // settlement date
	sheet.Cell("A2").SetTime(time.Date(2048, time.January, 1, 0, 0, 0, 0, time.UTC)) // maturity date
	sheet.Cell("A3").SetNumber(0.08)                                                 // coupon rate
	sheet.Cell("A4").SetNumber(0.09)                                                 // yield rate
	sheet.Cell("A5").SetNumber(2)                                                    // frequency of payments
	sheet.Cell("A6").SetNumber(0)                                                    // basis
	sheet.Cell("A7").SetString("07/01/2018")                                         // settlement date in string format
	sheet.Cell("A8").SetString("01/01/2048")                                         // maturity date in string format

	td := []testStruct{
		{`=DURATION(A1,A2,A3,A4,A5)`, `10.9191452815 ResultTypeNumber`},
		{`=DURATION(A1,A2,A3,A4,A5,)`, `10.9191452815 ResultTypeNumber`},
		{`=DURATION(A1,A2,A3,A4,A5,A6)`, `10.9191452815 ResultTypeNumber`},
		{`=DURATION(A7,A8,A3,A4,A5,A6)`, `10.9191452815 ResultTypeNumber`},
		{`=DURATION(A1,A2,A3,A4,A5,5)`, `#NUM! ResultTypeError`},
		{`=DURATION(A8,A7,A3,A4,A5,A6)`, `#NUM! ResultTypeError`},
	}

	ctx := sheet.FormulaContext()
	runTests(t, ctx, td)
}

func TestMduration(t *testing.T) {
	ss := spreadsheet.New()
	sheet := ss.AddSheet()

	sheet.Cell("A1").SetTime(time.Date(2008, time.January, 1, 0, 0, 0, 0, time.UTC)) // settlement date
	sheet.Cell("A2").SetTime(time.Date(2016, time.January, 1, 0, 0, 0, 0, time.UTC)) // maturity date
	sheet.Cell("A3").SetNumber(0.08)                                                 // coupon rate
	sheet.Cell("A4").SetNumber(0.09)                                                 // yield rate
	sheet.Cell("A5").SetNumber(2)                                                    // frequency of payments
	sheet.Cell("A6").SetNumber(0)                                                    // basis
	sheet.Cell("A7").SetString("01/01/2008")                                         // settlement date in string format
	sheet.Cell("A8").SetString("01/01/2016")                                         // maturity date in string format

	td := []testStruct{
		{`=MDURATION(A1,A2,A3,A4,A5)`, `5.73566981391 ResultTypeNumber`},
		{`=MDURATION(A1,A2,A3,A4,A5,)`, `5.73566981391 ResultTypeNumber`},
		{`=MDURATION(A1,A2,A3,A4,A5,A6)`, `5.73566981391 ResultTypeNumber`},
		{`=MDURATION(A7,A8,A3,A4,A5,A6)`, `5.73566981391 ResultTypeNumber`},
		{`=MDURATION(A1,A2,A3,A4,A5,5)`, `#NUM! ResultTypeError`},
		{`=MDURATION(A8,A7,A3,A4,A5,A6)`, `#NUM! ResultTypeError`},
	}

	ctx := sheet.FormulaContext()
	runTests(t, ctx, td)
}

func TestPduration(t *testing.T) {
	ss := spreadsheet.New()
	sheet := ss.AddSheet()

	td := []testStruct{
		{`=PDURATION(0.025,2000,2200)`, `3.85986616262 ResultTypeNumber`},
		{`=PDURATION(0,2000,2200)`, `#NUM! ResultTypeError`},
		{`=PDURATION(0.025,"2000",2200)`, `#VALUE! ResultTypeError`},
	}

	ctx := sheet.FormulaContext()
	runTests(t, ctx, td)
}

func TestRow(t *testing.T) {
	td := []testStruct{
		{`=ROW(A1)`, `1 ResultTypeNumber`},
		{`=ROW(A11)`, `11 ResultTypeNumber`},
		{`=ROW(B11)`, `11 ResultTypeNumber`},
	}

	ss := spreadsheet.New()
	sheet := ss.AddSheet()

	ctx := sheet.FormulaContext()

	runTests(t, ctx, td)
}

func TestRows(t *testing.T) {
	td := []testStruct{
		{`=ROWS(A1:E8)`, `8 ResultTypeNumber`},
		{`=ROWS(E8:A1)`, `#VALUE! ResultTypeError`},
	}

	ss := spreadsheet.New()
	sheet := ss.AddSheet()

	ctx := sheet.FormulaContext()

	runTests(t, ctx, td)
}

func TestLookup(t *testing.T) {
	ss := spreadsheet.New()
	sheet := ss.AddSheet()

	sheet.Cell("A1").SetNumber(1)
	sheet.Cell("A2").SetNumber(2)
	sheet.Cell("A3").SetNumber(3)

	sheet.Cell("B1").SetString("value1")
	sheet.Cell("B2").SetString("value2")
	sheet.Cell("B3").SetString("value3")

	td := []testStruct{
		{`=LOOKUP(2,A1:A3,B1:B3)`, `value2 ResultTypeString`},
		{`=LOOKUP(2,A1:B1,A2:B2)`, `#N/A ResultTypeError`},
		{`=LOOKUP(1,A1:B1,A2:B2)`, `2 ResultTypeNumber`},
	}

	ctx := sheet.FormulaContext()

	runTests(t, ctx, td)
}

func TestVlookup(t *testing.T) {
	ss := spreadsheet.New()
	sheet := ss.AddSheet()

	sheet.Cell("A1").SetNumber(100)
	sheet.Cell("A2").SetNumber(200)
	sheet.Cell("A3").SetNumber(300)
	sheet.Cell("A4").SetNumber(400)

	sheet.Cell("B1").SetString("value1")
	sheet.Cell("B2").SetString("value2")
	sheet.Cell("B3").SetString("value3")
	sheet.Cell("B4").SetString("value4")

	td := []testStruct{
		{`=VLOOKUP(150,A1:B4,2)`, `value1 ResultTypeString`},
		{`=VLOOKUP(250,A1:B4,2)`, `value2 ResultTypeString`},
		{`=VLOOKUP(250,A1:B4,2,)`, `value2 ResultTypeString`},
		{`=VLOOKUP(250,A1:B4,2,FALSE)`, `#N/A ResultTypeError`},
		{`=VLOOKUP(300,A1:B4,2,FALSE)`, `value3 ResultTypeString`},
	}

	ctx := sheet.FormulaContext()

	runTests(t, ctx, td)
}

func TestLarge(t *testing.T) {
	ss := spreadsheet.New()
	sheet := ss.AddSheet()

	sheet.Cell("A1").SetNumber(400)
	sheet.Cell("A2").SetNumber(300)
	sheet.Cell("A3").SetNumber(500)
	sheet.Cell("A4").SetNumber(100)
	sheet.Cell("A5").SetNumber(200)
	sheet.Cell("B1").SetString("abcde")
	sheet.Cell("B2").SetNumber(150)
	sheet.Cell("B3").SetNumber(350)
	sheet.Cell("B4").SetNumber(450)
	sheet.Cell("B5").SetNumber(250)

	td := []testStruct{
		{`=LARGE(A1:B5,4)`, `350 ResultTypeNumber`},
		{`=LARGE(A1:B5,0)`, `#NUM! ResultTypeError`},
		{`=LARGE(A1:B5,10)`, `#NUM! ResultTypeError`},
		{`=LARGE(A2:B2,2)`, `150 ResultTypeNumber`},
	}

	ctx := sheet.FormulaContext()

	runTests(t, ctx, td)
}

func TestSmall(t *testing.T) {
	ss := spreadsheet.New()
	sheet := ss.AddSheet()

	sheet.Cell("A1").SetNumber(400)
	sheet.Cell("A2").SetNumber(300)
	sheet.Cell("A3").SetNumber(500)
	sheet.Cell("A4").SetNumber(100)
	sheet.Cell("A5").SetNumber(200)
	sheet.Cell("B1").SetString("abcde")
	sheet.Cell("B2").SetNumber(150)
	sheet.Cell("B3").SetNumber(350)
	sheet.Cell("B4").SetNumber(450)
	sheet.Cell("B5").SetNumber(250)

	td := []testStruct{
		{`=SMALL(A1:B5,4)`, `250 ResultTypeNumber`},
		{`=SMALL(A1:B5,0)`, `#NUM! ResultTypeError`},
		{`=SMALL(A1:B5,10)`, `#NUM! ResultTypeError`},
		{`=SMALL(A2:B2,2)`, `300 ResultTypeNumber`},
	}

	ctx := sheet.FormulaContext()

	runTests(t, ctx, td)
}

func TestLower(t *testing.T) {
	ss := spreadsheet.New()
	sheet := ss.AddSheet()

	sheet.Cell("A1").SetNumber(400)
	sheet.Cell("A2").SetString("Hello")
	sheet.Cell("B1").SetString("World!")

	sheet.Cell("D1").SetFormulaArray("=LOWER(A1:B2)")
	sheet.RecalculateFormulas()

	td := []testStruct{
		{`=D1`, `400 ResultTypeArray`},
		{`=D2`, `hello ResultTypeString`},
		{`=E1`, `world! ResultTypeString`},
		{`=E2`, ` ResultTypeEmpty`},
	}

	ctx := sheet.FormulaContext()

	runTests(t, ctx, td)
}

func TestReplace(t *testing.T) {
	ss := spreadsheet.New()
	sheet := ss.AddSheet()

	sheet.Cell("A1").SetNumber(400)
	sheet.Cell("A2").SetString("Hello")
	sheet.Cell("B1").SetString("World!")

	sheet.Cell("D1").SetFormulaArray("=LOWER(A1:B2)")
	sheet.RecalculateFormulas()

	td := []testStruct{
		{`=REPLACE("Hello World!",7,5,"Earth")`, `Hello Earth! ResultTypeString`},
		{`=REPLACE("Hello World!",7,10,"Earth")`, `Hello Earth ResultTypeString`},
		{`=REPLACE("Hello World",30,10,"!")`, `Hello World! ResultTypeString`},
		{`=REPLACE("Hello World!",7,0,"new ")`, `Hello new World! ResultTypeString`},
	}

	ctx := sheet.FormulaContext()

	runTests(t, ctx, td)
}

func TestTextJoin(t *testing.T) {
	ss := spreadsheet.New()
	sheet := ss.AddSheet()

	sheet.Cell("A1").SetNumber(1)
	sheet.Cell("B1").SetString("2")
	sheet.Cell("A2").SetNumber(3)
	sheet.Cell("B2").SetString("4")
	sheet.Cell("A3").SetNumber(5)
	sheet.Cell("B3").SetString("6")
	sheet.Cell("A4").SetString("7")
	sheet.Cell("A6").SetString("8")
	sheet.Cell("A7").SetString("9")

	td := []testStruct{
		{`=TEXTJOIN(".",FALSE,A1)`, `1 ResultTypeString`},
		{`=TEXTJOIN("|",TRUE,A4:A7)`, `7|8|9 ResultTypeString`},
		{`=TEXTJOIN("|",FALSE,A4:A7)`, `7||8|9 ResultTypeString`},
		{`=TEXTJOIN(".",TRUE,A1:B2,A3:B3,A4,A5,A6,A7)`, `1.2.3.4.5.6.7.8.9 ResultTypeString`},
	}

	ctx := sheet.FormulaContext()

	runTests(t, ctx, td)
}

func TestIndex(t *testing.T) {
	ss := spreadsheet.New()
	sheet := ss.AddSheet()

	sheet.Cell("A1").SetNumber(1)
	sheet.Cell("B1").SetString("2")
	sheet.Cell("C1").SetNumber(3)
	sheet.Cell("B2").SetNumber(5)
	sheet.Cell("C2").SetString("6")
	sheet.Cell("A3").SetString("7")
	sheet.Cell("B3").SetString("8")
	sheet.Cell("C3").SetString("9")

	sheet.Cell("A4").SetFormulaRaw(`=INDEX(A1:C3,1,1)`)
	sheet.Cell("A5").SetFormulaArray(`=INDEX(A1:C3,2)`)
	sheet.Cell("A6").SetFormulaArray(`=INDEX(A1:C3,,2)`)

	sheet.RecalculateFormulas()

	td := []testStruct{
		{`=A4`, `1 ResultTypeNumber`},
		{`=A5`, `0 ResultTypeArray`},
		{`=B5`, `5 ResultTypeNumber`},
		{`=C5`, `6 ResultTypeNumber`},
		{`=A7`, `5 ResultTypeNumber`},
		{`=A8`, `8 ResultTypeNumber`},
	}

	ctx := sheet.FormulaContext()

	runTests(t, ctx, td)
}

func TestText(t *testing.T) {
	ss := spreadsheet.New()
	sheet := ss.AddSheet()

	td := []testStruct{
		{`=TEXT(A1,"0#.00")`, `00.00 ResultTypeString`},
		{`=TEXT(1,"0.00")`, `1.00 ResultTypeString`},
		{`=TEXT(12345678,"0.00E+000")`, `1.23E+007 ResultTypeString`},
		{`=TEXT(0.987654321,"0.000%")`, `98.765% ResultTypeString`},
		{`=TEXT(0.05,"# ??/??")`, `1/20 ResultTypeString`},
	}

	ctx := sheet.FormulaContext()

	runTests(t, ctx, td)
}

func TestSum(t *testing.T) {
	ss := spreadsheet.New()
	sheet := ss.AddSheet()

	sheet.Cell("A1").SetNumber(1)
	sheet.Cell("B1").SetNumber(2)
	sheet.Cell("C1").SetNumber(3)
	sheet.Cell("D1").SetNumber(4)
	sheet.Cell("E1").SetNumber(5)

	sheet.Cell("A2").SetNumber(1)
	sheet.Cell("B2").SetNumber(10)
	sheet.Cell("C2").SetNumber(100)
	sheet.Cell("D2").SetNumber(1000)
	sheet.Cell("E2").SetNumber(10000)

	ctx := sheet.FormulaContext()

	td := []testStruct{
		{`=SUM(A1:E1)`, `15 ResultTypeNumber`},
		{`=SUM(A2:C2,E2)`, `10111 ResultTypeNumber`},
	}

	runTests(t, ctx, td)
}

func TestIf(t *testing.T) {
	ss := spreadsheet.New()
	sheet := ss.AddSheet()

	sheet.Cell("A1").SetNumber(1)
	sheet.Cell("A2").SetNumber(2)
	sheet.Cell("A3").SetNumber(3)
	sheet.Cell("A4").SetNumber(4)
	sheet.Cell("A5").SetNumber(5)
	sheet.Cell("A6").SetFormulaArray(`IF(A1:A5>2,"passed","not passed")`)

	sheet.RecalculateFormulas()

	ctx := sheet.FormulaContext()

	td := []testStruct{
		{`=IF(A5=5,"five",""))`, `five ResultTypeString`},
		{`=A6)`, `not passed ResultTypeArray`},
		{`=A8)`, `passed ResultTypeString`},
		{`=A9)`, `passed ResultTypeString`},
	}

	runTests(t, ctx, td)
}

func TestAccrintm(t *testing.T) {
	ss := spreadsheet.New()
	sheet := ss.AddSheet()

	ctx := sheet.FormulaContext()

	td := []testStruct{
		{`=ACCRINTM(39539,39614,0.1,1000)`, `20.5555555555 ResultTypeNumber`},
		{`=ACCRINTM(39539,39614,0.1,1000,)`, `20.5555555555 ResultTypeNumber`},
		{`=ACCRINTM(39539,39614,0.1,1000,0)`, `20.5555555555 ResultTypeNumber`},
		{`=ACCRINTM(39539,39614,0.1,1000,1)`, `20.4918032786 ResultTypeNumber`},
		{`=ACCRINTM(39539,39614,0.1,1000,2)`, `20.8333333333 ResultTypeNumber`},
		{`=ACCRINTM(39539,39614,0.1,1000,3)`, `20.5479452054 ResultTypeNumber`},
		{`=ACCRINTM(39539,39614,0.1,1000,4)`, `20.5555555555 ResultTypeNumber`},
	}

	runTests(t, ctx, td)
}

func TestAmordegrc(t *testing.T) {
	ss := spreadsheet.New()
	sheet := ss.AddSheet()

	ctx := sheet.FormulaContext()

	td := []testStruct{
		{`=AMORDEGRC(2400,39679,39813,300,1,0.15)`, `776 ResultTypeNumber`},
		{`=AMORDEGRC(2400,39679,39813,300,1,0.15,)`, `776 ResultTypeNumber`},
		{`=AMORDEGRC(2400,39679,39813,300,1,0.15,0)`, `776 ResultTypeNumber`},
		{`=AMORDEGRC(2400,39679,39813,300,1,0.15,1)`, `776 ResultTypeNumber`},
		{`=AMORDEGRC(2400,39679,39813,300,1,0.15,2)`, `#NUM! ResultTypeError`},
	}

	runTests(t, ctx, td)
}

func TestAmorlinc(t *testing.T) {
	ss := spreadsheet.New()
	sheet := ss.AddSheet()

	ctx := sheet.FormulaContext()

	td := []testStruct{
		{`=AMORLINC(2400,39679,39813,300,1,0.15)`, `360 ResultTypeNumber`},
		{`=AMORLINC(2400,39679,39813,300,1,0.15,)`, `360 ResultTypeNumber`},
		{`=AMORLINC(2400,39679,39813,300,1,0.15,0)`, `360 ResultTypeNumber`},
		{`=AMORLINC(2400,39679,39813,300,1,0.15,1)`, `360 ResultTypeNumber`},
		{`=AMORLINC(2400,39679,39813,300,1,0.15,2)`, `#NUM! ResultTypeError`},
	}

	runTests(t, ctx, td)
}

func TestCoupdaybs(t *testing.T) {
	ss := spreadsheet.New()
	sheet := ss.AddSheet()

	ctx := sheet.FormulaContext()

	td := []testStruct{
		{`=COUPDAYBS(40568,40862,2)`, `70 ResultTypeNumber`},
		{`=COUPDAYBS(40568,40862,2,)`, `70 ResultTypeNumber`},
		{`=COUPDAYBS(40568,40862,2,0)`, `70 ResultTypeNumber`},
		{`=COUPDAYBS(40568,40862,2,1)`, `71 ResultTypeNumber`},
		{`=COUPDAYBS(40568,40862,2,2)`, `71 ResultTypeNumber`},
		{`=COUPDAYBS(40568,40862,2,3)`, `71 ResultTypeNumber`},
		{`=COUPDAYBS(40568,40862,2,4)`, `70 ResultTypeNumber`},
		{`=COUPDAYBS(40599,40862,2,0)`, `100 ResultTypeNumber`},
		{`=COUPDAYBS(40599,40862,2,1)`, `102 ResultTypeNumber`},
		{`=COUPDAYBS(40599,40862,2,2)`, `102 ResultTypeNumber`},
		{`=COUPDAYBS(40599,40862,2,3)`, `102 ResultTypeNumber`},
		{`=COUPDAYBS(40599,40862,2,4)`, `100 ResultTypeNumber`},
		{`=COUPDAYBS(40811,40862,2,0)`, `130 ResultTypeNumber`},
		{`=COUPDAYBS(40811,40862,2,1)`, `133 ResultTypeNumber`},
		{`=COUPDAYBS(40811,40862,2,2)`, `133 ResultTypeNumber`},
		{`=COUPDAYBS(40811,40862,2,3)`, `133 ResultTypeNumber`},
		{`=COUPDAYBS(40811,40862,2,4)`, `130 ResultTypeNumber`},
		{`=COUPDAYBS(40872,40568,2,1)`, `#NUM! ResultTypeError`},
	}

	runTests(t, ctx, td)
}

func TestCoupdays(t *testing.T) {
	ss := spreadsheet.New()
	sheet := ss.AddSheet()

	ctx := sheet.FormulaContext()

	td := []testStruct{
		{`=COUPDAYS(40964,41228,1)`, `360 ResultTypeNumber`},
		{`=COUPDAYS(40964,41228,1,)`, `360 ResultTypeNumber`},
		{`=COUPDAYS(40964,41228,1,0)`, `360 ResultTypeNumber`},
		{`=COUPDAYS(40964,41228,1,1)`, `366 ResultTypeNumber`},
		{`=COUPDAYS(40964,41228,1,2)`, `360 ResultTypeNumber`},
		{`=COUPDAYS(40964,41228,1,3)`, `365 ResultTypeNumber`},
		{`=COUPDAYS(40964,41228,1,4)`, `360 ResultTypeNumber`},
		{`=COUPDAYS(40964,41228,2,0)`, `180 ResultTypeNumber`},
		{`=COUPDAYS(40964,41228,2,1)`, `182 ResultTypeNumber`},
		{`=COUPDAYS(40964,41228,2,2)`, `180 ResultTypeNumber`},
		{`=COUPDAYS(40964,41228,2,3)`, `182.5 ResultTypeNumber`},
		{`=COUPDAYS(40964,41228,2,4)`, `180 ResultTypeNumber`},
		{`=COUPDAYS(40964,41228,4,0)`, `90 ResultTypeNumber`},
		{`=COUPDAYS(40964,41228,4,1)`, `90 ResultTypeNumber`},
		{`=COUPDAYS(40964,41228,4,2)`, `90 ResultTypeNumber`},
		{`=COUPDAYS(40964,41228,4,3)`, `91.25 ResultTypeNumber`},
		{`=COUPDAYS(40964,41228,4,4)`, `90 ResultTypeNumber`},
		{`=COUPDAYS(41228,40964,2,1)`, `#NUM! ResultTypeError`},
	}

	runTests(t, ctx, td)
}

func TestCoupdaysnc(t *testing.T) {
	ss := spreadsheet.New()
	sheet := ss.AddSheet()

	ctx := sheet.FormulaContext()

	td := []testStruct{
		{`=COUPDAYSNC(40933,41228,1)`, `290 ResultTypeNumber`},
		{`=COUPDAYSNC(40933,41228,1,)`, `290 ResultTypeNumber`},
		{`=COUPDAYSNC(40933,41228,1,0)`, `290 ResultTypeNumber`},
		{`=COUPDAYSNC(40933,41228,1,1)`, `295 ResultTypeNumber`},
		{`=COUPDAYSNC(40933,41228,1,2)`, `295 ResultTypeNumber`},
		{`=COUPDAYSNC(40933,41228,1,3)`, `295 ResultTypeNumber`},
		{`=COUPDAYSNC(40933,41228,1,4)`, `290 ResultTypeNumber`},
		{`=COUPDAYSNC(40933,41228,2,0)`, `110 ResultTypeNumber`},
		{`=COUPDAYSNC(40933,41228,2,1)`, `111 ResultTypeNumber`},
		{`=COUPDAYSNC(40933,41228,2,2)`, `111 ResultTypeNumber`},
		{`=COUPDAYSNC(40933,41228,2,3)`, `111 ResultTypeNumber`},
		{`=COUPDAYSNC(40933,41228,2,4)`, `110 ResultTypeNumber`},
		{`=COUPDAYSNC(40933,41228,4,0)`, `20 ResultTypeNumber`},
		{`=COUPDAYSNC(40933,41228,4,1)`, `21 ResultTypeNumber`},
		{`=COUPDAYSNC(40933,41228,4,2)`, `21 ResultTypeNumber`},
		{`=COUPDAYSNC(40933,41228,4,3)`, `21 ResultTypeNumber`},
		{`=COUPDAYSNC(40933,41228,4,4)`, `20 ResultTypeNumber`},
		{`=COUPDAYSNC(41228,40933,2,1)`, `#NUM! ResultTypeError`},
	}

	runTests(t, ctx, td)
}

func TestCoupncd(t *testing.T) {
	ss := spreadsheet.New()
	sheet := ss.AddSheet()

	ctx := sheet.FormulaContext()

	td := []testStruct{
		{`=COUPNCD(40568,40862,1)`, `40862 ResultTypeNumber`},
		{`=COUPNCD(40568,40862,1,)`, `40862 ResultTypeNumber`},
		{`=COUPNCD(40568,40862,1,0)`, `40862 ResultTypeNumber`},
		{`=COUPNCD(40568,40862,1,1)`, `40862 ResultTypeNumber`},
		{`=COUPNCD(40568,40862,2,1)`, `40678 ResultTypeNumber`},
		{`=COUPNCD(40568,40862,4,1)`, `40589 ResultTypeNumber`},
		{`=COUPNCD(40872,40568,2,1)`, `#NUM! ResultTypeError`},
	}

	runTests(t, ctx, td)
}

func TestCouppcd(t *testing.T) {
	ss := spreadsheet.New()
	sheet := ss.AddSheet()

	ctx := sheet.FormulaContext()

	td := []testStruct{
		{`=COUPPCD(40568,40862,2)`, `40497 ResultTypeNumber`},
		{`=COUPPCD(40568,40862,2,)`, `40497 ResultTypeNumber`},
		{`=COUPPCD(40568,40862,2,0)`, `40497 ResultTypeNumber`},
		{`=COUPPCD(40568,40862,2,1)`, `40497 ResultTypeNumber`},
		{`=COUPPCD(40872,40568,2,1)`, `#NUM! ResultTypeError`},
	}

	runTests(t, ctx, td)
}

func TestCoupnum(t *testing.T) {
	ss := spreadsheet.New()
	sheet := ss.AddSheet()

	ctx := sheet.FormulaContext()

	td := []testStruct{
		{`=COUPNUM(39107,39767,2)`, `4 ResultTypeNumber`},
		{`=COUPNUM(39107,39767,2,)`, `4 ResultTypeNumber`},
		{`=COUPNUM(39107,39767,2,0)`, `4 ResultTypeNumber`},
		{`=COUPNUM(39107,39767,2,1)`, `4 ResultTypeNumber`},
		{`=COUPNUM(39767,39107,2,1)`, `#NUM! ResultTypeError`},
	}

	runTests(t, ctx, td)
}

func TestCumipmt(t *testing.T) {
	ss := spreadsheet.New()
	sheet := ss.AddSheet()

	sheet.Cell("A1").SetNumber(0.09)
	sheet.Cell("A2").SetNumber(30)
	sheet.Cell("A3").SetNumber(125000)

	ctx := sheet.FormulaContext()

	td := []testStruct{
		{`=CUMIPMT(A1/12,A2*12,A3,13,24,0)`, `-11135.23213 ResultTypeNumber`},
		{`=CUMIPMT(A1/12,A2*12,A3,1,1,0)`, `-937.5 ResultTypeNumber`},
	}

	runTests(t, ctx, td)
}

func TestCumprinc(t *testing.T) {
	ss := spreadsheet.New()
	sheet := ss.AddSheet()

	sheet.Cell("A1").SetNumber(0.09)
	sheet.Cell("A2").SetNumber(30)
	sheet.Cell("A3").SetNumber(125000)

	ctx := sheet.FormulaContext()

	td := []testStruct{
		{`=CUMPRINC(A1/12,A2*12,A3,13,24,0)`, `-934.10712342 ResultTypeNumber`},
		{`=CUMPRINC(A1/12,A2*12,A3,1,1,0)`, `-68.27827118 ResultTypeNumber`},
	}

	runTests(t, ctx, td)
}

func TestDb(t *testing.T) {
	ss := spreadsheet.New()
	sheet := ss.AddSheet()

	sheet.Cell("A1").SetNumber(1000000)
	sheet.Cell("A2").SetNumber(100000)
	sheet.Cell("A3").SetNumber(6)

	ctx := sheet.FormulaContext()

	td := []testStruct{
		{`=DB(A1,A2,A3,1,7)`, `186083.333333 ResultTypeNumber`},
		{`=DB(A1,A2,A3,2,7)`, `259639.416666 ResultTypeNumber`},
		{`=DB(A1,A2,A3,3,7)`, `176814.44275 ResultTypeNumber`},
		{`=DB(A1,A2,A3,4,7)`, `120410.635512 ResultTypeNumber`},
		{`=DB(A1,A2,A3,5,7)`, `81999.6427841 ResultTypeNumber`},
		{`=DB(A1,A2,A3,6,7)`, `55841.756736 ResultTypeNumber`},
		{`=DB(A1,A2,A3,7,7)`, `15845.0984738 ResultTypeNumber`},
	}

	runTests(t, ctx, td)
}

func TestDdb(t *testing.T) {
	ss := spreadsheet.New()
	sheet := ss.AddSheet()

	sheet.Cell("A1").SetNumber(2400)
	sheet.Cell("A2").SetNumber(300)
	sheet.Cell("A3").SetNumber(10)

	ctx := sheet.FormulaContext()

	td := []testStruct{
		{`=DDB(A1,A2,A3*365,1)`, `1.31506849315 ResultTypeNumber`},
		{`=DDB(A1,A2,A3*12,1,2)`, `40 ResultTypeNumber`},
		{`=DDB(A1,A2,A3,1,2)`, `480 ResultTypeNumber`},
		{`=DDB(A1,A2,A3,2,1.5)`, `306. ResultTypeNumber`},
		{`=DDB(A1,A2,A3,10)`, `22.1225472 ResultTypeNumber`},
	}

	runTests(t, ctx, td)
}

func TestDisc(t *testing.T) {
	ss := spreadsheet.New()
	sheet := ss.AddSheet()

	sheet.Cell("A1").SetDate(time.Date(2018, 7, 1, 0, 0, 0, 0, time.UTC))
	sheet.Cell("A2").SetDate(time.Date(2048, 1, 1, 0, 0, 0, 0, time.UTC))
	sheet.Cell("A3").SetNumber(97.975)
	sheet.Cell("A4").SetNumber(100)

	ctx := sheet.FormulaContext()

	td := []testStruct{
		{`=DISC(A1,A2,A3,A4)`, `0.00068644067 ResultTypeNumber`},
		{`=DISC(A1,A2,A3,A4,)`, `0.00068644067 ResultTypeNumber`},
		{`=DISC(A1,A2,A3,A4,0)`, `0.00068644067 ResultTypeNumber`},
		{`=DISC(A1,A2,A3,A4,1)`, `0.00068638416 ResultTypeNumber`},
		{`=DISC(A1,A2,A3,A4,2)`, `0.00067650334 ResultTypeNumber`},
		{`=DISC(A1,A2,A3,A4,3)`, `0.00068589922 ResultTypeNumber`},
		{`=DISC(A1,A2,A3,A4,4)`, `0.00068644067 ResultTypeNumber`},
	}

	runTests(t, ctx, td)
}

func TestDollarde(t *testing.T) {
	ss := spreadsheet.New()
	sheet := ss.AddSheet()

	ctx := sheet.FormulaContext()

	td := []testStruct{
		{`=DOLLARDE(1.02,16)`, `1.125 ResultTypeNumber`},
		{`=DOLLARDE(1.1,32)`, `1.3125 ResultTypeNumber`},
		{`=DOLLARDE(-1.1,32)`, `-1.3125 ResultTypeNumber`},
	}

	runTests(t, ctx, td)
}

func TestDollarfr(t *testing.T) {
	ss := spreadsheet.New()
	sheet := ss.AddSheet()

	ctx := sheet.FormulaContext()

	td := []testStruct{
		{`=DOLLARFR(1.125,16)`, `1.02 ResultTypeNumber`},
		{`=DOLLARFR(1.125,32)`, `1.04 ResultTypeNumber`},
		{`=DOLLARFR(-1.125,32)`, `-1.04 ResultTypeNumber`},
	}

	runTests(t, ctx, td)
}

func TestEffect(t *testing.T) {
	ss := spreadsheet.New()
	sheet := ss.AddSheet()

	ctx := sheet.FormulaContext()

	td := []testStruct{
		{`=EFFECT(0.0525,4)`, `0.05354266737 ResultTypeNumber`},
		{`=EFFECT(0.1,4)`, `0.10381289062 ResultTypeNumber`},
		{`=EFFECT(0.1,4.5)`, `0.10381289062 ResultTypeNumber`},
		{`=EFFECT(0,4.5)`, `#NUM! ResultTypeError`},
		{`=EFFECT(0.1,0.5)`, `#NUM! ResultTypeError`},
		{`=EFFECT("Hello world",4)`, `#VALUE! ResultTypeError`},
	}

	runTests(t, ctx, td)
}

func TestFv(t *testing.T) {
	ss := spreadsheet.New()
	sheet := ss.AddSheet()

	ctx := sheet.FormulaContext()

	td := []testStruct{
		{`=FV(0.06/12,10,-200,-500,1)`, `2581.40337406 ResultTypeNumber`},
		{`=FV(0,12,-100,-1000,1)`, `2200 ResultTypeNumber`},
	}

	runTests(t, ctx, td)
}

func TestFvschedule(t *testing.T) {
	ss := spreadsheet.New()
	sheet := ss.AddSheet()

	sheet.Cell("A1").SetNumber(0.09)
	sheet.Cell("A2").SetNumber(0.11)
	sheet.Cell("A3").SetNumber(0.1)
	sheet.Cell("A4").SetBool(true)

	ctx := sheet.FormulaContext()

	td := []testStruct{
		{`=FVSCHEDULE(1,A1:A3)`, `1.33089 ResultTypeNumber`},
		{`=FVSCHEDULE(1,A1:A4)`, `#VALUE! ResultTypeError`},
	}

	runTests(t, ctx, td)
}

func TestIntrate(t *testing.T) {
	ss := spreadsheet.New()
	sheet := ss.AddSheet()

	sheet.Cell("A1").SetDate(time.Date(2008, 2, 15, 0, 0, 0, 0, time.UTC))
	sheet.Cell("A2").SetDate(time.Date(2008, 5, 15, 0, 0, 0, 0, time.UTC))
	sheet.Cell("A3").SetNumber(1000000)
	sheet.Cell("A4").SetNumber(1014420)

	ctx := sheet.FormulaContext()

	td := []testStruct{
		{`=INTRATE(A1,A2,A3,A4)`, `0.05768 ResultTypeNumber`},
		{`=INTRATE(A1,A2,A3,A4,)`, `0.05768 ResultTypeNumber`},
		{`=INTRATE(A1,A2,A3,A4,0)`, `0.05768 ResultTypeNumber`},
		{`=INTRATE(A1,A2,A3,A4,1)`, `0.05864133333 ResultTypeNumber`},
		{`=INTRATE(A1,A2,A3,A4,2)`, `0.05768 ResultTypeNumber`},
		{`=INTRATE(A1,A2,A3,A4,3)`, `0.05848111111 ResultTypeNumber`},
		{`=INTRATE(A1,A2,A3,A4,4)`, `0.05768 ResultTypeNumber`},
	}

	runTests(t, ctx, td)
}

func TestIpmt(t *testing.T) {
	ss := spreadsheet.New()
	sheet := ss.AddSheet()

	sheet.Cell("A1").SetNumber(0.1)
	sheet.Cell("A2").SetNumber(1)
	sheet.Cell("A3").SetNumber(3)
	sheet.Cell("A4").SetNumber(8000)

	ctx := sheet.FormulaContext()

	td := []testStruct{
		{`=IPMT(0.1/12,1,36,8000)`, `-66.666666666 ResultTypeNumber`},
		{`=IPMT(0.1,3,3,8000)`, `-292.4471299 ResultTypeNumber`},
		{`=IPMT(0.1/12,6,24,100000,1000000,0)`, `928.82357184 ResultTypeNumber`},
		{`=IPMT(0.1/12,6,24,100000,1000000,1)`, `921.147343973 ResultTypeNumber`},
		{`=IPMT(0.1/12,1,24,100000,1000000,1)`, `0 ResultTypeNumber`},
		{`=IPMT(0.1/12,1,24,100000,1000000,0)`, `-833.33333333 ResultTypeNumber`},
	}

	runTests(t, ctx, td)
}

func TestIrr(t *testing.T) {
	ss := spreadsheet.New()
	sheet := ss.AddSheet()

	sheet.Cell("A1").SetNumber(-70000)
	sheet.Cell("A2").SetNumber(12000)
	sheet.Cell("A3").SetNumber(15000)
	sheet.Cell("A4").SetNumber(18000)
	sheet.Cell("A5").SetNumber(21000)
	sheet.Cell("A6").SetNumber(26000)

	ctx := sheet.FormulaContext()

	td := []testStruct{
		{`=IRR(A1:A6)`, `0.08663094803 ResultTypeNumber`},
		{`=IRR(A1:A5)`, `-0.0212448482 ResultTypeNumber`},
		{`=IRR(A1:A4)`, `-0.1821374641 ResultTypeNumber`},
		{`=IRR(A1:A3,0.2)`, `-0.4435069413 ResultTypeNumber`},
	}

	runTests(t, ctx, td)
}

func TestIspmt(t *testing.T) {
	ss := spreadsheet.New()
	sheet := ss.AddSheet()

	sheet.Cell("A1").SetNumber(0.1)
	sheet.Cell("A2").SetNumber(4)
	sheet.Cell("A3").SetNumber(4000)

	ctx := sheet.FormulaContext()

	td := []testStruct{
		{`=ISPMT(0.1/12,6,2*12,100000)`, `-625 ResultTypeNumber`},
		{`=ISPMT(A1,0,A2,A3)`, `-400 ResultTypeNumber`},
		{`=ISPMT(A1,1,A2,A3)`, `-300 ResultTypeNumber`},
		{`=ISPMT(A1,2,A2,A3)`, `-200 ResultTypeNumber`},
		{`=ISPMT(A1,3,A2,A3)`, `-100 ResultTypeNumber`},
	}

	runTests(t, ctx, td)
}

func TestMirr(t *testing.T) {
	ss := spreadsheet.New()
	sheet := ss.AddSheet()

	sheet.Cell("A1").SetNumber(-120000)
	sheet.Cell("A2").SetNumber(39000)
	sheet.Cell("A3").SetNumber(30000)
	sheet.Cell("A4").SetNumber(21000)
	sheet.Cell("A5").SetNumber(37000)
	sheet.Cell("A6").SetNumber(46000)

	ctx := sheet.FormulaContext()

	td := []testStruct{
		{`=MIRR(A1:A6,0.1,0.12)`, `0.12609413036 ResultTypeNumber`},
		{`=MIRR(A1:A4,0.1,0.12)`, `-0.0480446552 ResultTypeNumber`},
		{`=MIRR(A1:A6,0.1,0.14)`, `0.13475911082 ResultTypeNumber`},
		{`=MIRR(A1:A6,0.2,0.14)`, `0.13475911082 ResultTypeNumber`},
		{`=MIRR(A1:A6,0.3,0.14)`, `0.13475911082 ResultTypeNumber`},
		{`=MIRR(A1:A6,0.4,0.14)`, `0.13475911082 ResultTypeNumber`},
		{`=MIRR(A1:A6,0,0.14)`, `0.13475911082 ResultTypeNumber`},
		{`=MIRR(A1:A6,-1,0.14)`, `0.13475911082 ResultTypeNumber`},
		{`=MIRR(A1:A6,0.1,-1)`, `#DIV/0! ResultTypeError`},
	}

	runTests(t, ctx, td)
}

func TestNominal(t *testing.T) {
	ss := spreadsheet.New()
	sheet := ss.AddSheet()

	ctx := sheet.FormulaContext()

	td := []testStruct{
		{`=NOMINAL(0.053543,4)`, `0.05250031986 ResultTypeNumber`},
		{`=NOMINAL(0.1,4)`, `0.09645475633 ResultTypeNumber`},
		{`=NOMINAL(0.1,4.5)`, `0.09645475633 ResultTypeNumber`},
		{`=NOMINAL(0,4.5)`, `#NUM! ResultTypeError`},
		{`=NOMINAL(0.1,0.5)`, `#NUM! ResultTypeError`},
		{`=NOMINAL("Hello world",4.5)`, `#VALUE! ResultTypeError`},
	}

	runTests(t, ctx, td)
}

func TestNper(t *testing.T) {
	ss := spreadsheet.New()
	sheet := ss.AddSheet()

	sheet.Cell("A1").SetNumber(0.12)
	sheet.Cell("A2").SetNumber(-100)
	sheet.Cell("A3").SetNumber(-1000)
	sheet.Cell("A4").SetNumber(10000)

	ctx := sheet.FormulaContext()

	td := []testStruct{
		{`=NPER(A1/12,A2,A3,A4,1)`, `59.6738656742 ResultTypeNumber`},
		{`=NPER(A1/12,A2,A3,A4)`, `60.0821228537 ResultTypeNumber`},
		{`=NPER(A1/12,A2,A3,A4,)`, `60.0821228537 ResultTypeNumber`},
		{`=NPER(A1/12,A2,A3,A4,0)`, `60.0821228537 ResultTypeNumber`},
		{`=NPER(A1/12,A2,A3)`, `-9.5785940398 ResultTypeNumber`},
		{`=NPER(A1/12,A2,A3,,)`, `-9.5785940398 ResultTypeNumber`},
		{`=NPER(A1/12,A2,A3,0,)`, `-9.5785940398 ResultTypeNumber`},
		{`=NPER(A1/12,A2,A3,,0)`, `-9.5785940398 ResultTypeNumber`},
	}

	runTests(t, ctx, td)
}

func TestNpv(t *testing.T) {
	ss := spreadsheet.New()
	sheet := ss.AddSheet()

	sheet.Cell("A1").SetNumber(0.1)
	sheet.Cell("A2").SetNumber(-10000)
	sheet.Cell("A3").SetNumber(3000)
	sheet.Cell("A4").SetNumber(4200)
	sheet.Cell("A5").SetNumber(6800)

	ctx := sheet.FormulaContext()

	td := []testStruct{
		{`=NPV(A1,A2,A3,A4,A5)`, `1188.44341233 ResultTypeNumber`},
		{`=NPV(A1,A2:A4,A5)`, `1188.44341233 ResultTypeNumber`},
		{`=NPV(A1,A2:A4,"Hello world",A5)`, `1188.44341233 ResultTypeNumber`},
		{`=NPV(0.12,12000,15000,18000,21000,24000)`, `62448.3625219 ResultTypeNumber`},
	}

	runTests(t, ctx, td)
}

func TestPmt(t *testing.T) {
	ss := spreadsheet.New()
	sheet := ss.AddSheet()

	sheet.Cell("A1").SetNumber(0.08)
	sheet.Cell("A2").SetNumber(10)
	sheet.Cell("A3").SetNumber(10000)
	sheet.Cell("A4").SetNumber(0.06)
	sheet.Cell("A5").SetNumber(18)
	sheet.Cell("A6").SetNumber(50000)

	ctx := sheet.FormulaContext()

	td := []testStruct{
		{`=PMT(A1/12,A2,A3)`, `-1037.0320893 ResultTypeNumber`},
		{`=PMT(A1/12,A2,A3,)`, `-1037.0320893 ResultTypeNumber`},
		{`=PMT(A1/12,A2,A3,0)`, `-1037.0320893 ResultTypeNumber`},
		{`=PMT(A1/12,A2,A3,1)`, `-1037.1291259 ResultTypeNumber`},
		{`=PMT(A4/12,A5*12,0,A6)`, `-129.08116086 ResultTypeNumber`},
		{`=PMT("A4/12",A5*12,0,A6)`, `#VALUE! ResultTypeError`},
	}

	runTests(t, ctx, td)
}

func TestPpmt(t *testing.T) {
	ss := spreadsheet.New()
	sheet := ss.AddSheet()

	ctx := sheet.FormulaContext()

	td := []testStruct{
		{`=PPMT(0.1/12,1,2*12,2000)`, `-75.623186008 ResultTypeNumber`},
		{`=PPMT(0.08,10,10,200000)`, `-27598.053462 ResultTypeNumber`},
		{`=PPMT(0.08,11,10,200000)`, `#NUM! ResultTypeError`},
		{`=PPMT("0.08%",10,10,200000)`, `#VALUE! ResultTypeError`},
	}

	runTests(t, ctx, td)
}

func TestPricedisc(t *testing.T) {
	ss := spreadsheet.New()
	sheet := ss.AddSheet()

	sheet.Cell("A1").SetDate(time.Date(2008, 2, 16, 0, 0, 0, 0, time.UTC))
	sheet.Cell("A2").SetDate(time.Date(2008, 3, 1, 0, 0, 0, 0, time.UTC))
	sheet.Cell("A3").SetNumber(0.0525)
	sheet.Cell("A4").SetNumber(100)

	ctx := sheet.FormulaContext()

	td := []testStruct{
		{`=PRICEDISC(A1,A2,A3,A4)`, `99.78125 ResultTypeNumber`},
		{`=PRICEDISC(A1,A2,A3,A4,)`, `99.78125 ResultTypeNumber`},
		{`=PRICEDISC(A1,A2,A3,A4,0)`, `99.78125 ResultTypeNumber`},
		{`=PRICEDISC(A1,A2,A3,A4,1)`, `99.7991803278 ResultTypeNumber`},
		{`=PRICEDISC(A1,A2,A3,A4,2)`, `99.7958333333 ResultTypeNumber`},
		{`=PRICEDISC(A1,A2,A3,A4,3)`, `99.7986301369 ResultTypeNumber`},
		{`=PRICEDISC(A1,A2,A3,A4,4)`, `99.78125 ResultTypeNumber`},
	}

	runTests(t, ctx, td)
}

func TestPv(t *testing.T) {
	ss := spreadsheet.New()
	sheet := ss.AddSheet()

	ctx := sheet.FormulaContext()

	td := []testStruct{
		{`=PV(0.08/12,20*12,500)`, `-59777.145851 ResultTypeNumber`},
		{`=PV(0.08/12,20*12,500,,)`, `-59777.145851 ResultTypeNumber`},
		{`=PV(0.08/12,20*12,500,0,)`, `-59777.145851 ResultTypeNumber`},
		{`=PV(0.08/12,20*12,500,,0)`, `-59777.145851 ResultTypeNumber`},
		{`=PV(0.08/12,20*12,500,0,0)`, `-59777.145851 ResultTypeNumber`},
		{`=PV(0.1/12,2*12,1000,10000)`, `-29864.950264 ResultTypeNumber`},
		{`=PV(0.1/12,2*12,1000,10000,)`, `-29864.950264 ResultTypeNumber`},
		{`=PV(0.1/12,2*12,1000,10000,0)`, `-29864.950264 ResultTypeNumber`},
		{`=PV(0.1/12,2*12,1000,10000,1)`, `-30045.540721 ResultTypeNumber`},
		{`=PV(0,2*12,1000,10000,1)`, `-34000 ResultTypeNumber`},
		{`=PV("hello world",2*12,1000,10000,1)`, `#VALUE! ResultTypeError`},
	}

	runTests(t, ctx, td)
}

func TestRate(t *testing.T) {
	ss := spreadsheet.New()
	sheet := ss.AddSheet()

	ctx := sheet.FormulaContext()

	td := []testStruct{
		{`=RATE(2*12,-1000,-10000,100000)`, `0.06517891177 ResultTypeNumber`},
		{`=RATE(2*12,-1000,-10000,100000,)`, `0.06517891177 ResultTypeNumber`},
		{`=RATE(2*12,-1000,-10000,100000,,)`, `0.06517891177 ResultTypeNumber`},
		{`=RATE(2*12,-1000,-10000,100000,0,0.1)`, `0.06517891177 ResultTypeNumber`},
		{`=RATE(2*12,-1000,-10000,100000,0,0.75)`, `0.06517891177 ResultTypeNumber`},
		{`=RATE(2*12,-1000,-10000,100000,0,0.065)`, `0.06517891177 ResultTypeNumber`},
		{`=RATE(2*12,-1000,-10000,100000,1,0.1)`, `0.06323958 ResultTypeNumber`},
	}

	runTests(t, ctx, td)
}

func TestReceived(t *testing.T) {
	ss := spreadsheet.New()
	sheet := ss.AddSheet()

	sheet.Cell("A1").SetDate(time.Date(2008, 2, 15, 0, 0, 0, 0, time.UTC))
	sheet.Cell("A2").SetDate(time.Date(2008, 5, 15, 0, 0, 0, 0, time.UTC))
	sheet.Cell("A3").SetNumber(1000000)
	sheet.Cell("A4").SetNumber(0.0575)

	ctx := sheet.FormulaContext()

	td := []testStruct{
		{`=RECEIVED(A1,A2,A3,A4)`, `1014584.6544 ResultTypeNumber`},
		{`=RECEIVED(A1,A2,A3,A4,)`, `1014584.6544 ResultTypeNumber`},
		{`=RECEIVED(A1,A2,A3,A4,0)`, `1014584.6544 ResultTypeNumber`},
		{`=RECEIVED(A1,A2,A3,A4,1)`, `1014342.13261 ResultTypeNumber`},
		{`=RECEIVED(A1,A2,A3,A4,2)`, `1014584.6544 ResultTypeNumber`},
		{`=RECEIVED(A1,A2,A3,A4,3)`, `1014381.99124 ResultTypeNumber`},
		{`=RECEIVED(A1,A2,A3,A4,4)`, `1014584.6544 ResultTypeNumber`},
	}

	runTests(t, ctx, td)
}

func TestRri(t *testing.T) {
	ss := spreadsheet.New()
	sheet := ss.AddSheet()

	ctx := sheet.FormulaContext()

	td := []testStruct{
		{`=RRI(96,10000,11000)`, `0.00099330737 ResultTypeNumber`},
		{`=RRI(8,10000,11000)`, `0.01198502414 ResultTypeNumber`},
		{`=RRI(0,10000,11000)`, `#NUM! ResultTypeError`},
		{`=RRI(8,0,11000)`, `#NUM! ResultTypeError`},
		{`=RRI(8,10000,-0.000001)`, `#NUM! ResultTypeError`},
	}

	runTests(t, ctx, td)
}

func TestOddlprice(t *testing.T) {
	ss := spreadsheet.New()
	sheet := ss.AddSheet()

	sheet.Cell("A1").SetDate(time.Date(2008, 2, 7, 0, 0, 0, 0, time.UTC))
	sheet.Cell("A2").SetDate(time.Date(2008, 6, 15, 0, 0, 0, 0, time.UTC))
	sheet.Cell("A3").SetDate(time.Date(2007, 10, 15, 0, 0, 0, 0, time.UTC))
	sheet.Cell("A4").SetNumber(0.0375)
	sheet.Cell("A5").SetNumber(0.0405)
	sheet.Cell("A6").SetNumber(100)
	sheet.Cell("A7").SetNumber(2)

	ctx := sheet.FormulaContext()

	td := []testStruct{
		{`=ODDLPRICE(A1,A2,A3,A4,A5,A6,A7,0)`, `99.8782860147 ResultTypeNumber`},
		{`=ODDLPRICE(A1,A2,A3,A4,A5,A6,A7,1)`, `99.8759395207 ResultTypeNumber`},
		{`=ODDLPRICE(A1,A2,A3,A4,A5,A6,A7,2)`, `99.8769016984 ResultTypeNumber`},
		{`=ODDLPRICE(A1,A2,A3,A4,A5,A6,A7,3)`, `99.8787957508 ResultTypeNumber`},
		{`=ODDLPRICE(A1,A2,A3,A4,A5,A6,A7,4)`, `99.8782860147 ResultTypeNumber`},
		{`=ODDLPRICE(A2,A1,A3,A4,A5,A6,A7,4)`, `#NUM! ResultTypeError`},
		{`=ODDLPRICE(A1,A3,A2,A4,A5,A6,A7,4)`, `#NUM! ResultTypeError`},
	}

	runTests(t, ctx, td)
}

func TestOddyield(t *testing.T) {
	ss := spreadsheet.New()
	sheet := ss.AddSheet()

	sheet.Cell("A1").SetDate(time.Date(2008, 4, 20, 0, 0, 0, 0, time.UTC))
	sheet.Cell("A2").SetDate(time.Date(2008, 6, 15, 0, 0, 0, 0, time.UTC))
	sheet.Cell("A3").SetDate(time.Date(2007, 12, 24, 0, 0, 0, 0, time.UTC))
	sheet.Cell("A4").SetNumber(0.0375)
	sheet.Cell("A5").SetNumber(99.875)
	sheet.Cell("A6").SetNumber(100)
	sheet.Cell("A7").SetNumber(2)

	ctx := sheet.FormulaContext()

	td := []testStruct{
		{`=ODDLYIELD(A1,A2,A3,A4,A5,A6,A7,0)`, `0.04519223562 ResultTypeNumber`},
		{`=ODDLYIELD(A1,A2,A3,A4,A5,A6,A7,1)`, `0.04517988549 ResultTypeNumber`},
		{`=ODDLYIELD(A1,A2,A3,A4,A5,A6,A7,2)`, `0.04503841511 ResultTypeNumber`},
		{`=ODDLYIELD(A1,A2,A3,A4,A5,A6,A7,3)`, `0.04515632373 ResultTypeNumber`},
		{`=ODDLYIELD(A1,A2,A3,A4,A5,A6,A7,4)`, `0.04519223562 ResultTypeNumber`},
		{`=ODDLYIELD(A2,A1,A3,A4,A5,A6,A7,4)`, `#NUM! ResultTypeError`},
		{`=ODDLYIELD(A1,A3,A2,A4,A5,A6,A7,4)`, `#NUM! ResultTypeError`},
	}

	runTests(t, ctx, td)
}

func TestPrice(t *testing.T) {
	ss := spreadsheet.New()
	sheet := ss.AddSheet()

	sheet.Cell("A1").SetDate(time.Date(2008, 2, 15, 0, 0, 0, 0, time.UTC))
	sheet.Cell("A2").SetDate(time.Date(2017, 11, 15, 0, 0, 0, 0, time.UTC))
	sheet.Cell("A3").SetNumber(0.0575)
	sheet.Cell("A4").SetNumber(0.065)
	sheet.Cell("A5").SetNumber(100)
	sheet.Cell("A6").SetNumber(2)

	ctx := sheet.FormulaContext()

	td := []testStruct{
		{`=PRICE(A1,A2,A3,A4,A5,A6,0)`, `94.6343616213 ResultTypeNumber`},
		{`=PRICE(A1,A2,A3,A4,A5,A6,1)`, `94.6354492078 ResultTypeNumber`},
		{`=PRICE(A1,A2,A3,A4,A5,A6,2)`, `94.6024171768 ResultTypeNumber`},
		{`=PRICE(A1,A2,A3,A4,A5,A6,3)`, `94.6435945482 ResultTypeNumber`},
		{`=PRICE(A1,A2,A3,A4,A5,A6,4)`, `94.6343616213 ResultTypeNumber`},
	}

	runTests(t, ctx, td)
}

func TestPricemat(t *testing.T) {
	ss := spreadsheet.New()
	sheet := ss.AddSheet()

	sheet.Cell("A1").SetDate(time.Date(2008, 2, 15, 0, 0, 0, 0, time.UTC))
	sheet.Cell("A2").SetDate(time.Date(2008, 4, 13, 0, 0, 0, 0, time.UTC))
	sheet.Cell("A3").SetDate(time.Date(2007, 11, 11, 0, 0, 0, 0, time.UTC))
	sheet.Cell("A4").SetNumber(0.061)
	sheet.Cell("A5").SetNumber(0.061)

	ctx := sheet.FormulaContext()

	td := []testStruct{
		{`=PRICEMAT(A1,A2,A3,A4,A5,0)`, `99.9844988755 ResultTypeNumber`},
		{`=PRICEMAT(A1,A2,A3,A4,A5,1)`, `99.9802978513 ResultTypeNumber`},
		{`=PRICEMAT(A1,A2,A3,A4,A5,2)`, `99.9841690643 ResultTypeNumber`},
		{`=PRICEMAT(A1,A2,A3,A4,A5,3)`, `99.9845977645 ResultTypeNumber`},
		{`=PRICEMAT(A1,A2,A3,A4,A5,4)`, `99.9844988755 ResultTypeNumber`},
	}

	runTests(t, ctx, td)
}

func TestSln(t *testing.T) {
	ss := spreadsheet.New()
	sheet := ss.AddSheet()

	ctx := sheet.FormulaContext()

	td := []testStruct{
		{`=SLN(30000,7500,10)`, `2250 ResultTypeNumber`},
		{`=SLN(30000,7500,0)`, `#DIV/0! ResultTypeError`},
		{`=SLN("hello world",7500,10)`, `#VALUE! ResultTypeError`},
	}

	runTests(t, ctx, td)
}

func TestSyd(t *testing.T) {
	ss := spreadsheet.New()
	sheet := ss.AddSheet()

	ctx := sheet.FormulaContext()

	td := []testStruct{
		{`=SYD(30000,7500,10,1)`, `4090.9090909 ResultTypeNumber`},
		{`=SYD(30000,7500,10,10)`, `409.09090909 ResultTypeNumber`},
		{`=SYD(30000,7500,10,11)`, `#NUM! ResultTypeError`},
		{`=SYD(30000,7500,0,0)`, `#NUM! ResultTypeError`},
		{`=SYD(30000,7500,10,0)`, `#NUM! ResultTypeError`},
		{`=SLN("hello world",7500,10,1)`, `#VALUE! ResultTypeError`},
	}

	runTests(t, ctx, td)
}

func TestTbilleq(t *testing.T) {
	ss := spreadsheet.New()
	sheet := ss.AddSheet()

	sheet.Cell("A1").SetDate(time.Date(2008, 3, 31, 0, 0, 0, 0, time.UTC))
	sheet.Cell("A2").SetDate(time.Date(2008, 6, 1, 0, 0, 0, 0, time.UTC))
	sheet.Cell("A3").SetNumber(0.0914)
	sheet.Cell("A4").SetDate(time.Date(2009, 4, 1, 0, 0, 0, 0, time.UTC))

	ctx := sheet.FormulaContext()

	td := []testStruct{
		{`=TBILLEQ(A1,A2,A3)`, `0.09415149356 ResultTypeNumber`},
		{`=TBILLEQ("A1",A2,A3)`, `#VALUE! ResultTypeError`},
		{`=TBILLEQ(A1,A2,0)`, `#NUM! ResultTypeError`},
		{`=TBILLEQ(A2,A1,A3)`, `#NUM! ResultTypeError`},
		{`=TBILLEQ(A1,A4,A3)`, `#NUM! ResultTypeError`},
	}

	runTests(t, ctx, td)
}

func TestTbillprice(t *testing.T) {
	ss := spreadsheet.New()
	sheet := ss.AddSheet()

	sheet.Cell("A1").SetDate(time.Date(2008, 3, 31, 0, 0, 0, 0, time.UTC))
	sheet.Cell("A2").SetDate(time.Date(2008, 6, 1, 0, 0, 0, 0, time.UTC))
	sheet.Cell("A3").SetNumber(0.09)
	sheet.Cell("A4").SetDate(time.Date(2009, 4, 1, 0, 0, 0, 0, time.UTC))

	ctx := sheet.FormulaContext()

	td := []testStruct{
		{`=TBILLPRICE(A1,A2,A3)`, `98.45 ResultTypeNumber`},
		{`=TBILLPRICE("A1",A2,A3)`, `#VALUE! ResultTypeError`},
		{`=TBILLPRICE(A1,A2,0)`, `#NUM! ResultTypeError`},
		{`=TBILLPRICE(A2,A1,A3)`, `#NUM! ResultTypeError`},
		{`=TBILLPRICE(A1,A4,A3)`, `#NUM! ResultTypeError`},
	}

	runTests(t, ctx, td)
}

func TestTbillyield(t *testing.T) {
	ss := spreadsheet.New()
	sheet := ss.AddSheet()

	sheet.Cell("A1").SetDate(time.Date(2008, 3, 31, 0, 0, 0, 0, time.UTC))
	sheet.Cell("A2").SetDate(time.Date(2008, 6, 1, 0, 0, 0, 0, time.UTC))
	sheet.Cell("A3").SetNumber(98.45)
	sheet.Cell("A4").SetDate(time.Date(2009, 4, 1, 0, 0, 0, 0, time.UTC))

	ctx := sheet.FormulaContext()

	td := []testStruct{
		{`=TBILLYIELD(A1,A2,A3)`, `0.09141696292 ResultTypeNumber`},
		{`=TBILLYIELD("A1",A2,A3)`, `#VALUE! ResultTypeError`},
		{`=TBILLYIELD(A1,A2,0)`, `#NUM! ResultTypeError`},
		{`=TBILLYIELD(A2,A1,A3)`, `#NUM! ResultTypeError`},
		{`=TBILLYIELD(A1,A4,A3)`, `#NUM! ResultTypeError`},
	}

	runTests(t, ctx, td)
}

func TestVdb(t *testing.T) {
	ss := spreadsheet.New()
	sheet := ss.AddSheet()

	sheet.Cell("A1").SetNumber(2400)
	sheet.Cell("A2").SetNumber(300)
	sheet.Cell("A3").SetNumber(10)

	ctx := sheet.FormulaContext()

	td := []testStruct{
		{`=VDB(A1,A2,A3*365,0,1)`, `1.31506849315 ResultTypeNumber`},
		{`=VDB(A1,A2,A3*12,0,1)`, `40 ResultTypeNumber`},
		{`=VDB(A1,A2,A3,0,1)`, `480 ResultTypeNumber`},
		{`=VDB(A1,A2,A3*12,6,18)`, `396.306053264 ResultTypeNumber`},
		{`=VDB(A1,A2,A3*12,6,18,)`, `210 ResultTypeNumber`},
		{`=VDB(A1,A2,A3*12,6,18,,1)`, `0 ResultTypeNumber`},
		{`=VDB(A1,A2,A3*12,6,18,1.5)`, `311.808936658 ResultTypeNumber`},
		{`=VDB(A1,A2,A3,0,0.875,1.5)`, `315 ResultTypeNumber`},
		{`=VDB(A1,A2,A3,0,0.875,,)`, `183.75 ResultTypeNumber`},
		{`=VDB(A1,A2,A3,0,0.875,,1)`, `0 ResultTypeNumber`},
	}

	runTests(t, ctx, td)
}

func TestYielddisc(t *testing.T) {
	ss := spreadsheet.New()
	sheet := ss.AddSheet()

	sheet.Cell("A1").SetDate(time.Date(2008, 2, 16, 0, 0, 0, 0, time.UTC))
	sheet.Cell("A2").SetDate(time.Date(2008, 3, 1, 0, 0, 0, 0, time.UTC))
	sheet.Cell("A3").SetNumber(99.795)
	sheet.Cell("A4").SetNumber(100)

	ctx := sheet.FormulaContext()

	td := []testStruct{
		{`=YIELDDISC(A1,A2,A3,A4)`, `0.04930106718 ResultTypeNumber`},
		{`=YIELDDISC(A1,A2,A3,A4,)`, `0.04930106718 ResultTypeNumber`},
		{`=YIELDDISC(A1,A2,A3,A4,0)`, `0.04930106718 ResultTypeNumber`},
		{`=YIELDDISC(A1,A2,A3,A4,1)`, `0.05370294818 ResultTypeNumber`},
		{`=YIELDDISC(A1,A2,A3,A4,2)`, `0.05282257198 ResultTypeNumber`},
		{`=YIELDDISC(A1,A2,A3,A4,3)`, `0.05355621882 ResultTypeNumber`},
		{`=YIELDDISC(A1,A2,A3,A4,4)`, `0.04930106718 ResultTypeNumber`},
	}

	runTests(t, ctx, td)
}

func TestXirr(t *testing.T) {
	ss := spreadsheet.New()
	sheet := ss.AddSheet()

	sheet.Cell("A1").SetNumber(-10000)
	sheet.Cell("A2").SetNumber(2750)
	sheet.Cell("A3").SetNumber(4250)
	sheet.Cell("A4").SetNumber(3250)
	sheet.Cell("A5").SetNumber(2750)
	sheet.Cell("B1").SetDate(time.Date(2008, 1, 1, 0, 0, 0, 0, time.UTC))
	sheet.Cell("B2").SetDate(time.Date(2008, 3, 1, 0, 0, 0, 0, time.UTC))
	sheet.Cell("B3").SetDate(time.Date(2008, 10, 30, 0, 0, 0, 0, time.UTC))
	sheet.Cell("B4").SetDate(time.Date(2009, 2, 15, 0, 0, 0, 0, time.UTC))
	sheet.Cell("B5").SetDate(time.Date(2009, 4, 1, 0, 0, 0, 0, time.UTC))

	ctx := sheet.FormulaContext()

	td := []testStruct{
		{`=XIRR(A1:A5,B1:B5,0.1)`, `0.37336253351 ResultTypeNumber`},
		{`=XIRR(A1:A5,B1:B5,"hello world")`, `#VALUE! ResultTypeError`},
		{`=XIRR(A2:A5,B2:B5,0.1)`, `#NUM! ResultTypeError`},
	}

	runTests(t, ctx, td)
}

func TestXnpv(t *testing.T) {
	ss := spreadsheet.New()
	sheet := ss.AddSheet()

	sheet.Cell("A1").SetNumber(-10000)
	sheet.Cell("A2").SetNumber(2750)
	sheet.Cell("A3").SetNumber(4250)
	sheet.Cell("A4").SetNumber(3250)
	sheet.Cell("A5").SetNumber(2750)
	sheet.Cell("B1").SetDate(time.Date(2008, 1, 1, 0, 0, 0, 0, time.UTC))
	sheet.Cell("B2").SetDate(time.Date(2008, 3, 1, 0, 0, 0, 0, time.UTC))
	sheet.Cell("B3").SetDate(time.Date(2008, 10, 30, 0, 0, 0, 0, time.UTC))
	sheet.Cell("B4").SetDate(time.Date(2009, 2, 15, 0, 0, 0, 0, time.UTC))
	sheet.Cell("B5").SetDate(time.Date(2009, 4, 1, 0, 0, 0, 0, time.UTC))

	ctx := sheet.FormulaContext()

	td := []testStruct{
		{`=XNPV(0.09,A1:A5,B1:B5)`, `2086.64760203 ResultTypeNumber`},
		{`=XNPV(-0.01,A1:A5,B1:B5)`, `#NUM! ResultTypeError`},
		{`=XNPV("hello world",A1:A5,B1:B5)`, `#VALUE! ResultTypeError`},
	}

	runTests(t, ctx, td)
}

func TestYield(t *testing.T) {
	ss := spreadsheet.New()
	sheet := ss.AddSheet()

	sheet.Cell("A1").SetDate(time.Date(2008, 2, 15, 0, 0, 0, 0, time.UTC))
	sheet.Cell("A2").SetDate(time.Date(2016, 11, 15, 0, 0, 0, 0, time.UTC))
	sheet.Cell("A3").SetNumber(0.0575)
	sheet.Cell("A4").SetNumber(95.04287)
	sheet.Cell("A5").SetNumber(100)
	sheet.Cell("A6").SetNumber(2)

	ctx := sheet.FormulaContext()

	td := []testStruct{
		{`=YIELD(A1,A2,A3,A4,A5,A6)`, `0.06500000688 ResultTypeNumber`},
		{`=YIELD(A1,A2,A3,A4,A5,A6,)`, `0.06500000688 ResultTypeNumber`},
		{`=YIELD(A1,A2,A3,A4,A5,A6,0)`, `0.06500000688 ResultTypeNumber`},
		{`=YIELD(A1,A2,A3,A4,A5,A6,1)`, `0.0650018206 ResultTypeNumber`},
		{`=YIELD(A1,A2,A3,A4,A5,A6,2)`, `0.06495005528 ResultTypeNumber`},
		{`=YIELD(A1,A2,A3,A4,A5,A6,3)`, `0.06501459236 ResultTypeNumber`},
		{`=YIELD(A1,A2,A3,A4,A5,A6,4)`, `0.06500000688 ResultTypeNumber`},
		{`=YIELD(A2,A1,A3,A4,A5,A6,4)`, `#NUM! ResultTypeError`},
		{`=YIELD(A1,A2,A3,A4,A5,A6,5)`, `#NUM! ResultTypeError`},
		{`=YIELD("hello world",A2,A3,A4,A5,A6,4)`, `#VALUE! ResultTypeError`},
	}

	runTests(t, ctx, td)
}

func TestYieldmat(t *testing.T) {
	ss := spreadsheet.New()
	sheet := ss.AddSheet()

	sheet.Cell("A1").SetDate(time.Date(2008, 3, 15, 0, 0, 0, 0, time.UTC))
	sheet.Cell("A2").SetDate(time.Date(2008, 11, 3, 0, 0, 0, 0, time.UTC))
	sheet.Cell("A3").SetDate(time.Date(2007, 11, 8, 0, 0, 0, 0, time.UTC))
	sheet.Cell("A4").SetNumber(0.0625)
	sheet.Cell("A5").SetNumber(100.0123)

	ctx := sheet.FormulaContext()

	td := []testStruct{
		{`=YIELDMAT(A1,A2,A3,A4,A5)`, `0.06095433369 ResultTypeNumber`},
		{`=YIELDMAT(A1,A2,A3,A4,A5,)`, `0.06095433369 ResultTypeNumber`},
		{`=YIELDMAT(A1,A2,A3,A4,A5,0)`, `0.06095433369 ResultTypeNumber`},
		{`=YIELDMAT(A1,A2,A3,A4,A5,1)`, `0.06096668564 ResultTypeNumber`},
		{`=YIELDMAT(A1,A2,A3,A4,A5,2)`, `0.06094805915 ResultTypeNumber`},
		{`=YIELDMAT(A1,A2,A3,A4,A5,3)`, `0.06096362992 ResultTypeNumber`},
		{`=YIELDMAT(A1,A2,A3,A4,A5,4)`, `0.06095433369 ResultTypeNumber`},
		{`=YIELDMAT(A2,A1,A3,A4,A5,4)`, `#NUM! ResultTypeError`},
		{`=YIELDMAT(A1,A2,A3,A4,A5,5)`, `#NUM! ResultTypeError`},
		{`=YIELDMAT("hello world",A2,A3,A4,A5,4)`, `#VALUE! ResultTypeError`},
	}

	runTests(t, ctx, td)
}

func TestMid(t *testing.T) {
	ss := spreadsheet.New()
	sheet := ss.AddSheet()

	ctx := sheet.FormulaContext()

	td := []testStruct{
		{`=MID("Fluid Flow",1,5)`, `Fluid ResultTypeString`},
		{`=MID("Fluid Flow",7,20)`, `Flow ResultTypeString`},
		{`=MID("Fluid Flow",20,5)`, ` ResultTypeString`},
		{`=MID("Fluid Flow",1,0)`, ` ResultTypeString`},
		{`=MID("Fluid Flow",0,5)`, `#VALUE! ResultTypeError`},
		{`=MID("Fluid Flow",7,-20)`, `#VALUE! ResultTypeError`},
	}

	runTests(t, ctx, td)
}

func TestSubstitute(t *testing.T) {
	ss := spreadsheet.New()
	sheet := ss.AddSheet()

	sheet.Cell("A1").SetString("Hello Earth Earth Earth")

	ctx := sheet.FormulaContext()

	td := []testStruct{
		{`=SUBSTITUTE(A1,"Earth","Krypton",1)`, `Hello Krypton Earth Earth ResultTypeString`},
		{`=SUBSTITUTE(A1,"Earth","Krypton",2)`, `Hello Earth Krypton Earth ResultTypeString`},
		{`=SUBSTITUTE(A1,"Earth","Krypton",3)`, `Hello Earth Earth Krypton ResultTypeString`},
		{`=SUBSTITUTE(A1,"Earth","Krypton",4)`, `Hello Earth Earth Earth ResultTypeString`},
		{`=SUBSTITUTE(A1,"Earth","Krypton")`, `Hello Krypton Krypton Krypton ResultTypeString`},
		{`=SUBSTITUTE(A1,"World","Krypton")`, `Hello Earth Earth Earth ResultTypeString`},
		{`=SUBSTITUTE(A1,"Earth","Krypton",0)`, `#VALUE! ResultTypeError`},
	}

	runTests(t, ctx, td)
}

func TestAnd(t *testing.T) {
	ss := spreadsheet.New()
	sheet := ss.AddSheet()

	ctx := sheet.FormulaContext()

	td := []testStruct{
		{`=AND(FALSE,FALSE)`, `0 ResultTypeNumber`},
		{`=AND(TRUE,FALSE)`, `0 ResultTypeNumber`},
		{`=AND(FALSE,TRUE)`, `0 ResultTypeNumber`},
		{`=AND(TRUE,TRUE)`, `1 ResultTypeNumber`},
	}

	runTests(t, ctx, td)
}

func TestIferror(t *testing.T) {
	ss := spreadsheet.New()
	sheet := ss.AddSheet()

	ctx := sheet.FormulaContext()

	td := []testStruct{
		{`=IFERROR("No error","ERROR")`, `No error ResultTypeString`},
		{`=IFERROR(1/0,"ERROR")`, `ERROR ResultTypeString`},
	}

	runTests(t, ctx, td)
}

func TestChar(t *testing.T) {
	ss := spreadsheet.New()
	sheet := ss.AddSheet()

	ctx := sheet.FormulaContext()

	td := []testStruct{
		{`=CHAR(65)`, `A ResultTypeString`},
		{`=CHAR(255)`, `ÿ ResultTypeString`},
		{`=CHAR(1000)`, `#VALUE! ResultTypeError`},
		{`=CHAR("invalid")`, `#VALUE! ResultTypeError`},
	}

	runTests(t, ctx, td)
}

func TestRound(t *testing.T) {
	ss := spreadsheet.New()
	sheet := ss.AddSheet()

	ctx := sheet.FormulaContext()

	td := []testStruct{
		{`=ROUND(2.14,1)`, `2.1 ResultTypeNumber`},
		{`=ROUND(2.16,1)`, `2.2 ResultTypeNumber`},
		{`=ROUND(-2.14,1)`, `-2.1 ResultTypeNumber`},
		{`=ROUND(-2.16,1)`, `-2.2 ResultTypeNumber`},
		{`=ROUND(21.5,-1)`, `20 ResultTypeNumber`},
		{`=ROUND(21.5,-2)`, `0 ResultTypeNumber`},
		{`=ROUND(-55.5,-1)`, `-60 ResultTypeNumber`},
		{`=ROUND(-55.5,-2)`, `-100 ResultTypeNumber`},
		{`=ROUND(-55.5,0)`, `-56 ResultTypeNumber`},
		{`=ROUND(-55.4,)`, `-55 ResultTypeNumber`},
		{`=ROUND(-55.4)`, `#VALUE! ResultTypeError`},
	}

	runTests(t, ctx, td)
}
