// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package formula_test

import (
	"fmt"
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
	wb, err := spreadsheet.Open("testdata/formulareference.xlsx")
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
					if got := result.Value(); got != cachedValue {
						t.Errorf("expected '%s', got '%s' for %s cell %s (%s) %s", cachedValue, got, sheet.Name(), cell.Reference(), cell.GetFormula(), result.ErrorMessage)
					} else {
						formulaCount++
					}

				}
			}
		}
	}
	t.Logf("evaluated %d formulas from reference sheet", formulaCount)
}
