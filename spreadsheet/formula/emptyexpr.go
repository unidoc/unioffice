// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased on https://unidoc.io.

package formula

import "github.com/unidoc/unioffice/spreadsheet/update"

// EmptyExpr is an empty expression.
type EmptyExpr struct {
}

// NewEmptyExpr constructs a new empty expression.
func NewEmptyExpr() Expression {
	return EmptyExpr{}
}

// Eval evaluates and returns the result of an empty expression.
func (e EmptyExpr) Eval(ctx Context, ev Evaluator) Result {
	return MakeEmptyResult()
}

// Reference returns an invalid reference for EmptyExpr.
func (e EmptyExpr) Reference(ctx Context, ev Evaluator) Reference {
	return ReferenceInvalid
}

// String returns an empty string for EmptyExpr.
func (e EmptyExpr) String() string {
	return ""
}

// Update returns the same object as updating sheet references does not affect EmptyExpr.
func (e EmptyExpr) Update(q *update.UpdateQuery) Expression {
	return e
}
