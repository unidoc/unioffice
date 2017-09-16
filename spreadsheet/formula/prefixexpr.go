// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package formula

import "fmt"

type PrefixExpr struct {
	pfx Expression
	exp Expression
}

func NewPrefixExpr(pfx, exp Expression) Expression {
	return &PrefixExpr{pfx, exp}
}

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

func (p PrefixExpr) Reference(ctx Context, ev Evaluator) Reference {
	return ReferenceInvalid
}
