// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package styles

import (
	"errors"

	"baliance.com/gooxml"
	sml "baliance.com/gooxml/schema/schemas.openxmlformats.org/spreadsheetml"
)

// StyleSheet is a document style sheet.
type StyleSheet struct {
	x *sml.StyleSheet
}

// NewStyleSheet constructs a new default stylesheet.
func NewStyleSheet() StyleSheet {
	ss := sml.NewStyleSheet()
	b := NewBorders()
	ss.Borders = b.X()
	ss.CellStyleXfs = sml.NewCT_CellStyleXfs()
	ss.CellXfs = sml.NewCT_CellXfs()
	ss.CellStyles = sml.NewCT_CellStyles()

	cs := sml.NewCT_CellStyle()
	cs.NameAttr = gooxml.String("Normal")
	cs.XfIdAttr = 0
	cs.BuiltinIdAttr = gooxml.Uint32(0)
	ss.CellStyles.CellStyle = append(ss.CellStyles.CellStyle, cs)
	ss.CellStyles.CountAttr = gooxml.Uint32(uint32(len(ss.CellStyles.CellStyle)))

	xf := sml.NewCT_Xf()
	xf.NumFmtIdAttr = gooxml.Uint32(0)
	xf.FontIdAttr = gooxml.Uint32(0)
	xf.FillIdAttr = gooxml.Uint32(0)
	xf.BorderIdAttr = gooxml.Uint32(0)
	ss.CellStyleXfs.Xf = append(ss.CellStyleXfs.Xf, xf)
	ss.CellStyleXfs.CountAttr = gooxml.Uint32(uint32(len(ss.CellStyleXfs.Xf)))

	fills := NewFills()
	ss.Fills = fills.X()
	fill := fills.AddPatternFill()
	fill.SetPattern(sml.ST_PatternTypeNone)
	fill = fills.AddPatternFill()
	fill.SetPattern(sml.ST_PatternTypeGray125)

	ss.Fonts = sml.NewCT_Fonts()

	s := StyleSheet{ss}
	fnt := s.AddFont()
	fnt.SetName("Calibri")
	fnt.SetSize(11)

	xf2 := sml.NewCT_Xf()
	*xf2 = *xf
	xf2.XfIdAttr = gooxml.Uint32(0)

	ss.CellXfs.Xf = append(ss.CellXfs.Xf, xf2)
	ss.CellXfs.CountAttr = gooxml.Uint32(uint32(len(ss.CellXfs.Xf)))
	return s
}

// X returns the innter XML entity for a stylesheet.
func (s StyleSheet) X() *sml.StyleSheet {
	return s.x
}

// AddFont adds a new empty font to the stylesheet.
func (s StyleSheet) AddFont() Font {
	font := sml.NewCT_Font()
	s.x.Fonts.Font = append(s.x.Fonts.Font, font)
	s.x.Fonts.CountAttr = gooxml.Uint32(uint32(len(s.x.Fonts.Font)))
	return NewFont(font, s.x)
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
		ret = append(ret, NewFont(f, s.x))
	}
	return ret
}

// AddCellStyle adds a new empty cell style to the stylesheet.
func (s StyleSheet) AddCellStyle() CellStyle {
	xf := sml.NewCT_Xf()
	s.x.CellXfs.Xf = append(s.x.CellXfs.Xf, xf)
	s.x.CellXfs.CountAttr = gooxml.Uint32(uint32(len(s.x.CellXfs.Xf)))
	return NewCellStyle(xf, s.x.CellXfs)
}

// Fills returns a Fills object that can be used to add/create/edit fills.
func (s StyleSheet) Fills() Fills {
	return Fills{s.x.Fills}
}
