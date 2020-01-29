// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased on https://unidoc.io.

package formula

// Evaluator is the interface for a formula evaluator.  This is needed so we can
// pass it to the spreadsheet to let it evaluate formula cells before returning
// the results.
// NOTE: in order to implement Evaluator without cache embed noCache in it.
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
