// Copyright 2018 FoxyUtils ehf. All rights reserved.
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

// CellMargins are the margins for an individual cell.
type CellMargins struct {
	x *wml.CT_TcMar
}

func setCellMarginPercent(w *wml.CT_TblWidth, pct float64) {
	w.TypeAttr = wml.ST_TblWidthPct
	w.WAttr = &wml.ST_MeasurementOrPercent{}
	w.WAttr.ST_DecimalNumberOrPercent = &wml.ST_DecimalNumberOrPercent{}
	w.WAttr.ST_DecimalNumberOrPercent.ST_UnqualifiedPercentage = unioffice.Int64(int64(pct * 50))
}
func setCellMargin(w *wml.CT_TblWidth, d measurement.Distance) {
	w.TypeAttr = wml.ST_TblWidthDxa
	w.WAttr = &wml.ST_MeasurementOrPercent{}
	w.WAttr.ST_DecimalNumberOrPercent = &wml.ST_DecimalNumberOrPercent{}
	w.WAttr.ST_DecimalNumberOrPercent.ST_UnqualifiedPercentage = unioffice.Int64(int64(d / measurement.Dxa))
}

// SetTopPct sets the cell top margin
func (c CellMargins) SetTopPct(pct float64) {
	c.x.Top = wml.NewCT_TblWidth()
	setCellMarginPercent(c.x.Top, pct)
}

// SetTop sets the cell top margin
func (c CellMargins) SetTop(d measurement.Distance) {
	c.x.Top = wml.NewCT_TblWidth()
	setCellMargin(c.x.Top, d)
}

// SetLeftPct sets the cell left margin
func (c CellMargins) SetLeftPct(pct float64) {
	c.x.Left = wml.NewCT_TblWidth()
	setCellMarginPercent(c.x.Left, pct)
}

// SetLeft sets the cell left margin
func (c CellMargins) SetLeft(d measurement.Distance) {
	c.x.Left = wml.NewCT_TblWidth()
	setCellMargin(c.x.Left, d)
}

// SetRightPct sets the cell right margin
func (c CellMargins) SetRightPct(pct float64) {
	c.x.Right = wml.NewCT_TblWidth()
	setCellMarginPercent(c.x.Right, pct)
}

// SetRight sets the cell right margin
func (c CellMargins) SetRight(d measurement.Distance) {
	c.x.Right = wml.NewCT_TblWidth()
	setCellMargin(c.x.Right, d)
}

// SetBottomPct sets the cell bottom margin
func (c CellMargins) SetBottomPct(pct float64) {
	c.x.Bottom = wml.NewCT_TblWidth()
	setCellMarginPercent(c.x.Bottom, pct)
}

// SetBottom sets the cell bottom margin
func (c CellMargins) SetBottom(d measurement.Distance) {
	c.x.Bottom = wml.NewCT_TblWidth()
	setCellMargin(c.x.Bottom, d)
}

// SetStartPct sets the cell start margin
func (c CellMargins) SetStartPct(pct float64) {
	c.x.Start = wml.NewCT_TblWidth()
	setCellMarginPercent(c.x.Start, pct)
}

// SetStart sets the cell start margin
func (c CellMargins) SetStart(d measurement.Distance) {
	c.x.Start = wml.NewCT_TblWidth()
	setCellMargin(c.x.Start, d)
}
