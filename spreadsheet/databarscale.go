// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheet

import (
	"baliance.com/gooxml"
	"baliance.com/gooxml/color"
	"baliance.com/gooxml/schema/soo/sml"
)

// DataBarScale is a colored scale that fills the cell with a background
// gradeint depending on the value.
type DataBarScale struct {
	x *sml.CT_DataBar
}

// X returns the inner wrapped XML type.
func (d DataBarScale) X() *sml.CT_DataBar {
	return d.x
}

// AddFormatValue adds a format value (databars require two).
func (d DataBarScale) AddFormatValue(t sml.ST_CfvoType, val string) {
	v := sml.NewCT_Cfvo()
	v.TypeAttr = t
	v.ValAttr = gooxml.String(val)
	d.x.Cfvo = append(d.x.Cfvo, v)
}

// SetColor sets teh color of the databar.
func (d DataBarScale) SetColor(c color.Color) {
	d.x.Color = sml.NewCT_Color()
	d.x.Color.RgbAttr = c.AsRGBAString()
}

// SetShowValue controls if the cell value is displayed.
func (d DataBarScale) SetShowValue(b bool) {
	d.x.ShowValueAttr = gooxml.Bool(b)
}

// SetMinLength sets the minimum bar length in percent.
func (d DataBarScale) SetMinLength(l uint32) {
	d.x.MinLengthAttr = gooxml.Uint32(l)
}

// SetMaxLength sets the maximum bar length in percent.
func (d DataBarScale) SetMaxLength(l uint32) {
	d.x.MaxLengthAttr = gooxml.Uint32(l)
}
