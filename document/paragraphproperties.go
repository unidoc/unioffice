// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased on https://unidoc.io.

package document

import (
	"fmt"

	"github.com/unidoc/unioffice"
	"github.com/unidoc/unioffice/color"
	"github.com/unidoc/unioffice/measurement"
	"github.com/unidoc/unioffice/schema/soo/ofc/sharedTypes"
	"github.com/unidoc/unioffice/schema/soo/wml"
)

// ParagraphProperties are the properties for a paragraph.
type ParagraphProperties struct {
	d *Document
	x *wml.CT_PPr
}

// X returns the inner wrapped XML type.
func (p ParagraphProperties) X() *wml.CT_PPr {
	return p.x
}

// SetSpacing sets the spacing that comes before and after the paragraph.
// Deprecated: See Spacing() instead which allows finer control.
func (p ParagraphProperties) SetSpacing(before, after measurement.Distance) {
	if p.x.Spacing == nil {
		p.x.Spacing = wml.NewCT_Spacing()
	}
	p.x.Spacing.BeforeAttr = &sharedTypes.ST_TwipsMeasure{}
	p.x.Spacing.BeforeAttr.ST_UnsignedDecimalNumber = unioffice.Uint64(uint64(before / measurement.Twips))
	p.x.Spacing.AfterAttr = &sharedTypes.ST_TwipsMeasure{}
	p.x.Spacing.AfterAttr.ST_UnsignedDecimalNumber = unioffice.Uint64(uint64(after / measurement.Twips))
}

// Spacing returns the paragraph spacing settings.
func (p ParagraphProperties) Spacing() ParagraphSpacing {
	if p.x.Spacing == nil {
		p.x.Spacing = wml.NewCT_Spacing()
	}
	return ParagraphSpacing{p.x.Spacing}
}

// SetAlignment controls the paragraph alignment
func (p ParagraphProperties) SetAlignment(align wml.ST_Jc) {
	if align == wml.ST_JcUnset {
		p.x.Jc = nil
	} else {
		p.x.Jc = wml.NewCT_Jc()
		p.x.Jc.ValAttr = align
	}
}

// Style returns the style for a paragraph, or an empty string if it is unset.
func (p ParagraphProperties) Style() string {
	if p.x.PStyle != nil {
		return p.x.PStyle.ValAttr
	}
	return ""
}

// SetStyle sets the style of a paragraph.
func (p ParagraphProperties) SetStyle(s string) {
	if s == "" {
		p.x.PStyle = nil
	} else {
		p.x.PStyle = wml.NewCT_String()
		p.x.PStyle.ValAttr = s
	}
}

// AddTabStop adds a tab stop to the paragraph.  It controls the position of text when using Run.AddTab()
func (p ParagraphProperties) AddTabStop(position measurement.Distance, justificaton wml.ST_TabJc, leader wml.ST_TabTlc) {
	if p.x.Tabs == nil {
		p.x.Tabs = wml.NewCT_Tabs()
	}
	tab := wml.NewCT_TabStop()
	tab.LeaderAttr = leader
	tab.ValAttr = justificaton
	tab.PosAttr.Int64 = unioffice.Int64(int64(position / measurement.Twips))
	p.x.Tabs.Tab = append(p.x.Tabs.Tab, tab)
}

// AddSection adds a new document section with an optional section break.  If t
// is ST_SectionMarkUnset, then no break will be inserted.
func (p ParagraphProperties) AddSection(t wml.ST_SectionMark) Section {
	p.x.SectPr = wml.NewCT_SectPr()
	if t != wml.ST_SectionMarkUnset {
		p.x.SectPr.Type = wml.NewCT_SectType()
		p.x.SectPr.Type.ValAttr = t
	}
	return Section{p.d, p.x.SectPr}
}

// SetHeadingLevel sets a heading level and style based on the level to a
// paragraph.  The default styles for a new gooxml document support headings
// from level 1 to 8.
func (p ParagraphProperties) SetHeadingLevel(idx int) {
	p.SetStyle(fmt.Sprintf("Heading%d", idx))
	if p.x.NumPr == nil {
		p.x.NumPr = wml.NewCT_NumPr()
	}
	p.x.NumPr.Ilvl = wml.NewCT_DecimalNumber()
	p.x.NumPr.Ilvl.ValAttr = int64(idx)
}

// SetKeepWithNext controls if this paragraph should be kept with the next.
func (p ParagraphProperties) SetKeepWithNext(b bool) {
	if !b {
		p.x.KeepNext = nil
	} else {
		p.x.KeepNext = wml.NewCT_OnOff()
	}
}

