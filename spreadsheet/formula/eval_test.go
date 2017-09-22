// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package formula_test

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"testing"

	"baliance.com/gooxml/spreadsheet"
	"baliance.com/gooxml/spreadsheet/formula"
)

func TestEval(t *testing.T) {
	td := []struct {
		Inp string
		Exp string
	}{
		{"TRUE", "1 ResultTypeNumber"},
		{"=FALSE", "0 ResultTypeNumber"},
		{"=1+2", "3 ResultTypeNumber"},
		{"=1+2+3", "6 ResultTypeNumber"},
		{"=1+2-3", "0 ResultTypeNumber"},
		{"=2*5", "10 ResultTypeNumber"},
		{"=2*5-3", "7 ResultTypeNumber"},
		{"=2*5+3*6-4", "24 ResultTypeNumber"},
		{"=2^10", "1024 ResultTypeNumber"},
		{"=1=1", "1 ResultTypeNumber"},
		{"=1<>1", "0 ResultTypeNumber"},
		{"=1<=5", "1 ResultTypeNumber"},
		{"=1>=5", "0 ResultTypeNumber"},
		{"=5*1>=5", "1 ResultTypeNumber"},
		{"=5*1>=5+1", "0 ResultTypeNumber"},
		{"=A1", "1.23 ResultTypeNumber"},
		{"A1", "1.23 ResultTypeNumber"},
		{"=A1+A1", "2.46 ResultTypeNumber"},
		{"=A1+B2", "3.23 ResultTypeNumber"},
		{"=SUM(1,2,3,4,5)", "15 ResultTypeNumber"},
		{"SUM(-2,-3,2,3,4)", "4 ResultTypeNumber"},
		{"SUM(B1:B3)", "6 ResultTypeNumber"},
		{"TRUE()", "1 ResultTypeNumber"},
		{`"test"`, "test ResultTypeString"},
		{`"te""st"`, `te"st ResultTypeString`},
	}

	ss := spreadsheet.New()
	sheet := ss.AddSheet()
	sheet.Cell("A1").SetNumber(1.23)
	sheet.Cell("B1").SetNumber(1)
	sheet.Cell("B2").SetNumber(2)
	sheet.Cell("B3").SetNumber(3)
	ctx := sheet.FormulaContext()

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

func TestReferenceSheet(t *testing.T) {
	testSheet("formulareference.xlsx", t)
}
func TestMacExcelSheet(t *testing.T) {
	testSheet("MacExcel365.xlsx", t)
}

func testSheet(fn string, t *testing.T) {
	// TODO: uncomment once we quit building on 1.8
	//t.Helper()
	wb, err := spreadsheet.Open("testdata/" + fn)
	if err != nil {
		t.Fatalf("error opening reference sheet: %s", err)
	}

	formulaCount := 0
	for _, sheet := range wb.Sheets() {
		for _, row := range sheet.Rows() {
			for _, cell := range row.Cells() {
				// the value should have a computed value for the formula that
				// Excel has cached
				cachedValue := cell.GetCachedFormulaResult()
				if cell.HasFormula() {
					cellFormula := formula.ParseString(cell.GetFormula())
					if cellFormula == nil {
						t.Errorf("error parsing formula %s", cell.GetFormula())
						continue
					}

					// so evaluating the formula in the context of the sheet,
					// should return the same results that Excel computed
					result := cellFormula.Eval(sheet.FormulaContext(), formula.NewEvaluator())
					if got := result.Value(); !cmpValue(got, cachedValue) {
						t.Errorf("expected '%s', got '%s' for %s cell %s (%s) %s", cachedValue, got, sheet.Name(), cell.Reference(), cell.GetFormula(), result.ErrorMessage)
					} else {
						formulaCount++
					}

				}
			}
		}
	}
	t.Logf("evaluated %d formulas from %s sheet", formulaCount, fn)
}

func cmpValue(l, r string) bool {
	if l == r {
		return true
	}
	lf, el := strconv.ParseFloat(l, 64)
	rf, er := strconv.ParseFloat(r, 64)
	if el == nil && er == nil {
		if math.Abs(lf-rf) < 1e-7 {
			return true
		}
	}
	return false
}

func TestArrayFormula(t *testing.T) {
	wb := spreadsheet.New()
	sheet := wb.AddSheet()
	_ = sheet
	sheet.Cell("A1").SetNumber(1)
	sheet.Cell("A2").SetNumber(2)
	sheet.Cell("A3").SetNumber(3)
	sheet.Cell("A4").SetNumber(4)

	sheet.Cell("B1").SetNumber(5)
	sheet.Cell("B2").SetNumber(6)
	sheet.Cell("B3").SetNumber(7)
	sheet.Cell("B4").SetNumber(8)

	// this tests the expansion of array results into surrounding cells
	sheet.Cell("C1").SetFormulaArray("TRANSPOSE(A1:B4)")
	sheet.RecalculateFormulas()

	if got := sheet.Cell("C1").GetFormattedValue(); got != "1" {
		t.Errorf("expected 1 in C1, got %s", got)
	}
	if got := sheet.Cell("D1").GetFormattedValue(); got != "2" {
		t.Errorf("expected 2 in D1, got %s", got)
	}
	if got := sheet.Cell("E1").GetFormattedValue(); got != "3" {
		t.Errorf("expected 3 in E1, got %s", got)
	}
	if got := sheet.Cell("F1").GetFormattedValue(); got != "4" {
		t.Errorf("expected 4 in F1, got %s", got)
	}

	if got := sheet.Cell("C2").GetFormattedValue(); got != "5" {
		t.Errorf("expected 5 in C2, got %s", got)
	}
	if got := sheet.Cell("D2").GetFormattedValue(); got != "6" {
		t.Errorf("expected 6 in D2, got %s", got)
	}
	if got := sheet.Cell("E2").GetFormattedValue(); got != "7" {
		t.Errorf("expected 7 in E2, got %s", got)
	}
	if got := sheet.Cell("F2").GetFormattedValue(); got != "8" {
		t.Errorf("expected 8 in F2, got %s", got)
	}

}
