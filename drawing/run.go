// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io.

package drawing

import "github.com/unidoc/unioffice/schema/soo/dml"

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
