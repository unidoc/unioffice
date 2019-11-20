// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased on https://unidoc.io.

package document

import (
	"github.com/unidoc/unioffice"
	"github.com/unidoc/unioffice/measurement"
	"github.com/unidoc/unioffice/schema/soo/wml"
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
	t.x.TblCellSpacing.WAttr.ST_DecimalNumberOrPercent.ST_UnqualifiedPercentage = unioffice.Int64(int64(pct * 50))
}

// SetCellSpacing sets the cell spacing within a table.
func (t TableProperties) SetCellSpacing(m measurement.Distance) {
	t.x.TblCellSpacing = wml.NewCT_TblWidth()
	t.x.TblCellSpacing.TypeAttr = wml.ST_TblWidthDxa
	t.x.TblCellSpacing.WAttr = &wml.ST_MeasurementOrPercent{}
	t.x.TblCellSpacing.WAttr.ST_DecimalNumberOrPercent = &wml.ST_DecimalNumberOrPercent{}
	t.x.TblCellSpacing.WAttr.ST_DecimalNumberOrPercent.ST_UnqualifiedPercentage = unioffice.Int64(int64(m / measurement.Dxa))
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
	t.x.TblW.WAttr.ST_DecimalNumberOrPercent.ST_UnqualifiedPercentage = unioffice.Int64(int64(pct * 50))
}

// SetLayout controls the table layout. wml.ST_TblLayoutTypeAutofit corresponds
// to "Automatically resize to fit contents" being checked, while
// wml.ST_TblLayoutTypeFixed corresponds to it being unchecked.
func (t TableProperties) SetLayout(l wml.ST_TblLayoutType) {
	// ST_TblLayoutTypeAutofit is the default
	if l == wml.ST_TblLayoutTypeUnset || l == wml.ST_TblLayoutTypeAutofit {
		t.x.TblLayout = nil
	} else {
		t.x.TblLayout = wml.NewCT_TblLayoutType()
		t.x.TblLayout.TypeAttr = l
	}
}

// SetAlignment sets the alignment of a table within the page.
func (t TableProperties) SetAlignment(align wml.ST_JcTable) {
	if align == wml.ST_JcTableUnset {
		t.x.Jc = nil
	} else {
		t.x.Jc = wml.NewCT_JcTable()
		t.x.Jc.ValAttr = align
	}
}

// SetWidth sets the table with to a specified width.
func (t TableProperties) SetWidth(d measurement.Distance) {
	t.x.TblW = wml.NewCT_TblWidth()
	t.x.TblW.TypeAttr = wml.ST_TblWidthDxa
	t.x.TblW.WAttr = &wml.ST_MeasurementOrPercent{}
	t.x.TblW.WAttr.ST_DecimalNumberOrPercent = &wml.ST_DecimalNumberOrPercent{}
	t.x.TblW.WAttr.ST_DecimalNumberOrPercent.ST_UnqualifiedPercentage = unioffice.Int64(int64(d / measurement.Twips))
}

// SetStyle sets the table style name.
func (t TableProperties) SetStyle(name string) {
	if name == "" {
		t.x.TblStyle = nil
	} else {
		t.x.TblStyle = wml.NewCT_String()
		t.x.TblStyle.ValAttr = name
	}
}

// TableLook returns the table look, or conditional formatting applied to a table style.
func (t TableProperties) TableLook() TableLook {
	if t.x.TblLook == nil {
		t.x.TblLook = wml.NewCT_TblLook()
	}
	return TableLook{t.x.TblLook}
}

// Borders allows manipulation of the table borders.
func (t TableProperties) Borders() TableBorders {
	if t.x.TblBorders == nil {
		t.x.TblBorders = wml.NewCT_TblBorders()
	}
	return TableBorders{t.x.TblBorders}
}
