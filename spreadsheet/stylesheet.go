// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheet

import (
	"errors"

	"github.com/unidoc/unioffice"
	"github.com/unidoc/unioffice/schema/soo/sml"
)

// StyleSheet is a document style sheet.
type StyleSheet struct {
	wb *Workbook
	x  *sml.StyleSheet
}

// NewStyleSheet constructs a new default stylesheet.
func NewStyleSheet(wb *Workbook) StyleSheet {
	ss := sml.NewStyleSheet()

	ss.CellStyleXfs = sml.NewCT_CellStyleXfs()
	ss.CellXfs = sml.NewCT_CellXfs()
	ss.CellStyles = sml.NewCT_CellStyles()

	cs := sml.NewCT_CellStyle()
	cs.NameAttr = unioffice.String("Normal")
	cs.XfIdAttr = 0
	cs.BuiltinIdAttr = unioffice.Uint32(0)
	ss.CellStyles.CellStyle = append(ss.CellStyles.CellStyle, cs)
	ss.CellStyles.CountAttr = unioffice.Uint32(uint32(len(ss.CellStyles.CellStyle)))

	xf := sml.NewCT_Xf()
	xf.NumFmtIdAttr = unioffice.Uint32(0)
	xf.FontIdAttr = unioffice.Uint32(0)
	xf.FillIdAttr = unioffice.Uint32(0)
	xf.BorderIdAttr = unioffice.Uint32(0)
	ss.CellStyleXfs.Xf = append(ss.CellStyleXfs.Xf, xf)
	ss.CellStyleXfs.CountAttr = unioffice.Uint32(uint32(len(ss.CellStyleXfs.Xf)))

	fills := NewFills()
	ss.Fills = fills.X()
	fill := fills.AddFill().SetPatternFill()
	fill.SetPattern(sml.ST_PatternTypeNone)
	fill = fills.AddFill().SetPatternFill()
	fill.SetPattern(sml.ST_PatternTypeGray125)

	ss.Fonts = sml.NewCT_Fonts()

	ss.Borders = sml.NewCT_Borders()

	s := StyleSheet{wb, ss}
	// default empty border
	s.AddBorder().InitializeDefaults()

	fnt := s.AddFont()
	fnt.SetName("Calibri")
	fnt.SetSize(11)

	xf2 := sml.NewCT_Xf()
	*xf2 = *xf
	xf2.XfIdAttr = unioffice.Uint32(0)

	ss.CellXfs.Xf = append(ss.CellXfs.Xf, xf2)
	ss.CellXfs.CountAttr = unioffice.Uint32(uint32(len(ss.CellXfs.Xf)))
	return s
}

// X returns the inner XML entity for a stylesheet.
func (s StyleSheet) X() *sml.StyleSheet {
	return s.x
}

// AddFont adds a new empty font to the stylesheet.
func (s StyleSheet) AddFont() Font {
	font := sml.NewCT_Font()
	s.x.Fonts.Font = append(s.x.Fonts.Font, font)
	s.x.Fonts.CountAttr = unioffice.Uint32(uint32(len(s.x.Fonts.Font)))
	return Font{font, s.x}
}

// AddBorder creates a new empty border that can be applied to a cell style.
func (s StyleSheet) AddBorder() Border {
	b := sml.NewCT_Border()
	s.x.Borders.Border = append(s.x.Borders.Border, b)
	s.x.Borders.CountAttr = unioffice.Uint32(uint32(len(s.x.Borders.Border)))
	return Border{b, s.x.Borders}
}

// RemoveFont removes a font from the style sheet.  It *does not* update styles that refer
// to this font.
func (s StyleSheet) RemoveFont(f Font) error {
	for i, sf := range s.x.Fonts.Font {
		if sf == f.X() {
			s.x.Fonts.Font = append(s.x.Fonts.Font[:i],
				s.x.Fonts.Font[i+1:]...)
			return nil
		}
	}
	return errors.New("font not found")
}

