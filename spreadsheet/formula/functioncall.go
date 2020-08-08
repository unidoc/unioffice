// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io.

package formula

import (
	"bytes"

	"github.com/unidoc/unioffice/spreadsheet/update"
)

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

// String returns a string representation of FunctionCall expression.
func (f FunctionCall) String() string {
	buf := bytes.Buffer{}
	buf.WriteString(f.name)
	buf.WriteString("(")
	lastArgIndex := len(f.args) - 1
	for argIndex, arg := range f.args {
		buf.WriteString(arg.String())
		if argIndex != lastArgIndex {
			buf.WriteString(",")
		}
	}
	buf.WriteString(")")
	return buf.String()
}

// Update updates the FunctionCall references after removing a row/column.
func (f FunctionCall) Update(q *update.UpdateQuery) Expression {
	newArgs := []Expression{}
	for _, arg := range f.args {
		newArg := arg.Update(q)
		newArgs = append(newArgs, newArg)
	}
	return FunctionCall{
		name: f.name,
		args: newArgs,
	}
}
