// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io.

package spreadsheet

import "github.com/unidoc/unioffice/schema/soo/sml"

type Table struct {
	x *sml.Table
}

// X returns the inner wrapped XML type.
func (t Table) X() *sml.Table {
	return t.x
}

// Name returns the name of the table
func (t Table) Name() string {
	if t.x.NameAttr != nil {
		return *t.x.NameAttr
	}
	return ""
}

// Reference returns the table reference (the cells within the table)
func (t Table) Reference() string {
	return t.x.RefAttr
}
