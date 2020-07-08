// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased on https://unidoc.io.

package document

import (
	"github.com/unidoc/unioffice"
	"github.com/unidoc/unioffice/color"
	"github.com/unidoc/unioffice/measurement"
	"github.com/unidoc/unioffice/schema/soo/ofc/sharedTypes"
	"github.com/unidoc/unioffice/schema/soo/wml"
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
	r.x.RFonts.AsciiAttr = unioffice.String(family)
	r.x.RFonts.HAnsiAttr = unioffice.String(family)
	r.x.RFonts.EastAsiaAttr = unioffice.String(family)
}

// SetSize sets the font size for a run.
func (r RunProperties) SetSize(size measurement.Distance) {
	r.x.Sz = wml.NewCT_HpsMeasure()
	r.x.Sz.ValAttr.ST_UnsignedDecimalNumber = unioffice.Uint64(uint64(size / measurement.HalfPoint))
	r.x.SzCs = wml.NewCT_HpsMeasure()
	r.x.SzCs.ValAttr.ST_UnsignedDecimalNumber = unioffice.Uint64(uint64(size / measurement.HalfPoint))
}

// SetKerning sets the run's font kerning.
func (r RunProperties) SetKerning(size measurement.Distance) {
	r.x.Kern = wml.NewCT_HpsMeasure()
	r.x.Kern.ValAttr.ST_UnsignedDecimalNumber = unioffice.Uint64(uint64(size / measurement.HalfPoint))
}

