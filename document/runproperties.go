// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package document

import (
	"baliance.com/gooxml"
	"baliance.com/gooxml/color"
	"baliance.com/gooxml/measurement"
	"baliance.com/gooxml/schema/soo/ofc/sharedTypes"
	"baliance.com/gooxml/schema/soo/wml"
)

// RunProperties controls run styling properties
type RunProperties struct {
	x *wml.CT_RPr
}

// X returns the inner wrapped XML type.
func (r RunProperties) X() *wml.CT_RPr {
	return r.x
}

// SetStyle sets the font size.
func (r RunProperties) SetStyle(style string) {
	if style == "" {
		r.x.RStyle = nil
	} else {
		r.x.RStyle = wml.NewCT_String()
		r.x.RStyle.ValAttr = style
	}
}

// SetFontFamily sets the Ascii & HAnsi fonly family for a run.
func (r RunProperties) SetFontFamily(family string) {
	if r.x.RFonts == nil {
		r.x.RFonts = wml.NewCT_Fonts()
	}
	r.x.RFonts.AsciiAttr = gooxml.String(family)
	r.x.RFonts.HAnsiAttr = gooxml.String(family)
	r.x.RFonts.EastAsiaAttr = gooxml.String(family)
}

// SetSize sets the font size for a run.
func (r RunProperties) SetSize(size measurement.Distance) {
	r.x.Sz = wml.NewCT_HpsMeasure()
	r.x.Sz.ValAttr.ST_UnsignedDecimalNumber = gooxml.Uint64(uint64(size / measurement.HalfPoint))
	r.x.SzCs = wml.NewCT_HpsMeasure()
	r.x.SzCs.ValAttr.ST_UnsignedDecimalNumber = gooxml.Uint64(uint64(size / measurement.HalfPoint))
}

// SetKerning sets the run's font kerning.
func (r RunProperties) SetKerning(size measurement.Distance) {
	r.x.Kern = wml.NewCT_HpsMeasure()
	r.x.Kern.ValAttr.ST_UnsignedDecimalNumber = gooxml.Uint64(uint64(size / measurement.HalfPoint))
}

// SetCharacterSpacing sets the run's Character Spacing Adjustment.
func (r RunProperties) SetCharacterSpacing(size measurement.Distance) {
	r.x.Spacing = wml.NewCT_SignedTwipsMeasure()
	r.x.Spacing.ValAttr.Int64 = gooxml.Int64(int64(size / measurement.Twips))
}

// Fonts returns the style's Fonts.
func (r RunProperties) Fonts() Fonts {
	if r.x.RFonts == nil {
		r.x.RFonts = wml.NewCT_Fonts()
	}
	return Fonts{r.x.RFonts}
}

// Color returns the style's Color.
func (r RunProperties) Color() Color {
	if r.x.Color == nil {
		r.x.Color = wml.NewCT_Color()
	}
	return Color{r.x.Color}
}

// SetUnderline controls underline for a run style.
func (r RunProperties) SetUnderline(style wml.ST_Underline, c color.Color) {
	if style == wml.ST_UnderlineUnset {
		r.x.U = nil
	} else {
		r.x.U = wml.NewCT_Underline()
		r.x.U.ColorAttr = &wml.ST_HexColor{}
		r.x.U.ColorAttr.ST_HexColorRGB = c.AsRGBString()
		r.x.U.ValAttr = style
	}
}

// IsBold returns true if the run has been set to bold.
func (r RunProperties) IsBold() bool {
	return r.x.B != nil
}

// SetBold sets the run to bold.
func (r RunProperties) SetBold(b bool) {
	if !b {
		r.x.B = nil
		r.x.BCs = nil
	} else {
		r.x.B = wml.NewCT_OnOff()
		r.x.BCs = wml.NewCT_OnOff()
	}
}

// IsItalic returns true if the run was set to bold.
func (r RunProperties) IsItalic() bool {
	if r.x == nil {
		return false
	}
	return r.x.I != nil
}

// SetItalic sets the run to italic.
func (r RunProperties) SetItalic(b bool) {
	if !b {
		r.x.I = nil
		r.x.ICs = nil
	} else {
		r.x.I = wml.NewCT_OnOff()
		r.x.ICs = wml.NewCT_OnOff()
	}
}

// SetAllCaps sets the run to all caps.
func (r RunProperties) SetAllCaps(b bool) {
	if !b {
		r.x.Caps = nil
	} else {
		r.x.Caps = wml.NewCT_OnOff()
	}
}

// SetSmallCaps sets the run to small caps.
func (r RunProperties) SetSmallCaps(b bool) {
	if !b {
		r.x.SmallCaps = nil
	} else {
		r.x.SmallCaps = wml.NewCT_OnOff()
	}
}

// SetStrikeThrough sets the run to strike-through.
func (r RunProperties) SetStrikeThrough(b bool) {
	if !b {
		r.x.Strike = nil
	} else {
		r.x.Strike = wml.NewCT_OnOff()
	}
}

// SetDoubleStrikeThrough sets the run to double strike-through.
func (r RunProperties) SetDoubleStrikeThrough(b bool) {
	if !b {
		r.x.Dstrike = nil
	} else {
		r.x.Dstrike = wml.NewCT_OnOff()
	}
}

// SetOutline sets the run to outlined text.
func (r RunProperties) SetOutline(b bool) {
	if !b {
		r.x.Outline = nil
	} else {
		r.x.Outline = wml.NewCT_OnOff()
	}
}

// SetShadow sets the run to shadowed text.
func (r RunProperties) SetShadow(b bool) {
	if !b {
		r.x.Shadow = nil
	} else {
		r.x.Shadow = wml.NewCT_OnOff()
	}
}

// SetEmboss sets the run to embossed text.
func (r RunProperties) SetEmboss(b bool) {
	if !b {
		r.x.Emboss = nil
	} else {
		r.x.Emboss = wml.NewCT_OnOff()
	}
}

// SetImprint sets the run to imprinted text.
func (r RunProperties) SetImprint(b bool) {
	if !b {
		r.x.Imprint = nil
	} else {
		r.x.Imprint = wml.NewCT_OnOff()
	}
}

// ClearColor clears the text color.
func (r RunProperties) ClearColor() {
	r.x.Color = nil
}

// SetColor sets the text color.
func (r RunProperties) SetColor(c color.Color) {
	r.x.Color = wml.NewCT_Color()
	r.x.Color.ValAttr.ST_HexColorRGB = c.AsRGBString()
}

// SetHighlight highlights text in a specified color.
func (r RunProperties) SetHighlight(c wml.ST_HighlightColor) {
	r.x.Highlight = wml.NewCT_Highlight()
	r.x.Highlight.ValAttr = c
}

// SetEffect sets a text effect on the run.
func (r RunProperties) SetEffect(e wml.ST_TextEffect) {
	if e == wml.ST_TextEffectUnset {
		r.x.Effect = nil
	} else {
		r.x.Effect = wml.NewCT_TextEffect()
		r.x.Effect.ValAttr = wml.ST_TextEffectShimmer
	}
}

// SetVerticalAlignment controls the vertical alignment of the run, this is used
// to control if text is superscript/subscript.
func (r RunProperties) SetVerticalAlignment(v sharedTypes.ST_VerticalAlignRun) {
	if v == sharedTypes.ST_VerticalAlignRunUnset {
		r.x.VertAlign = nil
	} else {
		r.x.VertAlign = wml.NewCT_VerticalAlignRun()
		r.x.VertAlign.ValAttr = v
	}
}
