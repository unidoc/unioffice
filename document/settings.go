// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package document

import (
	"baliance.com/gooxml"
	"baliance.com/gooxml/schema/soo/wml"
)

// Settings controls the document settings.
type Settings struct {
	x *wml.Settings
}

// NewSettings constructs a new empty Settings
func NewSettings() Settings {
	s := wml.NewSettings()
	s.Compat = wml.NewCT_Compat()
	stng := wml.NewCT_CompatSetting()
	stng.NameAttr = gooxml.String("compatibilityMode")
	stng.UriAttr = gooxml.String("http://schemas.microsoft.com/office/word")
	stng.ValAttr = gooxml.String("15")
	s.Compat.CompatSetting = append(s.Compat.CompatSetting, stng)
	return Settings{s}
}

// X returns the inner wrapped XML type.
func (s Settings) X() *wml.Settings {
	return s.x
}

// SetUpdateFieldsOnOpen controls if fields are recalculated upon opening the
// document. This is useful for things like a table of contents as the library
// only adds the field code and relies on Word/LibreOffice to actually compute
// the content.
func (s Settings) SetUpdateFieldsOnOpen(b bool) {
	if !b {
		s.x.UpdateFields = nil
	} else {
		s.x.UpdateFields = wml.NewCT_OnOff()
	}
}
