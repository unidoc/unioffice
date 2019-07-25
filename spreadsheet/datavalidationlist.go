// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheet

import (
	"strings"

	"github.com/unidoc/unioffice"
	"github.com/unidoc/unioffice/schema/soo/sml"
)

// DataValidationList is just a view on a DataValidation configured as a list.
// It presents a drop-down combo box for spreadsheet users to select values. The
// contents of the dropdown can either pull from a rang eof cells (SetRange) or
// specified directly (SetValues).
type DataValidationList struct {
	x *sml.CT_DataValidation
}

// SetRange sets the range that contains the possible values. This is incompatible with SetValues.
func (d DataValidationList) SetRange(cellRange string) {
	d.x.Formula1 = unioffice.String(cellRange)
	d.x.Formula2 = unioffice.String("0")
}

// SetValues sets the possible values. This is incompatible with SetRange.
func (d DataValidationList) SetValues(values []string) {
	d.x.Formula1 = unioffice.String("\"" + strings.Join(values, ",") + "\"")
	d.x.Formula2 = unioffice.String("0")
}
