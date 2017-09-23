// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package document

import (
	"baliance.com/gooxml"
	"baliance.com/gooxml/measurement"
	"baliance.com/gooxml/schema/soo/wml"
)

// TableProperties are the properties for a table within a document
type TableProperties struct {
	x *wml.CT_TblPr
}

// X returns the inner wrapped XML type.
func (t TableProperties) X() *wml.CT_TblPr {
	return t.x
}

// SetCellSpacingAuto sets the cell spacing within a table to automatic.
func (t TableProperties) SetCellSpacingAuto() {
	t.x.TblCellSpacing = wml.NewCT_TblWidth()
	t.x.TblCellSpacing.TypeAttr = wml.ST_TblWidthAuto
}

// SetCellSpacingPercent sets the cell spacing within a table to a percent width.
func (t TableProperties) SetCellSpacingPercent(pct float64) {
	t.x.TblCellSpacing = wml.NewCT_TblWidth()
	t.x.TblCellSpacing.TypeAttr = wml.ST_TblWidthPct
	t.x.TblCellSpacing.WAttr = &wml.ST_MeasurementOrPercent{}
	t.x.TblCellSpacing.WAttr.ST_DecimalNumberOrPercent = &wml.ST_DecimalNumberOrPercent{}
	t.x.TblCellSpacing.WAttr.ST_DecimalNumberOrPercent.ST_UnqualifiedPercentage = gooxml.Int64(int64(pct * 50))
}

// SetCellSpacing sets the cell spacing within a table.
func (t TableProperties) SetCellSpacing(m measurement.Distance) {
	t.x.TblCellSpacing = wml.NewCT_TblWidth()
	t.x.TblCellSpacing.TypeAttr = wml.ST_TblWidthDxa
	t.x.TblCellSpacing.WAttr = &wml.ST_MeasurementOrPercent{}
	t.x.TblCellSpacing.WAttr.ST_DecimalNumberOrPercent = &wml.ST_DecimalNumberOrPercent{}
	t.x.TblCellSpacing.WAttr.ST_DecimalNumberOrPercent.ST_UnqualifiedPercentage = gooxml.Int64(int64(m / measurement.Dxa))
}

// SetWidthAuto sets the the table width to automatic.
func (t TableProperties) SetWidthAuto() {
	t.x.TblW = wml.NewCT_TblWidth()
	t.x.TblW.TypeAttr = wml.ST_TblWidthAuto
}

// SetWidthPercent sets the table to a width percentage.
func (t TableProperties) SetWidthPercent(pct float64) {
	t.x.TblW = wml.NewCT_TblWidth()
	t.x.TblW.TypeAttr = wml.ST_TblWidthPct
	t.x.TblW.WAttr = &wml.ST_MeasurementOrPercent{}
	t.x.TblW.WAttr.ST_DecimalNumberOrPercent = &wml.ST_DecimalNumberOrPercent{}
	// percent value is measured in 1/50'th of a percent
	t.x.TblW.WAttr.ST_DecimalNumberOrPercent.ST_UnqualifiedPercentage = gooxml.Int64(int64(pct * 50))
}

// SetWidth sets the table with to a specified width.
func (t TableProperties) SetWidth(d measurement.Distance) {
	t.x.TblW = wml.NewCT_TblWidth()
	t.x.TblW.TypeAttr = wml.ST_TblWidthDxa
	t.x.TblW.WAttr = &wml.ST_MeasurementOrPercent{}
	t.x.TblW.WAttr.ST_DecimalNumberOrPercent = &wml.ST_DecimalNumberOrPercent{}
	t.x.TblW.WAttr.ST_DecimalNumberOrPercent.ST_UnqualifiedPercentage = gooxml.Int64(int64(d / measurement.Twips))
}

// Borders allows manipulation of the table borders.
func (t TableProperties) Borders() TableBorders {
	if t.x.TblBorders == nil {
		t.x.TblBorders = wml.NewCT_TblBorders()
	}
	return TableBorders{t.x.TblBorders}
}
