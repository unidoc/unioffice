// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package formula

import (
	"fmt"
	"strconv"
	"strings"

	"baliance.com/gooxml/spreadsheet/reference"
)

// Range is a range expression that when evaluated returns a list of Results.
type Range struct {
	from, to Expression
}

// NewRange constructs a new range.
func NewRange(from, to Expression) Expression {
	return Range{from, to}
}

// Eval evaluates the range returning a list of results or an error.
func (r Range) Eval(ctx Context, ev Evaluator) Result {
	from := r.from.Reference(ctx, ev)
	to := r.to.Reference(ctx, ev)
	if from.Type == ReferenceTypeCell && to.Type == ReferenceTypeCell {
		return resultFromCellRange(ctx, ev, from.Value, to.Value)
	}
	return MakeErrorResult("invalid range " + from.Value + " to " + to.Value)
}

func (r Range) Reference(ctx Context, ev Evaluator) Reference {
	from := r.from.Reference(ctx, ev)
	to := r.to.Reference(ctx, ev)
	if from.Type == ReferenceTypeCell && to.Type == ReferenceTypeCell {
		return MakeRangeReference(from.Value + ":" + to.Value)
	}
	return ReferenceInvalid
}

// TODO: move these somewhere to remove duplication
func ParseCellReference(s string) (col string, row uint32, err error) {
	s = strings.Replace(s, "$", "", -1)
	split := -1
lfor:
	for i := 0; i < len(s); i++ {
		switch {
		case s[i] >= '0' && s[i] <= '9':
			split = i
			break lfor
		}
	}
	switch split {
	case 0:
		return col, row, fmt.Errorf("no letter prefix in %s", s)
	case -1:
		return col, row, fmt.Errorf("no digits in %s", s)
	}

	col = s[0:split]
	r64, err := strconv.ParseUint(s[split:], 10, 32)
	row = uint32(r64)
	return col, row, err
}

func resultFromCellRange(ctx Context, ev Evaluator, from, to string) Result {
	fc, fr, fe := ParseCellReference(from)
	tc, tr, te := ParseCellReference(to)
	if fe != nil {
		return MakeErrorResult("unable to parse range " + from)
	}
	if te != nil {
		return MakeErrorResult("unable to parse range " + to)
	}
	bc := reference.ColumnToIndex(fc)
	ec := reference.ColumnToIndex(tc)
	arr := [][]Result{}
	for r := fr; r <= tr; r++ {
		args := []Result{}
		for c := bc; c <= ec; c++ {
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
