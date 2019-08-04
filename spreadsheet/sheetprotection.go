// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheet

import (
	"github.com/unidoc/unioffice"
	"github.com/unidoc/unioffice/schema/soo/sml"
)

type SheetProtection struct {
	x *sml.CT_SheetProtection
}

// X returns the inner wrapped XML type.
func (p SheetProtection) X() *sml.CT_SheetProtection {
	return p.x
}

// IsSheetLocked returns whether the sheet is locked.
func (p SheetProtection) IsSheetLocked() bool {
	return p.x.SheetAttr != nil && *p.x.SheetAttr
}

// LockSheet controls the locking of the sheet.
func (p SheetProtection) LockSheet(b bool) {
	if !b {
		p.x.SheetAttr = nil
	} else {
		p.x.SheetAttr = unioffice.Bool(true)
	}
}

// IsSheetLocked returns whether the sheet objects are locked.
func (p SheetProtection) IsObjectLocked() bool {
	return p.x.ObjectsAttr != nil && *p.x.ObjectsAttr
}

// LockObject controls the locking of the sheet objects.
func (p SheetProtection) LockObject(b bool) {
	if !b {
		p.x.ObjectsAttr = nil
	} else {
		p.x.ObjectsAttr = unioffice.Bool(true)
	}
}

// PasswordHash returns the hash of the workbook password.
func (p SheetProtection) PasswordHash() string {
	if p.x.PasswordAttr == nil {
		return ""
	}
	return *p.x.PasswordAttr
}

// SetPassword sets the password hash to a hash of the input password.
func (p SheetProtection) SetPassword(pw string) {
	p.SetPasswordHash(PasswordHash(pw))
}

// SetPasswordHash sets the password hash to the input.
func (p SheetProtection) SetPasswordHash(pwHash string) {
	p.x.PasswordAttr = unioffice.String(pwHash)
}
