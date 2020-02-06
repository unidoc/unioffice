// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased on https://unidoc.io.

package formula

import (
	"strconv"

	"github.com/unidoc/unioffice"
)

// Number is a nubmer expression.
type Number struct {
	v float64
}

// NewNumber constructs a new number expression.
func NewNumber(v string) Expression {
	f, err := strconv.ParseFloat(v, 64)
	if err != nil {
		unioffice.Log("error parsing formula number %s: %s", v, err)
	}
	return Number{f}
}

// Eval evaluates and returns a number.
func (n Number) Eval(ctx Context, ev Evaluator) Result {
	return MakeNumberResult(n.v)
}

// Reference returns an invalid reference for Number.
func (n Number) Reference(ctx Context, ev Evaluator) Reference {
	return ReferenceInvalid
}

// ToString returns a string representation of Number.
func (n Number) ToString() string {
	return strconv.FormatFloat(n.v, 'f', -1, 64)
}

// MoveLeft returns the same object as moving ranges to left does not affect Number.
func (n Number) MoveLeft(q *MoveQuery) Expression {
	return n
}
