// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package drawing

import "baliance.com/gooxml/schema/soo/dml"

// MakeRun constructs a new Run wrapper.
func MakeRun(x *dml.EG_TextRun) Run {
	return Run{x}
}

// Run is a run within a paragraph.
type Run struct {
	x *dml.EG_TextRun
}

// X returns the inner wrapped XML type.
func (r Run) X() *dml.EG_TextRun {
	return r.x
}

// SetText sets the run's text contents.
func (r Run) SetText(s string) {
	r.x.Br = nil
	r.x.Fld = nil
	if r.x.R == nil {
		r.x.R = dml.NewCT_RegularTextRun()
	}
	r.x.R.T = s
}

// Properties returns the run's properties.
func (r Run) Properties() RunProperties {
	if r.x.R == nil {
		r.x.R = dml.NewCT_RegularTextRun()
	}
	if r.x.R.RPr == nil {
		r.x.R.RPr = dml.NewCT_TextCharacterProperties()
	}
	return RunProperties{r.x.R.RPr}
}
