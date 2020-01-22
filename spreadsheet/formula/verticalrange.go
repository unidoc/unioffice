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

// Eval evaluates the range returning a list of results or an error.
func (r VerticalRange) Eval(ctx Context, ev Evaluator) Result {
	from, to := cellRefsFromVerticalRange(ctx, r.colFrom, r.colTo)
	return resultFromCellRange(ctx, ev, from, to)
}

func (r VerticalRange) Reference(ctx Context, ev Evaluator) Reference {
	return Reference{Type: ReferenceTypeVerticalRange, Value: fmt.Sprintf("%s:%s", r.colFrom, r.colTo)}
}

func cellRefsFromVerticalRange(ctx Context, colFrom, colTo string) (string, string) {
	from := colFrom + "1"
	lastRow := ctx.LastRow(colFrom)
	to := colTo + strconv.Itoa(lastRow)
	return from, to
}
