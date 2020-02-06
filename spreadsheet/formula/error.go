// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased on https://unidoc.io.

package formula

// Error is an error expression.
type Error struct {
	s string
}

// NewError constructs a new error expression from a string.
func NewError(v string) Expression {
	return Error{v}
}

// Eval evaluates and returns the result of an error expression.
func (e Error) Eval(ctx Context, ev Evaluator) Result {
	return MakeErrorResult(e.s)
}

// Reference returns an invalid reference for Error.
func (e Error) Reference(ctx Context, ev Evaluator) Reference {
	return ReferenceInvalid
}

// ToString returns an empty string for Error.
func (e Error) ToString() string {
	return ""
}

// MoveLeft returns the same object as moving ranges to left does not affect Error.
func (e Error) MoveLeft(q *MoveQuery) Expression {
	return e
}
