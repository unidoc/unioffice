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

// TableWidth controls width values in table settings.
type TableWidth struct {
	x *wml.CT_TblWidth
}

// NewTableWidth returns a newly intialized TableWidth
func NewTableWidth() TableWidth {
	return TableWidth{wml.NewCT_TblWidth()}
}

// X returns the inner wrapped XML type.
func (s TableWidth) X() *wml.CT_TblWidth {
	return s.x
}

// SetValue sets the width value.
func (s TableWidth) SetValue(m measurement.Distance) {
	s.x.WAttr = &wml.ST_MeasurementOrPercent{}
	s.x.WAttr.ST_DecimalNumberOrPercent = &wml.ST_DecimalNumberOrPercent{}
	s.x.WAttr.ST_DecimalNumberOrPercent.ST_UnqualifiedPercentage = gooxml.Int64(int64(m / measurement.Twips))
	s.x.TypeAttr = wml.ST_TblWidthDxa
}
