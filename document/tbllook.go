// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased on https://unidoc.io.

package document

import (
	st "github.com/unidoc/unioffice/schema/soo/ofc/sharedTypes"
	"github.com/unidoc/unioffice/schema/soo/wml"
)

// TableLook is the conditional formatting associated with a table style that
// has been assigned to a table.
type TableLook struct {
	x *wml.CT_TblLook
}

// X returns the inner wrapped XML type.
func (t TableLook) X() *wml.CT_TblLook {
	return t.x
}

// SetFirstColumn controls the conditional formatting for the first column in a table.
func (t TableLook) SetFirstColumn(on bool) {
	if !on {
		t.x.FirstColumnAttr = &st.ST_OnOff{}
		t.x.FirstColumnAttr.ST_OnOff1 = st.ST_OnOff1Off
	} else {
		t.x.FirstColumnAttr = &st.ST_OnOff{}
		t.x.FirstColumnAttr.ST_OnOff1 = st.ST_OnOff1On
	}
}

// SetFirstRow controls the conditional formatting for the first row in a table.
func (t TableLook) SetFirstRow(on bool) {
	if !on {
		t.x.FirstRowAttr = &st.ST_OnOff{}
		t.x.FirstRowAttr.ST_OnOff1 = st.ST_OnOff1Off
	} else {
		t.x.FirstRowAttr = &st.ST_OnOff{}
		t.x.FirstRowAttr.ST_OnOff1 = st.ST_OnOff1On
	}
}

// SetLastColumn controls the conditional formatting for the last column in a table.
func (t TableLook) SetLastColumn(on bool) {
	if !on {
		t.x.LastColumnAttr = &st.ST_OnOff{}
		t.x.LastColumnAttr.ST_OnOff1 = st.ST_OnOff1Off
	} else {
		t.x.LastColumnAttr = &st.ST_OnOff{}
		t.x.LastColumnAttr.ST_OnOff1 = st.ST_OnOff1On
	}
}

// SetLastRow controls the conditional formatting for the last row in a table.
// This is called the 'Total' row within Word.
func (t TableLook) SetLastRow(on bool) {
	if !on {
		t.x.LastRowAttr = &st.ST_OnOff{}
		t.x.LastRowAttr.ST_OnOff1 = st.ST_OnOff1Off
	} else {
		t.x.LastRowAttr = &st.ST_OnOff{}
		t.x.LastRowAttr.ST_OnOff1 = st.ST_OnOff1On
	}
}

// SetHorizontalBanding controls the conditional formatting for horizontal banding.
func (t TableLook) SetHorizontalBanding(on bool) {
	if !on {
		// inverted logic
		t.x.NoHBandAttr = &st.ST_OnOff{}
		t.x.NoHBandAttr.ST_OnOff1 = st.ST_OnOff1On
	} else {
		t.x.NoHBandAttr = &st.ST_OnOff{}
		t.x.NoHBandAttr.ST_OnOff1 = st.ST_OnOff1Off
	}
}

// SetVerticalBanding controls the conditional formatting for vertical banding.
func (t TableLook) SetVerticalBanding(on bool) {
	if !on {
		// inverted logic
		t.x.NoVBandAttr = &st.ST_OnOff{}
		t.x.NoVBandAttr.ST_OnOff1 = st.ST_OnOff1On
	} else {
		t.x.NoVBandAttr = &st.ST_OnOff{}
		t.x.NoVBandAttr.ST_OnOff1 = st.ST_OnOff1Off
	}
}
