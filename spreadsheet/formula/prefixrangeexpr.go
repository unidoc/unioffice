// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io.

package formula

import (
	"fmt"

	"github.com/unidoc/unioffice/spreadsheet/update"
)

// PrefixRangeExpr is a range expression that when evaluated returns a list of Results from a given sheet like Sheet1!A1:B4 (all cells from A1 to B4 from a sheet 'Sheet1').
type PrefixRangeExpr struct {
	pfx, from, to Expression
}

// NewPrefixRangeExpr constructs a new range with prefix.
func NewPrefixRangeExpr(pfx, from, to Expression) Expression {
	return PrefixRangeExpr{pfx, from, to}
}

// Eval evaluates a range with prefix returning a list of results or an error.
func (p PrefixRangeExpr) Eval(ctx Context, ev Evaluator) Result {
	pfx := p.pfx.Reference(ctx, ev)
	from := p.from.Reference(ctx, ev)
	to := p.to.Reference(ctx, ev)
	switch pfx.Type {
	case ReferenceTypeSheet:
		ref := prefixRangeReference(pfx, from, to)
		if from.Type == ReferenceTypeCell && to.Type == ReferenceTypeCell {
			if cached, found := ev.GetFromCache(ref); found {
				return cached
			} else {
				result := resultFromCellRange(ctx.Sheet(pfx.Value), ev, from.Value, to.Value)
				ev.SetCache(ref, result)
				return result
			}
		}
		return MakeErrorResult("invalid range " + ref)
	default:
		return MakeErrorResult(fmt.Sprintf("no support for reference type %s", pfx.Type))
	}
}

func prefixRangeReference(pfx, from, to Reference) string {
	return fmt.Sprintf("%s!%s:%s", pfx.Value, from.Value, to.Value)
}

// Reference returns a string reference value to a range with prefix.
func (p PrefixRangeExpr) Reference(ctx Context, ev Evaluator) Reference {
	pfx := p.pfx.Reference(ctx, ev)
	from := p.from.Reference(ctx, ev)
	to := p.to.Reference(ctx, ev)
	if pfx.Type == ReferenceTypeSheet && from.Type == ReferenceTypeCell && to.Type == ReferenceTypeCell {
		return MakeRangeReference(prefixRangeReference(pfx, from, to))
	}
	return ReferenceInvalid
}

// String returns a string representation of a range with prefix.
func (r PrefixRangeExpr) String() string {
	return fmt.Sprintf("%s!%s:%s", r.pfx.String(), r.from.String(), r.to.String())
}

// Update updates references in the PrefixRangeExpr after removing a row/column.
func (r PrefixRangeExpr) Update(q *update.UpdateQuery) Expression {
	new := r
	sheetName := r.pfx.String()
	if sheetName == q.SheetToUpdate {
		newQ := *q
		newQ.UpdateCurrentSheet = true
		new.from = r.from.Update(&newQ)
		new.to = r.to.Update(&newQ)
	}
	return new
}