// SetKeepOnOnePage controls if all lines in a paragraph are kept on the same
// page.
func (p ParagraphProperties) SetKeepOnOnePage(b bool) {
	if !b {
		p.x.KeepLines = nil
	} else {
		p.x.KeepLines = wml.NewCT_OnOff()
	}
}

// SetPageBreakBefore controls if there is a page break before this paragraph.
func (p ParagraphProperties) SetPageBreakBefore(b bool) {
	if !b {
		p.x.PageBreakBefore = nil
	} else {
		p.x.PageBreakBefore = wml.NewCT_OnOff()
	}
}

// SetWindowControl controls if the first or last line of the paragraph is
// allowed to dispay on a separate page.
func (p ParagraphProperties) SetWindowControl(b bool) {
	if !b {
		p.x.WidowControl = nil
	} else {
		p.x.WidowControl = wml.NewCT_OnOff()
	}
}

// SetFirstLineIndent controls the indentation of the first line in a paragraph.
func (p ParagraphProperties) SetFirstLineIndent(m measurement.Distance) {
	if p.x.Ind == nil {
		p.x.Ind = wml.NewCT_Ind()
	}
	if m == measurement.Zero {
		p.x.Ind.FirstLineAttr = nil
	} else {
		p.x.Ind.FirstLineAttr = &sharedTypes.ST_TwipsMeasure{}
		p.x.Ind.FirstLineAttr.ST_UnsignedDecimalNumber = unioffice.Uint64(uint64(m / measurement.Twips))
	}
}

// SetStartIndent controls the start indentation.
func (p ParagraphProperties) SetStartIndent(m measurement.Distance) {
	if p.x.Ind == nil {
		p.x.Ind = wml.NewCT_Ind()
	}
	if m == measurement.Zero {
		p.x.Ind.StartAttr = nil
	} else {
		p.x.Ind.StartAttr = &wml.ST_SignedTwipsMeasure{}
		p.x.Ind.StartAttr.Int64 = unioffice.Int64(int64(m / measurement.Twips))
	}
}

// SetEndIndent controls the end indentation.
func (p ParagraphProperties) SetEndIndent(m measurement.Distance) {
	if p.x.Ind == nil {
		p.x.Ind = wml.NewCT_Ind()
	}
	if m == measurement.Zero {
		p.x.Ind.EndAttr = nil
	} else {
		p.x.Ind.EndAttr = &wml.ST_SignedTwipsMeasure{}
		p.x.Ind.EndAttr.Int64 = unioffice.Int64(int64(m / measurement.Twips))
	}
}

// SetHangingIndent controls the indentation of the non-first lines in a paragraph.
func (p ParagraphProperties) SetHangingIndent(m measurement.Distance) {
	if p.x.Ind == nil {
		p.x.Ind = wml.NewCT_Ind()
	}
	if m == measurement.Zero {
		p.x.Ind.HangingAttr = nil
	} else {
		p.x.Ind.HangingAttr = &sharedTypes.ST_TwipsMeasure{}
		p.x.Ind.HangingAttr.ST_UnsignedDecimalNumber = unioffice.Uint64(uint64(m / measurement.Twips))
	}
}

// Bold returns true if paragraph font is bold.
func (p ParagraphProperties) Bold() bool {
	x := p.x.RPr
	return getBool(x.B) || getBool(x.BCs)
}

// Italic returns true if paragraph font is italic.
func (p ParagraphProperties) Italic() bool {
	x := p.x.RPr
	return getBool(x.I) || getBool(x.ICs)
}

// Caps returns true if paragraph font is capitalized.
func (p ParagraphProperties) Caps() bool {
	return getBool(p.x.RPr.Caps)
}

// Strike returns true if paragraph is striked.
func (p ParagraphProperties) Strike() bool {
	return getBool(p.x.RPr.Strike)
}

// DoubleStrike returns true if paragraph is double striked.
func (p ParagraphProperties) DoubleStrike() bool {
	return getBool(p.x.RPr.Dstrike)
}

// Outline returns true if paragraph outline is on.
func (p ParagraphProperties) Outline() bool {
	return getBool(p.x.RPr.Outline)
}

// Shadow returns true if paragraph shadow is on.
func (p ParagraphProperties) Shadow() bool {
	return getBool(p.x.RPr.Shadow)
}

// Emboss returns true if paragraph emboss is on.
func (p ParagraphProperties) Emboss() bool {
	return getBool(p.x.RPr.Emboss)
}

