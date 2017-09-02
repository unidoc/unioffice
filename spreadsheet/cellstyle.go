// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheet

import (
	"baliance.com/gooxml"
	sml "baliance.com/gooxml/schema/schemas.openxmlformats.org/spreadsheetml"
)

// CellStyle is a formatting style for a cell.  CellStyles are spreadsheet global
// and can be applied to cells across sheets.
type CellStyle struct {
	xf  *sml.CT_Xf
	xfs *sml.CT_CellXfs
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
		cs.xf.Alignment.WrapTextAttr = gooxml.Bool(true)
		cs.xf.ApplyAlignmentAttr = gooxml.Bool(true)
	}
}

// SetHorizontalAlignment sets the horizontal alignment of a cell style.
func (cs CellStyle) SetHorizontalAlignment(a sml.ST_HorizontalAlignment) {
	if cs.xf.Alignment == nil {
		cs.xf.Alignment = sml.NewCT_CellAlignment()
	}
	cs.xf.Alignment.HorizontalAttr = a
	cs.xf.ApplyAlignmentAttr = gooxml.Bool(true)
}

// SetVerticalAlignment sets the vertical alignment of a cell style.
func (cs CellStyle) SetVerticalAlignment(a sml.ST_VerticalAlignment) {
	if cs.xf.Alignment == nil {
		cs.xf.Alignment = sml.NewCT_CellAlignment()
	}
	cs.xf.ApplyAlignmentAttr = gooxml.Bool(true)
	cs.xf.Alignment.VerticalAttr = a
}

func (cs CellStyle) SetShrinkToFit(b bool) {
	if cs.xf.Alignment == nil {
		cs.xf.Alignment = sml.NewCT_CellAlignment()
	}
	cs.xf.ApplyAlignmentAttr = gooxml.Bool(true)
	if !b {
		cs.xf.Alignment.ShrinkToFitAttr = nil
	} else {
		cs.xf.Alignment.ShrinkToFitAttr = gooxml.Bool(b)
	}
}

func (cs CellStyle) ClearFont() {
	cs.xf.FontIdAttr = nil
	cs.xf.ApplyFontAttr = nil
}

func (cs CellStyle) SetFont(f Font) {
	cs.xf.FontIdAttr = gooxml.Uint32(f.Index())
	cs.xf.ApplyFontAttr = gooxml.Bool(true)
}

func (cs CellStyle) SetFill(f Fill) {
	cs.xf.FillIdAttr = gooxml.Uint32(f.Index())
	cs.xf.ApplyFillAttr = gooxml.Bool(true)
}

func (cs CellStyle) Index() uint32 {
	for i, xf := range cs.xfs.Xf {
		if cs.xf == xf {
			return uint32(i)
		}
	}
	return 0
}
