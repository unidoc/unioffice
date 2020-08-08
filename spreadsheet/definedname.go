// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io.

package spreadsheet

import "github.com/unidoc/unioffice/schema/soo/sml"
import "github.com/unidoc/unioffice"

// DefinedName is a named range, formula, etc.
type DefinedName struct {
	x *sml.CT_DefinedName
}

// X returns the inner wrapped XML type.
func (d DefinedName) X() *sml.CT_DefinedName {
	return d.x
}

// Name returns the name of the defined name.
func (d DefinedName) Name() string {
	return d.x.NameAttr
}

// Content returns the content of the defined range (the range in most cases)/
func (d DefinedName) Content() string {
	return d.x.Content
}

// SetContent sets the defined name content.
func (d DefinedName) SetContent(s string) {
	d.x.Content = s
}

// SetHidden marks the defined name as hidden.
func (d DefinedName) SetHidden(b bool) {
	d.x.HiddenAttr = unioffice.Bool(b)
}

// SetHidden marks the defined name as hidden.
func (d DefinedName) SetLocalSheetID(id uint32) {
	d.x.LocalSheetIdAttr = unioffice.Uint32(id)
}
