// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased on https://unidoc.io.

package formula

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// Evaluator is the interface for a formula evaluator.  This is needed so we can
// pass it to the spreadsheet to let it evaluate formula cells before returning
// the results.
type Evaluator interface {
	Eval(ctx Context, formula string) Result
	SetCache(key string, value Result)
	GetFromCache(key string) (Result, bool)
	LastEvalIsRef() bool
}

// NewEvaluator constructs a new defEval object which is the default formula evaluator.
func NewEvaluator() Evaluator {
	ev := &defEval{}
	ev.evCache = newEvCache()
	return ev
}

// defEval is the default formula evaluator which implements the Evaluator interface.
type defEval struct {
	evCache
	lastEvalIsRef bool
}

func (d *defEval) Eval(ctx Context, formula string) Result {
	expr := ParseString(formula)
	if expr != nil {
		d.addInfo(ctx, expr)
		return expr.Eval(ctx, d)
	}
	return MakeErrorResult(fmt.Sprintf("unable to parse formula %s", formula))
}

func (d *defEval) LastEvalIsRef() bool {
	return d.lastEvalIsRef
}

// addInfo adds information which is needed for some functions but is lost after evaluation. E.g. which arguments are actually references.
func (d *defEval) addInfo(ctx Context, expr Expression) {
	switch expr.(type) {
	case FunctionCall:
		switch expr.(FunctionCall).name {
		case "ISREF":
			for _, arg := range expr.(FunctionCall).args {
				switch arg.(type) {
				case CellRef:
					d.lastEvalIsRef = validateRef(arg.(CellRef))
					return
				case Range:
					switch arg.(Range).from.(type) {
					case CellRef:
						d.lastEvalIsRef = validateRef(arg.(Range).from.(CellRef))
						return
					}
					switch arg.(Range).to.(type) {
					case CellRef:
						d.lastEvalIsRef = validateRef(arg.(Range).to.(CellRef))
						return
					}
				default:
					d.lastEvalIsRef = false
				}
			}
		}
	}
}

var refRegexp *regexp.Regexp = regexp.MustCompile(`^([a-z]+)([0-9]+)$`)

func validateRef(cr CellRef) bool {
	if submatch := refRegexp.FindStringSubmatch(strings.ToLower(cr.s)); len(submatch) > 2 {
		col := submatch[1]
		row, err := strconv.Atoi(submatch[2])
		if err != nil { // for the case if the row number is bigger then int capacity
			return false
		}
		return row <= 1048576 && col <= "zz"
	}
	return false
}