// Fonts returns the list of fonts defined in the stylesheet.
func (s StyleSheet) Fonts() []Font {
	ret := []Font{}
	for _, f := range s.x.Fonts.Font {
		ret = append(ret, Font{f, s.x})
	}
	return ret
}

// AddCellStyle adds a new empty cell style to the stylesheet.
func (s StyleSheet) AddCellStyle() CellStyle {
	xf := sml.NewCT_Xf()
	s.x.CellXfs.Xf = append(s.x.CellXfs.Xf, xf)
	s.x.CellXfs.CountAttr = unioffice.Uint32(uint32(len(s.x.CellXfs.Xf)))
	return CellStyle{s.wb, xf, s.x.CellXfs}
}

// CellStyles returns the list of defined cell styles
func (s StyleSheet) CellStyles() []CellStyle {
	ret := []CellStyle{}
	for _, xf := range s.x.CellXfs.Xf {
		ret = append(ret, CellStyle{s.wb, xf, s.x.CellXfs})
	}
	return ret
}

// AddDifferentialStyle adds a new empty differential cell style to the stylesheet.
func (s StyleSheet) AddDifferentialStyle() DifferentialStyle {
	if s.x.Dxfs == nil {
		s.x.Dxfs = sml.NewCT_Dxfs()
	}
	dxf := sml.NewCT_Dxf()
	s.x.Dxfs.Dxf = append(s.x.Dxfs.Dxf, dxf)
	s.x.Dxfs.CountAttr = unioffice.Uint32(uint32(len(s.x.Dxfs.Dxf)))
	return DifferentialStyle{dxf, s.wb, s.x.Dxfs}
}

// GetOrCreateStandardNumberFormat gets or creates a cell style with a given
// standard format. This should only be used when you want to perform
// number/date/time formatting only.  Manipulating the style returned will cause
// all cells using style returned from this for a given format to be formatted.
func (s StyleSheet) GetOrCreateStandardNumberFormat(f StandardFormat) CellStyle {
	for _, cs := range s.CellStyles() {
		// found an existing number format
		if cs.HasNumberFormat() && cs.NumberFormat() == uint32(f) {
			return cs
		}
	}
	// need to create a new format
	cs := s.AddCellStyle()
	cs.SetNumberFormatStandard(f)
	return cs
}

// AddNumberFormat adds a new blank number format to the stylesheet.
func (s StyleSheet) AddNumberFormat() NumberFormat {
	if s.x.NumFmts == nil {
		s.x.NumFmts = sml.NewCT_NumFmts()
	}
	nf := sml.NewCT_NumFmt()
	// start our IDs at 200 so we can ensure we don't conflict with any
	// pre-defined formats
	nf.NumFmtIdAttr = uint32(200 + len(s.x.NumFmts.NumFmt))
	s.x.NumFmts.NumFmt = append(s.x.NumFmts.NumFmt, nf)
	s.x.NumFmts.CountAttr = unioffice.Uint32(uint32(len(s.x.NumFmts.NumFmt)))
	return NumberFormat{s.wb, nf}
}

// Fills returns a Fills object that can be used to add/create/edit fills.
func (s StyleSheet) Fills() Fills {
	return Fills{s.x.Fills}
}

func (s StyleSheet) GetCellStyle(id uint32) CellStyle {
	for i, f := range s.x.CellXfs.Xf {
		if uint32(i) == id {
			return CellStyle{s.wb, f, s.x.CellXfs}
		}
	}
	return CellStyle{}
}

func (s StyleSheet) GetNumberFormat(id uint32) NumberFormat {
	if id >= 0 && id < 50 {
		return CreateDefaultNumberFormat(StandardFormat(id))
	}
	for _, nf := range s.x.NumFmts.NumFmt {
		if nf.NumFmtIdAttr == id {
			return NumberFormat{s.wb, nf}
		}
	}
	return NumberFormat{}
}
