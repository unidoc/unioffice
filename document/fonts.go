// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package document

import "baliance.com/gooxml/schema/soo/wml"

// Fonts allows manipulating a style or run's fonts.
type Fonts struct {
	x *wml.CT_Fonts
}

// X returns the inner wrapped XML type.
func (f Fonts) X() *wml.CT_Fonts {
	return f.x
}

// SetASCIITheme sets the font ASCII Theme.
func (f Fonts) SetASCIITheme(t wml.ST_Theme) {
	f.x.AsciiThemeAttr = t
}

// SetEastAsiaTheme sets the font East Asia Theme.
func (f Fonts) SetEastAsiaTheme(t wml.ST_Theme) {
	f.x.EastAsiaThemeAttr = t
}

// SetHANSITheme sets the font H ANSI Theme.
func (f Fonts) SetHANSITheme(t wml.ST_Theme) {
	f.x.HAnsiThemeAttr = t
}

// SetCSTheme sets the font complex script theme.
func (f Fonts) SetCSTheme(t wml.ST_Theme) {
	f.x.CsthemeAttr = t
}
