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

// PrefixHorizontalRange is a range expression that when evaluated returns a list of Results from references like Sheet1!1:4 (all cells from rows 1 to 4 of sheet 'Sheet1').
type PrefixHorizontalRange struct {
	pfx            Expression
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

// Eval evaluates a horizontal range with prefix returning a list of results or an error.
func (r PrefixHorizontalRange) Eval(ctx Context, ev Evaluator) Result {
	pfx := r.pfx.Reference(ctx, ev)
	switch pfx.Type {
	case ReferenceTypeSheet:
		ref := r.horizontalRangeReference(pfx.Value)
		if cached, found := ev.GetFromCache(ref); found {
			return cached
		}
		c := ctx.Sheet(pfx.Value)
		from, to := cellRefsFromHorizontalRange(c, r.rowFrom, r.rowTo)
		result := resultFromCellRange(c, ev, from, to)
		ev.SetCache(ref, result)
		return result
	default:
		return MakeErrorResult(fmt.Sprintf("no support for reference type %s", pfx.Type))
	}
}

func (r PrefixHorizontalRange) horizontalRangeReference(sheetName string) string {
	return fmt.Sprintf("%s!%d:%d", sheetName, r.rowFrom, r.rowTo)
}

// Reference returns a string reference value to a horizontal range with prefix.
func (r PrefixHorizontalRange) Reference(ctx Context, ev Evaluator) Reference {
	pfx := r.pfx.Reference(ctx, ev)
	return Reference{Type: ReferenceTypeHorizontalRange, Value: r.horizontalRangeReference(pfx.Value)}
}

// String returns a string representation of a horizontal range with prefix.
func (r PrefixHorizontalRange) String() string {
	return fmt.Sprintf("%s!%d:%d", r.pfx.String(), r.rowFrom, r.rowTo)
}

// Update updates references in the PrefixHorizontalRange after removing a row/column.
func (r PrefixHorizontalRange) Update(q *update.UpdateQuery) Expression {
	return r
}
