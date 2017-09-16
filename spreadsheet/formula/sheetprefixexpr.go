// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package formula

type SheetPrefixExpr struct {
	sheet string
}

func NewSheetPrefixExpr(s string) Expression {
	return &SheetPrefixExpr{s}
}

func (s SheetPrefixExpr) Eval(ctx Context, ev Evaluator) Result {
	return MakeErrorResult("sheet prefix should never be evaluated")
}

func (s SheetPrefixExpr) Reference(ctx Context, ev Evaluator) Reference {
	return Reference{Type: ReferenceTypeSheet, Value: s.sheet}
}
