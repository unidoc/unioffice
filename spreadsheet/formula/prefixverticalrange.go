// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased on https://unidoc.io.

package formula

import (
	"fmt"
	"strings"

	"github.com/unidoc/unioffice/spreadsheet/update"
)

// PrefixVerticalRange is a range expression that when evaluated returns a list of Results from references like Sheet1!AA:IJ (all cells from columns AA to IJ of sheet 'Sheet1').
type PrefixVerticalRange struct {
	pfx            Expression
	colFrom, colTo string
}

// NewPrefixVerticalRange constructs a new full columns range with prefix.
func NewPrefixVerticalRange(pfx Expression, v string) Expression {
	sl := strings.Split(v, ":")
	if len(sl) != 2 {
		return nil
	}
	return PrefixVerticalRange{pfx, sl[0], sl[1]}
}

// Eval evaluates a vertical range with prefix returning a list of results or an error.
func (r PrefixVerticalRange) Eval(ctx Context, ev Evaluator) Result {
	pfx := r.pfx.Reference(ctx, ev)
	switch pfx.Type {
	case ReferenceTypeSheet:
		ref := r.verticalRangeReference(pfx.Value)
		if cached, found := ev.GetFromCache(ref); found {
			return cached
		}
		c := ctx.Sheet(pfx.Value)
		from, to := cellRefsFromVerticalRange(c, r.colFrom, r.colTo)
		result := resultFromCellRange(c, ev, from, to)
		ev.SetCache(ref, result)
		return result
	default:
		return MakeErrorResult(fmt.Sprintf("no support for reference type %s", pfx.Type))
	}
}

func (r PrefixVerticalRange) verticalRangeReference(sheetName string) string {
	return fmt.Sprintf("%s!%s:%s", sheetName, r.colFrom, r.colTo)
}

// Reference returns a string reference value to a vertical range with prefix.
func (r PrefixVerticalRange) Reference(ctx Context, ev Evaluator) Reference {
	pfx := r.pfx.Reference(ctx, ev)
	return Reference{Type: ReferenceTypeVerticalRange, Value: r.verticalRangeReference(pfx.Value)}
}

// String returns a string representation of a vertical range with prefix.
func (r PrefixVerticalRange) String() string {
	return fmt.Sprintf("%s!%s:%s", r.pfx.String(), r.colFrom, r.colTo)
}

// Update updates references in the PrefixVerticalRange after removing a row/column.
func (r PrefixVerticalRange) Update(q *update.UpdateQuery) Expression {
	if q.UpdateType == update.UpdateActionRemoveColumn {
		new := r
		sheetName := r.pfx.String()
		if sheetName == q.SheetToUpdate {
			columnIdx := q.ColumnIdx
			new.colFrom = updateColumnToLeft(r.colFrom, columnIdx)
			new.colTo = updateColumnToLeft(r.colTo, columnIdx)
		}
		return new
	}
	return r
}
