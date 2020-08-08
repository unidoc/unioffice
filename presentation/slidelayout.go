// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io.

package presentation

import (
	"github.com/unidoc/unioffice/schema/soo/pml"
)

// SlideLayout is a layout from which slides can be created.
type SlideLayout struct {
	x *pml.SldLayout
}

// X returns the inner wrapped XML type.
func (s SlideLayout) X() *pml.SldLayout {
	return s.x
}

// Type returns the type of the slide layout.
func (s SlideLayout) Type() pml.ST_SlideLayoutType {
	return s.x.TypeAttr
}

// Name returns the name of the slide layout.
func (s SlideLayout) Name() string {
	if s.x.CSld != nil && s.x.CSld.NameAttr != nil {
		return *s.x.CSld.NameAttr
	}
	return ""
}
