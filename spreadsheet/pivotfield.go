// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheet

import (
	"baliance.com/gooxml/schema/soo/sml"
)

type PivotField struct {
	x *sml.CT_PivotField
}

// X returns the inner wrapped XML type.
func (p PivotField) X() *sml.CT_PivotField {
	return p.x
}
