// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package document

import (
	"fmt"

	"baliance.com/gooxml"
	"baliance.com/gooxml/measurement"
	"baliance.com/gooxml/schema/schemas.openxmlformats.org/officeDocument/2006/sharedTypes"
	wml "baliance.com/gooxml/schema/schemas.openxmlformats.org/wordprocessingml"
)

// Styles is the document wide styles contained in styles.xml.
type Styles struct {
	x *wml.Styles
}

// NewStyles constructs a new empty Styles
func NewStyles() Styles {
	return Styles{wml.NewStyles()}
}

// X returns the inner wrapped XML type.
func (s Styles) X() *wml.Styles {
	return s.x
}

// Clear clears the styes.
func (s Styles) Clear() {
	s.x.DocDefaults = nil
	s.x.LatentStyles = nil
	s.x.Style = nil
}

// AddStyle adds a new empty style.
func (s Styles) AddStyle(styleID string, t wml.ST_StyleType, isDefault bool) Style {
	ss := wml.NewCT_Style()
	ss.TypeAttr = t
	if isDefault {
		ss.DefaultAttr = &sharedTypes.ST_OnOff{}
		ss.DefaultAttr.Bool = gooxml.Bool(isDefault)
	}
	ss.StyleIdAttr = gooxml.String(styleID)
	s.x.Style = append(s.x.Style, ss)
	return Style{ss}
}

// InitializeDefault constructs the default styles.
func (s Styles) InitializeDefault() {
	s.initializeDocDefaults()
	// Normal
	normal := s.AddStyle("Normal", wml.ST_StyleTypeParagraph, true)
	normal.SetName("Normal")
	normal.SetPrimaryStyle(true)

	// DefaultParagraphFont
	dpf := s.AddStyle("DefaultParagraphFont", wml.ST_StyleTypeCharacter, true)
	dpf.SetName("Default Paragraph Font")
	dpf.SetUISortOrder(1)
	dpf.SetSemiHidden(true)
	dpf.SetUnhideWhenUsed(true)

	// TableNormal
	tbl := s.AddStyle("TableNormal", wml.ST_StyleTypeTable, true)
	tbl.SetName("Normal Table")
	tbl.SetUISortOrder(99)
	tbl.SetSemiHidden(true)
	tbl.SetUnhideWhenUsed(true)
	tbl.X().TblPr = wml.NewCT_TblPrBase()

	tw := NewTableWidth()
	tbl.X().TblPr.TblInd = tw.X()
	tw.SetValue(0 * measurement.Dxa)

	tbl.X().TblPr.TblCellMar = wml.NewCT_TblCellMar()

	tw = NewTableWidth()
	tbl.X().TblPr.TblCellMar.Top = tw.X()
	tw.SetValue(0 * measurement.Dxa)

	tw = NewTableWidth()
	tbl.X().TblPr.TblCellMar.Bottom = tw.X()
	tw.SetValue(0 * measurement.Dxa)

	tw = NewTableWidth()
	tbl.X().TblPr.TblCellMar.Left = tw.X()
	tw.SetValue(108 * measurement.Dxa)

	tw = NewTableWidth()
	tbl.X().TblPr.TblCellMar.Right = tw.X()
	tw.SetValue(108 * measurement.Dxa)

	// NoList
	nbr := s.AddStyle("NoList", wml.ST_StyleTypeNumbering, true)
	nbr.SetName("No List")
	nbr.SetUISortOrder(1)
	nbr.SetSemiHidden(true)
	nbr.SetUnhideWhenUsed(true)

	// HeaderChar
	hc := s.AddStyle("HeaderChar", wml.ST_StyleTypeCharacter, true)
	hc.SetName("Header Char")
	hc.SetBasedOn(dpf.StyleID())
	hc.SetLinkedStyle("Header")
	hc.SetUISortOrder(99)

	// Header
	hdr := s.AddStyle("Header", wml.ST_StyleTypeParagraph, true)
	hdr.SetName("header")
	hdr.SetBasedOn(normal.StyleID())
	hdr.SetUISortOrder(1)
	hdr.SetSemiHidden(true)
	hdr.SetUnhideWhenUsed(true)
	hdr.SetLinkedStyle(hc.StyleID())
	hdr.ParagraphStyleProperties().AddTabStop(4680*measurement.Twips, wml.ST_TabJcCenter, wml.ST_TabTlcUnset)
	hdr.ParagraphStyleProperties().AddTabStop(9360*measurement.Twips, wml.ST_TabJcRight, wml.ST_TabTlcUnset)
	hdr.ParagraphStyleProperties().SetSpacing(0, 240*measurement.Twips)

	// FooterChar
	fc := s.AddStyle("FooterChar", wml.ST_StyleTypeCharacter, true)
	fc.SetName("Footer Char")
	fc.SetBasedOn(dpf.StyleID())
	fc.SetLinkedStyle("Footer")
	fc.SetUISortOrder(99)

	// Footer
	ftr := s.AddStyle("Footer", wml.ST_StyleTypeParagraph, true)
	ftr.SetName("footer")
	ftr.SetBasedOn(normal.StyleID())
	ftr.SetUISortOrder(1)
	ftr.SetSemiHidden(true)
	ftr.SetUnhideWhenUsed(true)
	ftr.SetLinkedStyle(fc.StyleID())
	ftr.ParagraphStyleProperties().AddTabStop(4680*measurement.Twips, wml.ST_TabJcCenter, wml.ST_TabTlcUnset)
	ftr.ParagraphStyleProperties().AddTabStop(9360*measurement.Twips, wml.ST_TabJcRight, wml.ST_TabTlcUnset)
	ftr.ParagraphStyleProperties().SetSpacing(0, 240*measurement.Twips)

	fontSizes := []measurement.Distance{16, 13, 12, 11, 11, 11, 11, 11, 11}
	spacing := []measurement.Distance{240, 40, 40, 40, 40, 40, 40, 40, 40}
	for i := 0; i < 9; i++ {
		id := fmt.Sprintf("Heading%d", i+1)

		hdngChar := s.AddStyle(id+"Char", wml.ST_StyleTypeCharacter, false)
		hdngChar.SetName(fmt.Sprintf("Heading %d Char", i+1))
		hdngChar.SetBasedOn(dpf.StyleID())
		hdngChar.SetLinkedStyle(id)
		hdngChar.SetUISortOrder(9 + i)
		hdngChar.RunStyle().SetSize(fontSizes[i] * measurement.Point)

		hdng := s.AddStyle(id, wml.ST_StyleTypeParagraph, false)
		hdng.SetName(fmt.Sprintf("heading %d", i+1))
		hdng.SetNextStyle(normal.StyleID())
		hdng.SetLinkedStyle(hdng.StyleID())
		hdng.SetUISortOrder(9 + i)
		hdng.SetPrimaryStyle(true)
		hdng.ParagraphStyleProperties().SetKeepNext(true)
		hdng.ParagraphStyleProperties().SetSpacing(spacing[i]*measurement.Twips, 0)
		hdng.ParagraphStyleProperties().SetOutlineLevel(i)
		hdng.RunStyle().SetSize(fontSizes[i] * measurement.Point)
	}

}

