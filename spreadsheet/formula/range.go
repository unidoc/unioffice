// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased on https://unidoc.io.

package formula

import (
	"fmt"

	"github.com/unidoc/unioffice/spreadsheet/reference"
	"github.com/unidoc/unioffice/spreadsheet/update"
)

// Range is a range expression that when evaluated returns a list of Results.
type Range struct {
	from, to Expression
}

// NewRange constructs a new range.
func NewRange(from, to Expression) Expression {
	return Range{from, to}
}

// Eval evaluates a range returning a list of results or an error.
func (r Range) Eval(ctx Context, ev Evaluator) Result {
	from := r.from.Reference(ctx, ev)
	to := r.to.Reference(ctx, ev)
	ref := rangeReference(from, to)
	if from.Type == ReferenceTypeCell && to.Type == ReferenceTypeCell {
		if cached, found := ev.GetFromCache(ref); found {
			return cached
		} else {
			result := resultFromCellRange(ctx, ev, from.Value, to.Value)
			ev.SetCache(ref, result)
			return result
		}
	}
	return MakeErrorResult("invalid range " + ref)
}

func rangeReference(from, to Reference) string {
	return fmt.Sprintf("%s:%s", from.Value, to.Value)
}

// Reference returns a string reference value to a range.
func (r Range) Reference(ctx Context, ev Evaluator) Reference {
	from := r.from.Reference(ctx, ev)
	to := r.to.Reference(ctx, ev)
	if from.Type == ReferenceTypeCell && to.Type == ReferenceTypeCell {
		return MakeRangeReference(rangeReference(from, to))
	}
	return ReferenceInvalid
}

func resultFromCellRange(ctx Context, ev Evaluator, from, to string) Result {
	fromRef, fe := reference.ParseCellReference(from)
	if fe != nil {
		return MakeErrorResult(fmt.Sprintf("unable to parse range %s: error %s", from, fe.Error()))
	}
	fc, fr := fromRef.ColumnIdx, fromRef.RowIdx

	toRef, te := reference.ParseCellReference(to)
	if te != nil {
		return MakeErrorResult(fmt.Sprintf("unable to parse range %s: error %s", to, te.Error()))
	}
	tc, tr := toRef.ColumnIdx, toRef.RowIdx

	arr := [][]Result{}
	for r := fr; r <= tr; r++ {
		args := []Result{}
		for c := fc; c <= tc; c++ {
			res := ctx.Cell(fmt.Sprintf("%s%d", reference.IndexToColumn(c), r), ev)
			args = append(args, res)
		}
		arr = append(arr, args)
	}
	// for a single row, just return a list
	if len(arr) == 1 {
		// single cell result
		if len(arr[0]) == 1 {
			return arr[0][0]
		}
		return MakeListResult(arr[0])
	}

	return MakeArrayResult(arr)
}

// String returns a string of a range.
func (r Range) String() string {
	return fmt.Sprintf("%s:%s", r.from.String(), r.to.String())
}

// Update updates references in the Range after removing a row/column.
func (r Range) Update(q *update.UpdateQuery) Expression {
	new := r
	if q.UpdateCurrentSheet {
		new.from = r.from.Update(q)
		new.to = r.to.Update(q)
	}
	return new
}
