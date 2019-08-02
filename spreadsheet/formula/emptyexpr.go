// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package formula

type EmptyExpr struct {
}

func NewEmptyExpr() Expression {
	return EmptyExpr{}
}

func (e EmptyExpr) Eval(ctx Context, ev Evaluator) Result {
	return MakeEmptyResult()
}

func (e EmptyExpr) Reference(ctx Context, ev Evaluator) Reference {
	return ReferenceInvalid
}