func (s Styles) initializeDocDefaults() {
	s.x.DocDefaults = wml.NewCT_DocDefaults()
	s.x.DocDefaults.RPrDefault = wml.NewCT_RPrDefault()
	s.x.DocDefaults.RPrDefault.RPr = wml.NewCT_RPr()

	base := wml.NewEG_RPrBase()
	s.x.DocDefaults.RPrDefault.RPr.EG_RPrBase = append(s.x.DocDefaults.RPrDefault.RPr.EG_RPrBase, base)
	base.RFonts = wml.NewCT_Fonts()
	base.RFonts.AsciiThemeAttr = wml.ST_ThemeMinorHAnsi
	base.RFonts.EastAsiaThemeAttr = wml.ST_ThemeMinorHAnsi
	base.RFonts.HAnsiThemeAttr = wml.ST_ThemeMinorHAnsi
	base.RFonts.CsthemeAttr = wml.ST_ThemeMinorBidi

	base = wml.NewEG_RPrBase()
	s.x.DocDefaults.RPrDefault.RPr.EG_RPrBase = append(s.x.DocDefaults.RPrDefault.RPr.EG_RPrBase, base)
	base.Sz = wml.NewCT_HpsMeasure()
	base.Sz.ValAttr.ST_UnsignedDecimalNumber = gooxml.Uint64(22)

	base = wml.NewEG_RPrBase()
	s.x.DocDefaults.RPrDefault.RPr.EG_RPrBase = append(s.x.DocDefaults.RPrDefault.RPr.EG_RPrBase, base)
	base.SzCs = wml.NewCT_HpsMeasure()
	base.SzCs.ValAttr.ST_UnsignedDecimalNumber = gooxml.Uint64(22)

	base = wml.NewEG_RPrBase()
	s.x.DocDefaults.RPrDefault.RPr.EG_RPrBase = append(s.x.DocDefaults.RPrDefault.RPr.EG_RPrBase, base)
	base.Lang = wml.NewCT_Language()
	base.Lang.ValAttr = gooxml.String("en-us")
	base.Lang.EastAsiaAttr = gooxml.String("en-us")
	base.Lang.BidiAttr = gooxml.String("ar-SA")

	s.x.DocDefaults.PPrDefault = wml.NewCT_PPrDefault()
	s.x.DocDefaults.PPrDefault.PPr = wml.NewCT_PPrGeneral()
	s.x.DocDefaults.PPrDefault.PPr.Spacing = wml.NewCT_Spacing()
	s.x.DocDefaults.PPrDefault.PPr.Spacing.AfterAttr = &sharedTypes.ST_TwipsMeasure{}
	s.x.DocDefaults.PPrDefault.PPr.Spacing.AfterAttr.ST_UnsignedDecimalNumber = gooxml.Uint64(160)
	s.x.DocDefaults.PPrDefault.PPr.Spacing.LineAttr = &wml.ST_SignedTwipsMeasure{}
	s.x.DocDefaults.PPrDefault.PPr.Spacing.LineAttr.Int64 = gooxml.Int64(259)
	s.x.DocDefaults.PPrDefault.PPr.Spacing.LineRuleAttr = wml.ST_LineSpacingRuleAuto
}

// Styles returns all styles.
func (s Styles) Styles() []Style {
	ret := []Style{}
	for _, s := range s.x.Style {
		ret = append(ret, Style{s})
	}
	return ret
}

// ParagraphStyles returns only the paragraph styles.
func (s Styles) ParagraphStyles() []Style {
	ret := []Style{}
	for _, s := range s.x.Style {
		if s.TypeAttr != wml.ST_StyleTypeParagraph {
			continue
		}
		ret = append(ret, Style{s})
	}
	return ret
}
