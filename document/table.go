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
	wml "baliance.com/gooxml/schema/schemas.openxmlformats.org/wordprocessingml"
)

type Table struct {
	d *Document
	x *wml.CT_Tbl
}

func (t Table) X() *wml.CT_Tbl {
	return t.x
}
func (t Table) ensurePr() {
	if t.x.TblPr == nil {
		t.x.TblPr = wml.NewCT_TblPr()
	}
}

func (t Table) AddRow() Row {
	c := wml.NewEG_ContentRowContent()
	t.x.EG_ContentRowContent = append(t.x.EG_ContentRowContent, c)
	tr := wml.NewCT_Row()
	c.Tr = append(c.Tr, tr)
	return Row{t.d, tr}
}

func (t Table) SetCellSpacing() {
	t.ensurePr()
	t.x.TblPr.TblCellSpacing = wml.NewCT_TblWidth()
}

func (t Table) SetWidthAuto() {
	t.ensurePr()
	t.x.TblPr.TblW = wml.NewCT_TblWidth()
	t.x.TblPr.TblW.TypeAttr = wml.ST_TblWidthAuto
}

func (t Table) SetWidthPercent(v float64) {
	t.ensurePr()
	t.x.TblPr.TblW = wml.NewCT_TblWidth()
	t.x.TblPr.TblW.TypeAttr = wml.ST_TblWidthPct
	t.x.TblPr.TblW.WAttr = &wml.ST_MeasurementOrPercent{}
	t.x.TblPr.TblW.WAttr.ST_DecimalNumberOrPercent = &wml.ST_DecimalNumberOrPercent{}
	// percent value is measured in 1/50'th of a percent
	t.x.TblPr.TblW.WAttr.ST_DecimalNumberOrPercent.ST_UnqualifiedPercentage = gooxml.Int64(int64(v * 50))
}

func (t Table) SetWidth(d measurement.Distance) {
	t.ensurePr()
	t.x.TblPr.TblW = wml.NewCT_TblWidth()
	t.x.TblPr.TblW.TypeAttr = wml.ST_TblWidthDxa
	t.x.TblPr.TblW.WAttr = &wml.ST_MeasurementOrPercent{}
	t.x.TblPr.TblW.WAttr.ST_DecimalNumberOrPercent = &wml.ST_DecimalNumberOrPercent{}
	t.x.TblPr.TblW.WAttr.ST_DecimalNumberOrPercent.ST_UnqualifiedPercentage = gooxml.Int64(int64(d / measurement.Twips))
}

func (t Table) Borders() Borders {
	if t.x.TblPr.TblBorders == nil {
		t.x.TblPr.TblBorders = wml.NewCT_TblBorders()
	}
	return Borders{t.x.TblPr.TblBorders}
}
