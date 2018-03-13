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
	"baliance.com/gooxml/schema/soo/wml"
)

// CellProperties are a table cells properties within a document.
type CellProperties struct {
	x *wml.CT_TcPr
}

// X returns the inner wrapped XML type.
func (c CellProperties) X() *wml.CT_TcPr {
	return c.x
}

// SetColumnSpan sets the number of Grid Columns Spanned by the Cell.  This is used
// to give the appearance of merged cells.
func (c CellProperties) SetColumnSpan(cols int) {
	if cols == 0 {
		c.x.GridSpan = nil
	} else {
		c.x.GridSpan = wml.NewCT_DecimalNumber()
		c.x.GridSpan.ValAttr = int64(cols)
	}
}

// SetVerticalMerge controls the vertical merging of cells.
func (c CellProperties) SetVerticalMerge(mergeVal wml.ST_Merge) {
	if mergeVal == wml.ST_MergeUnset {
		c.x.VMerge = nil
	} else {
		c.x.VMerge = wml.NewCT_VMerge()
		c.x.VMerge.ValAttr = mergeVal
	}
}

// SetWidthAuto sets the the cell width to automatic.
func (c CellProperties) SetWidthAuto() {
	c.x.TcW = wml.NewCT_TblWidth()
	c.x.TcW.TypeAttr = wml.ST_TblWidthAuto
}

// SetWidthPercent sets the cell to a width percentage.
func (c CellProperties) SetWidthPercent(pct float64) {
	c.x.TcW = wml.NewCT_TblWidth()
	c.x.TcW.TypeAttr = wml.ST_TblWidthPct
	c.x.TcW.WAttr = &wml.ST_MeasurementOrPercent{}
	c.x.TcW.WAttr.ST_DecimalNumberOrPercent = &wml.ST_DecimalNumberOrPercent{}
	// percent value is measured in 1/50'th of a percent
	c.x.TcW.WAttr.ST_DecimalNumberOrPercent.ST_UnqualifiedPercentage = gooxml.Int64(int64(pct * 50))
}

// SetWidth sets the cell width to a specified width.
func (c CellProperties) SetWidth(d measurement.Distance) {
	c.x.TcW = wml.NewCT_TblWidth()
	c.x.TcW.TypeAttr = wml.ST_TblWidthDxa
	c.x.TcW.WAttr = &wml.ST_MeasurementOrPercent{}
	c.x.TcW.WAttr.ST_DecimalNumberOrPercent = &wml.ST_DecimalNumberOrPercent{}
	c.x.TcW.WAttr.ST_DecimalNumberOrPercent.ST_UnqualifiedPercentage = gooxml.Int64(int64(d / measurement.Twips))
}

// SetShading controls the cell shading.
func (c CellProperties) SetShading(shd wml.ST_Shd, foreground, fill color.Color) {
	if shd == wml.ST_ShdUnset {
		c.x.Shd = nil
	} else {
		c.x.Shd = wml.NewCT_Shd()
		c.x.Shd.ValAttr = shd
		c.x.Shd.ColorAttr = &wml.ST_HexColor{}
		if foreground.IsAuto() {
			c.x.Shd.ColorAttr.ST_HexColorAuto = wml.ST_HexColorAutoAuto
		} else {
			c.x.Shd.ColorAttr.ST_HexColorRGB = foreground.AsRGBString()
		}
		c.x.Shd.FillAttr = &wml.ST_HexColor{}
		if fill.IsAuto() {
			c.x.Shd.FillAttr.ST_HexColorAuto = wml.ST_HexColorAutoAuto
		} else {
			c.x.Shd.FillAttr.ST_HexColorRGB = fill.AsRGBString()
		}
	}
}

// SetVerticalAlignment sets the vertical alignment of content within a table cell.
func (c CellProperties) SetVerticalAlignment(align wml.ST_VerticalJc) {
	if align == wml.ST_VerticalJcUnset {
		c.x.VAlign = nil
	} else {
		c.x.VAlign = wml.NewCT_VerticalJc()
		c.x.VAlign.ValAttr = align
	}
}

// Borders allows controlling individual cell borders.
func (c CellProperties) Borders() CellBorders {
	if c.x.TcBorders == nil {
		c.x.TcBorders = wml.NewCT_TcBorders()
	}
	return CellBorders{c.x.TcBorders}
}