// RightToLeft returns true if paragraph text goes from right to left.
func (p ParagraphProperties) RightToLeft() bool {
	return getBool(p.x.RPr.Rtl)
}

// RStyle returns the name of character style.
// It is defined here http://officeopenxml.com/WPstyleCharStyles.php
func (p ParagraphProperties) RStyle() string {
	if p.x.RPr.RStyle != nil {
		return p.x.RPr.RStyle.ValAttr
	}
	return ""
}

// Font returns the name of paragraph font family.
func (p ParagraphProperties) Font() string {
	if fonts := p.x.RPr.RFonts; fonts != nil {
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

// EastAsiaFont returns the name of paragraph font family for East Asia.
func (p ParagraphProperties) EastAsiaFont() string {
	if fonts := p.x.RPr.RFonts; fonts != nil {
		if fonts.EastAsiaAttr != nil {
			return *fonts.EastAsiaAttr
		}
	}
	return ""
}

// GetColor returns the color.Color object representing the run color.
func (p ParagraphProperties) GetColor() color.Color {
	if c := p.x.RPr.Color; c != nil {
		valAttr := c.ValAttr
		if valAttr.ST_HexColorRGB != nil {
			return color.FromHex(*valAttr.ST_HexColorRGB)
		}
	}
	return color.Color{}
}

// CharacterSpacingValue returns the value of characters spacing in twips (1/20 of point).
func (p ParagraphProperties) CharacterSpacingValue() int64 {
	if spacing := p.x.RPr.Spacing; spacing != nil {
		valAttr := spacing.ValAttr
		if valAttr.Int64 != nil {
			return *valAttr.Int64
		}
	}
	return int64(0)
}

// CharacterSpacingMeasure returns paragraph characters spacing with its measure which can be mm, cm, in, pt, pc or pi.
func (p ParagraphProperties) CharacterSpacingMeasure() string {
	if spacing := p.x.RPr.Spacing; spacing != nil {
		valAttr := spacing.ValAttr
		if valAttr.ST_UniversalMeasure != nil {
			return *valAttr.ST_UniversalMeasure
		}
	}
	return ""
}

// SizeValue returns the value of paragraph font size in points.
func (p ParagraphProperties) SizeValue() float64 {
	if sz := p.x.RPr.Sz; sz != nil {
		valAttr := sz.ValAttr
		if valAttr.ST_UnsignedDecimalNumber != nil {
			return float64(*valAttr.ST_UnsignedDecimalNumber) / 2
		}
	}
	return 0.0
}

// SizeMeasure returns font with its measure which can be mm, cm, in, pt, pc or pi.
func (p ParagraphProperties) SizeMeasure() string {
	if sz := p.x.RPr.Sz; sz != nil {
		valAttr := sz.ValAttr
		if valAttr.ST_PositiveUniversalMeasure != nil {
			return *valAttr.ST_PositiveUniversalMeasure
		}
	}
	return ""
}

// ComplexSizeValue returns the value of paragraph font size for complex fonts in points.
func (p ParagraphProperties) ComplexSizeValue() float64 {
	if szCs := p.x.RPr.SzCs; szCs != nil {
		valAttr := szCs.ValAttr
		if valAttr.ST_UnsignedDecimalNumber != nil {
			return float64(*valAttr.ST_UnsignedDecimalNumber) / 2
		}
	}
	return 0.0
}

// ComplexSizeMeasure returns font with its measure which can be mm, cm, in, pt, pc or pi.
func (p ParagraphProperties) ComplexSizeMeasure() string {
	if szCs := p.x.RPr.SzCs; szCs != nil {
		valAttr := szCs.ValAttr
		if valAttr.ST_PositiveUniversalMeasure != nil {
			return *valAttr.ST_PositiveUniversalMeasure
		}
	}
	return ""
}

// Underline returns the type of paragraph underline.
func (p ParagraphProperties) Underline() wml.ST_Underline {
	if underline := p.x.RPr.U; underline != nil {
		return underline.ValAttr
	}
	return 0
}

// UnderlineColor returns the hex color value of paragraph underline.
func (p ParagraphProperties) UnderlineColor() string {
	if underline := p.x.RPr.U; underline != nil {
		color := underline.ColorAttr
		if color != nil && color.ST_HexColorRGB != nil {
			return *color.ST_HexColorRGB
		}
	}
	return ""
}

// VerticalAlign returns the value of paragraph vertical align.
func (p ParagraphProperties) VerticalAlignment() sharedTypes.ST_VerticalAlignRun {
	if vertAlign := p.x.RPr.VertAlign; vertAlign != nil {
		return vertAlign.ValAttr
	}
	return 0
}
