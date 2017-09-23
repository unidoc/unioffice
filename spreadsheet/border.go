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

// Border is a cell border configuraton.
type Border struct {
	x       *sml.CT_Border
	borders *sml.CT_Borders
}

// X returns the inner wrapped XML type.
func (b Border) X() *sml.CT_Border {
	return b.x
}

// InitializeDefaults initializes a border to its defaulte empty values.
func (b Border) InitializeDefaults() {
	b.x.Left = sml.NewCT_BorderPr()
	b.x.Bottom = sml.NewCT_BorderPr()
	b.x.Right = sml.NewCT_BorderPr()
	b.x.Top = sml.NewCT_BorderPr()
	b.x.Diagonal = sml.NewCT_BorderPr()
}

// Index returns the index of the border for use with a cell style.
func (b Border) Index() uint32 {
	for i, bx := range b.borders.Border {
		if bx == b.x {
			return uint32(i)
		}
	}
	return 0
}

func (b Border) SetLeft(style sml.ST_BorderStyle, c color.Color) {
	if b.x.Left == nil {
		b.x.Left = sml.NewCT_BorderPr()
	}
	b.x.Left.Color = sml.NewCT_Color()
	b.x.Left.Color.RgbAttr = c.AsRGBAString()
	b.x.Left.StyleAttr = style
}

func (b Border) SetRight(style sml.ST_BorderStyle, c color.Color) {
	if b.x.Right == nil {
		b.x.Right = sml.NewCT_BorderPr()
	}
	b.x.Right.Color = sml.NewCT_Color()
	b.x.Right.Color.RgbAttr = c.AsRGBAString()
	b.x.Right.StyleAttr = style
}

func (b Border) SetTop(style sml.ST_BorderStyle, c color.Color) {
	if b.x.Top == nil {
		b.x.Top = sml.NewCT_BorderPr()
	}
	b.x.Top.Color = sml.NewCT_Color()
	b.x.Top.Color.RgbAttr = c.AsRGBAString()
	b.x.Top.StyleAttr = style
}

func (b Border) SetBottom(style sml.ST_BorderStyle, c color.Color) {
	if b.x.Bottom == nil {
		b.x.Bottom = sml.NewCT_BorderPr()
	}
	b.x.Bottom.Color = sml.NewCT_Color()
	b.x.Bottom.Color.RgbAttr = c.AsRGBAString()
	b.x.Bottom.StyleAttr = style
}

func (b Border) SetDiagonal(style sml.ST_BorderStyle, c color.Color, up, down bool) {
	if b.x.Diagonal == nil {
		b.x.Diagonal = sml.NewCT_BorderPr()
	}
	b.x.Diagonal.Color = sml.NewCT_Color()
	b.x.Diagonal.Color.RgbAttr = c.AsRGBAString()
	b.x.Diagonal.StyleAttr = style

	if up {
		b.x.DiagonalUpAttr = gooxml.Bool(true)
	}
	if down {
		b.x.DiagonalDownAttr = gooxml.Bool(true)
	}
}
