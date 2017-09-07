// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheet

import sml "baliance.com/gooxml/schema/schemas.openxmlformats.org/spreadsheetml"

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
