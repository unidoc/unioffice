// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io.

package document

import "github.com/unidoc/unioffice/schema/soo/wml"

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