// SetCharacterSpacing sets the run's Character Spacing Adjustment.
func (r RunProperties) SetCharacterSpacing(size measurement.Distance) {
	r.x.Spacing = wml.NewCT_SignedTwipsMeasure()
	r.x.Spacing.ValAttr.Int64 = unioffice.Int64(int64(size / measurement.Twips))
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

// BoldValue returns the precise nature of the bold setting (unset, off or on).
func (r RunProperties) BoldValue() OnOffValue {
	return convertOnOff(r.x.B)
}

// IsBold returns true if the run has been set to bold.
func (r RunProperties) IsBold() bool {
	return r.BoldValue() == OnOffValueOn
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

// ItalicValue returns the precise nature of the italic setting (unset, off or on).
func (r RunProperties) ItalicValue() OnOffValue {
	return convertOnOff(r.x.I)
}

// IsItalic returns true if the run has been set to italics.
func (r RunProperties) IsItalic() bool {
	return r.ItalicValue() == OnOffValueOn
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

// Bold returns true if run font is bold.
func (r RunProperties) Bold() bool {
	x := r.x
	return getBool(x.B) || getBool(x.BCs)
}

// Italic returns true if run font is italic.
func (r RunProperties) Italic() bool {
	x := r.x
	return getBool(x.I) || getBool(x.ICs)
}

// Caps returns true if run font is capitalized.
func (r RunProperties) Caps() bool {
	return getBool(r.x.Caps)
}

// Strike returns true if run is striked.
func (r RunProperties) Strike() bool {
	return getBool(r.x.Strike)
}

// DoubleStrike returns true if run is double striked.
func (r RunProperties) DoubleStrike() bool {
	return getBool(r.x.Dstrike)
}

// Outline returns true if run outline is on.
func (r RunProperties) Outline() bool {
	return getBool(r.x.Outline)
}

// Shadow returns true if run shadow is on.
func (r RunProperties) Shadow() bool {
	return getBool(r.x.Shadow)
}

// Emboss returns true if run emboss is on.
func (r RunProperties) Emboss() bool {
	return getBool(r.x.Emboss)
}

// RightToLeft returns true if run text goes from right to left.
func (r RunProperties) RightToLeft() bool {
	return getBool(r.x.Rtl)
}

// RStyle returns the name of character style.
// It is defined here http://officeopenxml.com/WPstyleCharStyles.php
func (r RunProperties) RStyle() string {
	if r.x.RStyle != nil {
		return r.x.RStyle.ValAttr
	}
	return ""
}

// Font returns the name of run font family.
func (r RunProperties) Font() string {
	if fonts := r.x.RFonts; fonts != nil {
		if fonts.AsciiAttr != nil {
			return *fonts.AsciiAttr
		} else if fonts.HAnsiAttr != nil {
			return *fonts.HAnsiAttr
		} else if fonts.CsAttr != nil {
			return *fonts.CsAttr
		}
	}
	return ""
}

// EastAsiaFont returns the name of run font family for East Asia.
func (r RunProperties) EastAsiaFont() string {
	if fonts := r.x.RFonts; fonts != nil {
		if fonts.EastAsiaAttr != nil {
			return *fonts.EastAsiaAttr
		}
	}
	return ""
}

// GetColor returns the color.Color object representing the run color.
func (r RunProperties) GetColor() color.Color {
	if c := r.x.Color; c != nil {
		valAttr := c.ValAttr
		if valAttr.ST_HexColorRGB != nil {
			return color.FromHex(*valAttr.ST_HexColorRGB)
		}
	}
	return color.Color{}
}

// CharacterSpacingValue returns the value of run's characters spacing in twips (1/20 of point).
func (r RunProperties) CharacterSpacingValue() int64 {
	if spacing := r.x.Spacing; spacing != nil {
		valAttr := spacing.ValAttr
		if valAttr.Int64 != nil {
			return *valAttr.Int64
		}
	}
	return int64(0)
}

// CharacterSpacingMeasure returns paragraph characters spacing with its measure which can be mm, cm, in, pt, pc or pi.
func (r RunProperties) CharacterSpacingMeasure() string {
	if spacing := r.x.Spacing; spacing != nil {
		valAttr := spacing.ValAttr
		if valAttr.ST_UniversalMeasure != nil {
			return *valAttr.ST_UniversalMeasure
		}
	}
	return ""
}

// SizeValue returns the value of run font size in points.
func (r RunProperties) SizeValue() float64 {
	if sz := r.x.Sz; sz != nil {
		valAttr := sz.ValAttr
		if valAttr.ST_UnsignedDecimalNumber != nil {
			return float64(*valAttr.ST_UnsignedDecimalNumber) / 2
		}
	}
	return 0.0
}

// SizeMeasure returns font with its measure which can be mm, cm, in, pt, pc or pi.
func (r RunProperties) SizeMeasure() string {
	if sz := r.x.Sz; sz != nil {
		valAttr := sz.ValAttr
		if valAttr.ST_PositiveUniversalMeasure != nil {
			return *valAttr.ST_PositiveUniversalMeasure
		}
	}
	return ""
}

// ComplexSizeValue returns the value of run font size for complex fonts in points.
func (r RunProperties) ComplexSizeValue() float64 {
	if szCs := r.x.SzCs; szCs != nil {
		valAttr := szCs.ValAttr
		if valAttr.ST_UnsignedDecimalNumber != nil {
			return float64(*valAttr.ST_UnsignedDecimalNumber) / 2
		}
	}
	return 0.0
}

// ComplexSizeMeasure returns font with its measure which can be mm, cm, in, pt, pc or pi.
func (r RunProperties) ComplexSizeMeasure() string {
	if szCs := r.x.SzCs; szCs != nil {
		valAttr := szCs.ValAttr
		if valAttr.ST_PositiveUniversalMeasure != nil {
			return *valAttr.ST_PositiveUniversalMeasure
		}
	}
	return ""
}

// Underline returns the type of run underline.
func (r RunProperties) Underline() wml.ST_Underline {
	if underline := r.x.U; underline != nil {
		return underline.ValAttr
	}
	return 0
}

// UnderlineColor returns the hex color value of run underline.
func (r RunProperties) UnderlineColor() string {
	if underline := r.x.U; underline != nil {
		color := underline.ColorAttr
		if color != nil && color.ST_HexColorRGB != nil {
			return *color.ST_HexColorRGB
		}
	}
	return ""
}

// VerticalAlign returns the value of run vertical align.
func (r RunProperties) VerticalAlignment() sharedTypes.ST_VerticalAlignRun {
	if vertAlign := r.x.VertAlign; vertAlign != nil {
		return vertAlign.ValAttr
	}
	return 0
}
