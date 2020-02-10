// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased on https://unidoc.io.

package formula

import "github.com/unidoc/unioffice/spreadsheet/update"

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

// Eval evaluates a horizontal range returning a list of results or an error.
func (r HorizontalRange) Eval(ctx Context, ev Evaluator) Result {
	key := r.horizontalRangeReference()
	if cached, found := ev.GetFromCache(key); found {
		return cached
	}
	from, to := cellRefsFromHorizontalRange(ctx, r.rowFrom, r.rowTo)
	res := resultFromCellRange(ctx, ev, from, to)
	ev.SetCache(key, res)
	return res
}

func (r HorizontalRange) horizontalRangeReference() string {
	return fmt.Sprintf("%d:%d", r.rowFrom, r.rowTo)
}

// Reference returns a string reference value to a horizontal range.
func (r HorizontalRange) Reference(ctx Context, ev Evaluator) Reference {
	return Reference{Type: ReferenceTypeHorizontalRange, Value: r.horizontalRangeReference()}
}

func cellRefsFromHorizontalRange(ctx Context, rowFrom, rowTo int) (string, string) {
	from := "A" + strconv.Itoa(rowFrom)
	lastColumn := ctx.LastColumn(rowFrom, rowTo)
	to := lastColumn + strconv.Itoa(rowTo)
	return from, to
}

// String returns a string representation of a horizontal range.
func (r HorizontalRange) String() string {
	return r.horizontalRangeReference()
}

// Update updates the horizontal range references after removing a row/column.
func (r HorizontalRange) Update(q *update.UpdateQuery) Expression {
	return r
}
