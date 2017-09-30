// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package formula

import (
	"strconv"

	"baliance.com/gooxml"
)

type Bool struct {
	b bool
}

func NewBool(v string) Expression {
	b, err := strconv.ParseBool(v)
	if err != nil {
		gooxml.Log("error parsing formula bool %s: %s", v, err)
	}
	return Bool{b}
}

func (b Bool) Eval(ctx Context, ev Evaluator) Result {
	return MakeBoolResult(b.b)
}

func (b Bool) Reference(ctx Context, ev Evaluator) Reference {
	return ReferenceInvalid
}
