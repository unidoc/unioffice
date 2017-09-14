// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package formula

import "fmt"

// Evaluator is the interface for a formula evaluator.  This is needed so we can
// pass it to the spreadsheet to let it evaluate formula cells before returning
// the results.
type Evaluator interface {
	Eval(ctx Context, formula string) Result
}

func NewEvaluator() Evaluator {
	return &defEval{}
}

type defEval struct {
}

func (d *defEval) Eval(ctx Context, formula string) Result {
	expr := ParseString(formula)
	if expr != nil {
		return expr.Eval(ctx, d)
	}
	return MakeErrorResult(fmt.Sprintf("unable to parse formula %s", formula))
}
