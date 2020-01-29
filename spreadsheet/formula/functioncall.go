// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased on https://unidoc.io.

package formula

// FunctionCall is a function call expression.
type FunctionCall struct {
	name string
	args []Expression
}

// NewFunction constructs a new function call expression.
func NewFunction(name string, args []Expression) Expression {
	return FunctionCall{name, args}
}

// Eval evaluates and returns the result of a function call.
func (f FunctionCall) Eval(ctx Context, ev Evaluator) Result {
	fn := LookupFunction(f.name)
	if fn != nil {
		args := make([]Result, len(f.args))
		for i, a := range f.args {
			args[i] = a.Eval(ctx, ev)
			args[i].Ref = a.Reference(ctx, ev)
		}
		return fn(args)
	}
	fnx := LookupFunctionComplex(f.name)
	if fnx != nil {
		args := make([]Result, len(f.args))
		for i, a := range f.args {
			args[i] = a.Eval(ctx, ev)
			args[i].Ref = a.Reference(ctx, ev)
		}
		return fnx(ctx, ev, args)
	}

	return MakeErrorResult("unknown function " + f.name)
}

// Reference returns an invalid reference for FunctionCall.
func (f FunctionCall) Reference(ctx Context, ev Evaluator) Reference {
	return ReferenceInvalid
}
