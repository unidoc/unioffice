// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased on https://unidoc.io.

package formula

import "fmt"

// Range is a range expression that when evaluated returns a list of Results.
type PrefixRangeExpr struct {
	pfx, from, to Expression
}

// NewRange constructs a new range.
func NewPrefixRangeExpr(pfx, from, to Expression) Expression {
	return PrefixRangeExpr{pfx, from, to}
}

func (p PrefixRangeExpr) Eval(ctx Context, ev Evaluator) Result {
	pfx := p.pfx.Reference(ctx, ev)
	from := p.from.Reference(ctx, ev)
	to := p.to.Reference(ctx, ev)
	switch pfx.Type {
	case ReferenceTypeSheet:
		if from.Type == ReferenceTypeCell && to.Type == ReferenceTypeCell {
			return resultFromCellRange(ctx.Sheet(pfx.Value), ev, from.Value, to.Value)
		}
		return MakeErrorResult("invalid range " + from.Value + " to " + to.Value)
	default:
		return MakeErrorResult(fmt.Sprintf("no support for reference type %s", pfx.Type))
	}
}

func (p PrefixRangeExpr) Reference(ctx Context, ev Evaluator) Reference {
	pfx := p.pfx.Reference(ctx, ev)
	from := p.from.Reference(ctx, ev)
	to := p.to.Reference(ctx, ev)
	if pfx.Type == ReferenceTypeSheet && from.Type == ReferenceTypeCell && to.Type == ReferenceTypeCell {
		return MakeRangeReference(pfx.Value + "!" + from.Value + ":" + to.Value)
	}
	return ReferenceInvalid
}
