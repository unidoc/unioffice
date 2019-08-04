// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package formula

type Negate struct {
	e Expression
}

func NewNegate(e Expression) Expression {
	return Negate{e}
}

func (n Negate) Eval(ctx Context, ev Evaluator) Result {
	r := n.e.Eval(ctx, ev)
	if r.Type == ResultTypeNumber {
		return MakeNumberResult(-r.ValueNumber)
	}
	return MakeErrorResult("NEGATE expected number argument")
}

func (n Negate) Reference(ctx Context, ev Evaluator) Reference {
	return ReferenceInvalid
}
