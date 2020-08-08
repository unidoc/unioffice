// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io.

package formula

import "fmt"

// defEval is the default formula evaluator which implements the Evaluator interface.
type defEval struct {
	evCache
	lastEvalIsRef bool
}

// Eval evaluates and returns the result of a formula.
func (d *defEval) Eval(ctx Context, formula string) Result {
	expr := ParseString(formula)
	if expr != nil {
		d.checkLastEvalIsRef(ctx, expr)
		result := expr.Eval(ctx, d)
		return result
	}
	return MakeErrorResult(fmt.Sprintf("unable to parse formula %s", formula))
}

// LastEvalIsRef returns if last evaluation with the evaluator was a reference.
func (d *defEval) LastEvalIsRef() bool {
	return d.lastEvalIsRef
}

// checkLastEvalIsRef adds information which is needed for some functions but is lost after evaluation. E.g. which arguments are actually references.
func (d *defEval) checkLastEvalIsRef(ctx Context, expr Expression) {
	switch expr.(type) {
	case FunctionCall:
		switch expr.(FunctionCall).name {
		case "ISREF":
			for _, arg := range expr.(FunctionCall).args {
				switch arg.(type) {
				case CellRef, Range, HorizontalRange, VerticalRange, NamedRangeRef, PrefixExpr, PrefixRangeExpr, PrefixHorizontalRange, PrefixVerticalRange:
					evResult := arg.Eval(ctx, d)
					d.lastEvalIsRef = !(evResult.Type == ResultTypeError && evResult.ValueString == "#NAME?")
				default:
					d.lastEvalIsRef = false
				}
			}
		}
	}
}
