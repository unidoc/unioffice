// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package formula

import (
	"fmt"
	"strings"
)

// NamedRangeRef is a reference to a named range
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
	switch ref.Type {
	case ReferenceTypeCell:
		return ev.Eval(ctx, ref.Value)
	case ReferenceTypeRange:
		// should look like "A2:C5"
		sp := strings.Split(ref.Value, ":")
		if len(sp) == 2 {
			return resultFromCellRange(ctx, ev, sp[0], sp[1])
		}
		return MakeErrorResult(fmt.Sprintf("unsuppported named range value %s", ref.Value))
	}
	return MakeErrorResult(fmt.Sprintf("unsuppported reference type %s", ref.Type))
}

func (n NamedRangeRef) Reference(ctx Context, ev Evaluator) Reference {
	return Reference{Type: ReferenceTypeNamedRange, Value: n.s}
}
