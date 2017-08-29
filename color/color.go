// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package color

import "fmt"

// Color is a 24 bit color that can be converted to
// internal ECMA-376 formats as needed.
type Color struct {
	a, r, g, b uint8
	isAuto     bool
}

// RGB constructs a new RGB color with a given red, green and blue value.
func RGB(r, g, b uint8) Color {
	return Color{255, r, g, b, false}
}

// IsAuto returns true if the color is the 'Auto' type.  If the
// field doesn't support an Auto color, then black is used.
func (c Color) IsAuto() bool {
	return c.isAuto
}

// AsRGBString is used by the various wrappers to return a pointer
// to a string containing a six digit hex RGB value.
func (c Color) AsRGBString() *string {
	v := fmt.Sprintf("%02x%02x%02x", c.r, c.g, c.b)
	return &v
}

// AsRGBAString is used by the various wrappers to return a pointer
// to a string containing a six digit hex RGB value.
func (c Color) AsRGBAString() *string {
	v := fmt.Sprintf("%02x%02x%02x%02x", c.a, c.r, c.g, c.b)
	return &v
}
