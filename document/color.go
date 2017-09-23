// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package document

import (
	"fmt"

	"baliance.com/gooxml/color"
	"baliance.com/gooxml/schema/soo/wml"
)

// Color controls the run or styles color.
type Color struct {
	x *wml.CT_Color
}

// X returns the inner wrapped XML type.
func (c Color) X() *wml.CT_Color {
	return c.x
}

// SetColor sets a specific color or auto.
func (c Color) SetColor(v color.Color) {
	if v.IsAuto() {
		c.x.ValAttr.ST_HexColorAuto = wml.ST_HexColorAutoAuto
		c.x.ValAttr.ST_HexColorRGB = nil
	} else {
		c.x.ValAttr.ST_HexColorAuto = wml.ST_HexColorAutoUnset
		c.x.ValAttr.ST_HexColorRGB = v.AsRGBString()
	}
}

// SetThemeColor sets the color from the theme.
func (c Color) SetThemeColor(t wml.ST_ThemeColor) {
	c.x.ThemeColorAttr = t
}

// SetThemeShade sets the shade based off the theme color.
func (c Color) SetThemeShade(s uint8) {
	shd := fmt.Sprintf("%02x", s)
	c.x.ThemeShadeAttr = &shd
}
