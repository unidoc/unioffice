// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io.

package formula

import "github.com/unidoc/unioffice/spreadsheet/update"

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

// String returns a string representation for Negate.
func (n Negate) String() string {
	return "-" + n.e.String()
}

// Update updates references in the Negate after removing a row/column.
func (n Negate) Update(q *update.UpdateQuery) Expression {
	return Negate{n.e.Update(q)}
}
