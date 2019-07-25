// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package formula

type ConstArrayExpr struct {
	data [][]Expression
}

func NewConstArrayExpr(data [][]Expression) Expression {
	return &ConstArrayExpr{data}
}

func (c ConstArrayExpr) Eval(ctx Context, ev Evaluator) Result {
	res := [][]Result{}
	for _, row := range c.data {
		r := []Result{}
		for _, col := range row {
			r = append(r, col.Eval(ctx, ev))
		}
		res = append(res, r)
	}
	return MakeArrayResult(res)
}

func (c ConstArrayExpr) Reference(ctx Context, ev Evaluator) Reference {
	return ReferenceInvalid
}
