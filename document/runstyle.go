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

// RunStyle controls run styling properties
type RunStyle struct {
	x *wml.CT_RPr
}

// X returns the inner wrapped XML type.
func (r RunStyle) X() *wml.CT_RPr {
	return r.x
}

// SetSize sets the font size for a run.
func (r RunStyle) SetSize(size measurement.Distance) {
	var sz *wml.CT_HpsMeasure
	var szCs *wml.CT_HpsMeasure

	for _, b := range r.x.EG_RPrBase {
		if b.Sz != nil {
			sz = b.Sz
		}
		if b.SzCs != nil {
			szCs = b.SzCs
		}
	}
	if sz == nil {
		b := wml.NewEG_RPrBase()
		b.Sz = wml.NewCT_HpsMeasure()
		sz = b.Sz
		if szCs == nil {
			b.SzCs = wml.NewCT_HpsMeasure()
			szCs = b.SzCs
		}
		r.x.EG_RPrBase = append(r.x.EG_RPrBase, b)
	}
	if szCs == nil {
		b := wml.NewEG_RPrBase()
		b.SzCs = wml.NewCT_HpsMeasure()
		szCs = b.SzCs
		r.x.EG_RPrBase = append(r.x.EG_RPrBase, b)
	}

	sz.ValAttr.ST_UnsignedDecimalNumber = gooxml.Uint64(uint64(size / measurement.HalfPoint))
	szCs.ValAttr.ST_UnsignedDecimalNumber = gooxml.Uint64(uint64(size / measurement.HalfPoint))
}
