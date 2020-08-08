// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io.

package document

import (
	"github.com/unidoc/unioffice"
	"github.com/unidoc/unioffice/measurement"
	"github.com/unidoc/unioffice/schema/soo/ofc/sharedTypes"
	"github.com/unidoc/unioffice/schema/soo/wml"
)

// RowProperties are the properties for a row within a table
type RowProperties struct {
	x *wml.CT_TrPr
}

// SetHeight allows controlling the height of a row within a table.
func (r RowProperties) SetHeight(ht measurement.Distance, rule wml.ST_HeightRule) {
	if rule == wml.ST_HeightRuleUnset {
		r.x.TrHeight = nil
	} else {
		htv := wml.NewCT_Height()
		htv.HRuleAttr = rule
		htv.ValAttr = &sharedTypes.ST_TwipsMeasure{}
		htv.ValAttr.ST_UnsignedDecimalNumber = unioffice.Uint64(uint64(ht / measurement.Twips))
		r.x.TrHeight = []*wml.CT_Height{htv}
	}
}
