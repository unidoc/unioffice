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
	"baliance.com/gooxml/schema/soo/ofc/sharedTypes"
	"baliance.com/gooxml/schema/soo/wml"
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
	s.initializeStyleDefaults()
}
func (s Styles) initializeStyleDefaults() {
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

	// Title
	titleChar := s.AddStyle("TitleChar", wml.ST_StyleTypeCharacter, false)
	titleChar.SetName("Title Char")
	titleChar.SetBasedOn(dpf.StyleID())
	titleChar.SetLinkedStyle("Title")
	titleChar.SetUISortOrder(10)
	titleChar.RunProperties().Fonts().SetASCIITheme(wml.ST_ThemeMajorAscii)
	titleChar.RunProperties().Fonts().SetEastAsiaTheme(wml.ST_ThemeMajorEastAsia)
	titleChar.RunProperties().Fonts().SetHANSITheme(wml.ST_ThemeMajorHAnsi)
	titleChar.RunProperties().Fonts().SetCSTheme(wml.ST_ThemeMajorBidi)
	titleChar.RunProperties().SetSize(28 * measurement.Point)
	titleChar.RunProperties().SetKerning(14 * measurement.Point)
	titleChar.RunProperties().SetCharacterSpacing(-10 * measurement.Twips)

	titlePara := s.AddStyle("Title", wml.ST_StyleTypeParagraph, false)
	titlePara.SetName("Title")
	titlePara.SetBasedOn(normal.StyleID())
	titlePara.SetNextStyle(normal.StyleID())
	titlePara.SetLinkedStyle(titleChar.StyleID())
	titlePara.SetUISortOrder(10)
	titlePara.SetPrimaryStyle(true)
	titlePara.ParagraphProperties().SetContextualSpacing(true)
	titlePara.RunProperties().Fonts().SetASCIITheme(wml.ST_ThemeMajorAscii)
	titlePara.RunProperties().Fonts().SetEastAsiaTheme(wml.ST_ThemeMajorEastAsia)
	titlePara.RunProperties().Fonts().SetHANSITheme(wml.ST_ThemeMajorHAnsi)
	titlePara.RunProperties().Fonts().SetCSTheme(wml.ST_ThemeMajorBidi)
	titlePara.RunProperties().SetSize(28 * measurement.Point)
	titlePara.RunProperties().SetKerning(14 * measurement.Point)
	titlePara.RunProperties().SetCharacterSpacing(-10 * measurement.Twips)

	// TableNormal
	tbl := s.AddStyle("TableNormal", wml.ST_StyleTypeTable, false)
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
	nbr := s.AddStyle("NoList", wml.ST_StyleTypeNumbering, false)
	nbr.SetName("No List")
	nbr.SetUISortOrder(1)
	nbr.SetSemiHidden(true)
	nbr.SetUnhideWhenUsed(true)

	fontSizes := []measurement.Distance{16, 13, 12, 11, 11, 11, 11, 11, 11}
	spacing := []measurement.Distance{240, 40, 40, 40, 40, 40, 40, 40, 40}
	for i := 0; i < 9; i++ {
		id := fmt.Sprintf("Heading%d", i+1)

		hdngChar := s.AddStyle(id+"Char", wml.ST_StyleTypeCharacter, false)
		hdngChar.SetName(fmt.Sprintf("Heading %d Char", i+1))
		hdngChar.SetBasedOn(dpf.StyleID())
		hdngChar.SetLinkedStyle(id)
		hdngChar.SetUISortOrder(9 + i)
		hdngChar.RunProperties().SetSize(fontSizes[i] * measurement.Point)

		hdng := s.AddStyle(id, wml.ST_StyleTypeParagraph, false)
		hdng.SetName(fmt.Sprintf("heading %d", i+1))
		hdng.SetNextStyle(normal.StyleID())
		hdng.SetLinkedStyle(hdng.StyleID())
		hdng.SetUISortOrder(9 + i)
		hdng.SetPrimaryStyle(true)
		hdng.ParagraphProperties().SetKeepNext(true)
		hdng.ParagraphProperties().SetSpacing(spacing[i]*measurement.Twips, 0)
		hdng.ParagraphProperties().SetOutlineLevel(i)
		hdng.RunProperties().SetSize(fontSizes[i] * measurement.Point)
	}
}

func (s Styles) initializeDocDefaults() {
	s.x.DocDefaults = wml.NewCT_DocDefaults()
	s.x.DocDefaults.RPrDefault = wml.NewCT_RPrDefault()
	s.x.DocDefaults.RPrDefault.RPr = wml.NewCT_RPr()

	rpr := RunStyleProperties{s.x.DocDefaults.RPrDefault.RPr}
	rpr.SetSize(12 * measurement.Point)
	rpr.Fonts().SetASCIITheme(wml.ST_ThemeMajorAscii)
	rpr.Fonts().SetEastAsiaTheme(wml.ST_ThemeMajorEastAsia)
	rpr.Fonts().SetHANSITheme(wml.ST_ThemeMajorHAnsi)
	rpr.Fonts().SetCSTheme(wml.ST_ThemeMajorBidi)

	rpr.X().Lang = wml.NewCT_Language()
	rpr.X().Lang.ValAttr = gooxml.String("en-US")
	rpr.X().Lang.EastAsiaAttr = gooxml.String("en-US")
	rpr.X().Lang.BidiAttr = gooxml.String("ar-SA")

	s.x.DocDefaults.PPrDefault = wml.NewCT_PPrDefault()
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
