// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheet

import "baliance.com/gooxml/spreadsheet/formula"

func newEvalContext(s *Sheet) *evalContext {
	return &evalContext{s, make(map[string]struct{})}
}

type evalContext struct {
	s          *Sheet
	evaluating map[string]struct{}
}

func (e *evalContext) Cell(ref string, ev formula.Evaluator) formula.Result {
	c := e.s.Cell(ref)

	// if we have a formula, evaluate it
	if c.HasFormula() {
		if _, ok := e.evaluating[ref]; ok {
			// recursively evaluating, so bail out
			return formula.MakeErrorResult("recursion detected during evaluation of " + ref)
		}
		e.evaluating[ref] = struct{}{}
		res := ev.Eval(e, c.GetFormula())
		delete(e.evaluating, ref)
		return res
	}

	if c.IsNumber() {
		v, _ := c.GetValueAsNumber()
		return formula.MakeNumberResult(v)
	} else if c.IsBool() {
		v, _ := c.GetValueAsBool()
		return formula.MakeBoolResult(v)
	} else {
		v, _ := c.GetValue()
		return formula.MakeStringResult(v)
	}
	// TODO: handle this properly
	// return formula.MakeErrorResult()
}
