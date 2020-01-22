// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased on https://unidoc.io.

package formula

import (
	"fmt"
	"strconv"
	"strings"
)

// PrefixHorizontalRange is a range expression that when evaluated returns a list of Results from references like Sheet1!1:4 (all cells from rows 1 to 4 of sheet 'Sheet1').
type PrefixHorizontalRange struct {
	pfx Expression
	rowFrom, rowTo int
}

// NewPrefixHorizontalRange constructs a new full rows range with prefix.
func NewPrefixHorizontalRange(pfx Expression, v string) Expression {
	sl := strings.Split(v, ":")
	if len(sl) != 2 {
		return nil
	}
	from, _ := strconv.Atoi(sl[0])
	to, _ := strconv.Atoi(sl[1])
	return PrefixHorizontalRange{pfx, from, to}
}

// Eval evaluates the range returning a list of results or an error.
func (r PrefixHorizontalRange) Eval(ctx Context, ev Evaluator) Result {
	pfx := r.pfx.Reference(ctx, ev)
	switch pfx.Type {
	case ReferenceTypeSheet:
		c := ctx.Sheet(pfx.Value)
		from, to := cellRefsFromHorizontalRange(c, r.rowFrom, r.rowTo)
		return resultFromCellRange(c, ev, from, to)
	default:
		return MakeErrorResult(fmt.Sprintf("no support for reference type %s", pfx.Type))
	}
}

func (r PrefixHorizontalRange) Reference(ctx Context, ev Evaluator) Reference {
	pfx := r.pfx.Reference(ctx, ev)
	return Reference{Type: ReferenceTypeHorizontalRange, Value: fmt.Sprintf("%s!%d:%d", pfx.Value, r.rowFrom, r.rowTo)}
}
