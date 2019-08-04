// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package formula

import "strings"

type String struct {
	s string
}

func NewString(v string) Expression {
	// Excel escapes quotes within a string by repeating them
	v = strings.Replace(v, `""`, `"`, -1)
	return String{v}
}

func (s String) Eval(ctx Context, ev Evaluator) Result {
	return MakeStringResult(s.s)
}

func (s String) Reference(ctx Context, ev Evaluator) Reference {
	return ReferenceInvalid
}
