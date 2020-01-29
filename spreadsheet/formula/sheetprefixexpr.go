// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased on https://unidoc.io.

package formula

// SheetPrefixExpr is a reference to a sheet like Sheet1! (reference to sheet 'Sheet1').
type SheetPrefixExpr struct {
	sheet string
}

// NewSheetPrefixExpr constructs a new prefix expression.
func NewSheetPrefixExpr(s string) Expression {
	return &SheetPrefixExpr{s}
}

// Eval evaluates and returns the result of a sheet expression.
func (s SheetPrefixExpr) Eval(ctx Context, ev Evaluator) Result {
	return MakeErrorResult("sheet prefix should never be evaluated")
}

// Reference returns a string reference value to a sheet.
func (s SheetPrefixExpr) Reference(ctx Context, ev Evaluator) Reference {
	return Reference{Type: ReferenceTypeSheet, Value: s.sheet}
}
