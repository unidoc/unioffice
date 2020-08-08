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

// ColorScale colors a cell background based off of the cell value.
type ColorScale struct {
	x *sml.CT_ColorScale
}

// X returns the inner wrapped XML type.
func (c ColorScale) X() *sml.CT_ColorScale {
	return c.x
}

// AddFormatValue adds a format value to be used to determine the cell background.
func (c ColorScale) AddFormatValue(t sml.ST_CfvoType, val string) {
	v := sml.NewCT_Cfvo()
	v.TypeAttr = t
	v.ValAttr = unioffice.String(val)
	c.x.Cfvo = append(c.x.Cfvo, v)
}

// AddGradientStop adds a color gradient stop.
func (c ColorScale) AddGradientStop(color color.Color) {
	clr := sml.NewCT_Color()
	clr.RgbAttr = color.AsRGBAString()
	c.x.Color = append(c.x.Color, clr)
}
