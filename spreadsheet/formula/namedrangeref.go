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

// NamedRangeRef is a reference to a named range.
type NamedRangeRef struct {
	s string
}

// NewNamedRangeRef constructs a new named range reference.
func NewNamedRangeRef(v string) Expression {
	return NamedRangeRef{v}
}

// Eval evaluates and returns the result of the NamedRangeRef reference.
func (n NamedRangeRef) Eval(ctx Context, ev Evaluator) Result {
	ref := ctx.NamedRange(n.s)
	refValue := ref.Value
	if cached, found := ev.GetFromCache(refValue); found {
		return cached
	}
	sl := strings.Split(refValue, "!")
	if len(sl) != 2 {
		return MakeErrorResult(fmt.Sprintf("unsupported named range value %s", refValue))
	}
	sheetCtx := ctx.Sheet(sl[0])
	sp := strings.Split(sl[1], ":")
	switch len(sp) {
	case 1:
		result := ev.Eval(sheetCtx, sp[0])
		ev.SetCache(refValue, result)
		return result
	case 2:
		// should look like "A2:C5"
		result := resultFromCellRange(sheetCtx, ev, sp[0], sp[1])
		ev.SetCache(refValue, result)
		return result
	}
	return MakeErrorResult(fmt.Sprintf("unsupported reference type %s", ref.Type))
}

// Reference returns a string reference value to a named range.
func (n NamedRangeRef) Reference(ctx Context, ev Evaluator) Reference {
	return Reference{Type: ReferenceTypeNamedRange, Value: n.s}
}

// String returns a string representation of a named range.
func (n NamedRangeRef) String() string {
	return n.s
}

// Update returns the same object as updating sheet references does not affect named ranges.
func (n NamedRangeRef) Update(q *update.UpdateQuery) Expression {
	return n
}
