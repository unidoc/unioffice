// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheet

import "baliance.com/gooxml/schema/soo/sml"

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
