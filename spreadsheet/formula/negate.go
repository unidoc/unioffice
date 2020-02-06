// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased on https://unidoc.io.

package formula

// Negate is a negate expression like -A1.
type Negate struct {
	e Expression
}

// NewNegate constructs a new negate expression.
func NewNegate(e Expression) Expression {
	return Negate{e}
}

// Eval evaluates and returns the result of a Negate expression.
func (n Negate) Eval(ctx Context, ev Evaluator) Result {
	r := n.e.Eval(ctx, ev)
	if r.Type == ResultTypeNumber {
		return MakeNumberResult(-r.ValueNumber)
	}
	return MakeErrorResult("NEGATE expected number argument")
}

// Reference returns an invalid reference for Negate.
func (n Negate) Reference(ctx Context, ev Evaluator) Reference {
	return ReferenceInvalid
}

// ToString returns a string representation for Negate.
func (n Negate) ToString() string {
	return "-" + n.e.ToString()
}

// MoveLeft makes the Negate moved left after removing a column.
func (n Negate) MoveLeft(q *MoveQuery) Expression {
	return Negate{n.e.MoveLeft(q)}
}
