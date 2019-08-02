// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package common

import "github.com/unidoc/unioffice/schema/soo/dml"

// Theme is a drawingml theme.
type Theme struct {
	x *dml.Theme
}

// NewTheme constructs a new theme.
func NewTheme() Theme {
	return Theme{dml.NewTheme()}
}

// X returns the inner wrapped XML type.
func (t Theme) X() *dml.Theme {
	return t.x
}
