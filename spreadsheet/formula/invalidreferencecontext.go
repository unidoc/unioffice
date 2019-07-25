// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package formula

// InvalidReferenceContext is a Context that can be used when evaluating an
// invalid reference (e.g. referencing a non-existent sheet).  It implements
// Context safely, but returns error results.
var InvalidReferenceContext = &ivr{}

type ivr struct {
}

func (i *ivr) Cell(ref string, ev Evaluator) Result {
	return MakeErrorResult("invalid reference")
}

func (i *ivr) NamedRange(ref string) Reference {
	return ReferenceInvalid
}

func (i *ivr) Sheet(name string) Context {
	return i
}

func (i *ivr) SetOffset(col, row uint32) {

}
