// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io.

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
