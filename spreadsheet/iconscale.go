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

// IconScale maps values to icons.
type IconScale struct {
	x *sml.CT_IconSet
}

// X returns the inner wrapped XML type.
func (c IconScale) X() *sml.CT_IconSet {
	return c.x
}

// SetIcons sets the icon set to use for display.
func (c IconScale) SetIcons(t sml.ST_IconSetType) {
	c.x.IconSetAttr = t
}

// AddFormatValue adds a format value to be used in determining which icons to display.
func (c IconScale) AddFormatValue(t sml.ST_CfvoType, val string) {
	v := sml.NewCT_Cfvo()
	v.TypeAttr = t
	v.ValAttr = unioffice.String(val)
	c.x.Cfvo = append(c.x.Cfvo, v)
}
