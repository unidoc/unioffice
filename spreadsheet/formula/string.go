// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased on https://unidoc.io.

package formula

import (
	"strings"

	"github.com/unidoc/unioffice/spreadsheet/update"
)

// String is a string expression.
type String struct {
	s string
}

// NewString constructs a new string expression.
func NewString(v string) Expression {
	// Excel escapes quotes within a string by repeating them
	v = strings.Replace(v, `""`, `"`, -1)
	return String{v}
}

// Eval evaluates and returns a string.
func (s String) Eval(ctx Context, ev Evaluator) Result {
	return MakeStringResult(s.s)
}

// Reference returns an invalid reference for String.
func (s String) Reference(ctx Context, ev Evaluator) Reference {
	return ReferenceInvalid
}

// String returns a string representation of String.
func (s String) String() string {
	return `"` + s.s + `"`
}

// Update returns the same object as updating sheet references does not affect String.
func (s String) Update(q *update.UpdateQuery) Expression {
	return s
}
