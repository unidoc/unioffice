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

type Number struct {
	v float64
}

func NewNumber(v string) Expression {
	f, err := strconv.ParseFloat(v, 64)
	if err != nil {
		gooxml.Log("error parsing formula number %s: %s", v, err)
	}
	return Number{f}
}

func (n Number) Eval(ctx Context, ev Evaluator) Result {
	return MakeNumberResult(n.v)
}

func (n Number) Reference(ctx Context, ev Evaluator) Reference {
	return ReferenceInvalid
}
