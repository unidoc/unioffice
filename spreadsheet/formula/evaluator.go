// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package formula

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

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
	isRef bool
	booleans []bool
}

func (d *defEval) Eval(ctx Context, formula string) Result {
	expr := ParseString(formula)
	if expr != nil {
		d.addInfo(ctx, expr)
		return expr.Eval(ctx, d)
	}
	return MakeErrorResult(fmt.Sprintf("unable to parse formula %s", formula))
}

//addInfo adds information which is needed for some functions but is lost after evaluation. E.g. which zeroes and ones were actually booleans before evaluation and which arguments are actually references.
func (d *defEval) addInfo(ctx Context, expr Expression) {
	switch expr.(type) {
	case FunctionCall:
		switch expr.(FunctionCall).name {
		case "ISREF":
			for _, arg := range expr.(FunctionCall).args {
				switch arg.(type) {
				case CellRef:
					d.isRef = validateRef(arg.(CellRef))
					return
				case Range:
					switch arg.(Range).from.(type) {
					case CellRef:
						d.isRef = validateRef(arg.(Range).from.(CellRef))
						return
					}
					switch arg.(Range).to.(type) {
					case CellRef:
						d.isRef = validateRef(arg.(Range).to.(CellRef))
						return
					}
				}
			}
		case "CONCAT", "_xlfn.CONCAT", "CONCATENATE":
			d.booleans = []bool{}
			for _, arg := range expr.(FunctionCall).args {
				switch arg.(type) {
					case CellRef:
						cr := arg.(CellRef).s
						d.booleans = append(d.booleans, ctx.IsBool(cr))
				}
			}
		}
	}
}

var refRegexp *regexp.Regexp = regexp.MustCompile(`^([a-z]+)([0-9]+)$`)

func validateRef(cr CellRef) bool {
	if submatch := refRegexp.FindStringSubmatch(strings.ToLower(cr.s)); len(submatch) > 2 {
		col := submatch[1]
		row, _ := strconv.Atoi(submatch[2])
		return row <= 1048576 && col <= "zz"
	}
	return false
}
