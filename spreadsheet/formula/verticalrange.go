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

	"github.com/unidoc/unioffice/spreadsheet/update"
)

// VerticalRange is a range expression that when evaluated returns a list of Results from references like AA:IJ (all cells from columns AA to IJ).
type VerticalRange struct {
	colFrom, colTo string
}

// NewVerticalRange constructs a new full columns range.
func NewVerticalRange(v string) Expression {
	sl := strings.Split(v, ":")
	if len(sl) != 2 {
		return nil
	}
	return VerticalRange{sl[0], sl[1]}
}

// Eval evaluates a vertical range returning a list of results or an error.
func (r VerticalRange) Eval(ctx Context, ev Evaluator) Result {
	key := r.verticalRangeReference()
	if cached, found := ev.GetFromCache(key); found {
		return cached
	}
	from, to := cellRefsFromVerticalRange(ctx, r.colFrom, r.colTo)
	res := resultFromCellRange(ctx, ev, from, to)
	ev.SetCache(key, res)
	return res
}

func (r VerticalRange) verticalRangeReference() string {
	return fmt.Sprintf("%s:%s", r.colFrom, r.colTo)
}

// Reference returns a string reference value to a vertical range.
func (r VerticalRange) Reference(ctx Context, ev Evaluator) Reference {
	return Reference{Type: ReferenceTypeVerticalRange, Value: r.verticalRangeReference()}
}

func cellRefsFromVerticalRange(ctx Context, colFrom, colTo string) (string, string) {
	from := colFrom + "1"
	lastRow := ctx.LastRow(colFrom)
	to := colTo + strconv.Itoa(lastRow)
	return from, to
}

// String returns a string representation of a vertical range.
func (r VerticalRange) String() string {
	return r.verticalRangeReference()
}

// Update updates references in the VerticalRange after removing a row/column.
func (r VerticalRange) Update(q *update.UpdateQuery) Expression {
	if q.UpdateType == update.UpdateActionRemoveColumn {
		new := r
		if q.UpdateCurrentSheet {
			columnIdx := q.ColumnIdx
			new.colFrom = updateColumnToLeft(r.colFrom, columnIdx)
			new.colTo = updateColumnToLeft(r.colTo, columnIdx)
		}
		return new
	}
	return r
}
