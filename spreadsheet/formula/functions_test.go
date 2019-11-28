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
	Input string
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

	td := []testStruct{
		{`=ISLOGICAL(A1)`, `1 ResultTypeNumber`},
		{`=ISLOGICAL(A2)`, `1 ResultTypeNumber`},
		{`=ISLOGICAL(A3)`, `0 ResultTypeNumber`},
		{`=ISLOGICAL(A4)`, `0 ResultTypeNumber`},
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
	sheet.Cell("A6").SetString("A1")

	td := []testStruct{
		{`A1`, `1 ResultTypeNumber`},
		{`A2`, `1 ResultTypeNumber`},
		{`A3`, `0 ResultTypeNumber`},
		{`A4`, `0 ResultTypeNumber`},
		{`A5`, `1 ResultTypeNumber`},
		{`A6`, `A1 ResultTypeString`},
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
