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

// HorizontalRange is a range expression that when evaluated returns a list of Results from references like 1:4 (all cells from rows 1 to 4).
type HorizontalRange struct {
	rowFrom, rowTo int
}

// NewHorizontalRange constructs a new full rows range.
func NewHorizontalRange(v string) Expression {
	sl := strings.Split(v, ":")
	if len(sl) != 2 {
		return nil
	}
	from, _ := strconv.Atoi(sl[0])
	to, _ := strconv.Atoi(sl[1])
	return HorizontalRange{from, to}
}

// Eval evaluates the range returning a list of results or an error.
func (r HorizontalRange) Eval(ctx Context, ev Evaluator) Result {
	from, to := cellRefsFromHorizontalRange(ctx, r.rowFrom, r.rowTo)
	return resultFromCellRange(ctx, ev, from, to)
}

func (r HorizontalRange) Reference(ctx Context, ev Evaluator) Reference {
	return Reference{Type: ReferenceTypeHorizontalRange, Value: fmt.Sprintf("%d:%d", r.rowFrom, r.rowTo)}
}

func cellRefsFromHorizontalRange(ctx Context, rowFrom, rowTo int) (string, string) {
	from := "A" + strconv.Itoa(rowFrom)
	lastColumn := ctx.LastColumn(rowFrom, rowTo)
	to := lastColumn + strconv.Itoa(rowTo)
	return from, to
}
