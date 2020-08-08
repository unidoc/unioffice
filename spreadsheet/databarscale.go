// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io.

package spreadsheet

import (
	"github.com/unidoc/unioffice"
	"github.com/unidoc/unioffice/color"
	"github.com/unidoc/unioffice/schema/soo/sml"
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
	v.ValAttr = unioffice.String(val)
	d.x.Cfvo = append(d.x.Cfvo, v)
}

// SetColor sets teh color of the databar.
func (d DataBarScale) SetColor(c color.Color) {
	d.x.Color = sml.NewCT_Color()
	d.x.Color.RgbAttr = c.AsRGBAString()
}

// SetShowValue controls if the cell value is displayed.
func (d DataBarScale) SetShowValue(b bool) {
	d.x.ShowValueAttr = unioffice.Bool(b)
}

// SetMinLength sets the minimum bar length in percent.
func (d DataBarScale) SetMinLength(l uint32) {
	d.x.MinLengthAttr = unioffice.Uint32(l)
}

// SetMaxLength sets the maximum bar length in percent.
func (d DataBarScale) SetMaxLength(l uint32) {
	d.x.MaxLengthAttr = unioffice.Uint32(l)
}
