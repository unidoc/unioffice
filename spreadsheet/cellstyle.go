// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheet

import (
	"github.com/unidoc/unioffice"
	"github.com/unidoc/unioffice/schema/soo/sml"
)

// CellStyle is a formatting style for a cell.  CellStyles are spreadsheet global
// and can be applied to cells across sheets.
type CellStyle struct {
	wb  *Workbook
	xf  *sml.CT_Xf
	xfs *sml.CT_CellXfs
}

// HasNumberFormat returns true if the cell style has a number format applied.
func (cs CellStyle) HasNumberFormat() bool {
	return cs.xf.NumFmtIdAttr != nil && cs.xf.ApplyNumberFormatAttr != nil &&
		*cs.xf.ApplyNumberFormatAttr
}

// NumberFormat returns the number format that the cell style uses, or zero if
// it is not set.
func (cs CellStyle) NumberFormat() uint32 {
	if cs.xf.NumFmtIdAttr == nil {
		return 0
	}
	return *cs.xf.NumFmtIdAttr
}

// ClearNumberFormat removes any number formatting from the style.
func (cs CellStyle) ClearNumberFormat() {
	cs.xf.NumFmtIdAttr = nil
	cs.xf.ApplyNumberFormatAttr = nil
}

// SetNumberFormatStandard sets the format based off of the ECMA 376 standard formats.  These
// formats are standardized and don't need to be defined in the styles.
func (cs CellStyle) SetNumberFormatStandard(s StandardFormat) {
	cs.xf.NumFmtIdAttr = unioffice.Uint32(uint32(s))
	cs.xf.ApplyNumberFormatAttr = unioffice.Bool(true)
}

func (cs CellStyle) SetNumberFormat(s string) {
	nf := cs.wb.StyleSheet.AddNumberFormat()
	nf.SetFormat(s)
	cs.xf.ApplyNumberFormatAttr = unioffice.Bool(true)
	cs.xf.NumFmtIdAttr = unioffice.Uint32(nf.ID())
}

// Wrapped returns true if the cell will wrap text.
func (cs CellStyle) Wrapped() bool {
	if cs.xf.Alignment == nil {
		return false
	}
	if cs.xf.Alignment.WrapTextAttr == nil {
		return false
	}
	return *cs.xf.Alignment.WrapTextAttr
}

// SetWrapped configures the cell to wrap text.
func (cs CellStyle) SetWrapped(b bool) {
	if cs.xf.Alignment == nil {
		cs.xf.Alignment = sml.NewCT_CellAlignment()
	}
	if !b {
		cs.xf.Alignment.WrapTextAttr = nil
	} else {
		cs.xf.Alignment.WrapTextAttr = unioffice.Bool(true)
		cs.xf.ApplyAlignmentAttr = unioffice.Bool(true)
	}
}

// SetHorizontalAlignment sets the horizontal alignment of a cell style.
func (cs CellStyle) SetHorizontalAlignment(a sml.ST_HorizontalAlignment) {
	if cs.xf.Alignment == nil {
		cs.xf.Alignment = sml.NewCT_CellAlignment()
	}
	cs.xf.Alignment.HorizontalAttr = a
	cs.xf.ApplyAlignmentAttr = unioffice.Bool(true)
}

// SetRotation configures the cell to be rotated.
func (cs CellStyle) SetRotation(deg uint8) {
	if cs.xf.Alignment == nil {
		cs.xf.Alignment = sml.NewCT_CellAlignment()
	}
	cs.xf.ApplyAlignmentAttr = unioffice.Bool(true)
	cs.xf.Alignment.TextRotationAttr = unioffice.Uint8(deg)
}

// SetVerticalAlignment sets the vertical alignment of a cell style.
func (cs CellStyle) SetVerticalAlignment(a sml.ST_VerticalAlignment) {
	if cs.xf.Alignment == nil {
		cs.xf.Alignment = sml.NewCT_CellAlignment()
	}
	cs.xf.ApplyAlignmentAttr = unioffice.Bool(true)
	cs.xf.Alignment.VerticalAttr = a
}

func (cs CellStyle) SetShrinkToFit(b bool) {
	if cs.xf.Alignment == nil {
		cs.xf.Alignment = sml.NewCT_CellAlignment()
	}
	cs.xf.ApplyAlignmentAttr = unioffice.Bool(true)
	if !b {
		cs.xf.Alignment.ShrinkToFitAttr = nil
	} else {
		cs.xf.Alignment.ShrinkToFitAttr = unioffice.Bool(b)
	}
}

// ClearFont clears any font configuration from the cell style.
func (cs CellStyle) ClearFont() {
	cs.xf.FontIdAttr = nil
	cs.xf.ApplyFontAttr = nil
}

// SetFont applies a font to a cell style.  The font is referenced by its
// index so modifying the font afterward will affect all styles that reference
// it.
func (cs CellStyle) SetFont(f Font) {
	cs.xf.FontIdAttr = unioffice.Uint32(f.Index())
	cs.xf.ApplyFontAttr = unioffice.Bool(true)
}

// SetBorder applies a border to a cell style.  The border is referenced by its
// index so modifying the border afterward will affect all styles that reference
// it.
func (cs CellStyle) SetBorder(b Border) {
	cs.xf.BorderIdAttr = unioffice.Uint32(b.Index())
	cs.xf.ApplyBorderAttr = unioffice.Bool(true)
}

// ClearBorder clears any border configuration from the cell style.
func (cs CellStyle) ClearBorder() {
	cs.xf.BorderIdAttr = nil
	cs.xf.ApplyBorderAttr = nil
}

// SetFill applies a fill to a cell style.  The fill is referenced by its index
// so modifying the fill afterward will affect all styles that reference it.
func (cs CellStyle) SetFill(f Fill) {
	cs.xf.FillIdAttr = unioffice.Uint32(f.Index())
	cs.xf.ApplyFillAttr = unioffice.Bool(true)
}

// ClearFill clears any fill configuration from the cell style.
func (cs CellStyle) ClearFill() {
	cs.xf.FillIdAttr = nil
	cs.xf.ApplyFillAttr = nil
}

func (cs CellStyle) Index() uint32 {
	for i, xf := range cs.xfs.Xf {
		if cs.xf == xf {
			return uint32(i)
		}
	}
	return 0
}
