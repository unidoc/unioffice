// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased on https://unidoc.io.

package formula

import "fmt"

// PrefixExpr is an expression containing reference to another sheet like Sheet1!A1 (the value of the cell A1 from sheet 'Sheet1').
type PrefixExpr struct {
	pfx Expression
	exp Expression
}

// NewPrefixExpr constructs an expression with prefix.
func NewPrefixExpr(pfx, exp Expression) Expression {
	return &PrefixExpr{pfx, exp}
}

// Eval evaluates and returns an expression with prefix.
func (p PrefixExpr) Eval(ctx Context, ev Evaluator) Result {
	ref := p.pfx.Reference(ctx, ev)
	switch ref.Type {
	case ReferenceTypeSheet:
		sheetCtx := ctx.Sheet(ref.Value)
		return p.exp.Eval(sheetCtx, ev)
	default:
		return MakeErrorResult(fmt.Sprintf("no support for reference type %s", ref.Type))
	}
}

// Reference returns a string reference value to an expression with prefix.
func (p PrefixExpr) Reference(ctx Context, ev Evaluator) Reference {
	pfx := p.pfx.Reference(ctx, ev)
	exp := p.exp.Reference(ctx, ev)
	if pfx.Type == ReferenceTypeSheet && exp.Type == ReferenceTypeCell {
		return Reference{Type: ReferenceTypeCell, Value: pfx.Value + "!" + exp.Value}
	}
	return ReferenceInvalid
}

// ToString returns a string representation of PrefixExpr.
func (p PrefixExpr) ToString() string {
	return fmt.Sprintf("%s!%s", p.pfx.ToString(), p.exp.ToString())
}

// MoveLeft makes the PrefixExpr moved left after removing a column.
func (p PrefixExpr) MoveLeft(q *MoveQuery) Expression {
	new := p
	sheetName := p.pfx.ToString()
	if sheetName == q.SheetToMove {
		newQ := *q
		newQ.MoveCurrentSheet = true
		new.exp = p.exp.MoveLeft(&newQ)
	}
	return new
}
