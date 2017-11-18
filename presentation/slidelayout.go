// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package presentation

import (
	"baliance.com/gooxml/schema/soo/pml"
)

// SlideLayout
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
